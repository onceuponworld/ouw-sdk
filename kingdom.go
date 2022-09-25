package ouwsdk

import (
	"fmt"
	"log"
)


type Kingdom struct {
	Name									string      `redis:"name" json:"name"`
	Capital               bool        `redis:"capital" json:"capital"`
	MaleRatio							int					`redis:"maleRatio" json:"maleRatio"`
	MedianAge							int					`redis:"medianAge" json:"medianAge"`
	BirthRate             int					`redis:"birthRate" json:"birthRate"`
	DeathRate      				int					`redis:"deathRate" json:"deathRate"`
	Population            int					`redis:"population" json:"population"`
	Land                  int         `redis:"land" json:"land"`
	Wealth                int         `redis:"wealth" json:"wealth"`
	Trees                 int         `redis:"trees" json:"trees"`
	Rocks                 int         `redis:"rocks" json:"rocks"`
	Cows                 	int         `redis:"cows" json:"cows"`
	TaxRate               int         `redis:"taxRate" json:"taxRate"`
	ConscriptAge    			int         `redis:"conscriptAge" json:"conscriptAge"`
	Municipals            []Municipal `redis"municipals" json:"municipals"`
}


func KingdomAdd(k Kingdom) {

	key := fmt.Sprintf("%s:%s", KEY_KINGDOM, k.Name)

	err := Store.HSet(ctx, key,
		FIELD_POPULATION, k.Population,
		FIELD_LAND, k.Land,
		FIELD_TREES, k.Trees,
		FIELD_ROCKS, k.Rocks,
		FIELD_WEALTH, k.Wealth)

	if err != nil {
		log.Println(err)
	}

	for _, m := range k.Municipals {

		key := fmt.Sprintf("%s:%s:%s", KEY_KINGDOM, k.Name, KEY_MUNICIPALS)

		SetAdd(key, m.Name)
	
	}

} // KingdomAdd


func KingdomGet(k string) *Kingdom {

	kingdom := Kingdom{}

	err := Store.HGetAll(ctx, k).Scan(&kingdom)

	if err != nil {
		
		log.Println(err)
		return nil

	} else {
		return &kingdom
	}

} // KingdomGet
