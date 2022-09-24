package red

func (p *Personnage) Gobelin() {
	p.goblin.name = "Goblin d'entrainement"
	p.goblin.point_de_vie_max = 4000
	p.goblin.point_de_vie_actuel = 4000
	p.goblin.resistance_physique = 10
	p.goblin.resistance_magique = 10
	p.goblin.attaque = 1

}

func (p *Personnage) Vers_geant() {
	p.vers_geant.name = "Vers géant"
	p.vers_geant.point_de_vie_max = 110
	p.vers_geant.point_de_vie_actuel = 110
	p.vers_geant.resistance_physique = 15
	p.vers_geant.resistance_magique = 10
	p.vers_geant.attaque = 6

}

func (p *Personnage) Troll() {
	p.troll.name = "Troll"
	p.troll.point_de_vie_max = 60
	p.troll.point_de_vie_actuel = 60
	p.troll.resistance_physique = 10
	p.troll.resistance_magique = 20
	p.troll.attaque = 14

}
