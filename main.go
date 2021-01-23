package main

import (
	"github.com/JokeCiCi/goim/ctrl"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", ctrl.Login)
	http.HandleFunc("/user/register", ctrl.Register)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
