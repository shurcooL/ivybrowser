ivybrowser
==========

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/ivybrowser.svg)](https://pkg.go.dev/github.com/shurcooL/ivybrowser)

ivy in the browser.

Inspired by the iOS and Android ports of Rob Pike's ivy,
I ported it to run in web browsers using GopherJS compiler.

Installation
------------

```sh
go install github.com/shurcooL/ivybrowser@latest
```

To run ivy in the browser, you'll need [GopherJS](https://github.com/gopherjs/gopherjs#installation-and-usage) compiler.

The quickest way is to run:

```sh
gopherjs serve
```

And then visit <http://localhost:8080/github.com/shurcooL/ivybrowser> in your browser. The package will be compiled on the fly and served over HTTP.

Alternatively, you can `cd` into this directory and run:

```sh
gopherjs build
```

That will build `ivybrowser.js` file. You can now open `index.html` in a browser.

License
-------

-	[BSD-style License](LICENSE)
