// 实现一个监听服务随机端口（一个随机的可用端口）
// 参考：https://dev.to/clavinjune/listening-to-random-available-port-in-go-6bl
package main

import (
	"log"
	"net"
	"net/http"
)

func createListener() (l net.Listener, close func()) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	return l, func() {
		_ = l.Close()
	}
}

func main() {
	listener, close := createListener()
	defer close()

	http.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 添加你的代码
	}))

	log.Panicln("listing at", listener.Addr().(*net.TCPAddr).Port)
	http.Serve(listener, nil)
}
