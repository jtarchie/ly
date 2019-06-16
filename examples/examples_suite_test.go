package examples_test

import (
	"os/exec"
	"testing"

	"github.com/MakeNowJust/heredoc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestExamples(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Examples Suite")
}

var pathToLY string

var _ = BeforeSuite(func() {
	var err error
	pathToLY, err = gexec.Build("github.com/jtarchie/ly/ly")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = Describe("Examples", func() {
	It("has the correct output for variables.lua", func() {
		command := exec.Command(pathToLY, "--config=variable.lua")
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))
		value := heredoc.Doc(`
      false_bool: false
      float: 123.123
      integer: 123
      list:
      - item1
      - item2
      list_with_a_map:
      - key1: value1
        key2: value2
      - item2
      nested_map:
        key1: value1
        key2: value2
      nullz: null
      string: value
      string_quoted: '#value'
      true_bool: true
		`)
		Expect(session.Out.Contents()).To(MatchYAML(value))
	})

	It("has the correct JSON output for variables.lua", func() {
		command := exec.Command(pathToLY, "--config", "variable.lua", "--format", "json")
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))
		value := `
		{
			"false_bool": false,
			"float": 123.123,
			"integer": 123,
			"list": [
				"item1",
				"item2"
			],
			"list_with_a_map": [
				{
					"key1": "value1",
					"key2": "value2"
				},
				"item2"
			],
			"nested_map": {
				"key1": "value1",
				"key2": "value2"
			},
			"nullz": null,
			"string": "value",
			"string_quoted": "#value",
			"true_bool": true
		}
		`
		Expect(session.Out.Contents()).To(MatchJSON(value))
	})
})
