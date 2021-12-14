package utility

import (
	"fmt"
	"strings"
	"unicode"
	"bufio"
	"io"
	"strconv"
)

var currentScheme SchemeType
var keywords ImportantWordsType

func JoinLines(linee []string) string {
	var sb strings.Builder
	for _, v := range linee {
		sb.WriteString(v + "\n")
	}
	return sb.String()
}
func ReadLines(in io.Reader) []string {
	scanner := bufio.NewScanner(in)
	var text_lines []string
	for scanner.Scan() {
		text_lines = append(text_lines, scanner.Text())
	}
	return text_lines
}
func ReadLine(in io.Reader) string {
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	return scanner.Text()
}
func RowIndex(num_righe int) int {
	var inp_ind string

	if num_righe == 0 {
		return 0
	}

	for {
		fmt.Print(":")
		fmt.Scan(&inp_ind)
		// è composto tutto da cifre?
		for _, v := range []rune(inp_ind) {
			if !unicode.IsNumber(v) {
				continue
			}
		}
		num, err := strconv.Atoi(inp_ind)
		if err != nil {
			continue
		} else {
			if num >= 0 && num < num_righe {
				return num
			}
		}
	}
}
func IsStringNumber(s string) bool {
	for _, v := range []rune(s) {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}
func IsStringaRealNumber(s string) bool {
	for _, v := range []rune(s) {
		if !(unicode.IsNumber(v) || v == '.') {
			return false
		}
	}
	return true
}

// funzione che restituisce se la stringa contiene un numero reale
/*
	nel caso si abbia ad esempio in c++ un'assegnamento ad un numero reale si avrà il seguente caso:
		int x = 10;
	e 10 non sarà considerato numero reale quindi non evidenziato
	devo creare una funzione che prenda e controlli se prima di tutto la stringa contiene un numero reale
	poi devo ciclare nella stampa e utilizzare il color code giusto per stampare i singoli caratteri del numero
	poi arrivato ad un carattere che non è un numero allora mi fermo
	se vi è un numero attaccato (nello stesso token) ad un punto e virgola allora è valido
	in tutti gli altri casi no -> non evidenziare quindi restituisco falso
*/
func TokenContainsValidNumber(s string) bool {
	valid := true
	for _, v := range []rune(s) {
		if !(unicode.IsNumber(v) || v == '.' || v == ';') {
			valid = false
		}
	}
	return valid
}
func SliceContains(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

func SetScheme(newScheme SchemeType){
	currentScheme = newScheme
}
func SetWords(wordSet ImportantWordsType){
	keywords = wordSet
}

func PrintLines(lines []string) {
	
	// stampo riga per riga con l'indice a sinistra
	for i := 0; i < len(lines); i++ {
		// per evidenziare le parole:
		/*
			separo ogni linea in tokens divisi dallo spazio
			se il token corrente è incluso in una delle slice di keywords lo coloro con il colore adatto
			altrimenti rimane con il colore standard
		*/
		// SE sono state aperte delle virgolette continuo a stampare con il colore delle stringhe letterali
		// fino a quando incontro una parola che contiene delle virgolette
		// nel testo fra virgolette NON ci deve essere evidenziazione sintattica
		printing_string_literal := false
		fmt.Print(currentScheme.Reset)
		fmt.Print(currentScheme.Background)
		fmt.Print(i, "\t: ")
		fmt.Print(currentScheme.Base)
		tokens := strings.Split(lines[i], " ")
		for _, v := range tokens {
			// controllo, se il token corrente contiene un " allora devo stampare la parte fino a quello e poi metto printing_string_literal a vero
			disable_print := false
			if indice_start := strings.Index(v, "\""); indice_start != -1 {
				// controllo se è un token che inizia e finisce con le virgolette 
				if ([]rune(v))[0] == ([]rune(v))[len([]rune(v))-1] &&  string(([]rune(v))[0]) == "\""{
					fmt.Print(currentScheme.StringsLiterals, v, currentScheme.Reset)
					disable_print = true
				} else{
				if !printing_string_literal {
					printing_string_literal = true
					// quello che avviene quando becco la prima virgoletta
					// stampo quello che c'è prima delle virgolette
					fmt.Print(v[:indice_start])
					// stampo quello che c'è dopo
					fmt.Print(currentScheme.StringsLiterals)
					fmt.Print((v[indice_start:]))
					fmt.Print(" ")
					disable_print = true
					}else{
						printing_string_literal = false
						fmt.Print(v[:indice_start+1])
						// resetta il colore e mette printing_string_literal a false
						fmt.Print(currentScheme.Reset)
						fmt.Print((v[indice_start+1:]))
						disable_print = true
						fmt.Print(" ")
				}
				}
			}

			if !printing_string_literal {
				switch {
				case SliceContains(keywords.Keywords, v):
					// stampo l'escape code per il colore delle keyword
					fmt.Print(currentScheme.Keywords)

				case SliceContains(keywords.Types, v):
					// stampo l'escape code per il colore delle keyword
					fmt.Print(currentScheme.Types)
				case TokenContainsValidNumber(v):
					fmt.Print(currentScheme.Numbers)
				default:
					fmt.Print(currentScheme.Base)
				}
			}
			// stampo il token
			if !disable_print{
			fmt.Print(v)
			}
			// resetto il colore
			if !printing_string_literal {
				fmt.Print(currentScheme.Reset)
			}
			fmt.Print(" ")
		}
		// vado a capo
		fmt.Println()
	}
}
func WriteSlice(writer *bufio.Writer, linee []string) {
	for _, v := range linee {
		v += "\n"
		fmt.Fprint(writer, v)
	}
	writer.Flush()
}
func AddLine(linee []string, lineaDaAggiungere string, pos int) []string {
	if len(linee) == pos {
		return append(linee, lineaDaAggiungere)
	}
	linee = append(linee[:pos+1], linee[pos:]...)
	linee[pos] = lineaDaAggiungere
	return linee
}
func RemoveLine(linee []string, pos int) []string {
	linee = append(linee[:pos], linee[pos+1:]...)
	return linee
}
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}