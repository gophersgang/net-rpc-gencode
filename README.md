# net-rpc-gencode
This library provides the same functions as net/rpc/jsonrpc but for communicating with [gencode](https://github.com/andyleap/gencode) instead. The library is modeled directly after the Go standard library so it should be easy to use and obvious.

See the [GoDoc]() for API documentation.

> according to my test: [Golang Serializer Benchmark Comparison](https://github.com/smallnest/gosercomp), gencode is very faster than other serializers.