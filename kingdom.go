package ouwsdk

import (
	"fmt"
	"log"
)


type Kingdom struct {
	Name									string      `redis:"name"`
	Capital               bool        `redis:"capital"`
	MaleRatio							int					`redis:"maleRatio"`
	MedianAge							int					`redis:"medianAge"`
	BirthRate             int					`redis:"birthRate"`
	DeathRate      				int					`redis:"deathRate"`
	Population            int					`redis:"population"`
	Land                  int         `redis:"land"`
	Wealth                int         `redis:"wealth"`
	Trees                 int         `redis:"trees"`
	Rocks                 int         `redis:"rocks"`
	Cows                 	int         `redis:"cows"`
	TaxRate               int         `redis:"taxRate"`
	ConscriptAge    			int         `redis:"conscriptAge"`
	Municipals            []Municipal `redis"municipals"`
}


func KingdomAdd(k Kingdom) {

	key := fmt.Sprintf("%s:%s", KEY_KINGDOM, k.Name)

	err := Store.HSet(ctx, key,
		FIELD_POPULATION, k.Population,
		FIELD_LAND, k.Land,
		FIELD_TREES, k.Trees,
		FIELD_ROCKS, k.Rocks,
		FIELD_WEALTH, k.Wealth
	)


	if err != nil {
		log.Println(err)
	}

	for _, m := range k.Municipals {

		key := fmt.Sprintf("%s:%s:%s", KEY_KINGDOM, k.Name, KEY_MUNICIPALS)

		SetAdd(key, m.Name)
	
	}

} // KingdomAdd
