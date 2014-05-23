## Clipboard for Go ([go-qml](http://godoc.org/gopkg.in/qml.v0) applications)

Uses Qt's facilities to store or load text in/out of the clipboard.

### Installation:

    $ go get github.com/xlab/clipboard

### Platforms:

All that supported by go-qml. Requires installed `gopkg.in/qml.v0` package.

### Doc: 

* http://godoc.org/github.com/xlab/clipboard

### Additional info:

The package [atotto/clipboard](https://github.com/atotto/clipboard) is currently incompatible with go-qml machinery, see [Issue #26](https://github.com/go-qml/qml/issues/26).
This package utilizes a CGO wrapper in order to use `QApplication::clipboard`. Also the benchmarks are better:

```
BenchmarkReadAll	   1000000	      2527 ns/op
BenchmarkWriteAll	   10000	    107819 ns/op
ok  	github.com/xlab/clipboard	3.909s
```

```
BenchmarkReadAll	     100	  20067708 ns/op
BenchmarkWriteAll	     100	  19592142 ns/op
ok  	github.com/atotto/clipboard	4.119s
```