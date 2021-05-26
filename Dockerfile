FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go mod tidy
