package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
	"gopkg.in/olivere/elastic.v3"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/build"
	"github.com/gorilla/mux"
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

/*GetDataIndex lists all data and or queries them
Return all entries
curl -XGET 'http://HOST:PORT/v1/data

Can search for specific fields
curl -XGET 'http://HOST:PORT/v1/data?key=val&key2=val2'
*/
func GetDataIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	var searchResult *elastic.SearchResult
	query := elastic.NewBoolQuery()
	if len(req.URL.Query()) > 0 {
		for k, v := range req.URL.Query() {
			query = query.Must(elastic.NewTermQuery(k, v[0]))
		}
		searchResult, err = Elastic.Search().Index("ammonitrix").Query(query).Do()
	} else {
		searchResult, err = Elastic.Search().Index("ammonitrix").Query(elastic.NewMatchAllQuery()).Do()
	}
	if searchResult.Hits != nil && err == nil {
		fmt.Printf("Found a total of %d results\n", searchResult.Hits.TotalHits)
		var dataResponse []config.ElasticData
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index
			var d config.ElasticData
			d.Unmarshal(*hit.Source)
			fmt.Printf(fmt.Sprintf("[DEBUG] %s", d.Print()))
			dataResponse = append(dataResponse, d)
		}
		response := map[string]interface{}{"Result": "Success", "Data": dataResponse}
		fullResponse, _ := json.Marshal(response)
		fmt.Fprintln(w, string(fullResponse))
	} else if searchResult.Hits == nil {
		response := map[string]interface{}{"Result": "No results"}
		fullResponse, _ := json.Marshal(response)
		fmt.Println(w, string(fullResponse))
	} else if err != nil {
		response := map[string]interface{}{"Result": "Fail"}
		fullResponse, _ := json.Marshal(response)
		fmt.Println(w, string(fullResponse))
	}
}

/*GetData gets V1 data
Return entry
curl -XGET 'http://HOST:PORT/v1/data/NAME,
equivalent to curl -XGET 'http://HOST:PORT/v1/data?name=NAME'

curl -XGET 'http://HOST:PORT/v1/data/NAME,
*/
func GetData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//uses gorillatoolkit lib
	vars := mux.Vars(req)
	name := string(vars["name"])

	var err error
	var searchResult *elastic.SearchResult
	query := elastic.NewTermQuery("name", name)
	searchResult, err = Elastic.Search().Index("ammonitrix").Query(query).Do()
	if searchResult.Hits != nil && err == nil {
		fmt.Printf("Found a total of %d results\n", searchResult.Hits.TotalHits)
		var dataResponse []config.ElasticData
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index
			var d config.ElasticData
			d.Unmarshal(*hit.Source)
			fmt.Printf(fmt.Sprintf("[DEBUG] %s", d.Print()))
			dataResponse = append(dataResponse, d)
		}
		response := map[string]interface{}{"Result": "Success", "Data": dataResponse}
		fullResponse, _ := json.Marshal(response)
		fmt.Fprintln(w, string(fullResponse))
	} else {
		response := map[string]interface{}{"Result": "Fail"}
		fullResponse, _ := json.Marshal(response)
		fmt.Println(w, string(fullResponse))
	}

}

/*GetDataName gets V1 data
curl -XGET 'http://HOST:PORT/v1/data/NAME, equivalent to curl -XGET 'http://HOST:PORT/v1/data?name=NAME'
*/
func GetDataName(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"Result": "Not implemented"}
	fullResponse, _ := json.Marshal(response)
	fmt.Println(w, string(fullResponse))
}
