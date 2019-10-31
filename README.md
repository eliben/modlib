Demonstrates the structure of a Go module on GitHub, with multiple packages and
command-line tools.

For more details see https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/

To use the packages in this module, add an import in Go code and call it:

```
import "github.com/eliben/modlib"
import "github.com/eliben/modlib/clientlib"

// ... later
s := modlib.Config()
h := clientlib.Hello()
```

To use the `modlib-server` binary, do:

```
$ go get github.com/eliben/modlib/cmd/modlib-server
$ modlib-server
```

All packages in this module are importable by other modules, except for packages
located in the `internal` directory. These can be used from within the module
itself, but cannot be imported from the outside.

To run all tests in this module:

```
go test ./...
```
