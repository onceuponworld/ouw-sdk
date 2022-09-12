package ouwsdk


type User struct {
	//ID 								string			`redis:"id"`
	//Email         		string      `redis:"email"`
	Name          			string      `redis:"name"`
	Pass 								string			`redis:"password"`
	//Job          			string      `redis:"job"`
	Municipal     			string      `redis:"municipal"`
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
	Bear								int					`redis:"bear"`
	Crane								int					`redis:"crane"`
	Drunken							int					`redis:"drunken"`
	IronFist						int					`redis:"ironfist"`
	Leopard							int					`redis:"leopard"`
	Mantis							int					`redis:"mantis"`
	Monkey							int					`redis:"monkey"`
	Snake								int					`redis:"snake"`
	Tiger								int					`redis:"tiger"`
}
