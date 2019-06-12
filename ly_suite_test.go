package ly_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLuaYaml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LuaYaml Suite")
}
