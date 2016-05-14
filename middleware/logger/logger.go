package logger

import (
	"io"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/logger"
)

// Options are the options of the logger middlweare
// contains 5 bools
// Latency, Status, IP, Method, Path
// if set to true then these will print
type Options struct {
	Latency bool
	Status  bool
	IP      bool
	Method  bool
	Path    bool
}

// DefaultOptions returns an options which all properties are true
func DefaultOptions() Options {
	return Options{true, true, true, true, true}
}

type loggerMiddleware struct {
	*logger.Logger
	options Options
}

// a poor  and ugly implementation of a logger but no need to worry about this at the moment
func (l *loggerMiddleware) Serve(ctx *iris.Context) {
	//all except latency to string
	var date, status, ip, method, path string
	var latency time.Duration
	var startTime, endTime time.Time
	path = ctx.PathString()
	method = ctx.MethodString()

	if l.options.Latency {
		startTime = time.Now()
	}

	ctx.Next()
	if l.options.Latency {
		//no time.Since in order to format it well after
		endTime = time.Now()
		date = endTime.Format("01/02 - 15:04:05")
		latency = endTime.Sub(startTime)
	}

	if l.options.Status {
		status = strconv.Itoa(ctx.Response.StatusCode())
	}

	if l.options.IP {
		ip = ctx.RemoteAddr()
	}

	if !l.options.Method {
		method = ""
	}

	if !l.options.Path {
		path = ""
	}

	//finally print the logs
	if l.options.Latency {
		l.Printf("%s %v %4v %s %s %s", date, status, latency, ip, method, path)
	} else {
		l.Printf("%s %v %s %s %s", date, status, ip, method, path)
	}

}

func newLoggerMiddleware(writer io.Writer, prefix string, flag int, options ...Options) *loggerMiddleware {
	l := &loggerMiddleware{Logger: logger.Custom(writer, prefix, flag)}

	if len(options) > 0 {
		l.options = options[0]
	} else {
		l.options = DefaultOptions()
	}

	return l
}

//all bellow are just for flexibility

// DefaultHandler returns the logger middleware with the default settings
func DefaultHandler(options ...Options) iris.Handler {
	return newLoggerMiddleware(logger.Output, "", 0)
}

// Default returns the logger middleware as HandlerFunc with the default settings
func Default(options ...Options) iris.HandlerFunc {
	return DefaultHandler(options...).Serve
}

// CustomHandler returns the logger middleware with customized settings
// accepts 3 parameters
// first parameter is the writer (io.Writer)
// second parameter is the prefix of which the message will follow up
// third parameter is the logger.Options
func CustomHandler(writer io.Writer, prefix string, flag int, options ...Options) iris.Handler {
	return newLoggerMiddleware(writer, prefix, flag, options...)
}

// Custom returns the logger middleware as HandlerFunc with customized settings
// accepts 3 parameters
// first parameter is the writer (io.Writer)
// second parameter is the prefix of which the message will follow up
// third parameter is the logger.Options
func Custom(writer io.Writer, prefix string, flag int, options ...Options) iris.HandlerFunc {
	return CustomHandler(writer, prefix, flag, options...).Serve
}
