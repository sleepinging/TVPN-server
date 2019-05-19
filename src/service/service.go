package service

import "github.com/bitly/go-simplejson"

//GenRes gen res
func GenRes(res int, msg string) (str string) {
	js := simplejson.New()
	js.Set("res", res)
	js.Set("msg", msg)
	b, _ := js.MarshalJSON()
	str = string(b)
	return
}
