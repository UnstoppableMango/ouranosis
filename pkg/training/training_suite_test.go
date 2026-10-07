package training_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTraining(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Training Suite")
}
