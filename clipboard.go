// Package clipboard uses Qt's facilities to store or load text in/out of system's clipboard
package clipboard

// #cgo LDFLAGS: -lstdc++
// #cgo pkg-config: Qt5Core Qt5Widgets
//
// #include <stdlib.h>
// #include "capi.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
	"gopkg.in/qml.v0"
)

type Clipboard struct {
	engine *qml.Engine
	clip   *qml.Common
}

func New(engine *qml.Engine) *Clipboard {
	return &Clipboard{
		engine: engine,
		clip:   qml.CommonOf(C.getClipboard(), engine),
	}
}

// ReadAll reads string from QApplication::clipboard
func (c *Clipboard) ReadAll() (text string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("panic: %v", x))
		}
	}()
	cstring := C.getText(unsafe.Pointer(c.clip.Addr()))
	text = C.GoString(cstring)
	return
}

// WriteAll writes string to QApplication::clipboard
func (c *Clipboard) WriteAll(text string) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("panic: %v", x))
		}
	}()
	cstring := C.CString(text)
	defer C.free(unsafe.Pointer(cstring))
	C.setText(unsafe.Pointer(c.clip.Addr()), cstring)
	return
}
