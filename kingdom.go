package ouwsdk


type Kingdom struct {
	Name									string      `redis:"name"`
	Capital               bool        `redis:"capital"`
	MaleRatio							int					`redis:"maleRatio"`
	MedianAge							int					`redis:"medianAge"`
	BirthRate             int					`redis:"birthRate"`
	DeathRate      				int					`redis:"deathRate"`
	Population            int					`redis:"population"`
	TaxRate               int         `redis:"taxRate"`
	ConscriptAge    			int         `redis:"conscriptAge"`
}
