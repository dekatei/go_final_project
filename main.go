package main

import (
	"fmt"
	"net/http"
	//"path/filepath"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "" {
		http.ServeFile(res, req, "./web/index.html")
	} else {
		http.ServeFile(res, req, "./web"+req.URL.Path)
	}

}

func main() {
	fmt.Println("Запускаем сервер")
	http.HandleFunc(`/`, mainHandle)
	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}
