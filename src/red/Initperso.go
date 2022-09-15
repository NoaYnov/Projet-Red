package red

type Personnage struct {

	age int
	nom string
	classe string
	niveau int
	force int
	point_de_vie_max int
	point_de_vie_actuel int
	resistance_physique int
	resistance_magique int
	inventaire string

}

func (p Personnage) InitPerso(){
    p.age = 32
    p.nom = "Legolas"
    p.classe = "Elfe"
    p.niveau = 25
    p.force = 10
    p.point_de_vie_max = 100
    p.point_de_vie_actuel = 100
    p.resistance_physique = 50
    p.resistance_magique = 50
    p.inventaire = ("3potions de soin")

}
