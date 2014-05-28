package Jibber_Jabber_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestJibberJabber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JibberJabber Suite")
}
