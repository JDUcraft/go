package infrastructure

import (
	supportThumbnailDomain "../domain"
	"encoding/json"
	"net/http"
)

func HandleController(supportThumbnailRepository supportThumbnailDomain.SupportThumbnailRepository) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		results := supportThumbnailRepository.GetSupportThumbnails()
		w.Header().Set("Content-Type", "application/json")
		errJson := json.NewEncoder(w).Encode(results)
		if errJson != nil {
			panic(errJson)
		}
	}
}

