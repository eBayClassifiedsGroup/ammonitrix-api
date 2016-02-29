package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/build"
)

//Root path output
func Root(w http.ResponseWriter, req *http.Request) {
	versionString := fmt.Sprintf("Ammonitrix Build Version %s", build.Version)
	fmt.Fprintln(w, versionString)
}

//GetIndex is the root of V1 API
func GetIndex(w http.ResponseWriter, req *http.Request) {
	esVersion, err := Elastic.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		esVersion = "Unknown"
	}
	w.Header().Set("Content-Type", "application/json")
	var result map[string]string
	result = make(map[string]string)
	result["Ammonitrix version"] = build.Version
	result["API version"] = "1"
	result["Elastic Search version"] = esVersion
	b, _ := json.Marshal(result)
	fmt.Fprintln(w, string(b))
}
