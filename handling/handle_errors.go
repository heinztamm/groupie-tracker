package handling

import (
	"fmt"
	"net/http"
)

/*  Handle404 sets the HTTP status code to 404 (Not Found) and sends a corresponding message to the client. */
func Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Not Found")
}

func Handle400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "400 Bad Request")
}
func Handle500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "500 Internal Server Error")
}
