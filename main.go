package main

import "net/http"

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}
func main() {

	//下面是三种处理方法

	//方式一
	hhm := helloHandler{}
	http.Handle("/hello", &hhm)

	//方式二
	ahm := aboutHandler{}
	http.HandleFunc("/about", ahm.ServeHTTP)

	//方式三：这里底层是类型转换，将这个匿名函数转换成了HandlerFunc类型
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	//创建一个HTTP服务器
	http.ListenAndServe("localhost:8080", nil)

}
