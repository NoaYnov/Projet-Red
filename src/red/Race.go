package red

type Race struct {
	nom_de_race string
}

func (p *Personnage) Elfe() {
	p.classe = "Elfe"
	p.niveau = 1
	p.force = 15
	p.point_de_vie_max = 110
	p.point_de_vie_actuel = 110
	p.resistance_physique = 5
	p.resistance_magique = 5
	
}

func (p *Personnage) Orc() {
	p.classe = "Orc"
	p.niveau = 1
	p.force = 10
	p.point_de_vie_max = 150
	p.point_de_vie_actuel = 150
	p.resistance_physique = 10
	p.resistance_magique = 5
}

func (p *Personnage) Demon() {
	p.classe = "Demon"
	p.niveau = 1
	p.force = 15
	p.point_de_vie_max = 200
	p.point_de_vie_actuel = 200
	p.resistance_physique = 15
	p.resistance_magique = 5
}

func (p *Personnage) Humain() {
	p.classe = "Humain"
	p.niveau = 1
	p.force = 10
	p.point_de_vie_max = 100
	p.point_de_vie_actuel = 100
	p.resistance_physique = 5
	p.resistance_magique = 5
}

func (p *Personnage) Nain() {
	p.classe = "Nain"
	p.niveau = 1
	p.force = 10
	p.point_de_vie_max = 100
	p.point_de_vie_actuel = 100
	p.resistance_physique = 10
	p.resistance_magique = 5
}
