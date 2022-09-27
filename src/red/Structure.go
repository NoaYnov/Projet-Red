package red

type Mob struct {
	name                string
	point_de_vie_max    int
	point_de_vie_actuel int
	resistance_physique int
	resistance_magique  int
	attaque             int
	initiative          int
}

type Personnage struct {
	age                 string
	nom                 string
	race                string
	classe              string
	niveau              int
	experience_actuel   int
	experience_max      int
	force               int
	initiative          int
	point_de_vie_max    int
	point_de_vie_actuel int
	resistance_physique int
	resistance_magique  int
	skill               map[string]int
	inventaire          map[string]int
	inventaireArmor     map[string]string
	money               int
	armor               Armor
	equiphelmet         string
	equipchestplate     string
	equippants          string
	equipboots          string
	equipgloves         string
	liminventaire       int
	limslot             int
	mana_max            int
	mana_actuel         int
	exit                bool
	goblin              Mob
	versgeant           Mob
	troll               Mob
	gorgone             Mob
}

type Armor struct {
	helmet     []string
	boots      []string
	gloves     []string
	pants      []string
	chestplate []string
}
