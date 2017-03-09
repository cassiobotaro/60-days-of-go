package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	txt "text/template"
)

func main() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
		<a href="/search?q={{.Q}}">{{.Q}}</a>
	</body>
</html>`
	// tpl should be parseable or a panic is raised.
	var t = template.Must(template.New("webpage").Parse(tpl))
	data := struct {
		Title     string
		Items     []string
		Q         string
		Something string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
		// It will be escaped according your position at template
		// look line 19 and the result printed
		Q:         "O'Reilly",
		Something: "away from txt",
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nTemplate from file\n")
	// parse a file to create a template

	t2, err := txt.ParseFiles("tmpl0.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := t2.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}
	// A simple template definition to test our function.
	const templateText = `{{title .}}`
	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
	const (
		//list of names, broken lines
		master = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		// template redefine block list
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		// to use Join inside template
		funcs     = txt.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	// Parse mater template
	masterTmpl, err := txt.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	// clone master template but with another block list
	overlayTmpl, err := txt.Must(masterTmpl.Clone()).Parse(overlay)
	//  Parse parses text as a template body for t. Named template definitions ({{define ...}} or {{block ...}} statements) in text define additional templates associated with t and are removed from the definition of t itself.

	if err != nil {
		log.Fatal(err)
	}
	// print master
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	// print overlay
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}

	// to see more: https://golang.org/pkg/html/template
	// and: https://golang.org/pkg/text/template
}
