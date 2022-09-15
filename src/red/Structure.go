package red

type Personnage struct {

	age int
	nom string
	race string
	classe string
	niveau int
	force int
	point_de_vie_max int
	point_de_vie_actuel int
	resistance_physique int
	resistance_magique int
	inventaire map[string]int

}

