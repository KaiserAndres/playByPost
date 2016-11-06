package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.String()[7:]
		file, err := os.Open("static/" + fileName)
		fileText := make([]byte, 2048, 2048)

		if err != nil {
			fmt.Println(err.Error())
		}

		file.Read(fileText)
		w.Write(fileText)
	})

	fmt.Print(http.ListenAndServe(":8080", nil))
}
