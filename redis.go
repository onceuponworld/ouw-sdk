package ouwsdk

import (
	"log"
)


func checkExists(k string, e string) bool {
				
	res, err := rds.SIsMember(ctx, k, e).Result()
	
	if err != nil {
		
		log.Println(err)

		return false

	} else {
		log.Println(res)
	}

	return true

} // checkExists
