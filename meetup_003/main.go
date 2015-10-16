package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

type ShortenedLink struct {
	ShortenedURL   string
	AvailableLinks map[string]string
}

var linksMap = make(map[string]string)
var mutex sync.RWMutex

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createLink(l string) string {
	mutex.Lock()
	defer mutex.Unlock()

	linkHash := randSeq()
	linksMap[linkHash] = l
	return linkHash
}

func createLinkHandler(w http.ResponseWriter, r *http.Request) {
	//Take parameter and create link
	link := r.URL.Query()["URL"][0]
	w.WriteHeader(201)

	shortenedLink := createLink(link)
	basicTemplate().Execute(w, ShortenedLink{
		ShortenedURL:   shortenedLink,
		AvailableLinks: linksMap,
	})
}

func basicTemplate() *template.Template {
	return template.Must(template.New("index").Parse(`
    <html>
      <body>
        <form action="create_link">
          {{if .ShortenedURL}}
            <p>Your link is <a href="{{.ShortenedURL}}">here</a></p>
          {{end}}
          URL: <input type="URL" name="URL"><br>
          <input type="submit" value="Shorten">
          {{if .AvailableLinks }}
            <h3>Existing links:</h3>
            <ul>
            {{ range $key, $value := .AvailableLinks }}
              <li><a href="{{ $key }}">{{ $key }} ({{ $value }})</a></li>
            {{ end }}
            </ul>
          {{ end }}
        </form>
      </body>
    </html>
`))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path[1:]) > 0 {
		mutex.RLock()
		defer mutex.RUnlock()

		_, exists := linksMap[r.URL.Path[1:]]

		if exists {
			w.Header().Set("Location", linksMap[r.URL.Path[1:]])
			w.WriteHeader(302)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("Not found any shortened link with this hash! :("))
		}
	} else {
		w.WriteHeader(200)
		basicTemplate().Execute(w, ShortenedLink{
			AvailableLinks: linksMap,
		})
	}
}

func configuredPort() string {
	portArg := "8080"
	if len(os.Args) == 2 {
		portArg = os.Args[1]
	}

	return portArg
}

func main() {
	port := configuredPort()

	http.HandleFunc("/create_link", createLinkHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
