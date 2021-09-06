package repl

import (
	"bufio"
	"fmt"
	"io"
	"layng/evaluator"
	"layng/lexer"
	"layng/object"
	"layng/parser"

	"github.com/fatih/color"
)

const PROMPT = ">> "
const EXIT = "exit"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	errorColor := color.New(color.FgRed)
	responseColor := color.New(color.FgCyan)

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
			printParserErrors(out, errorColor, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if _, isError := evaluated.(*object.Error); isError {
				_, _ = errorColor.Fprintf(out, evaluated.Inspect())
			} else {
				_, _ = responseColor.Fprintf(out, evaluated.Inspect())
			}
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errorColor *color.Color, errors []string) {
	errorColor.Fprintf(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = errorColor.Fprintf(out, "\t"+msg+"\n")
	}
}
