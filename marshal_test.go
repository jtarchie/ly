package ly_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/jtarchie/ly"
	lua "github.com/yuin/gopher-lua"
)

var _ = Describe("Marhsal", func() {
	DescribeTable("handles different tables successfully", func(source string, finalContents string) {
		l := lua.NewState()
		defer l.Close()

		err := l.DoString(source)
		Expect(err).NotTo(HaveOccurred())

		table := l.ToTable(-1)
		Expect(table).NotTo(BeNil())

		contents, err := ly.Marshal(table)
		Expect(err).NotTo(HaveOccurred())
		Expect(contents).To(MatchYAML(finalContents))
	},
		Entry("boolean truth", "return {bool=true}", "bool: true"),
		Entry("boolean false", "return {bool=false}", "bool: false"),
		Entry("int", "return {num=1}", "num: 1"),
		Entry("float", "return {num=1.1}", "num: 1.1"),
		Entry("string", `return {str="123"}`, `str: "123"`),
		Entry("array", `return {arr={1, 2, "3"}}`, `arr: [1,2,"3"]`),
		Entry("nested map", `return {a={b={c=1}}}`, `a: {b: {c: 1}}`),
	)

	DescribeTable("errors with tables", func(source string, errMsg string) {
		l := lua.NewState()
		defer l.Close()

		err := l.DoString(source)
		Expect(err).NotTo(HaveOccurred())

		table := l.ToTable(-1)
		Expect(table).NotTo(BeNil())

		_, err = ly.Marshal(table)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring(errMsg))
	},
		Entry("a circular reference", `a={a=1}; a["b"]=a; return a`, "cannot encode recursively nested tables to YAML"),
		Entry("a sparse array", `return {[0]=1, [1]=2}`, "cannot encode sparse array"),
		Entry("a funtion", `return {a=function() end}`, "cannot encode"),
		Entry("a boolean as a key", `return {[true]="1", a=1}`, "cannot encode mixed or invalid key types"),
	)
})
