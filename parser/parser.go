package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/iris-net/hack_assembler/parser/command"
	"github.com/iris-net/hack_assembler/parser/comp"
	"github.com/iris-net/hack_assembler/parser/dest"
	"github.com/iris-net/hack_assembler/parser/jump"
)

type AssemblyParser struct {
	AssemblyLines []string
	Cursor        int
}

// load input and get ready to parse it
func NewAssemblyParser(path string) (AssemblyParser, error) {
	a := AssemblyParser{Cursor: -1}
	err := a.load(path)
	if err != nil {
		return AssemblyParser{}, err
	}
	return a, nil
}

// load input file
func (a *AssemblyParser) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	commands := make([]string, 0, 100)
	for scanner.Scan() {
		cmd := a.removeNotCommand(scanner.Text())
		if len(cmd) > 0 {
			commands = append(commands, cmd)
		}
	}

	a.AssemblyLines = commands

	return nil
}

// removeNotCommand removes not command descriptions
func (a AssemblyParser) removeNotCommand(text string) (ret string) {
	r := regexp.MustCompile(`(\s|\t|//(\w|\W)*)`)
	ret = r.ReplaceAllString(text, "")

	return ret
}

func (a *AssemblyParser) ResetCursor() {
	a.Cursor = -1
}

// HasMoreCommands returns whether there are more commands in the input
func (a AssemblyParser) HasMoreCommands() bool {
	return a.Cursor+1 < len(a.AssemblyLines)
}

// Advance reads the next command from the input and makes it the command.
func (a *AssemblyParser) Advance() {
	a.Cursor += 1
}

// CommandType returns the type of the current command
func (a AssemblyParser) CommandType() (command.Type, error) {
	line := a.AssemblyLines[a.Cursor]

	r := regexp.MustCompile(`^@`)
	if r.MatchString(line) {
		return command.A, nil
	}

	r = regexp.MustCompile(`^(AMD|AD|AM|MD|A|D|M)=`)
	if r.MatchString(line) {
		return command.C, nil
	}

	r = regexp.MustCompile(`;(JGT|JEQ|JGE|JLT|JNE|JLE|JMP)`)
	if r.MatchString(line) {
		return command.C, nil
	}

	r = regexp.MustCompile(`^\((\w|\W)+\)`)
	if r.MatchString(line) {
		return command.L, nil
	}

	return command.Unknown, fmt.Errorf("[%s] unknown command type", a.AssemblyLines[a.Cursor])
}

// HasSymbol checks whether this command has Symbol or not
func (a AssemblyParser) HasSymbol() (bool, error) {
	cType, err := a.CommandType()
	if err != nil {
		return false, err
	}

	if cType != command.A && cType != command.L {
		return false, fmt.Errorf("[%s] this isn't A or L command type", a.AssemblyLines[a.Cursor])
	}

	line := a.AssemblyLines[a.Cursor]

	r := regexp.MustCompile(`(@|\(|\))`)
	ret := r.ReplaceAllString(line, "")

	_, err = strconv.Atoi(ret)

	return err != nil, nil
}

// Symbol returns the symbol of the current command.
func (a AssemblyParser) Symbol() (string, error) {
	cType, err := a.CommandType()
	if err != nil {
		return "", err
	}

	if cType != command.A && cType != command.L {
		return "", fmt.Errorf("[%s] this isn't A or L command type", a.AssemblyLines[a.Cursor])
	}

	line := a.AssemblyLines[a.Cursor]

	r := regexp.MustCompile(`(@|\(|\))`)
	ret := r.ReplaceAllString(line, "")

	return ret, nil
}

// DirectAddress return direct memory address
func (a AssemblyParser) DirectAddress() (int, error) {
	cType, err := a.CommandType()
	if err != nil {
		return 0, err
	}

	if cType != command.A && cType != command.L {
		return 0, fmt.Errorf("[%s] this isn't A or L command type", a.AssemblyLines[a.Cursor])
	}

	line := a.AssemblyLines[a.Cursor]

	r := regexp.MustCompile(`(@|\(|\))`)
	ret := r.ReplaceAllString(line, "")

	address, err := strconv.Atoi(ret)
	if err != nil {
		return 0, fmt.Errorf("[%s] this command don't specify any direct address", a.AssemblyLines[a.Cursor])
	}

	return address, nil
}

// Dest returns the dest mnemonic in the curent C-command
func (a AssemblyParser) Dest() (dest.Mnemonic, error) {
	cType, err := a.CommandType()
	if err != nil {
		return dest.Unknown, err
	}

	if cType != command.C {
		return dest.Unknown, fmt.Errorf("[%s] this isn't C command type", a.AssemblyLines[a.Cursor])
	}

	line := a.AssemblyLines[a.Cursor]

	r := regexp.MustCompile(`^(AMD|AD|AM|MD|A|D|M)=`)
	matches := r.FindStringSubmatch(line)
	if len(matches) == 0 {
		return dest.Null, nil
	}

	d := dest.NewMnemonic(matches[1])

	return d, nil
}

// Comp returns the comp mnemonic in the curent C-command
func (a AssemblyParser) Comp() (comp.Mnemonic, error) {
	cType, err := a.CommandType()
	if err != nil {
		return comp.Unknown, err
	}

	if cType != command.C {
		return comp.Unknown, fmt.Errorf("[%s] this isn't C command type", a.AssemblyLines[a.Cursor])
	}

	line := a.AssemblyLines[a.Cursor]

	pattern := `(D\|A|D\|M|D&A|D&M|A\-D|M\-D|D\-A|D\-M|D\+A|D\+M|A\-1|M\-1|D\-1|A\+1|M\+1|D\+1|\-A|\-M|\-D|!A|!M|!D|A|M|D|\-1|1|0)`

	r := regexp.MustCompile(fmt.Sprintf(`=%s$`, pattern))
	matches := r.FindStringSubmatch(line)
	if len(matches) == 0 {
		r = regexp.MustCompile(fmt.Sprintf(`^%s;`, pattern))
		matches = r.FindStringSubmatch(line)

		if len(matches) == 0 {
			return comp.Unknown, fmt.Errorf("[%s] unknown comp command. %s", a.AssemblyLines[a.Cursor], line)
		}
	}

	d := comp.NewMnemonic(matches[1])

	return d, nil
}

// Jump returns the jump mnemonic in the current C-Command
func (a AssemblyParser) Jump() (jump.Mnemonic, error) {
	cType, err := a.CommandType()
	if err != nil {
		return jump.Unknown, err
	}

	if cType != command.C {
		return jump.Unknown, fmt.Errorf("[%s] this isn't C command type", a.AssemblyLines[a.Cursor])
	}

	line := a.AssemblyLines[a.Cursor]

	r := regexp.MustCompile(`;(JGT|JEQ|JGE|JLT|JNE|JLE|JMP)`)
	matches := r.FindStringSubmatch(line)
	if len(matches) == 0 {
		return jump.Null, nil
	}

	j := jump.NewMnemonic(matches[1])

	return j, nil
}
