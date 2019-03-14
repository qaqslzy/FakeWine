package fakewine

/**
*
* @author Liu Weiyi
* @date 2019-03-14 15:58
 */

const defaultMultipartMemory = 32 << 20 // 32 MB

type Handle func(ctx *Context)

type Fengine struct {
	GRouter
}
