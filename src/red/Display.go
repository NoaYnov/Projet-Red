package red


import "fmt"

func (p *Personnage) Display() {
	fmt.Println("Age :", p.age)
	fmt.Println("Nom :", p.nom)
	fmt.Println("Race :", p.race)
	fmt.Println("Classe :", p.classe)
	fmt.Println("Niveau :", p.niveau)
	fmt.Println("Force :", p.force)
	fmt.Println("PV max:", p.point_de_vie_max)
	fmt.Println("PV actuel :", p.point_de_vie_actuel)
	fmt.Println("Resistance Physique :", p.resistance_physique)
	fmt.Println("Resistance Magique :", p.resistance_magique)
	fmt.Println("-----------------------")
}
