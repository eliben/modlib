To use the packages in this module, add an import in Go code and call it:

```
import "github.com/eliben/modlib"
import "github.com/eliben/modlib/otherlib"

// ... later
s := modlib.Foo1()
v := otherlib.Bob1()
```

To use the `modcmd` binary, do:

```
$ go get github.com/eliben/modlib/cmd/modcmd
$ modcmd
```
