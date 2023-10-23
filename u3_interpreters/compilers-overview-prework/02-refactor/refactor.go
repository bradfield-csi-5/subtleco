package main

import (
	"bytes"
	"log"
	"os"
	"sort"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

const src string = `package foo

import (
	"fmt"
	"time"
)

func baz() {
	fmt.Println("Hello, world!")
}

type A int

const b = "testing"

func bar() {
	fmt.Println(time.Now())
}`

// Moves all top-level functions to the end, sorted in alphabetical order.
// The "source file" is given as a string (rather than e.g. a filename).
func SortFunctions(src string) (string, error) {
	f, err := decorator.Parse(src)
	if err != nil {
		panic(err)
	}

	var declarations []*dst.GenDecl
	var functions []*dst.FuncDecl

	// Map out AST
	for _, node := range f.Decls {
		switch nodeT := node.(type) {
		case *dst.FuncDecl:
			functions = append(functions, nodeT)
		case *dst.GenDecl:
			declarations = append(declarations, nodeT)
		}
	}

	// Sort SortFunctions
	sort.Slice(functions, func(i, j int) bool {
		return functions[i].Name.Name < functions[j].Name.Name
	})

	// Assign statements
	for i := range f.Decls {
		if i < len(declarations) {
			f.Decls[i] = declarations[i]
		} else {
			pos := i - len(declarations)
			f.Decls[i] = functions[pos]
		}
	}

	var buf bytes.Buffer
	if err := decorator.Fprint(&buf, f); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	f, err := decorator.Parse(src)
	if err != nil {
		log.Fatal(err)
	}

	// Print AST
	err = dst.Fprint(os.Stdout, f, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Convert AST back to source
	err = decorator.Print(f)
	if err != nil {
		log.Fatal(err)
	}
}
