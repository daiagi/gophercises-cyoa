package main

import (
	_ "fmt"
	"gophersices/cyoa/cyoa"
	"html/template"
	"net/http"
	"os"
)




func main() {

	f, err:= os.Open("gopher.json")
	if err != nil {
		panic(err)

	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	defer f.Close()


	tpl := template.Must(template.New("").Parse("hello"))
	h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))
	http.ListenAndServe(":3000", h)

}
