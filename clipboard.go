// Package clipboard uses Qt's facilities to store or load text in/out of system's clipboard
package clipboard

// #cgo LDFLAGS: -lstdc++
// #cgo pkg-config: Qt5Core Qt5Widgets
//
// #include "clipboard.h"
import "C"
import (
	"errors"
	"fmt"

	"gopkg.in/qml.v1"
)

type Clipboard struct {
	engine *qml.Engine
	helper qml.Common
}

func New(engine *qml.Engine) *Clipboard {
	return &Clipboard{
		engine: engine,
		helper: *qml.CommonOf(C.clipboardHelper(), engine),
	}
}

// ReadAll reads string from QApplication::clipboard
func (c *Clipboard) ReadAll() (text string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("panic: %v", x))
		}
	}()
	text = c.helper.Call("getText").(string)
	return
}

// WriteAll writes string to QApplication::clipboard
func (c *Clipboard) WriteAll(text string) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("panic: %v", x))
		}
	}()
	c.helper.Call("setText", text)
	return
}
