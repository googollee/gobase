package gobase

// CodeHash needs to be set by ldflag, as the hash of the code repo which is compiled at. It uses for tracing source code/version from the repo.
// Example: `go build -ldflags="-X 'gobase.CodeHash='`git rev-parse --short HEAD`" ./src`
var CodeHash string = "no_hash"
