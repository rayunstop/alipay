package main

import (
	"github.com/alipay-sdk/gateway"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"
	log.Printf("alipay service is running on %s", addr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	http.HandleFunc("/service/gateway.do", gateway.GatewayService)
	log.Fatal(http.ListenAndServe(addr, nil))

}
