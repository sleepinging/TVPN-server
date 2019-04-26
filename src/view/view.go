package view

import (
	"client"
	"fmt"
	"net/http"
)

func OnlineHandler(w http.ResponseWriter, r *http.Request) {
	_ = client.GetAllClient(10)
	fmt.Fprintln(w, `{"name":"twt","ip":"192.168.1.1"}`)
}

func StartHTTPServer(port int) {

	http.HandleFunc("/getonline", OnlineHandler)

	http.Handle("/", http.FileServer(http.Dir("html")))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
