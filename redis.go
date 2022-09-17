package ouwsdk

import (
	"log"

	"github.com/go-redis/redis/v8"
)


func initRedis(h string, p string) *redis.Client {
	
	address := Addr(h, p)

  return redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})

} // InitReinitRedisdis


func SetExist(k string, e string) bool {
				
	res, err := Rds.SIsMember(ctx, k, e).Result()
	
	if err != nil {
		
		log.Println(err)

		return false

	} else {
		log.Println(res)
	}

	return true

} // SetExist


func SetAdd(k string, v string) {

	err := Rds.SAdd(ctx, k, v)

	if err != nil {
		log.Println(err)
	}

} // SetAdd


func MapAdd(k string) {

} // MapAdd
