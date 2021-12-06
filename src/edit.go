package main

/*
editor di testo, permette di editare una riga per volta inserendo riga e contenuto voluto.
supporta inoltre l'evidenziazione della sintassi (dove per ogni linguaggio deve essere implementata)
*/

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"unicode"
	"strconv"
	"strings"
	"encoding/json"
)

type importantWords_type struct {
	Types []string 
	Keywords []string 
} 
var keywords importantWords_type

type colors_type struct {
	/*
	"base": "\u001b[31m",
    "reset": "\u001b[0m",
    "keywords": "\u001b[34m",
    "types": "\u001b[35m"
	*/
	Base string
	Reset string
	Keywords string
	Types string
}
var colormap colors_type

func JoinLinee(linee []string) string{
	var sb strings.Builder
	for _,v := range linee {
		sb.WriteString(v + "\n")
	}
	return sb.String()
}
func LeggiLinee(in io.Reader)  []string {
	scanner := bufio.NewScanner(in)
	var text_lines []string 
	for scanner.Scan(){
		text_lines = append(text_lines, scanner.Text())
	}
	return text_lines
}
func LeggiLinea(in io.Reader) string{
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	return  scanner.Text()
}
func IndiceRiga(num_righe int) int {
	var inp_ind string

	
	if num_righe == 0{
		return 0
	}
	

	for {
		fmt.Print(":")
		fmt.Scan(&inp_ind);
		// è composto tutto da cifre?
		for _,v := range []rune(inp_ind) {
			if !unicode.IsNumber(v){
				continue
			}
		}
		num, err := strconv.Atoi(inp_ind)
		if err != nil{
			continue
		}else{
			if num >= 0 && num < num_righe  {
				return num
			}
		}
	} 
}
func IsStringaNumero(s string) bool {
	for _,v := range []rune(s) {
		if !unicode.IsNumber(v){
			return false
		}
	}
	return true
}
func SliceContains(s []string, v string) bool {
    for _, a := range s {
        if a ==  v{
            return true
        }
    }
    return false
}

func PrintLines(lines []string){
	// stampo riga per riga con l'indice a sinistra
	for i:=0; i < len(lines); i++{
		// per evidenziare le parole:
		/*
			separo ogni linea in tokens divisi dallo spazio
			se il token corrente è incluso in una delle slice di keywords lo coloro con il colore adatto
			altrimenti rimane con il colore standard
		*/
		fmt.Print(i	, "\t: ")
		tokens := strings.Split(lines[i], " ")
		for _,v := range tokens{
			
			switch{
			case SliceContains(keywords.Keywords, v):
				// stampo l'escape code per il colore delle keyword
				fmt.Print(colormap.Keywords)
			
			case SliceContains(keywords.Types, v):
				// stampo l'escape code per il colore delle keyword
				fmt.Print(colormap.Types)
			default:
				fmt.Print(colormap.Base)
			}
			// stampo il token 
			fmt.Print(v + " ")
			// resetto il colore
			fmt.Print(colormap.Reset)
		}
		// vado a capo
		fmt.Println()
	}
}
func ScriviSlice(writer *bufio.Writer, linee []string){
	for _,v := range linee{
		v += "\n"
		fmt.Fprint(writer, v)
	}
	writer.Flush()
}
func AggiungiLinea(linee []string, lineaDaAggiungere string, pos int) []string{
	if len(linee) == pos{
		return append(linee, lineaDaAggiungere)
	}
	linee = append(linee[:pos+1], linee[pos:]...)
	linee[pos] = lineaDaAggiungere
	return linee
}
func EliminaLinea(linee []string, pos int) []string{
	linee = append(linee[:pos], linee[pos+1:]...)
	return linee
}
func pulisciSchermo(){
	fmt.Print("\033[H\033[2J")
}

func main(){
	// prendi nome file da argomenti
	var file_name string
	if len(os.Args) > 1{
		file_name = os.Args[1]
	}else{
		fmt.Println("Usage: edit filename")
		return
	}

	// se il file ha un'estensione cerco nella cartella keywords per le keyword di quel linguaggio
	// questo perchè almeno poi nella stampa posso evidenziarle
	if strings.Index(file_name, ".") != -1{
		// carica le keyword nella struttura dati dal file json corrispondente
		keyword_path := "config/keywords/" + strings.Split(file_name,".")[1] +".json"
		// SE il file esiste (i linguaggi possono anche non essere supportati)
		// carico il json contenuto nel file in una struttura adeguata
		// altrimenti metto in quella struttura nil -> così posso controllare nella stampa
		_, err := os.Stat(keyword_path)
		if err == nil{
			// carico il file
			f, err := os.Open(keyword_path)
			if err == nil {
				// carico tutto il file in una sola stringa
				tmp_string := strings.Join(LeggiLinee(f), "")
				err = json.Unmarshal([]byte(tmp_string), &keywords)
			}
			f.Close()
		}
	}
	// carico i colori
	var color_path string = "config/colors.json"
	_, err := os.Stat(color_path)
	if err == nil{
		// carico il file
		f, err := os.Open(color_path)
		if err == nil {
			// carico tutto il file in una sola stringa
			tmp_string := strings.Join(LeggiLinee(f), "")
			err = json.Unmarshal([]byte(tmp_string), &colormap)
		}
		f.Close()
	}
	
	var lines []string
	_, err = os.Stat(file_name)
	if err == nil{
		// il file esiste? se si leggo il contenuto altrimenti creo il file
		f, err := os.OpenFile(file_name, os.O_RDONLY|os.O_CREATE, 0755)
		if err != nil{
			fmt.Println("Error opening / creating file:", err)
		}
		// ora il file è creato / stato aperto 

		lines = LeggiLinee(f)
		
		f.Close()
	}
	// ora ho letto tutte le linee
	var quit bool
	// CICLO OPERAZIONI PARTE DA QUA
	for !quit {
		// pulisco lo schermo e scrivo a schermo tutte le linee
		pulisciSchermo()
		PrintLines(lines)
		// prendo il comando in input
		/*
			comandi disponibili: 
				i x n -> con n e x numero 
					inserisce n linee vuote dopo la posizione x
		*/
		// prendo in input l'indice della riga che voglio modificare
		var com string
		fmt.Print(":")
		com = LeggiLinea(os.Stdin)
		switch {
		case []rune(com)[0] == 'b':
			// devo controllare che vi siano x e n
			cmd_args := strings.Split(com, " ")
			if len(cmd_args) == 3 && IsStringaNumero(cmd_args[1]) && IsStringaNumero(cmd_args[2]){
				// prendo i numeri come interi per comodità
				var x, n int
				x, _ = strconv.Atoi(cmd_args[1])
				n, _= strconv.Atoi(cmd_args[2])
				// eseguo il comando
				for i:= 0; i < n; i++{
					lines = AggiungiLinea(lines, " ", x)
				}
			}
		case []rune(com)[0] == 'i':
			// devo controllare che vi siano x e n
			cmd_args := strings.Split(com, " ")
			
			// tutto cioò che viene messo dopo il secondo spazio viene preso come testo da inserire
			if len(cmd_args) >= 3 && IsStringaNumero(cmd_args[1]){
				
				// prendo posizione come stringa e il testo da inserire nella posizione come stringa
				stringa_da_inserire := strings.Join(strings.Split(com, " ")[2:], " ")
				i, _ := strconv.Atoi(cmd_args[1])
				fmt.Println(len(lines)-1, " i:", i)
				// se la riga selezionata non è inclusa nella lunghezza corrente allora inserisco righe bianche fino a quando lo è
				if i > len(lines){
					lun := len(lines)
					for j := 0; j < lun; j++{
						lines = AggiungiLinea(lines, " ", len(lines)-1)
					}
				}
				// eseguo il comando
				lines = AggiungiLinea(lines, stringa_da_inserire, i)
				
			}
		case []rune(com)[0] == 'e':
			// devo controllare che vi siano x e n
			cmd_args := strings.Split(com, " ")
			// tutto cioò che viene messo dopo il secondo spazio viene preso come testo da inserire
			if len(cmd_args) >= 3 && IsStringaNumero(cmd_args[1]){
				// prendo posizione come stringa e il testo da inserire nella posizione come stringa
				stringa_da_inserire := strings.Join(strings.Split(com, " ")[2:], " ")
				i, _ := strconv.Atoi(cmd_args[1])
				fmt.Println("stringa_da_inserire")
				fmt.Println(stringa_da_inserire)
				// eseguo il comando
				lines[i] = stringa_da_inserire
			}
		case []rune(com)[0] == 'D':
			// elimina la riga nell'indice cmd_args[1]
			cmd_args := strings.Split(com, " ")
			if len(cmd_args) == 2 && IsStringaNumero(cmd_args[1]){
				i, _ := strconv.Atoi(cmd_args[1])
				lines = EliminaLinea(lines, i)
			}
		case []rune(com)[0] == 'd':
			// mette una riga vuota alla posizione cmd_args[1]
			cmd_args := strings.Split(com, " ")
			if len(cmd_args) == 2 && IsStringaNumero(cmd_args[1]){
				i, _ := strconv.Atoi(cmd_args[1])
				lines[i] = " "
			}
		
		case []rune(com)[0] == 'q':
			quit = true
			
		}
		fmt.Println(lines)
		// fmt.Println("Insert the new value for line", indice, ":")
		// ora ho all'interno di indice la riga voluta:
		// leggo da input del testo (con la funz LeggiTesto)
		// e lo metto all'interno dell'elemento di indice indice della slice

		// stringa_input := JoinLinee(LeggiLinee(os.Stdin))
		// lines[indice] = stringa_input
	}
	// CICLO OPERAZIONI FINISCE QUA

	// riapro file in scrittura stavolta
	f, err := os.OpenFile(file_name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil{
		fmt.Println("Error opening / creating file:", err)
	}
	// ora scrivo la slice nel file
	ScriviSlice(bufio.NewWriter(f), lines)
	// pulisco schermo
	fmt.Print("\033[H\033[2J")	
}