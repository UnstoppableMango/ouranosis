package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/olivere/vite"
	"github.com/spf13/pflag"
	"github.com/unmango/go/cli"
)

var indexTemplate = `
<head>
    <meta charset="UTF-8" />
    <title>My Go Application</title>
    {{ .Vite.Tags }}
</head>
<body>
    <div id="app"></div>
</body>
`

func main() {
	dev := pflag.BoolP("dev", "", false, "Development mode")

	viteFragment, err := vite.HTMLFragment(vite.Config{
		FS:    os.DirFS("frontend/dist"),
		IsDev: *dev,
	})
	if err != nil {
		cli.Fail(err)
	}

	tmpl := template.Must(template.New("index").Parse(indexTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := map[string]any{
			"Vite": viteFragment,
		}

		if err = tmpl.Execute(w, pageData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
