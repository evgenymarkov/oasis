package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evgenymarkov/oasis/examples/basic/internal/api"
)

func main() {
	api := api.NewAPI(http.NewServeMux())
	openapiDocument := api.Document()
	apiSpecification, _ := json.MarshalIndent(openapiDocument, "", "  ")
	fmt.Println(string(apiSpecification))
}
