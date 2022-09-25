package ouwsdk

import (
	"log"

	"github.com/go-redis/redis/v8"
)


func initRedis() {
	
	address := addr(app.Store.Host, app.Store.Port)

  Store = redis.NewClient(&redis.Options{
		Addr: address,
		Password: EMPTY_STR,
		DB: 0,
	})

} // InitRedis


func SetExist(k string, e string) bool {
				
	res, err := Store.SIsMember(ctx, k, e).Result()
	
	if err != nil {
		
		log.Println(err)

		return false

	} else {
		return res
	}

} // SetExist


func SetAdd(k string, v string) {

	err := Store.SAdd(ctx, k, v)

	if err != nil {
		log.Println(err)
	}

} // SetAdd


func SetGetAll(k string) []string {
				
	var all []string

	res, err := Store.SMembers(ctx, k).Result()
	
	if err != nil {
		log.Println(err)
	} else {
		
		for _, v := range res {
			all = append(all, v)
		}
		
	}

	return all

} // SetGetAll


func MapAdd(k string, o interface{}) {

	log.Println(o)

} // MapAdd
