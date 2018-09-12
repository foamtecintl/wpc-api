package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/foamtecintl/wpc-api/repository"
	"github.com/foamtecintl/wpc-api/router"
)

var config []byte

func main() {
	readConfig()
	repository.InitDB(getConfig("database"))
	api := router.SetupAPI(router.NewRouter(getConfig("secretKey")))
	fmt.Println("server start... => http://localhost:" + getConfig("port"))
	log.Fatal(http.ListenAndServe(":"+getConfig("port"), api.MakeHandler()))
}

func readConfig() {
	dataFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	config = dataFile
}

func getConfig(key string) string {
	var objmap map[string]*json.RawMessage
	json.Unmarshal(config, &objmap)
	var value string
	json.Unmarshal(*objmap[key], &value)
	return value
}
