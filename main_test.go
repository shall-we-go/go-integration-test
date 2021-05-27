// +build integration

package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var API_END_POINT = os.Getenv("API_END_POINT")

var _ = Describe("Integration Test", func() {

	readyCh := make(chan bool)
	go isReady(readyCh)
	<-readyCh

	When("GET /helloworld", func() {
		It("can get msg from Redis", func() {
			expect := "Hello World"
			res, err := http.Get(API_END_POINT + "/helloworld")
			Expect(err).To(BeNil())
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			Expect(err).To(BeNil())
			Expect(string(body)).To(Equal(expect))
		})
	})
})

func isReady(readyCh chan bool) {
	_, err := http.Get(API_END_POINT + "/healthcheck")
	if err == nil {
		readyCh <- true
		return
	}
	fmt.Println("waiting server to be ready...")
	time.Sleep(time.Second)
	isReady(readyCh)
}
