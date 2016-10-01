**Run instructions are as follows:**
(This assumes that you have installed golang)

- `go get ./...`
- `go build`
- `go test -v ./...`

If you don't have Golang, you are welcome to use the Docker run guide:
- `docker build -t hash .`
- `docker run hash`

This implementation of a hash table uses bucketing, as opposed to probing, this is generally a better idea for fixed size hash tables.

`hashtable_test.go` contains the tests for `hashtable.go`