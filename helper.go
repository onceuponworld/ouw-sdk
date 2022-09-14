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


func Addr(h string, p string, t int) string {

	var host, port string

	if t == APP_SERVER_WEB {

		host = DEFAULT_HOST
		port = DEFAULT_PORT
	
	} else if t == APP_SERVER_REDIS {

		host = DEFAULT_REDIS_HOST
		port = DEFAULT_REDIS_PORT
		
	}


	if len(h) != 0 {		
		host = h
	}

	if len(p) != 0 {
		port = p
	}

	return fmt.Sprintf("%s:%s", host, port)

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
