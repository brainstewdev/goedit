package main

import (
	"bufio"
	"os"
	"fmt"
)
// programma che legge nome del file da argomenti e scrive quello che legge da bufio

func LeggiTesto()  []string{
	scanner := bufio.NewScanner(os.Stdin)
	var text_lines []string 
	for scanner.Scan(){
		text_lines = append(text_lines, scanner.Text())
	}
	return text_lines
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
	// crea il file
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// imposta la chiusura di f quando viene terminato main
	defer f.Close()
	// prendi slice di testo letto
	lines := LeggiTesto();
	// scrivi sul file la riga
	for _,v := range lines{
		fmt.Fprint(f, v + "\n")
	}
}