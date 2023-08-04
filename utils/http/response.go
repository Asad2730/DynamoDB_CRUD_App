package Http

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *response {
	return &response{
		Status: status,
		Result: data,
	}
}

func (res *response) bytes() []byte {
	data, _ := json.Marshal(res)
	return data
}

func (res *response) string() string {
	return string(res.bytes())
}

func (res *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(res.Status)
	_, _ = w.Write(res.bytes())
	log.Print(res.string())
}

// 200
func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusOK).sendResponse(w, r)
}

// 204
func StatusNoContent(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusNoContent).sendResponse(w, r)
}

// 400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

// 404
func StatusNotFond(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

// 405
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w, r)
}

// 409
func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusConflict).sendResponse(w, r)
}

// 500
func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusInternalServerError).sendResponse(w, r)
}
