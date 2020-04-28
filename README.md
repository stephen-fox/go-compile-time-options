# Compile time options in Go applications
This repository demonstrates a simple pattern for implementing compile time
options in Go applications. For further information, please refer to
[my blog post](https://www.0x2d13.net/posts/go-compile-time-options/) on
the subject.

## Running
The following `go` commands can be used to run the example:

```
go run cmd/myapp/main.go
go run -tags debug cmd/myapp/main.go
```
