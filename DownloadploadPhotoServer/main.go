package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	buf, err := os.ReadFile(id+".png")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)

}

func main() {
	// Upload route
	//http.HandleFunc("/upload", uploadHandler)
	//http.HandleFunc("/download", downloadHandler)

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/upload", uploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/download/{id:[0-9]+}", downloadHandler).Methods(http.MethodGet)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!!"))
	})

	//Listen on port 8080
	http.ListenAndServe(":8181", router)
}
