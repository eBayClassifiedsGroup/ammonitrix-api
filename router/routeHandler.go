package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/build"
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
	"github.com/gorilla/mux"
	elastigo "github.com/mattbaird/elastigo/lib"
)

//Root path output
func Root(w http.ResponseWriter, req *http.Request) {
	versionString := fmt.Sprintf("Ammonitrix Build Version %s", build.Version)
	fmt.Fprintln(w, versionString)
}

//GetIndex is the root of V1 API
func GetIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result map[string]string
	result = make(map[string]string)
	result["Ammonitrix version"] = build.Version
	result["API version"] = "1"
	b, _ := json.Marshal(result)
	fmt.Fprintln(w, string(b))
}

/*GetDataIndex lists all data and or queries them
Can search for specific fields
curl -XGET 'http://HOST:PORT/v1/data?key=val&key2=val2'
*/
func GetDataIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var search map[string]interface{}
	search = make(map[string]interface{})
	var query []string
	for k, v := range req.URL.Query() {
		//FIXME: should not be index 0
		query = append(query, fmt.Sprintf("%s:%s", k, v[0]))
	}
	queryString := strings.Join(query, "&")
	search["q"] = queryString
	out, _ := Elastic.SearchUri("ammonitrix", "event", search)
	b, _ := FormatOutput(out.Hits.Hits)
	fmt.Fprintln(w, string(b))
}

/*GetData gets V1 data
curl -XGET 'http://HOST:PORT/v1/data/NAME, equivalent to curl -XGET 'http://HOST:PORT/v1/data?name=NAME'
*/
func GetData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//uses gorillatoolkit lib
	vars := mux.Vars(req)
	dataID := string(vars["dataID"])
	nameQuery := fmt.Sprintf("name:%s", dataID)
	//elastigo query
	var search map[string]interface{}
	search = make(map[string]interface{})
	search["q"] = nameQuery
	out, _ := Elastic.SearchUri("ammonitrix", "event", search)
	b, _ := FormatOutput(out.Hits.Hits)
	fmt.Fprintln(w, string(b))
}

/*FormatOutput properly formats output of search result hits */
func FormatOutput(hits []elastigo.Hit) ([]byte, error) {
	if len(hits) >= 1 {
		var results []config.ElasticData
		for k := range hits {
			var result config.ElasticData
			result.Unmarshal(*hits[k].Source)
			results = append(results, result)
		}
		b, _ := json.Marshal(results)
		return b, nil
	}
	var results map[string]string
	results = make(map[string]string)
	results["results"] = "No match"
	b, _ := json.Marshal(results)
	return b, nil
}
