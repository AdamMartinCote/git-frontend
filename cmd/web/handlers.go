package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func (app *AppConfig) render(w http.ResponseWriter, r *http.Request) {
	filePrefix, _ := filepath.Abs("./cmd/web/templates")
	t := "index.gohtml"
	files := []string{
		fmt.Sprintf("%s/%s", filePrefix, t),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	context := struct{ Port string }{
		Port: app.Port,
	}
	if err := tmpl.Execute(w, context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type jsonResponse struct {
	Error bool `json:"error"`
	Data  any  `json:"data,omitempty"`
}

func (app *AppConfig) commits(w http.ResponseWriter, _ *http.Request) {
	commits := GetHistory(app.GitRoot, 50)
	payload := jsonResponse{
		Error: false,
		Data:  commits,
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(out)
}
