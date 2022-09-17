package ouwsdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)


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


func InitRedis(h string, p string) *redis.Client {
	
	address := Addr(h, p)

	return redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})

} // InitRedis


func Addr(h string, p string) string {
	return fmt.Sprintf("%s:%s", h, p)
} // Addr


func ParseConfig(f string, a *AppConfig, store bool) {

	_, err := os.Stat(f)

	if err != nil || os.IsNotExist(err) {
		log.Fatal(err)
	} else {

		buf, err := ioutil.ReadFile(f)

		if err != nil {
			log.Fatal(err)
		} else {

			err := json.Unmarshal(buf, &a)

			if err != nil {
				log.Println(err)
			} else {

				if len(a.Host) == 0 || len(a.Port) == 0 || (store &&
					(len(a.Store.Host) == 0 || len(a.Store.Port) == 0)) {
					log.Fatal("Error: config file invalid, please fix.")
				}

			} 

		}
	
	}

} // ParseConfig
