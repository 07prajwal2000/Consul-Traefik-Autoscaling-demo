package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type JsonMap map[string]any

func main() {
	port := os.Getenv("PORT")
	appId := os.Getenv("APP_ID")
	if port == "" {
		port = "3000"
	}
	if appId == "" {
		appId = strconv.Itoa(os.Getpid())
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(JsonMap{
			"hello": "world",
			"appId": appId,
		})
		w.Write(data)
	})
	fmt.Println("listening on 0.0.0.0:" + port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
