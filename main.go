package main

import (
	"net/http"
	"os"
)

func main() {
	serverPort := ":8000"

	fileServer :=  http.FileServer(http.Dir("./static"))

	if os.Getenv("PORT") != "" {
		serverPort = os.Getenv("PORT")
	}

	http.HandleFunc("/", fileServer.ServeHTTP)


	panic(http.ListenAndServe(serverPort, nil))
}
