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


func MunicipalMapKey(kid string, mid string) string {
	return fmt.Sprintf("%s:%s:%s", KEY_KINGDOM, kid, mid)
} // MunicipalMapKey


func MunicipalSetKey(kid string) string {
	return fmt.Sprintf("%s:%s:%s", KEY_KINGDOMS, kid, KEY_MUNICIPALS)
} // MunicipalSetKey


func addMunicipal() {

} // addMunicipal


func updateMunicipal(m Municipal) {

} // updateMunicipal
