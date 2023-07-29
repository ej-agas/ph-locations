package handlers

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	EmptyResponse(w, http.StatusNotFound)
}
