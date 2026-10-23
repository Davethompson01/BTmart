package handler

import (
	"bsmart/config"
	"encoding/json"
	"io"
	"net/http"
// )

// func (c *config.Apiconfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	d, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Opps", http.StatusBadRequest)
// 		return
// 	}

// }

// func (r *http.Request) ToJson(w io.Writer) error{
// 	encode := json.NewEncoder(w)
// 	return encode.Encode(r)
// }