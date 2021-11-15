package assembler

import (
	"fmt"
	"log"
	"os"

	"github.com/iris-net/hack_assembler/parser"
	"github.com/iris-net/hack_assembler/parser/command"
	symboltable "github.com/iris-net/hack_assembler/symbol_table"
)

type Assembler struct {
	asmParser parser.AssemblyParser
	symbols   symboltable.SymbolTable
	binaries  []string
}

func NewAssembler(inputFile string) (*Assembler, error) {
	p, err := parser.NewAssemblyParser(inputFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Assembler{asmParser: p}, nil
}

func (a *Assembler) Reset() {
	a.binaries = make([]string, 0, len(a.asmParser.AssemblyLines))
	a.symbols = symboltable.NewSymbolTable()
}

func (a Assembler) GetParser() parser.AssemblyParser {
	return a.asmParser
}

func (a Assembler) GetSymbols() symboltable.SymbolTable {
	return a.symbols
}

func (a Assembler) GetBinaries() []string {
	return a.binaries
}

func (a *Assembler) Execute() error {
	a.Reset()

	// load assembly to create symbol table
	err := a.SetupSymbolTable()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// load assembly again to generate binary codes with refering the symbol table
	err = a.Parse()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// SetupSymbolTable setup symbol table with loading assembly
func (a Assembler) SetupSymbolTable() error {
	rom := 0
	for a.asmParser.HasMoreCommands() {
		a.asmParser.Advance()

		cmdType, err := a.asmParser.CommandType()
		if err != nil {
			log.Fatal(err)
			return err
		}

		if cmdType == command.L {
			symbol, err := a.asmParser.Symbol()
			if err != nil {
				log.Fatal(err)
				return err
			}
			if a.symbols.Contains(symbol) {
				err = fmt.Errorf("[%s] %s has already used", a.asmParser.AssemblyLines[a.asmParser.Cursor], symbol)
				log.Fatal(err)
				return err
			}
			a.symbols.AddEntry(symbol, rom)
		} else {
			rom += 1
		}
	}

	return nil
}

// Parse generate binary codes from assembly
func (a *Assembler) Parse() error {
	ram := 16
	binaryStr := ""
	a.asmParser.ResetCursor()
	for a.asmParser.HasMoreCommands() {
		a.asmParser.Advance()

		cmdType, err := a.asmParser.CommandType()
		if err != nil {
			log.Fatal(err)
			return err
		}
		if cmdType == command.Unknown {
			err = fmt.Errorf("[%s] unknown command type", a.asmParser.AssemblyLines[a.asmParser.Cursor])
			log.Fatal(err)
			return err
		}

		switch cmdType {
		case command.L:
			continue
		case command.A:
			hasSymbol, err := a.asmParser.HasSymbol()
			if err != nil {
				log.Fatal(err)
				return err
			}

			address := 0
			if hasSymbol {
				symbol, err := a.asmParser.Symbol()
				if err != nil {
					log.Fatal(err)
					return err
				}

				if !a.symbols.Contains(symbol) {
					a.symbols.AddEntry(symbol, ram)
					ram += 1
				}

				address = a.symbols.GetAddress(symbol)
			} else {
				address, err = a.asmParser.DirectAddress()
				if err != nil {
					log.Fatal(err)
					return err
				}
			}

			binaryStr = fmt.Sprintf("0%015b", address)
		case command.C:
			comp, err := a.asmParser.Comp()
			if err != nil {
				log.Fatal(err)
				return err
			}

			dest, err := a.asmParser.Dest()
			if err != nil {
				log.Fatal(err)
				return err
			}

			jump, err := a.asmParser.Jump()
			if err != nil {
				log.Fatal(err)
				return err
			}

			binaryStr = fmt.Sprintf("111%s%s%s", comp.Binary(), dest.Binary(), jump.Binary())
		}

		a.binaries = append(a.binaries, binaryStr)
	}

	return nil
}

// ExportBinaryCode exports binary code to the given file path
func (a Assembler) ExportBinaryCode(outFile string) error {
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	for _, b := range a.binaries {
		_, err := f.WriteString(b + "\n")
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
