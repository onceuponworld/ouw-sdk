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
	Males						Demography							`redis:"males"`
	Females					Demography							`redis:"females"`
	Latitude				int											`redis:"latitude"`
	Longitude       int											`redis:"longitude"`
	Magistrate      string									`redis:"magistrate"`
	Wealth          int											`redis:"wealth"`
	Supply          map[string] Resource		`redis:"supply"`
}


func addMunicipal() {

} // addMunicipal
