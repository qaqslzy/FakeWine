package fakewine

import "net/http"

/**
*
* @author Liu Weiyi
* @date 2019-03-14 16:08
 */

type Context struct {
	request  *http.Request
	reponsew repinseWriter
	path     string
	handles  HandlersChain
}

func (ctx *Context) JSON(code int, obj interface{}) {

}
