package view

import (
	"client"
	"encoding/json"
	"fmt"
	"net/http"
)

func OnlineHandler(w http.ResponseWriter, r *http.Request) {
	cs := client.GetAllClient(10)
	data, _ := json.Marshal(cs)
	fmt.Fprintln(w, string(data))
}

func StartHTTPServer(port int) {

	http.HandleFunc("/getonline", OnlineHandler)

	http.Handle("/", http.FileServer(http.Dir("html")))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
