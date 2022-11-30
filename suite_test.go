package loggregator_test

import (
	"io"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/grpclog"

	"testing"
)

func TestGoLoggregator(t *testing.T) {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(GinkgoWriter, io.Discard, io.Discard))
	log.SetOutput(GinkgoWriter)

	RegisterFailHandler(Fail)
	RunSpecs(t, "GoLoggregator Suite")
}
