package render

import (
	"encoding/json"
	"net/http"
)

/**
*
* @author Liu Weiyi
* @date 2019-03-15 15:24
 */

type Render interface {
	Render(http.ResponseWriter) error
	WriteContentType(w http.ResponseWriter)
}

type JSON struct {
	Data interface{}
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func (r JSON) Render(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

func (r JSON) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	header["Content-Type"] = jsonContentType
}

func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	JSON{}.WriteContentType(w)
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonObj)
	return err
}
