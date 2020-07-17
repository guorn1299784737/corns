package http

import (
	"net/http"
)

func get(request *http.Request, name string) string {
	query := request.URL.Query()
	if name != "" {
		return query.Get(name)
	}

	return ""
}

func post()  {

}
