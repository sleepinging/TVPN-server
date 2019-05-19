package service

import "github.com/bitly/go-simplejson"
import "net/http"
import "io/ioutil"

//GenRes gen res
func GenRes(res int, msg string) (str string) {
	js := simplejson.New()
	js.Set("res", res)
	js.Set("msg", msg)
	b, _ := js.MarshalJSON()
	str = string(b)
	return
}

//Parse2Json 解析为json
func Parse2Json(r *http.Request) (js *simplejson.Json, err error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	js, err = simplejson.NewJson(data)
	return
}
