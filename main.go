package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const ()

func main() {
	template := `
<html>
<head>
    <title>Rozwiązanie zadania PZC 2.1</title>
</head>
<body>
%s
</body>
`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		now := time.Now()
		isoFormat := now.Format("2006-01-02T15:04:05-07:00")
		fmt.Println(isoFormat)
		html := fmt.Sprintf(template, isoFormat)
		w.Write([]byte(html))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
