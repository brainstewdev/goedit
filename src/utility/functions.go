package utility

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
	"os"
	"regexp"
	"golang.org/x/term"
)

var currentScheme SchemeType
var keywords ImportantWordsType
var terminalLines, terminalCols int

var tokenExpr *regexp.Regexp

func Initialize(){
	lines, cols, ok := terminalSize(int(os.Stdin.Fd()));
	if ok == nil {
		fmt.Println("tutto ok, nessun errore, ho ottenuto", lines, cols);
		terminalLines, terminalCols = lines, cols
		tokenExpr = regexp.MustCompile(`[\w!]+|.`)
	} else {
		terminalLines, terminalCols = 0, 0
		fmt.Println("errore nell'inizializzazione: ", ok);
	}
}
func terminalSize(fd int) (cols, lines int, ok  error) {

	if term.IsTerminal(fd){
		lines, cols, ok = term.GetSize(fd);				
	}else{
		lines, cols = 0, 0;
	}
	return;
}
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

func SetScheme(newScheme SchemeType) {
	currentScheme = newScheme
}
func SetWords(wordSet ImportantWordsType) {
	keywords = wordSet
}

// funzione che imposta il colore da stampare per l'output
func SetColor(rgb string, background bool) {
	//ESC[ 38;2;⟨r⟩;⟨g⟩;⟨b⟩ m Select RGB foreground color
	escape := "\033"
	if !background {
		fmt.Print(escape + "[38;2;" + rgb + "m")
	} else {
		fmt.Print(escape + "[48;2;" + rgb + "m")
	}
}

// function which returns the tokens in the line
func GetTokens(line string) []string{
	return tokenExpr.FindAllString(line, -1)
}

func PrintLines(lines []string, from int) {
	// stampo riga per riga con l'indice a sinistra
	for i := from; i-from < terminalLines-1; i++ {
		if i >= len(lines) {
			fmt.Print("-")
		} else {
			tokens := GetTokens(lines[i])
			var printingString bool

			SetColor(currentScheme.Reset, false)
			SetColor(currentScheme.Background, true)

			// print the line number
			fmt.Print(i, "\t: ")

			// print each token
			for _, v := range tokens {
				switch{
				case SliceContains(keywords.Keywords, v):
					// the token is a keyword, set the color for the keywords
					SetColor(currentScheme.Keywords, false)
				case SliceContains(keywords.Types, v):
					SetColor(currentScheme.Types, false)
				case TokenContainsValidNumber(v):
					SetColor(currentScheme.Numbers, false)
				case v == "\"":
					// if is already printing a string just change the state of printingString
					// otherwise set printingString to true
					if !printingString{
						SetColor(currentScheme.StringsLiterals, false)
						printingString = true
					} else{
						printingString = false
					}
				default:
					if !printingString{
						SetColor(currentScheme.Base, false)
					}
				}
				// print the token
				fmt.Print(v)
				if !printingString{
					SetColor(currentScheme.Base, false)
				}
			}
		}	
		// vado a capo
		fmt.Print("\n\r");
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
func Move(lines []string, start_line *int) {
	// read a character directly from the keyboard.
	// if the character is j -> go up a line
	// if the character is k -> go down a line
	// if the character is ESC return to the normal cicle
	// fd 0 is stdin
	state, err := term.MakeRaw(0)
	if err != nil {
		return
	}
	defer func() {
		if err := term.Restore(0, state); err != nil {
			SetColor("255;0;0", false)
			fmt.Println("Error restoring terminal state!")
		}
	}()

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err != nil {
			break
		}
		if r == 'j' {
			if *start_line > 0 {
				*start_line--
			}
		}
		if r == 'k' {
			if *start_line < len(lines) {
				*start_line++
			}
		}

		if r == 'q' {
			break
		}

		PrintLines(lines, *start_line)

	}
}
