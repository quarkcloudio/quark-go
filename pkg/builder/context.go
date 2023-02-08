package builder

import (
	"net/http"
)

// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	Request  *http.Request
	Writer   http.ResponseWriter
	fullPath string
}

func NewContext(Request *http.Request, Writer http.ResponseWriter) *Context {
	return &Context{
		Request: Request,
		Writer:  Writer,
	}
}

// FullPath returns a matched route full path. For not found routes
// returns an empty string.
//     router.GET("/user/:id", func(c *gin.Context) {
//         c.FullPath() == "/user/:id" // true
//     })
func (c *Context) FullPath() string {
	return c.fullPath
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (c *Context) Method() string {
	return c.Request.Method
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (c *Context) Write(p []byte) {

	c.Writer.Write(p)
}
