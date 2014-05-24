## Clipboard for Go ([go-qml](http://godoc.org/gopkg.in/qml.v0) applications)

Uses Qt's facilities to store or load text in/out of the clipboard. This branch utilizes `qml.Common` to make calls to QClipboard wrapper. This way is a little slower (than [master](https://github.com/xlab/clipboard/tree/master)), but code is much cleaner and tends to be safer.

### Installation:

    $ go get gopkg.in/xlab/clipboard.v1

### Platforms:

All that supported by go-qml. Requires installed `gopkg.in/qml.v0` package.

### Doc: 

* http://godoc.org/gopkg.in/xlab/clipboard.v1

### Additional info:

The package [atotto/clipboard](https://github.com/atotto/clipboard) is currently incompatible with go-qml machinery, see [Issue #26](https://github.com/go-qml/qml/issues/26).
This package utilizes CGO and `qml.Common` in order to use `QApplication::clipboard`. Also the benchmarks are better:

```
BenchmarkReadAll	  100000	     16090 ns/op
BenchmarkWriteAll	   10000	    271159 ns/op
ok  	gopkg.in/xlab/clipboard.v1	4.723s
```

```
BenchmarkReadAll	     100	  20067708 ns/op
BenchmarkWriteAll	     100	  19592142 ns/op
ok  	github.com/atotto/clipboard	4.119s
```