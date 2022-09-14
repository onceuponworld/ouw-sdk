package ouwsdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)


func InitRedis(a string) *redis.Client {
	
	return redis.NewClient(&redis.Options{
		Addr: a,
		Password: "",
		DB: 0,
	})

} // InitRedis


func Addr(h string, p string) string {
	return fmt.Sprintf("%s:%s", h, p)
} // Addr


func ParseConfig(f string, a *AppConfig) {

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
			}

		}
	
	}

} // ParseConfig
