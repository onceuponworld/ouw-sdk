package ouwsdk


type Demography struct {
	U10							int
	U20							int
	U30							int
	U40							int
	U50							int
	U60							int
	U70							int
	U80							int
	U90							int
	U100						int	
}

type Municipal struct {
	Name						string									`redis:"name" json:"name"`
	Males						Demography							`redis:"males" json:"males"`
	Females					Demography							`redis:"females" json:"females"`
	Latitude				int											`redis:"latitude" json:"latitude"`
	Longitude       int											`redis:"longitude" json:"longitude"`
	Magistrate      string									`redis:"magistrate" json:"magistrate"`
	Wealth          int											`redis:"wealth" json:"wealth"`
	Supply          map[string] Resource		`redis:"supply" json:"supply"`
}


func addMunicipal() {

} // addMunicipal


func updateMunicipal(m Municipal) {

} // updateMunicipal
