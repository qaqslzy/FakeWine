package fakewine

import "net/http"

/**
*
* @author Liu Weiyi
* @date 2019-03-14 23:29
 */

type repinseWriter struct {
	http.ResponseWriter
	size   int
	status int
}
