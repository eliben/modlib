Demonstrates the structure of a Go module on GitHub, with multiple packages and
command-line tools.

To use the packages in this module, add an import in Go code and call it:

```
import "github.com/eliben/modlib"
import "github.com/eliben/modlib/clientlib"

// ... later
s := modlib.Config()
clientlib.Hello()
```

To use the `server` binary, do:

```
$ go get github.com/eliben/modlib/cmd/server
$ server
```

All packages in this module are importable by other modules, except for packages
located in the `internal` directory. These can only be used from within the
module itself, but cannot be imported from the outside.
