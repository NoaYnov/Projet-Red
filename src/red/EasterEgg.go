package red

func (p *Personnage) SUS() { // Fonction qui permet de devenir un GIGA CHAD
	p.inventaire = map[string]int{"potion de vie": 999, "Main de Dieu": 1, "tissu": 300}
	p.classe = "XXgameurXX"
	p.race = "GIGA CHAD"
	p.niveau = 1
	p.force = 9999999999
	p.point_de_vie_max = 9999999999
	p.point_de_vie_actuel = 9999999999
	p.resistance_physique = 9999999999
	p.resistance_magique = 9999999999
	p.money = 9999999999
}
