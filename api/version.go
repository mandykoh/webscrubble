package api

import (
	"encoding/json"
	"net/http"

	"github.com/mandykoh/scrubble"
	"github.com/mandykoh/webscrubble"
)

func (e Endpoints) Version(w http.ResponseWriter, r *http.Request) {
	info := struct {
		Version       string `json:"version"`
		EngineVersion string `json:"engineVersion"`
	}{
		webscrubble.Version,
		scrubble.Version,
	}

	versionJson, _ := json.Marshal(&info)

	w.Header().Set("content-type", "application/json")
	w.Write(versionJson)
}
