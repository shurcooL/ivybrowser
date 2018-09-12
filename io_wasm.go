// +build js,wasm

package main

import (
	"io"
	"syscall/js"
)

var document = js.Global().Get("document")

func init() {
	// The default os.Stdout, os.Stderr are printed to browser's console, which isn't a friendly interface.
	// Create an implementation of stdin, stdout, stderr that use a <input> and <pre> html elements.
	stdin = NewReader(document.Call("getElementById", "input"))
	stdout = NewWriter(document.Call("getElementById", "output"))
	stderr = NewWriter(document.Call("getElementById", "output"))

	// Send a copy of stdin to stdout (like in most terminals).
	stdin = io.TeeReader(stdin, stdout)

	// When console is clicked, focus the input element.
	// TODO: Make it possible/friendlier to copy the text from stdout...
	document.Call("getElementById", "console").Call("addEventListener", "click", js.NewEventCallback(js.PreventDefault, func(js.Value) {
		document.Call("getElementById", "input").Call("focus")
	}))
}

// NewReader takes an <input> element and makes an io.Reader out of it.
func NewReader(input js.Value) io.Reader {
	r := &reader{
		in: make(chan []byte, 8),
	}
	input.Call("addEventListener", "keydown", js.NewEventCallback(0, func(ke js.Value) {
		if ke.Get("keyCode").Int() == '\r' {
			r.in <- []byte(input.Get("value").String() + "\n")
			input.Set("value", "")
			// TODO: Can't call preventDefault here; what to do? See issue https://golang.org/issue/26045.
		}
	}))
	return r
}

type reader struct {
	pending []byte
	in      chan []byte // This channel is never closed here, so no need to detect it and return io.EOF.
}

func (r *reader) Read(p []byte) (n int, err error) {
	if len(r.pending) == 0 {
		r.pending = <-r.in
	}
	n = copy(p, r.pending)
	r.pending = r.pending[n:]
	return n, nil
}

// NewWriter takes a <pre> element and makes an io.Writer out of it.
func NewWriter(pre js.Value) io.Writer {
	return &writer{pre: pre}
}

type writer struct {
	pre js.Value
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.pre.Set("textContent", w.pre.Get("textContent").String()+string(p))
	return len(p), nil
}
