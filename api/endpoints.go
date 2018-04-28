package api

import (
	"net/http"
	"encoding/json"
)

type Endpoints struct {
}

func (e Endpoints) Version(w http.ResponseWriter, r *http.Request) {
	versionJson, _ := json.Marshal(struct {
		Major string `json:"major"`
		Minor string `json:"minor"`
		Revision string `json:"revision"`
	}{
		"0", "0", "0",
	})

	w.Write(versionJson)
}