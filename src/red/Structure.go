package red

type Mob struct {
	name                string
	point_de_vie_max    int
	point_de_vie_actuel int
	resistance_physique int
	resistance_magique  int
	attaque             int
}

type Personnage struct {
	age                 int
	nom                 string
	race                string
	classe              string
	niveau              int
	force               int
	point_de_vie_max    int
	point_de_vie_actuel int
	resistance_physique int
	resistance_magique  int
	skill               []string
	inventaire          map[string]int
	inventaireArmor     map[string]string
	money               int
	armor               Armor
	equiphelmet         string
	equipchestplate     string
	equippants          string
	equipboots          string
	equipgloves         string
	goblin              Mob
	vers_geant          Mob
	troll               Mob
}

type Armor struct {
	helmet     []string
	boots      []string
	gloves     []string
	pants      []string
	chestplate []string
}

type Deplacement struct {
	Direction string
}
