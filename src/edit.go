package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"unicode"
	"strconv"
	"strings"
)
// programma che legge nome del file da argomenti e scrive quello che legge da bufio

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



func PrintLines(lines []string){
	// stampo riga per riga con l'indice a sinistra
	for i:=0; i < len(lines); i++{
		fmt.Print(i, "\t: ", lines[i] + "\n")
	}
}

func ScriviSlice(writer *bufio.Writer, linee []string){
	for _,v := range linee{
		fmt.Println([]byte(v))
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

func main(){
	// prendi nome file da argomenti
	var file_name string
	if len(os.Args) > 1{
		file_name = os.Args[1]
	}else{
		fmt.Println("Usage: edit filename")
		return
	}
	// il file esiste? se si leggo il contenuto altrimenti creo il file
	f, err := os.OpenFile(file_name, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil{
		fmt.Println("Error opening / creating file:", err)
	}
	// ora il file è creato / stato aperto 
	var lines []string
	lines = LeggiLinee(f)
	// ora ho letto tutte le linee
	var quit bool
	// CICLO OPERAZIONI PARTE DA QUA
	for !quit {
		// pulisco lo schermo e scrivo a schermo tutte le linee
		// fmt.Print("\033[H\033[2J")
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
		case []rune(com)[0] == 'i':
			// devo controllare che vi siano x e n
			cmd_args := strings.Split(com, " ")
			fmt.Println("dentro il case switch: ", cmd_args)
			
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

	f.Close()
	// riapro file in scrittura stavolta
	f, err = os.OpenFile(file_name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil{
		fmt.Println("Error opening / creating file:", err)
	}
	// ora scrivo la slice nel file
	ScriviSlice(bufio.NewWriter(f), lines)	
}