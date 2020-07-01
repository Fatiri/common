package message

import (
	"encoding/json"
	"net/http"
)

//JSON give response
func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}
