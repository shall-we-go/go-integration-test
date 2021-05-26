package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("test-reports/junit/main.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Main Suite", []Reporter{junitReporter})
}
