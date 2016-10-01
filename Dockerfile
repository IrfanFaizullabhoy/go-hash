FROM golang:1.7.1-alpine

COPY . $GOPATH/src/github.com/IrfanFaizullabhoy/go-hash

WORKDIR $GOPATH/src/github.com/IrfanFaizullabhoy/go-hash

ENTRYPOINT ["go", "test", "-v"]