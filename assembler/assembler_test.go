package assembler_test

import (
	"bufio"
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"

	"github.com/iris-net/hack_assembler/assembler"
)

const TestFile = "../test.asm"
const TestHackFile = "../test.hack"

var _ = Describe("Assembler", func() {
	Context("NewAssembler()", func() {
		When("load valid file path", func() {
			It("Successfully load", func() {
				_, err := assembler.NewAssembler(TestFile)

				Expect(err).To(gomega.BeNil())
			})
		})
	})

	Context("SetupSymbolTable()", func() {
		When("load valid assembly code", func() {
			It("setup symbol table", func() {
				a, _ := assembler.NewAssembler(TestFile)
				a.Reset()

				err := a.SetupSymbolTable()
				Expect(err).To(gomega.BeNil())

				expected := map[string]int{
					"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4,
					"R0": 0, "R1": 1, "R2": 2, "R3": 3, "R4": 4, "R5": 5, "R6": 6, "R7": 7, "R8": 8, "R9": 9, "R10": 10, "R11": 11, "R12": 12, "R13": 13, "R14": 14, "R15": 15,
					"SCREEN": 16384, "KBD": 24576,
					"OUTPUT_FIRST": 10, "OUTPUT_D": 12, "INFINITE_LOOP": 14,
				}

				ret := a.GetSymbols()
				Expect(len(ret.Table)).To(gomega.Equal(len(expected)))

				for k, v := range expected {
					value := ret.Table[k]
					Expect(value).To(gomega.Equal(v), k)
				}
			})
		})
	})

	Context("Parse()", func() {
		When("load valid assembly code", func() {
			It("parse assembly code", func() {
				a, _ := assembler.NewAssembler(TestFile)
				a.Reset()

				a.SetupSymbolTable()

				err := a.Parse()
				Expect(err).To(gomega.BeNil())

				file, err := os.Open(TestHackFile)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()

				ret := a.GetBinaries()

				expected := make([]string, 0, len(ret))
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					expected = append(expected, scanner.Text())
				}

				Expect(ret).To(gomega.Equal(expected))
			})
		})
	})
})
