package ouwsdk

type Resource struct {
	Name						string			`redis:"name"`
	Description			string			`redis:"description"`
	Quality					int					`redis:"quality"`
	Quantity				int					`redis:"quantity"`
	IsEdible        bool				`redis:"isEdible"`
	IsPoisonous     bool        `redis:"isPoisonous"`
}
