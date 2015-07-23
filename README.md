## Clipboard for Go ([go-qml](http://godoc.org/gopkg.in/qml.v1) applications)

Uses Qt's facilities to store or load text in/out of the clipboard. This branch utilizes `qml.Common` to make calls to QClipboard wrapper. This way is a little slower (than [master](https://github.com/xlab/clipboard/tree/master)), but code is much cleaner and tends to be safer.

### Installation:

See the [requirements](https://github.com/go-qml/qml/tree/master) of go-qml. If they're met and `PKG_CONFIG_PATH` is set:

    $ go get gopkg.in/xlab/clipboard.v2

### Platforms:

This clipboard package can be used on all platforms that are supported by go-qml. Requires installed `gopkg.in/qml.v1` package (N.B. `clipboard.v1` depends on `qml.v0` respectively).

### Doc: 

* http://godoc.org/gopkg.in/xlab/clipboard.v2

### Additional info:

The package [atotto/clipboard](https://github.com/atotto/clipboard) might be incompatible with go-qml machinery, see progress on [Issue #26](https://github.com/go-qml/qml/issues/26).
This package utilizes CGO and `qml.Common` in order to use `QApplication::clipboard`. Also the benchmarks are better:

```
BenchmarkReadAll      200000          9770 ns/op
BenchmarkWriteAll      10000        185354 ns/op
ok      github.com/xlab/clipboard.v2   5.079s
```

```
BenchmarkReadAll	     100	  20067708 ns/op
BenchmarkWriteAll	     100	  19592142 ns/op
ok  	github.com/atotto/clipboard	4.119s
```

License is [MIT](http://xlab.mit-license.org).
