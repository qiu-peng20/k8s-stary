package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	str := []byte(fmt.Sprintf("hostName is %v", s))
	_, _ = w.Write(str)
}

func main() {
	http.HandleFunc("/", indexHandle)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
