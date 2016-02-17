package iris

import (
	"testing"
	"strings"
)

type testApiUsersHandler struct {
	Handler `get:"/api/users/:userId(int)" template:"user.html"`
}

func (t *testApiUsersHandler) Handle(ctx *Context, renderer *Renderer) {}

type testStructedRoute struct {
	handler Handler
	expectedMethods []string
	expectedPath string
	expectedTemplateFilename string
}


var structedTests = [...]testStructedRoute{{
	handler: new(testApiUsersHandler),
	expectedMethods: []string{"GET"},
	expectedPath : "/api/users/:userId(int)",
	expectedTemplateFilename : "/user.html",
}}

func TestRouterRegisterHandler(t *testing.T) {
	iris := New()
	for _,sr := range structedTests{
		route,err := iris.Router.RegisterHandler(sr.handler)
		
		if err != nil {
			t.Fatal("Error on RegisterHandler: "+ err.Error())
		}else {
			if !slicesAreEqual(sr.expectedMethods,route.methods){
				t.Fatal("Error on compare Methods: "+strings.Join(sr.expectedMethods,",")+" != "+strings.Join(route.methods,","))
			}
			
			if sr.expectedPath != route.path {
				t.Fatal("Error on compare Path: "+sr.expectedPath+" != "+route.path)
			}
			
			if templatesDirectory+sr.expectedTemplateFilename != route.templates.filesTemp[0] {
				t.Fatal("Error on compare Template filename: "+templatesDirectory+sr.expectedTemplateFilename+" != "+route.templates.filesTemp[0])
			}
			
		}
	}
	
}

func slicesAreEqual(s1, s2 []string) bool {

    if s1 == nil && s2 == nil { 
        return true; 
    }

    if s1 == nil || s2 == nil { 
        return false; 
    }

    if len(s1) != len(s2) {
        return false
    }

    for i := range s1 {
        if s1[i] != s2[i] {
            return false
        }
    }

    return true
}
