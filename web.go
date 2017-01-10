package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct{
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, m.message)
}

func main(){
	mux := http.NewServeMux()
	mh1 := &messageHandler{" welcome to go !"}
	mux.Handle("/welcome",mh1)
	
	mh2 := &messageHandler{"go is cool~"}
	mux.Handle("/message",mh2)

	log.Println("Listening ...")
	http.ListenAndServe(":19090",mux)
}
