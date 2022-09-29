package red

import "fmt"

func (p *Personnage) VersGeantTurn(nbr_tour int) { // Fonction qui gère le tour du Vers Géant
	if p.exit == true {
		return
	}
	fmt.Println("C'est le tour du vers geant")
	fmt.Println("Le vers geant vous attaque")
	if nbr_tour%3 == 0 {
		p.versgeant.resistance_physique *= 2
		p.point_de_vie_actuel -= p.versgeant.attaque * (1 - p.resistance_physique/100)
		fmt.Println("le versgeant devient 2 fois plus résistant physiquement")
		fmt.Println("Vous avez perdu", p.versgeant.attaque*(1-p.resistance_physique/100), "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel*(1-p.resistance_physique/100), "points de vie")
	}
	if nbr_tour%3 != 0 {
		p.versgeant.resistance_physique = 15
		p.point_de_vie_actuel -= p.versgeant.attaque * (1 - p.resistance_physique/100)
		fmt.Println("Vous avez perdu", p.versgeant.attaque*(1-p.resistance_physique/100), "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel*(1-p.resistance_physique/100), "points de vie")
	}
}
