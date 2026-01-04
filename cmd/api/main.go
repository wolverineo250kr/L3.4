package main

import (
    "image-processor/internal/kafka"
    "image-processor/internal/meta"
    "io"
    "net/http"

    "github.com/google/uuid"
)

func upload(w http.ResponseWriter, r *http.Request) {
    file, _, _ := r.FormFile("file")
    defer file.Close()

    id := uuid.New().String()
    meta.Create(id)

    dst, _ := io.Create("storage/originals/" + id)
    io.Copy(dst, file)

    kafka.Send(id)
    w.Write([]byte(id))
}

func status(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    img, ok := meta.Get(id)
    if !ok {
        w.WriteHeader(404)
        return
    }
    w.Write([]byte(string(img.Status)))
}

func main() {
    http.HandleFunc("/upload", upload)
    http.HandleFunc("/image/{id}", status)
    http.Handle("/", http.FileServer(http.Dir("web")))

    http.ListenAndServe(":8080", nil)
}
