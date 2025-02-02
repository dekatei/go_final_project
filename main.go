package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	var filePath string

	if req.URL.Path == "/" {
		filePath = filepath.Join("web", "index.html")
	} else {
		filePath = filepath.Join("web", req.URL.Path)
	}
	http.ServeFile(res, req, filePath)

}

func main() {
	fmt.Println("Запускаем сервер")
	/*http.HandleFunc(`/`, mainHandle)
	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")*/
	go func() {
		err := http.ListenAndServe(":7540", nil)
		if err != nil {
			panic(err)
		}
	}()
}
