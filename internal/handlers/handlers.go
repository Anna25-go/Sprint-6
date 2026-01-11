package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../index.html")
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, fileHeader, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	translate, err := service.Translate(string(content))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(fileHeader.Filename)
	name := time.Now().UTC().Format("2006-01-02_15-04-05") + ext

	if err := os.WriteFile(name, []byte(translate), 0755); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = res.Write([]byte(translate))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
