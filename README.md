# REST Helper in Go

<!-- > 2020-11-23T04:39:24+08:00 -->

A helper library
for treed routers and corresponding routes of RESTful APIs.


## I. Introduction

> [ [*Pending* Documentations](#) ]

### 1. Get Started

Choose one implementation of the libraries in **sec: 3. Related Libraries** below and get to use.

### 2. Basic Concepts
<!-- ## II. Basic Concepts -->

- [x] Treed Routers
- [x] Grouped Routes
- [ ] Request Helper( and Customization)
- [ ] Logical Flow Controller
    - Use common control-flows for your business logics to avoid logical duplicates.
- [x] Response Helper( and Customization)
    - A unified place to handle errors with context of router background.
- [x] Custom API Handler (Response Helper Wrapper)
	- You may use different other than `func(ctx *gin.Context)` or `func(ctx iris.Context)`.

### 3. Related Libraries

The implementations of this library for some HTTP web frameworks in golang.

- [agin](http://github.com/goslib/agin)
    - [A](http://github.com/goslib/agin) RESTful API helper for [Gin](https://github.com/gin-gonic/gin).



## II. Features

### 1. Handlers Chaining or Path Reuse

Literally, you don't *have to* duplicate the path of a *REST resource*
for its `GET` handler, `POST` handler, and other handlers.

```go
	// ...
	router := gin.Default()
	router.GET("/resource-path", getting)
	router.POST("/resource-path", posting)
	// ...
	router.DELETE("/resource-path", deleting)
	// ...
```

Instead of the demo above, we got this:

```go
	rest.RoutePath("/resource-path").
		Get("", "", getHandler).
		Post().
		End()
```


### 2. Unified Error Handling with Context of Route Environment

You got a unified place to handle errors, and additionally,
the route background is right in the way to use with.

```go
func (m *ResponseHelper) endInternalServerError(err error, label string, values []interface{}) *ResponseBundle {
	m.Context.Status(http.StatusInternalServerError)
	m.errorout("[AGIN/500] ["+m.Route.GetTag()+"] [INTERNAL_SERVER_ERROR]", err, label, values)
	return nil
}

func (m *ResponseHelper) EndInternalServerError(err error, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, "", values)
}

func (m *ResponseHelper) InternalServerError(err error, label string, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, label, values)
}

// ---------- Other Alias ---------- //

func (m *ResponseHelper) InternalDatabaseError(err error, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, "Inner Database Error", values)
}

func (m *ResponseHelper) InternalServicesError(err error, values ...interface{}) *ResponseBundle {
	return m.endInternalServerError(err, "Inner Services Error", values)
}
```

### 3. Tailoring a Custom Route Handler of Your Own

You are empowered to use a different format of route handler,
other than the default `func(ctx *gin.Context)` in [Gin](https://github.com/gin-gonic/gin)
or`func(ctx iris.Context)` in [iris](https://github.com/kataras/iris].

> A wrapper takes your custom route handler and receives the chained route as background context,
and finally take cares of the right request. 

You may creat your own wrapper for the response helper, like the demo shown below, to use a different format of handler.

```go
func NewGinResponseHelperWrapper(
	handler func(ctx *Context, res *ResponseHelper) *ResponseBundle,
) func(env *Route) func(ctx *Context) {
	// @see: [utils.go@github/gin#nameOfFunction()](https://github.com/gin-gonic/gin/blob/7742ff50e0a05d079a0c468ccfbf7c6ecfe2414b/utils.go#L123)
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return func(env *Route) func(ctx *Context) {
		env.HandlerName = name
		return func(ctx *Context) {
			res := NewResponseHelper(env, ctx)
			handler(ctx, res)
		}
	}
}
```
