package ouwsdk

import (
	"log"
)


func CheckExists(k string, e string) bool {
				
	res, err := Rds.SIsMember(Ctx, k, e).Result()
	
	if err != nil {
		
		log.Println(err)

		return false

	} else {
		log.Println(res)
	}

	return true

} // CheckExists
