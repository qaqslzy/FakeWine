package fakewine

import (
	"net/http"
	"sync"
)

/**
*
* @author Liu Weiyi
* @date 2019-03-14 15:58
 */

const defaultMultipartMemory = 32 << 20 // 32 MB

type Handle func(ctx *Context)

type HandlersChain []Handle

type Fengine struct {
	GRouter
	pool sync.Pool
}

func (engine *Fengine) addRoute(path string, handle ...Handle) {

}

func (engine *Fengine) ServeHTTP(http.ResponseWriter, *http.Request) {
	c := engine.pool.Get().(*Context)
	// TODO 重置 Context

	// TODO 开始对请求进行响应

	engine.pool.Put(c)
}

func (engine *Fengine) Run(address string) error {
	// TODO 对地址进行修改
	return http.ListenAndServe(address, engine)
}

func New() *Fengine {
	return &Fengine{}
}
