FROM golang:latest

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.2

COPY . /go/src/bits.chrsm.org/actions-test
WORKDIR /go/src/bits.chrsm.org/actions-test

RUN ls -lha
RUN go mod download && go mod vendor

CMD [ "golangci-lint", "run" ]
