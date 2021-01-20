package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"cyoa"
)

func main() {
	fileName := flag.String("f", "data.json", "file json name")
	flag.Parse()

	adventureRepo, err := cyoa.NewAdventureRepository(*fileName)
	if err != nil {
		log.Fatalln(err)
	}

	for _, adventure := range adventureRepo.Adventures {
		http.HandleFunc("/"+adventure.Slug, cyoa.TemplateHandler(adventure))
	}
	http.Handle("/", http.RedirectHandler("/intro", http.StatusTemporaryRedirect))

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
