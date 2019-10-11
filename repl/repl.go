package repl

import (
	"Interpreter/lexer"
	"Interpreter/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

// Start tokenizes source code and prints the tokens
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		// print all tokens until an EOF is encountered 
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
