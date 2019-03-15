package fakewine

import (
	"FakeWine/fakewine/render"
	"net/http"
)

/**
*
* @author Liu Weiyi
* @date 2019-03-14 16:08
 */

type Context struct {
	request *http.Request
	writer  responseWriter
	path    string
	handles HandlersChain
	params  Params
	index   int
	engine  *Fengine
}

func (ctx *Context) Render(code int, r render.Render) {
	ctx.Status(code)

	if err := r.Render(ctx.writer); err != nil {
		panic(err)
	}
}

func (ctx *Context) Status(code int) {
	ctx.writer.writeHeader(code)
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.Render(code, render.JSON{Data: obj})
}

func (ctx *Context) File(path string) {
	http.ServeFile(ctx.writer, ctx.request, path)
}

func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < len(ctx.handles) {
		ctx.handles[ctx.index](ctx)
		ctx.index++
	}
}

func (ctx *Context) reset() {
	ctx.handles = nil
	ctx.params = nil
	ctx.index = -1
}
