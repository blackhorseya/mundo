package model

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"text/template"
)

var (
	templates    = make(map[string]*template.Template)
	templatesDir = "templates"
)

//go:embed templates/*
var f embed.FS

const (
	ExplainTemplate = "explain.tmpl"
)

func init() {
	tmplFiles, err := fs.ReadDir(f, templatesDir)
	if err != nil {
		log.Fatal("read templates dir failed", err)
	}

	var pt *template.Template
	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err = template.ParseFS(f, templatesDir+"/"+tmpl.Name())
		if err != nil {
			log.Fatal("parse template failed", err)
		}

		templates[tmpl.Name()] = pt
	}
}

// GetPromptByString get template by string.
func GetPromptByString(name string, data map[string]any) (string, error) {
	tmpl, ok := templates[name]
	if !ok {
		return "", fmt.Errorf("template %s not found", name)
	}

	var tpl bytes.Buffer
	err := tmpl.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
