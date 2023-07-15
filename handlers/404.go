package handlers

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, ResponseMessage{StatusCode: 404, Message: "route not found"}, 404)
}
