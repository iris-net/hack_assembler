package parser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/iris-net/hack_assembler/parser"
	"github.com/iris-net/hack_assembler/parser/command"
	"github.com/iris-net/hack_assembler/parser/comp"
	"github.com/iris-net/hack_assembler/parser/dest"
	"github.com/iris-net/hack_assembler/parser/jump"
)

const TestFile = "../test.asm"

var _ = Describe("Parser", func() {
	Context("NewAssemblyParser()", func() {
		When("Assembly file is loaded", func() {
			It("Successfully load", func() {
				p, err := parser.NewAssemblyParser(TestFile)

				expects := []string{"@R0", "D=M", "@R1", "D=D-M", "@OUTPUT_FIRST", "D;JGT", "@R1", "D=M", "@OUTPUT_D", "0;JMP", "(OUTPUT_FIRST)", "@R0", "D=M", "(OUTPUT_D)", "@R2", "M=D", "(INFINITE_LOOP)", "@INFINITE_LOOP", "0;JMP"}

				Expect(p.AssemblyLines).To(Equal(expects))
				Expect(err).To(BeNil())
			})
		})
		When("Assembly file is not existed", func() {
			It("Error", func() {
				_, err := parser.NewAssemblyParser("notfound.asm")

				Expect(err != nil).To(BeTrue())
			})
		})
	})

	Context("CommandType()", func() {
		When("All assembly commands are valid", func() {
			It("return command type", func() {
				p, err := parser.NewAssemblyParser(TestFile)
				Expect(err).To(BeNil())

				expects := []command.Type{command.A, command.C, command.A, command.C, command.A, command.C, command.A, command.C, command.A, command.C, command.L, command.A, command.C, command.L, command.A, command.C, command.L, command.A, command.C}

				for _, ex := range expects {
					p.Advance()
					ct, err := p.CommandType()
					Expect(ct).To(Equal(ex))
					Expect(err).To(BeNil())
				}
			})
		})
	})

	Context("CommandType()", func() {
		When("All assembly commands are valid", func() {
			It("return command types", func() {
				p, err := parser.NewAssemblyParser(TestFile)
				Expect(err).To(BeNil())

				expects := []command.Type{command.A, command.C, command.A, command.C, command.A, command.C, command.A, command.C, command.A, command.C, command.L, command.A, command.C, command.L, command.A, command.C, command.L, command.A, command.C}

				for _, ex := range expects {
					p.Advance()
					ct, err := p.CommandType()
					Expect(ct).To(Equal(ex))
					Expect(err).To(BeNil())
				}
			})
		})
	})

	Context("Symbol()", func() {
		When("All assembly commands are valid", func() {
			It("return symbols", func() {
				p, err := parser.NewAssemblyParser(TestFile)
				Expect(err).To(BeNil())

				expects := []string{"R0", "", "R1", "", "OUTPUT_FIRST", "", "R1", "", "OUTPUT_D", "", "OUTPUT_FIRST", "R0", "", "OUTPUT_D", "R2", "", "INFINITE_LOOP", "INFINITE_LOOP", ""}

				for _, ex := range expects {

					p.Advance()
					s, err := p.Symbol()
					Expect(s).To(Equal(ex))

					if len(ex) == 0 {
						Expect(err != nil).To(BeTrue())
					} else {
						Expect(err).To(BeNil())
					}
				}
			})
		})
	})

	Context("Dest()", func() {
		When("All assembly commands are valid", func() {
			It("return dest types", func() {
				p, err := parser.NewAssemblyParser(TestFile)
				Expect(err).To(BeNil())

				expects := []dest.Mnemonic{dest.Unknown, dest.D, dest.Unknown, dest.D, dest.Unknown, dest.Null, dest.Unknown, dest.D, dest.Unknown, dest.Null, dest.Unknown, dest.Unknown, dest.D, dest.Unknown, dest.Unknown, dest.M, dest.Unknown, dest.Unknown, dest.Null}

				for _, ex := range expects {
					p.Advance()
					d, err := p.Dest()
					Expect(d).To(Equal(ex))

					if d == dest.Unknown {
						Expect(err != nil).To(BeTrue())
					} else {
						Expect(err).To(BeNil())
					}
				}
			})
		})
	})

	Context("Comp()", func() {
		When("All assembly commands are valid", func() {
			It("return comp types", func() {
				p, err := parser.NewAssemblyParser(TestFile)
				Expect(err).To(BeNil())

				expects := []comp.Mnemonic{comp.Unknown, comp.M, comp.Unknown, comp.DMinusM, comp.Unknown, comp.D, comp.Unknown, comp.M, comp.Unknown, comp.Zero, comp.Unknown, comp.Unknown, comp.M, comp.Unknown, comp.Unknown, comp.D, comp.Unknown, comp.Unknown, comp.Zero}

				for _, ex := range expects {
					p.Advance()
					c, err := p.Comp()
					Expect(c).To(Equal(ex))

					if c == comp.Unknown {
						Expect(err != nil).To(BeTrue())
					} else {
						Expect(err).To(BeNil())
					}
				}
			})
		})
	})

	Context("Jump()", func() {
		When("All assembly commands are valid", func() {
			It("return dest types", func() {
				p, err := parser.NewAssemblyParser(TestFile)
				Expect(err).To(BeNil())

				expects := []jump.Mnemonic{jump.Unknown, jump.Null, jump.Unknown, jump.Null, jump.Unknown, jump.JGT, jump.Unknown, jump.Null, jump.Unknown, jump.JMP, jump.Unknown, jump.Unknown, jump.Null, jump.Unknown, jump.Unknown, jump.Null, jump.Unknown, jump.Unknown, jump.JMP}

				for _, ex := range expects {
					p.Advance()
					j, err := p.Jump()
					Expect(j).To(Equal(ex))

					if j == jump.Unknown {
						Expect(err != nil).To(BeTrue())
					} else {
						Expect(err).To(BeNil())
					}
				}
			})
		})
	})
})
