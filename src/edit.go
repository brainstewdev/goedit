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

func ScriviSlice(writer *bufio.Writer, linee []string){
	for _,v := range linee{
		fmt.Println([]byte(v))
		fmt.Fprint(writer, v)
	}
	writer.Flush()
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
	// pulisco lo schermo e scrivo a schermo tutte le linee
	fmt.Print("\033[H\033[2J")
	// stampo riga per riga con l'indice a sinistra
	for i:=0; i < len(lines); i++{
		fmt.Print(i, "\t: ", lines[i] + "\n")
	}
	// prendo in input l'indice della riga che voglio modificare
	var indice int
	indice = IndiceRiga(len(lines))
	fmt.Println("Insert the new value for line", indice, ":")
	// ora ho all'interno di indice la riga voluta:
	// leggo da input del testo (con la funz LeggiTesto)
	// e lo metto all'interno dell'elemento di indice indice della slice

	stringa_input := JoinLinee(LeggiLinee(os.Stdin))
	lines[indice] = stringa_input
	for _, v := range lines {
		fmt.Println(v)
	}

	f.Close()
	// riapro file in scrittura stavolta
	f, err = os.OpenFile(file_name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil{
		fmt.Println("Error opening / creating file:", err)
	}
	// ora scrivo la slice nel file
	ScriviSlice(bufio.NewWriter(f), lines)	
}