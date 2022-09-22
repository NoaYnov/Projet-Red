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
	skill []string
	inventaire map[string]int
	inventaireArmor map[string]string
	money int
	armor Armor
	equiphelmet string
	equipchestplate string
	equippants string
	equipboots string
	equipgloves string


}

type Armor struct{
	helmet []string
	boots []string
	gloves []string
	pants []string
	chestplate []string

}
