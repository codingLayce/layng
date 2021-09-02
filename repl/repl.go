package repl

import (
	"bufio"
	"fmt"
	"io"
	"layng/evaluator"
	"layng/lexer"
	"layng/parser"
)

const PROMPT = ">> "
const EXIT = "exit"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	_, _ = fmt.Fprintf(out, "Type %s to leave\n\n", EXIT)

	for {
		_, _ = fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == EXIT {
			return
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
