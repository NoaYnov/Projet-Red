package red

func (p *Personnage) Gobelin() {
	p.goblin.name = "Goblin d'entrainement"
	p.goblin.point_de_vie_max = 4000
	p.goblin.point_de_vie_actuel = 4000
	p.goblin.resistance_physique = 10
	p.goblin.resistance_magique = 10
	p.goblin.attaque = 1
	p.goblin.initiative = 10

}

func (p *Personnage) VersGeant() {
	p.versgeant.name = "Vers géant"
	p.versgeant.point_de_vie_max = 110
	p.versgeant.point_de_vie_actuel = 110
	p.versgeant.resistance_physique = 15
	p.versgeant.resistance_magique = 10
	p.versgeant.attaque = 6
	p.versgeant.initiative = 10

}

func (p *Personnage) Troll() {
	p.troll.name = "Troll"
	p.troll.point_de_vie_max = 75
	p.troll.point_de_vie_actuel = 75
	p.troll.resistance_physique = 10
	p.troll.resistance_magique = 20
	p.troll.attaque = 12
	p.troll.initiative = 10

}

func (p *Personnage) Gorgone() {
	p.gorgone.name = "Gorgone"
	p.gorgone.point_de_vie_max = 1500
	p.gorgone.point_de_vie_actuel = 1500
	p.gorgone.resistance_physique = 20
	p.gorgone.resistance_magique = 20
	p.gorgone.attaque = 20
	p.gorgone.initiative = 100

}
