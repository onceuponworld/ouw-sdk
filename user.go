package ouwsdk


type CharAttributes struct {
	IsFemale        		bool      	`redis:"isFemale"`
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
	Dagger				MilitaryAttribute		`redis:"dagger"`
	Sword					MilitaryAttribute		`redis:"sword"`
	Scimitar			MilitaryAttribute		`redis:"scimitar"`
	Spear					MilitaryAttribute		`redis:"spear"`
	HookedSpear		MilitaryAttribute		`redis:"hookedSpear"`
	Staff					MilitaryAttribute		`redis:"staff"`
	DragonBlade		MilitaryAttribute		`redis:"dragonBlade"`
	Club					MilitaryAttribute		`redis:"club"`
	Hammer				MilitaryAttribute		`redis:"hammer"`
	Axe						MilitaryAttribute		`redis:"axe"`
	Halberd				MilitaryAttribute		`redis:"halberd"`
	Bow						MilitaryAttribute		`redis:"bow"`
	CrossBow			MilitaryAttribute		`redis:"crossbow"`
	HorseBow			MilitaryAttribute		`redis:"horseBow"`
	Horse					MilitaryAttribute		`redis:"horse"`
	Chariot				MilitaryAttribute		`redis:"chariot"`
	Shield				MilitaryAttribute		`redis:"shield"`
}


type User struct {
	ID 						string			`json:"id"`
	Email         string      `json:"email"`
	Pass 					string			`json:"password"`
	Name          string      `json:"name"`
	Job          	string      `json:"job"`
	Municipal     string      `json:"municipal"`
	Attributes    CharAttributes
	Combat      	MilitaryAttributes
}
