package red

func (p *Personnage) SUS(){
	p.inventaire = map[string]int{"potion de vie": 999, "Main de Dieu": 1}
	p.classe = "XXgamerXX"
	p.race = "GIGA CHAD"
	p.niveau = 1
	p.force = 9999999999
	p.point_de_vie_max = 9999999999
	p.point_de_vie_actuel = 9999999999
	p.resistance_physique = 9999999999
	p.resistance_magique = 9999999999
}