package ouwsdk


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
	//ConscriptAge    			int         `redis:"conscriptAge"`
}
