package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func handleNoAuth(w http.ResponseWriter, r *http.Request) {
	urlPath := html.EscapeString(r.URL.Path)
	filename := strings.TrimPrefix(urlPath, "/")
	filePath := string(cfg.Handler.RootDirectory + "/" + filename)
	_, err := os.Stat(filePath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorMessage{Code: "404", Message: http.StatusText(404)})
		log.Println("NOT FOUND (" + urlPath + ")")
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename='%s'", filepath.Base(filePath)))
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	http.ServeFile(w, r, filePath)
}
