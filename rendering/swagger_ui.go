package rendering

import (
	"bytes"
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"strings"
	"text/template"
)

type SwaggerUIConfig struct {
	BaseURL   string
	PageTitle string
}

type indexRenderData struct {
	BaseURL string
	Title   string
}

var (
	//go:embed static
	staticFilesFS embed.FS

	//go:embed templates/index.tmpl
	indexTemplateRaw string
)

const (
	staticFilesDir    = "static"
	indexTemplateName = "swagger-ui.index.tmpl"
)

func NewSwaggerUIHandler(config SwaggerUIConfig) http.HandlerFunc {
	var (
		indexBytes     []byte
		indexRenderErr error
	)

	indexTemplate, indexReadErr := template.New(indexTemplateName).Parse(indexTemplateRaw)

	return func(response http.ResponseWriter, _ *http.Request) {
		if indexReadErr != nil {
			message := errors.Join(errDocumentRender, indexReadErr).Error()
			http.Error(response, message, http.StatusInternalServerError)

			return
		}

		if indexBytes == nil && indexRenderErr == nil {
			indexBuffer := bytes.Buffer{}
			indexRenderErr = indexTemplate.Execute(&indexBuffer, indexRenderData{
				BaseURL: config.BaseURL,
				Title:   config.PageTitle,
			})

			if indexRenderErr == nil {
				indexBytes = indexBuffer.Bytes()
			}
		}

		if indexRenderErr != nil {
			message := errors.Join(errDocumentRender, indexRenderErr).Error()
			http.Error(response, message, http.StatusInternalServerError)

			return
		}

		response.Header().Set("Content-Type", "text/html")
		response.WriteHeader(http.StatusOK)
		_, _ = response.Write(indexBytes)
	}
}

func GetSwaggerUIStaticFiles() []string {
	files := make([]string, 0)

	_ = fs.WalkDir(
		staticFilesFS,
		staticFilesDir,
		func(path string, d fs.DirEntry, _ error) error {
			if d.IsDir() {
				return nil
			}

			file := strings.TrimPrefix(path, staticFilesDir+"/")
			files = append(files, file)

			return nil
		},
	)

	return files
}

func NewSwaggerUIStaticHandler(file string) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		http.ServeFileFS(response, request, staticFilesFS, staticFilesDir+"/"+file)
	}
}
