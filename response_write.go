package FakeWine

import "net/http"

/**
*
* @author Liu Weiyi
* @date 2019-03-14 23:29
 */

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

const noWritten = -1

func (r *responseWriter) reset(w http.ResponseWriter) {
	r.ResponseWriter = w
	r.size = -1
	r.status = http.StatusOK
}

func (r *responseWriter) writeHeader(code int) {
	r.status = code
}

func (r *responseWriter) writeHeaderNow() {
	if !r.Written() {
		r.ResponseWriter.WriteHeader(r.status)
	}
}

func (w *responseWriter) Written() bool {
	return w.size != noWritten
}
