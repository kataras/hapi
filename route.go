package iris

import (
	"net/http"
	"strings"
)

// Route contains its middleware, handler, pattern , it's path string, http methods and a template cache
// Used to determinate which handler on which path must call
// Used on router.go
type Route struct {

	//Middleware
	MiddlewareSupporter
	//mu            sync.RWMutex
	methods    []string
	pathPrefix string // this is to make a little faster the match, before regexp Match runs, it's the path before the first path parameter :
	//the pathPrefix is with the last / if parameters exists.
	parts       []string //stores the string path AFTER the pathPrefix, without the pathPrefix. no need to that but no problem also.
	fullpath    string   // need only on parameters.Params(...)
	handler     Handler
	isReady     bool
	templates   *TemplateCache //this is passed to the Renderer
	httpErrors  *HTTPErrors    //the only need of this is to pass into the Context, in order to  developer get the ability to perfom emit errors (eg NotFound) directly from context
	hasWildcard bool
}

// newRoute creates, from a path string, handler and optional http methods and returns a new route pointer
func newRoute(registedPath string, handler Handler) *Route {
	r := &Route{handler: handler}
	hasPathParameters := false
	firstPathParamIndex := strings.IndexByte(registedPath, ParameterStartByte)
	if firstPathParamIndex != -1 {
		r.pathPrefix = registedPath[:firstPathParamIndex] ///api/users  to /api/users/
		hasPathParameters = true
	} else {
		//check for only for* , here no path parameters registed.
		firstPathParamIndex = strings.IndexByte(registedPath, MatchEverythingByte)

		if firstPathParamIndex != -1 {
			if firstPathParamIndex <= 1 { // set to '*' to pathPrefix if no prefix exists except the slash / if any [Get("/*",..) or Get("*",...]
				//has no prefix just *
				r.pathPrefix = MatchEverything
				r.hasWildcard = true
			} else { //if firstPathParamIndex == len(registedPath)-1 { // it's the last
				//has some prefix and sufix of *
				r.pathPrefix = registedPath[:firstPathParamIndex] //+1
				r.hasWildcard = true
				hasPathParameters = true
			}

		} else {
			//else no path parameter or match everything symbol so use the whole path as prefix it will be faster at the check for static routes too!
			r.pathPrefix = registedPath
		}

	}

	if hasPathParameters || (hasPathParameters && r.hasWildcard) {
		r.parts = strings.Split(registedPath[len(r.pathPrefix):], "/")
		r.fullpath = registedPath //we need this only to take Params so set it if has path parameters.
	}

	return r
}

// containsMethod determinates if this route contains a http method
func (r *Route) containsMethod(method string) bool {
	for _, m := range r.methods {
		if m == method {
			return true
		}
	}

	return false
}

// Methods adds methods to its registed http methods
func (r *Route) Methods(methods ...string) *Route {
	if r.methods == nil {
		r.methods = make([]string, 0)
	}
	r.methods = append(r.methods, methods...)
	return r
}

// Method SETS a method to its registed http methods, overrides the previous methods registed (if any)
func (r *Route) Method(method string) *Route {
	r.methods = []string{HTTPMethods.GET}
	return r
}

// match determinates if this route match with the request, returns bool as first value and PathParameters as second value, if any
func (r *Route) match(urlPath string) bool {
	if r.pathPrefix == MatchEverything {
		return true
	}
	if r.pathPrefix == urlPath {
		//it's route without path parameters or * symbol, and if the request url has prefix of it  and it's the same as the whole preffix which is the path itself returns true without checking for regexp pattern
		//so it's just a path without named parameters
		return true

	} else if r.parts != nil {
		partsLen := len(r.parts)
		// count the slashes after the prefix, we start from one because we will have at least one slash.
		reqPartsLen := 1
		s := urlPath[len(r.pathPrefix):]
		for i := 0; i < len(s); i++ {
			if s[i] == SlashByte {
				reqPartsLen++
			}
		}

		//if request has more parts than the route, but the route has finish with * symbol then it's wildcard
		//maybe it's a little confusing , why dont u use just reqPartsLen < partsLen return false ? it doesnt work this way :)
		if reqPartsLen >= partsLen && r.hasWildcard { // r.parts[partsLen-1][0] == MatchEverythingByte { // >= and no != because we check for matchEveryting *
			return true
		} else if reqPartsLen != partsLen {
			return false
		}

		return true

	} else {
		return false
	}

}

// Template creates (if not exists) and returns the template cache for this route
func (r *Route) Template() *TemplateCache {
	if r.templates == nil {
		r.templates = NewTemplateCache()
	}
	return r.templates
}

// prepare prepares the route's handler , places it to the last middleware , handler acts like a middleware too.
// Runs once before the first ServeHTTP
func (r *Route) prepare() {
	//r.mu.Lock()
	//look why on router ->HandleFunc defer r.mu.Unlock()
	//but wait... do we need locking here?
	if r.handler != nil {
		convertedMiddleware := MiddlewareHandlerFunc(func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
			r.handler.run(r, res, req)
			next(res, req)
		})

		r.Use(convertedMiddleware)
	}

	//here if no methods are defined at all, then use GET by default.
	if r.methods == nil {
		r.methods = []string{HTTPMethods.GET}
	}

	r.isReady = true
}

// ServeHTTP serves this route and it's middleware
func (r *Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if r.isReady == false && r.handler != nil {
		r.prepare()
	}
	r.middleware.ServeHTTP(res, req)
}
