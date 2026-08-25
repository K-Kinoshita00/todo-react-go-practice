package main

import (
	"net/http"
	"fmt"
)

func main() {
	// 参考: https://tech.yappli.io/entry/2022/05/16/Go%E3%81%AEnet/http%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%82%92%E7%90%86%E8%A7%A3%E3%81%99%E3%82%8B

	healthReq := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok"}`)
	}

	http.HandleFunc("/health", healthReq)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
