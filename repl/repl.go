package repl

import (
	"bufio"
	"fmt"
	"io"
	"layng/lexer"
	"layng/token"
)

const PROMPT = ">> "
const EXIT = "exit"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	_, _ = fmt.Fprintf(out, "Type %s to leave\n\n", EXIT)

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == EXIT {
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}