package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http server

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	queryParm := r.URL.Query()
	name := queryParm.Get("name")
	age := queryParm.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
