// Copyright (c) 2016, Gerasimos Maropoulos
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//	  this list of conditions and the following disclaimer
//    in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse
//    or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER AND CONTRIBUTOR, GERASIMOS MAROPOULOS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package iris

import (
	"html/template"

	"github.com/kataras/iris/logger"
	"github.com/kataras/iris/server"
)

// DefaultIris in order to use iris.Get(...,...) we need a default Iris on the package level
var DefaultIris *Iris

// init the only one.
func init() {
	DefaultIris = New()
}

// DefaultConfig returns the default iris.Config for the Iris
func DefaultConfig() *IrisConfig {
	return &IrisConfig{
		PathCorrection:     true,
		MaxRequestBodySize: -1,
		Log:                true,
		Profile:            false,
		ProfilePath:        DefaultProfilePath,
		Render: &RenderConfig{
			Directory:                 "templates",
			Asset:                     nil,
			AssetNames:                nil,
			Layout:                    "",
			Extensions:                []string{".html"},
			Funcs:                     []template.FuncMap{},
			Delims:                    Delims{"{{", "}}"},
			Charset:                   DefaultCharset,
			IndentJSON:                false,
			IndentXML:                 false,
			PrefixJSON:                []byte(""),
			PrefixXML:                 []byte(""),
			HTMLContentType:           "text/html",
			IsDevelopment:             false,
			UnEscapeHTML:              false,
			StreamingJSON:             false,
			RequirePartials:           false,
			DisableHTTPErrorRendering: false,
		}}
}

// Listen starts the standalone http server
// which listens to the addr parameter which as the form of
// host:port or just port
//
// It panics on error if you need a func to return an error use the ListenWithErr
// ex: iris.Listen(":8080")
func Listen(addr string) {
	DefaultIris.Listen(addr)
}

// ListenWithErr starts the standalone http server
// which listens to the addr parameter which as the form of
// host:port or just port
//
// It returns an error you are responsible how to handle this
// if you need a func to panic on error use the Listen
// ex: log.Fatal(iris.ListenWithErr(":8080"))
func ListenWithErr(addr string) error {
	return DefaultIris.ListenWithErr(addr)
}

// ListenTLS Starts a https server with certificates,
// if you use this method the requests of the form of 'http://' will fail
// only https:// connections are allowed
// which listens to the addr parameter which as the form of
// host:port or just port
//
// It panics on error if you need a func to return an error use the ListenTLSWithErr
// ex: iris.ListenTLS(":8080","yourfile.cert","yourfile.key")
func ListenTLS(addr string, certFile, keyFile string) {
	DefaultIris.ListenTLS(addr, certFile, keyFile)
}

// ListenTLSWithErr Starts a https server with certificates,
// if you use this method the requests of the form of 'http://' will fail
// only https:// connections are allowed
// which listens to the addr parameter which as the form of
// host:port or just port
//
// It returns an error you are responsible how to handle this
// if you need a func to panic on error use the ListenTLS
// ex: log.Fatal(iris.ListenTLSWithErr(":8080","yourfile.cert","yourfile.key"))
func ListenTLSWithErr(addr string, certFile, keyFile string) error {
	return DefaultIris.ListenTLSWithErr(addr, certFile, keyFile)
}

// Close is used to close the net.Listener of the standalone http server which has already running via .Listen
func Close() { DefaultIris.Close() }

// Router implementation

// Party is just a group joiner of routes which have the same prefix and share same middleware(s) also.
// Party can also be named as 'Join' or 'Node' or 'Group' , Party chosen because it has more fun
func Party(rootPath string) IParty {
	return DefaultIris.Party(rootPath)
}

// Handle registers a route to the server's router
func Handle(method string, registedPath string, handlers ...Handler) {
	DefaultIris.Handle(method, registedPath, handlers...)
}

// HandleFunc registers a route with a method, path string, and a handler
func HandleFunc(method string, path string, handlersFn ...HandlerFunc) {
	DefaultIris.HandleFunc(method, path, handlersFn...)
}

// HandleAnnotated registers a route handler using a Struct implements iris.Handler (as anonymous property)
// which it's metadata has the form of
// `method:"path"` and returns the route and an error if any occurs
// handler is passed by func(urstruct MyStruct) Serve(ctx *Context) {}
func HandleAnnotated(irisHandler Handler) error {
	return DefaultIris.HandleAnnotated(irisHandler)
}

// Use appends a middleware to the route or to the router if it's called from router
func Use(handlers ...Handler) {
	DefaultIris.Use(handlers...)
}

// UseFunc same as Use but it accepts/receives ...HandlerFunc instead of ...Handler
// form of acceptable: func(c *iris.Context){//first middleware}, func(c *iris.Context){//second middleware}
func UseFunc(handlersFn ...HandlerFunc) {
	DefaultIris.UseFunc(handlersFn...)
}

// Get registers a route for the Get http method
func Get(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Get(path, handlersFn...)
}

// Post registers a route for the Post http method
func Post(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Post(path, handlersFn...)
}

// Put registers a route for the Put http method
func Put(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Put(path, handlersFn...)
}

// Delete registers a route for the Delete http method
func Delete(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Delete(path, handlersFn...)
}

// Connect registers a route for the Connect http method
func Connect(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Connect(path, handlersFn...)
}

// Head registers a route for the Head http method
func Head(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Head(path, handlersFn...)
}

// Options registers a route for the Options http method
func Options(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Options(path, handlersFn...)
}

// Patch registers a route for the Patch http method
func Patch(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Patch(path, handlersFn...)
}

// Trace registers a route for the Trace http methodd
func Trace(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Trace(path, handlersFn...)
}

// Any registers a route for ALL of the http methods (Get,Post,Put,Head,Patch,Options,Connect,Delete)
func Any(path string, handlersFn ...HandlerFunc) {
	DefaultIris.Any(path, handlersFn...)
}

// Static serves a directory
// accepts three parameters
// first parameter is the request url path (string)
// second parameter is the system directory (string)
// third parameter is the level (int) of stripSlashes
// * stripSlashes = 0, original path: "/foo/bar", result: "/foo/bar"
// * stripSlashes = 1, original path: "/foo/bar", result: "/bar"
// * stripSlashes = 2, original path: "/foo/bar", result: ""
func Static(requestPath string, systemPath string, stripSlashes int) {
	DefaultIris.Static(requestPath, systemPath, stripSlashes)
}

// StaticFS registers a route which serves a system directory
// it generates an index page to view the directory's files
func StaticFS(requestPath string, systemPath string, stripSlashes int) {
	DefaultIris.StaticFS(requestPath, systemPath, stripSlashes)
}

// OnError Registers a handler for a specific http error status
func OnError(httpStatus int, handler HandlerFunc) {
	DefaultIris.OnError(httpStatus, handler)
}

// EmitError executes the handler of the given error http status code
func EmitError(httpStatus int, ctx *Context) {
	DefaultIris.EmitError(httpStatus, ctx)
}

// OnNotFound sets the handler for http status 404,
// default is a response with text: 'Not Found' and status: 404
func OnNotFound(handlerFunc HandlerFunc) {
	DefaultIris.OnNotFound(handlerFunc)
}

// OnPanic sets the handler for http status 500,
// default is a response with text: The server encountered an unexpected condition which prevented it from fulfilling the request. and status: 500
func OnPanic(handlerFunc HandlerFunc) {
	DefaultIris.OnPanic(handlerFunc)
}

// ***********************
// Export DefaultIris's  exported properties
// ***********************

// Server returns the DefaultIris.Server
func Server() *server.Server {
	return DefaultIris.Server
}

// Plugins returns the plugin container,  DefaultIris.Plugins
func Plugins() *PluginContainer {
	return DefaultIris.Plugins
}

// Config returns the DefaultIris.Config
func Config() *IrisConfig {
	return DefaultIris.Config
}

// Logger returns the DefaultIris.Logger
func Logger() *logger.Logger {
	return DefaultIris.Logger
}

// SetMaxRequestBodySize Maximum request body size.
//
// The server rejects requests with bodies exceeding this limit.
//
// By default request body size is unlimited.
func SetMaxRequestBodySize(size int) {
	DefaultIris.SetMaxRequestBodySize(size)
}

// SetRenderConfig sets the Config.Render, can be setted before server's listen, not after.
func SetRenderConfig(renderCfg *RenderConfig) {
	DefaultIris.SetRenderConfig(renderCfg)
}
