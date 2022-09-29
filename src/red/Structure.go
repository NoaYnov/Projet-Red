package red

type Mob struct { // Structure des monstres
	name                string
	point_de_vie_max    int
	point_de_vie_actuel int
	resistance_physique int
	resistance_magique  int
	attaque             int
	initiative          int
}

type Personnage struct { // Structure du personnage et incluant d'autres éléments
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
	mortgoblin          bool
	mortvers            bool
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
}

type Armor struct { // Structure des armures
	helmet     []string
	boots      []string
	gloves     []string
	pants      []string
	chestplate []string
}
