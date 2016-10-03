// +build js

package main

import (
	"io"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func init() {
	// The default os.Stdout, os.Stderr are printed to browser's console, which isn't a friendly interface.
	// Create an implementation of stdin, stdout, stderr that use a <input> and <pre> html elements.
	stdin = NewReader(document.GetElementByID("input").(*dom.HTMLInputElement))
	stdout = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement))
	stderr = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement))

	// Send a copy of stdin to stdout (like in most terminals).
	stdin = io.TeeReader(stdin, stdout)

	// When console is clicked, focus the input element.
	// TODO: Make it possible/friendlier to copy the text from stdout...
	document.GetElementByID("console").AddEventListener("click", false, func(event dom.Event) {
		document.GetElementByID("input").(dom.HTMLElement).Focus()
		event.PreventDefault()
	})
}

// NewReader takes an <input> element and makes an io.Reader out of it.
func NewReader(input *dom.HTMLInputElement) io.Reader {
	r := &reader{
		in: make(chan []byte, 8),
	}
	input.AddEventListener("keydown", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		if ke.KeyCode == '\r' {
			r.in <- []byte(input.Value + "\n")
			input.Value = ""
			ke.PreventDefault()
		}
	})
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
func NewWriter(pre *dom.HTMLPreElement) io.Writer {
	return &writer{pre: pre}
}

type writer struct {
	pre *dom.HTMLPreElement
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.pre.SetTextContent(w.pre.TextContent() + string(p))
	return len(p), nil
}
