ivybrowser
==========

[![Build Status](https://travis-ci.org/shurcooL/ivybrowser.svg?branch=master)](https://travis-ci.org/shurcooL/ivybrowser) [![GoDoc](https://godoc.org/github.com/shurcooL/ivybrowser?status.svg)](https://godoc.org/github.com/shurcooL/ivybrowser)

ivy in the browser.

Inspired by the iOS and Android ports of Rob Pike's ivy, I ported it to the web using GopherJS compiler.

Installation
------------

```bash
go get -u github.com/shurcooL/ivybrowser
GOARCH=js go get -u -d github.com/shurcooL/ivybrowser
```

To run ivy in the browser, you'll need [GopherJS compiler](https://github.com/gopherjs/gopherjs#installation-and-usage).

The quickest way is to run:

```bash
gopherjs serve
```

And then visit http://localhost:8080/github.com/shurcooL/ivybrowser in your browser. The package will be compiled on the fly and served over HTTP.

Alternatively, you can `cd` into this directory and run:

```bash
gopherjs build
```

That will build `ivybrowser.js` file. You can now open `index.html` in a browser.

License
-------

-	[BSD-style License](LICENSE)
