package subreaper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSubreaper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Subreaper Suite")
}
