package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/mdeheij/logwrap"
	uuid "github.com/satori/go.uuid"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewV4()
	file, err := os.Create("./f/" + id.String())
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

	log.Notice("Uploaded file was", n, "bytes and saved as ", id.String())
	w.Write([]byte(fmt.Sprintf("%s", id.String())))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":5050", nil)
}
