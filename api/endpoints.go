package api

import (
	"net/http"
	"encoding/json"
	"github.com/mandykoh/webscrubble"
)

type Endpoints struct {
}

func (e Endpoints) Version(w http.ResponseWriter, r *http.Request) {
	versionJson, _ := json.Marshal(struct {
		Major string `json:"major"`
		Minor string `json:"minor"`
		Revision string `json:"revision"`
	}{
		webscrubble.VersionMajor,
		webscrubble.VersionMinor,
		webscrubble.VersionRevision,
	})

	w.Write(versionJson)
}
