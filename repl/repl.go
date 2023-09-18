package repl

import (
	"bufio"
	"fmt"
	"io"

	// "github.com/Corralitz/cocolang-go/compiler"
	"github.com/Corralitz/cocolang-go/lexer"
	// "github.com/Corralitz/cocolang-go/object"
	"github.com/Corralitz/cocolang-go/parser"
	"github.com/Corralitz/cocolang-go/vm"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalSize)
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	for {
		fmt.Printf(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		// comp := compiler.NewWithState(symbolTable, constants)
		// err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation fialed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode fialed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
