package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/corona10/goimagehash"
)

type myHandler struct{}

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, r.Method)

	// POST /phash  binary
	if r.Method == "POST" && r.URL.Path == "/phash" {
		defer r.Body.Close() // 必须关闭，避免内存泄漏

		img, m, err := image.Decode(r.Body)
		if m == "" { // 图片格式 jpeg、png、gif
			log.Println("decode error:", m, "--", err)
			w.WriteHeader(415) // StatusUnsupportedMediaType
			w.Write([]byte("unsupport media type"))
			return
		}
		if err != nil {
			log.Println("decode error:", m, "--", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		hash, err := goimagehash.PerceptionHash(img)
		if err != nil {
			log.Println("phash error:", m, "--", err)
			w.WriteHeader(500)
			w.Write([]byte("file phash error"))
			return
		}

		log.Println(hash.ToString())

		w.Write([]byte(strings.Split(hash.ToString(), ":")[1]))
		return
	}
	w.WriteHeader(404)
	w.Write([]byte("not found"))
}

func main() {
	s := &http.Server{
		Addr:           ":6789",
		Handler:        &myHandler{},
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1024 * 1024 * 20, // 20MB
	}
	log.Println("http server start at", s.Addr)
	log.Fatal(s.ListenAndServe())
}
