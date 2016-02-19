package router

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/build"
	"github.com/eBayClassifiedsGroup/ammonitrix-api/config"
	"github.com/gorilla/mux"
)

//Elastic Connection Singleton
var Elastic = &config.DefaultConfig.Elastic

//Root path output
func Root(w http.ResponseWriter, req *http.Request) {
	versionString := fmt.Sprintf("Ammonitrix Build Version %s", build.Version)
	fmt.Fprintln(w, versionString)
}

//GetIndex is the root of V1 API
func GetIndex(w http.ResponseWriter, req *http.Request) {
	versionString := fmt.Sprintf("Ammonitrix API Version 1, Build Version %s", build.Version)
	fmt.Fprintln(w, versionString)
}

/*GetDataIndex lists all data and or queries them
Can search for specific fields
curl -XGET 'http://HOST:PORT/v1/data?key=val&key2=val2'
*/
func GetDataIndex(w http.ResponseWriter, req *http.Request) {
	getParams := fmt.Sprintf("GET params: %s", req.URL.Query())
	s := []string{}
	for k, v := range req.URL.Query() {
		newTerms := strings.Join([]string{k, ":", v[0]}, "")
		s = append(s, newTerms)
	}
	searches := strings.Join(s, "&")
	fmt.Fprintln(w, getParams)
	fmt.Fprintln(w, searches)
	url := fmt.Sprintf("http://%s%s/%s/_search?q=%s", Elastic.Host, Elastic.Port, Elastic.IndexName, searches)
	r, err := http.Get(url)
	if err != nil || r.StatusCode >= 400 {
		fmt.Fprintln(w, "[ERROR] Couldn't search Elastic")
		log.Println("[ERROR] Couldn't search Elastic")
	} else {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintln(w, string(body))
	}
}

/*GetData gets V1 data
curl -XGET 'http://HOST:PORT/v1/data/NAME, equivalent to curl -XGET 'http://HOST:PORT/v1/data?name=NAME'
*/
func GetData(w http.ResponseWriter, req *http.Request) {
	//uses gorillatoolkit lib
	vars := mux.Vars(req)
	dataID := string(vars["dataID"])
	nameQuery := fmt.Sprintf("name:%s", dataID)
	fmt.Fprintln(w, nameQuery)
	url := fmt.Sprintf("http://%s%s/%s/_search?q=%s", Elastic.Host, Elastic.Port, Elastic.IndexName, nameQuery)
	r, err := http.Get(url)
	if err != nil || r.StatusCode >= 400 {
		fmt.Fprintln(w, "[ERROR] Couldn't search Elastic")
		log.Println("[ERROR] Couldn't search Elastic")
	} else {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintln(w, string(body))
	}
}
