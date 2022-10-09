package ouwsdk


import (
	"fmt"
)


type Municipal struct {
	Name						string									`redis:"name" json:"name"`
	Latitude				int											`redis:"latitude" json:"latitude"`
	Longitude       int											`redis:"longitude" json:"longitude"`
	Magistrate      string									`redis:"magistrate" json:"magistrate"`
	Wealth          int											`redis:"wealth" json:"wealth"`
	Supply          map[string] Resource		`redis:"supply" json:"supply"`
}


func MunicipalKey(kid string, mid string) string {
	return fmt.Sprintf("%s:%s:%s", KEY_KINGDOM, kid, mid)
} // MunicipalKey


func MunicipalsKey(kid string) string {
	return fmt.Sprintf("%s:%s:%s", KEY_KINGDOM, kid, KEY_MUNICIPALS)
} // MunicipalsKey


func addMunicipal(m Municipal) {

} // addMunicipal


func updateMunicipal(m Municipal) {

} // updateMunicipal
