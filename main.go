package main

import (
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
)

type HelloMessage struct {
	Message string `json:"message"`
	Random []byte `json:"random"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		randbytes := make([]byte, 512)
		rand.Read(randbytes)
		result, _ := json.Marshal(HelloMessage{
			Message: "Hello World!",
			Random: randbytes,
		})
		w.Write(result)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
