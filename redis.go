package ouwsdk

import (
	"log"
)


func SetExist(k string, e string) bool {
				
	res, err := Rds.SIsMember(Ctx, k, e).Result()
	
	if err != nil {
		
		log.Println(err)

		return false

	} else {
		log.Println(res)
	}

	return true

} // SetExist


func SetAdd(k string, v string) {

	err := Rds.SAdd(Ctx, k, v)

	if err != nil {
		log.Println(err)
	}

} // SetAdd


func MapAdd(k string) {

} // MapAdd
