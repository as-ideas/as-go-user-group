package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	urlMap map[string]string
	mutex  sync.RWMutex
)

func init() {
	urlMap = map[string]string{}
}

func index(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()

	tpl, _ := template.New("index").Parse(`
<html>
<body>
<form method="POST" action="/shorturl">
    URL: <input type="text" name="url">
    <input type="submit">
</form>

<ul>
    {{range $shortUrl, $longUrl := .Urls}}
    <li>
        <a href="/open?key={{$shortUrl}}" target="_blank">{{$longUrl}}</a>
    </li>
    {{end}}
</ul>

</body>
</html>
`)

	tpl.Execute(w, struct {
		Urls map[string]string
	}{urlMap})
}

func sanitizeUrl(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	} else {
		return url
	}
}

func addUrl(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	url := r.PostFormValue("url")
	if url == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	url = sanitizeUrl(url)

	key := fmt.Sprint(len(urlMap))
	urlMap[key] = url
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func openUrl(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok {
		http.Error(w, "Missing query param 'key'", http.StatusBadRequest)
		return
	}

	key := keys[0]
	if url, ok := urlMap[key]; ok {
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "No url found with key "+key, http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/shorturl", addUrl)
	http.HandleFunc("/open", openUrl)
	http.HandleFunc("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
