package utility

import (
	"encoding/json"
	"net/http"
)

//format response in json format
func ResponseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}