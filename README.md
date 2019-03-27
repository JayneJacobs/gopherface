# gopherface.jaynejacobs.com
# gopherface
Go prevents cross-sight scripting attacks
GopherJS converts Go code into JavaScript code

GopherJS compatibility table:

https://github.com/gopherjs/gopherjs/blob/master/doc/packages.md

```sh
go get -u github.com/gopherjs/gopherjs
```

* Build Js source file
    
```sh
gopherjs build
```

* Build and minify JS source file:

```sh
gopherjs build -m
```
    
   Compiles Go packages and start Http server
    
```sh
    gopherjs serve
```

   * use int instead of int8/16/32/64

   * Use float64 not float32

*JavaScript:*

```js
element = document.getElementById("navConteiner");
```
GopherJS:

```go
element := js.Global.Get("document").Call("getElementById", "navContainer")
```

* DOM Binding:

```go
element := dom.GetWindow().Document().GetElementByID("navContainer")
```

* Alias calls to js.Global and Dom

```go
var JS = js.Global
var D = dom.GetWindow().Document()
```
### Changing CSS Style

* JavaScript

```js
element = document.GetElementByID("modalLayer");
element.style.display = "none";
```

* GopherJS:
```go
var JS = js.Global
element := js.Global
element := JS.Get("document").Call("getElementById", "modalLayer")
element.Get("style").Set(display", "none")
```

Move .js generated files to /static in project root

# Making XHR Calls
* XHR object
  
  ### XML HTTP Request Object

  Data retreieved from a URL without causing full refresh
  xhr binding
  to install
  
  ```sh
  go get -u honnef.co/go/js/xhr 
```
* build file by referencing the root directory
* run the app from the root of the project