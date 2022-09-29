package red

func (p *Personnage) Gobelin() { // initialise le Gobelin
	p.goblin.name = "Goblin d'entrainement"
	p.goblin.point_de_vie_max = 4000
	p.goblin.point_de_vie_actuel = 4000
	p.goblin.resistance_physique = 10
	p.goblin.resistance_magique = 10
	p.goblin.attaque = 1
	p.goblin.initiative = 10
}

func (p *Personnage) VersGeant() { // initialise le Vers Géant
	p.versgeant.name = "Vers géant"
	p.versgeant.point_de_vie_max = 110
	p.versgeant.point_de_vie_actuel = 110
	p.versgeant.resistance_physique = 15
	p.versgeant.resistance_magique = 10
	p.versgeant.attaque = 6
	p.versgeant.initiative = 26
}
