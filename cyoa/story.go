package cyoa

import (
	"encoding/json"
	_ "fmt"
	"html/template"
	"io"
	"net/http"
)

func init() {
	tpl = template.Must(template.New("").Parse(dfltTamplateHandler))

}

var tpl *template.Template

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil

}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type HandlerOpt func(*handler)

func WithTemplate (t *template.Template) HandlerOpt {
	return func(h *handler) { h.t = t }
}

func NewHandler(s Story, opt ...HandlerOpt) http.Handler {
	h := handler{s, tpl}
	for _, optFunc := range opt {
		optFunc(&h)
	}
	return h
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" || path == "" {
		path = "/intro"
	}
	path = path[1:]
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			http.Error(w, "Something Went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)

}

var dfltTamplateHandler = `<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose you own adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
        <li>
            <a href="/{{.Chapter}}">{{.Text}}</a>
        </li>
        {{end}}
    </ul>
    
</body>
</html>
`
