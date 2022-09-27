package red

import "fmt"

func (p *Personnage) VersGeantTurn(nbr_tour int) {

	fmt.Println("C'est le tour du vers geant")
	fmt.Println("Le vers geant vous attaque")

	if nbr_tour%3 == 0 {
		p.versgeant.resistance_physique *= 2
		p.point_de_vie_actuel -= p.versgeant.attaque
		fmt.Println("le versgeant devient 2 fois plus résistant physiquement")
		fmt.Println("Vous avez perdu", p.versgeant.attaque, "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
	if nbr_tour%3 != 0 {
		p.versgeant.resistance_physique = 15
		p.point_de_vie_actuel -= p.versgeant.attaque
		fmt.Println("Vous avez perdu", p.versgeant.attaque, "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
}
