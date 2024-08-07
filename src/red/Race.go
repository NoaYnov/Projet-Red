package red

func (p *Personnage) Elfe() { // Selection de la race Elfe
	p.race = "Elfe"
	p.niveau = 1
	p.force = 15
	p.point_de_vie_max = 110
	p.point_de_vie_actuel = 110
	p.resistance_physique = 5
	p.resistance_magique = 5
}

func (p *Personnage) Orc() { // Selection de la race Orc
	p.race = "Orc"
	p.niveau = 1
	p.force = 10
	p.point_de_vie_max = 150
	p.point_de_vie_actuel = 150
	p.resistance_physique = 10
	p.resistance_magique = 5
}

func (p *Personnage) Demon() { // Selection de la race Demon
	p.race = "Demon"
	p.niveau = 1
	p.force = 15
	p.point_de_vie_max = 200
	p.point_de_vie_actuel = 200
	p.resistance_physique = 15
	p.resistance_magique = 5
}

func (p *Personnage) Humain() { // Selection de la race Humain
	p.race = "Humain"
	p.niveau = 1
	p.force = 10
	p.point_de_vie_max = 100
	p.point_de_vie_actuel = 100
	p.resistance_physique = 5
	p.resistance_magique = 5
}

func (p *Personnage) Nain() { // Selection de la race Nain
	p.race = "Nain"
	p.niveau = 1
	p.force = 10
	p.point_de_vie_max = 100
	p.point_de_vie_actuel = 100
	p.resistance_physique = 10
	p.resistance_magique = 5
}
