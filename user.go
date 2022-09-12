package ouwsdk


type CharAttributes struct {
	Female        			bool      	`redis:"female"`
	Age           			int         `redis:"age"`
	Height        			int         `redis:"height"`
	Family        			string      `redis:"family"`
	Kingdom         		string      `redis:"kingdom"`
	Experience      		int         `redis:"experience"`
	Luck            		int         `redis:"luck"`
	Virtue      				int         `redis:"virtue"`
	Loyalty         		int         `redis:"loyalty"`
	Popularity    			int         `redis:"popularity"`
	Health          		int         `redis:"health"`
	Energy       				int         `redis:"energy"`
	Intelligence    		int         `redis:"intelligence"`
	Wisdom       				int         `redis:"wisdom"`
	Integrity       		int         `redis:"integrity"`
	Dexterity       		int         `redis:"dexterity"`
	Charisma        		int         `redis:"charisma"`
	Attractiveness    	int         `redis:"attractiveness"`
	Strength          	int         `redis:"strength"`
	Stamina           	int         `redis:"stamina"`
	Balance            	int         `redis:"balance"`
	Temperance          int         `redis:"temperance"`
	Recovery            int         `redis:"recovery"`
	Innovation          int         `redis:"innovation"`
}


type MilitaryAttribute struct {
	Size								int 				`redis:"size"`
	Rating          		int         `redis:"rating"`
}


type MilitaryAttributes struct {
	Dagger				int		`redis:"dagger"`
	Sword					int		`redis:"sword"`
	Scimitar			int		`redis:"scimitar"`
	Spear					int		`redis:"spear"`
	HookedSpear		int		`redis:"hookedspear"`
	Staff					int		`redis:"staff"`
	DragonBlade		int		`redis:"dragonblade"`
	Club					int		`redis:"club"`
	Hammer				int		`redis:"hammer"`
	Axe						int		`redis:"axe"`
	Halberd				int		`redis:"halberd"`
	Bow						int		`redis:"bow"`
	CrossBow			int		`redis:"crossbow"`
	HorseBow			int		`redis:"horsebow"`
	Horse					int		`redis:"horse"`
	Chariot				int		`redis:"chariot"`
	Shield				int		`redis:"shield"`
}


type MartialArtsAttributes struct {
	Bear					int		`redis:"bear"`
	Crane					int		`redis:"crane"`
	Drunken				int		`redis:"drunken"`
	IronFist			int		`redis:"ironfist"`
	Leopard				int		`redis:"leopard"`
	Mantis				int		`redis:"mantis"`
	Monkey				int		`redis:"monkey"`
	Snake					int		`redis:"snake"`
	Tiger					int		`redis:"tiger"`
}


type User struct {
	ID 						string			`redis:"id"`
	//Email         string      `redis:"email"`
	Pass 					string			`redis:"password"`
	Name          string      `redis:"name"`
	//Job          	string      `redis:"job"`
	Municipal     string      `redis:"municipal"`
	Attributes    CharAttributes
	Combat      	MilitaryAttributes
}
