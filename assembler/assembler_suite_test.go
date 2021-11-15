package assembler_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAssembler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assembler Suite")
}
