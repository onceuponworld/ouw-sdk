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


func KingdomKey(n string) string {
	return fmt.Sprintf("%s:%s", KEY_KINGDOM, n)
} // KingdomKey


func KingdomsKey() string {
	return KEY_KINGDOMS
} // KingdomsKey


func KingdomAdd(k Kingdom) {

	key := fmt.Sprintf("%s:%s", KEY_KINGDOM, k.Name)

	err := Store.HSet(ctx, key,
		FIELD_NAME, k.Name,
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

	key := KingdomKey(k)

	err := Store.HGetAll(ctx, key).Scan(&kingdom)

	if err != nil {
		
		log.Println(err)
		return nil

	} else {
		return &kingdom
	}

} // KingdomGet


func KingdomGetAll() []Kingdom {

	var kingdoms []Kingdom

	keys := SetGetAll(KingdomsKey())

	for _, k := range keys {

		kingdom := KingdomGet(k)

		kingdoms = append(kingdoms, *kingdom)

	}

	return kingdoms

} // KingdomGetAll


func KingdomUpdate(k Kingdom) {

	key := KingdomKey(k.Name)

	err := Store.HSet(ctx, key,
		FIELD_NAME, k.Name,
		FIELD_POPULATION, k.Population,
		FIELD_LAND, k.Land,
		FIELD_WEALTH, k.Wealth,
		FIELD_TREES, k.Trees,
		FIELD_ROCKS, k.Rocks,
		FIELD_COWS, k.Cows)

	if err != nil {
		log.Println(err)
	}

} // KingdomUpdate
