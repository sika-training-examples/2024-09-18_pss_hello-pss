package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()

	prefix := ""
	prefixEnv := os.Getenv("PREFIX")
	if prefixEnv != "" {
		prefix = prefixEnv + " "
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1><center>%sHello PSS.sk! %s</center></h1>\n", prefix, hostname)
	})

	fmt.Println("Listen on 0.0.0.0:8000, see: http://127.0.0.1:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
