package ouwsdk

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func Init(store bool) {

	flag.Parse()

	parseConfig(store)

	initRedis()

} // Init


func CheckStr(p string, l int) bool {

	if len(p) < l {
		return false
	} else {
		return true
	}

} // CheckStr


func SendErr(m string, w http.ResponseWriter) {

	w.WriteHeader(http.StatusBadRequest)

	em := ResponseErr{
		Msg: m,
	}

	j, err := json.Marshal(em)

	if err != nil {
		log.Println(err)
	} else {
		w.Write(j)
	}

} // SendErr


func GetAddr() string {
  return addr(app.Host, app.Port)
} // GetAddr


func GetStoreAddr() string {
	return addr(app.Store.Host, app.Store.Port)
} // GetStoreAddr


func ToJson(s interface{}) []byte {

	j, err := json.Marshal(s)

	if err != nil {
		log.Println(err)
		return nil
	} else {
		return j
	}

} // ToJson


func addr(h string, p string) string {
	return fmt.Sprintf("%s:%s", h, p)
} // addr


func parseConfig(store bool) {

	_, err := os.Stat(*conf)

	if err != nil || os.IsNotExist(err) {
		log.Fatal(err)
	} else {

		buf, err := ioutil.ReadFile(*conf)

		if err != nil {
			log.Fatal(err)
		} else {

			err := json.Unmarshal(buf, &app)

			if err != nil {
				log.Println(err)
			} else {

				if len(app.Host) == 0 || len(app.Port) == 0 || (store &&
					(len(app.Store.Host) == 0 || len(app.Store.Port) == 0)) {
					log.Fatal("Error: config file invalid, please fix.")
				}

			} 

		}
	
	}

} // parseConfig
