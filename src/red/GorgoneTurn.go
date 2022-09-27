package red

import "fmt"

func (p *Personnage) GorgoneTurn(nbr_tour int) {

	fmt.Println("C'est le tour de la Gorgone")
	fmt.Println("La Gorgone vous attaque")

	if nbr_tour%3 == 0 {
		p.point_de_vie_actuel -= p.gorgone.attaque * 2

		fmt.Println("la Gorgone vous infigle le double de ses dégats")
		fmt.Println("Vous avez perdu", p.gorgone.attaque, "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
	if nbr_tour%2 == 0 {
		p.gorgone.resistance_magique *= 2
		p.point_de_vie_actuel -= p.troll.attaque
		fmt.Println("la Gorgone devient 2 fois plus résistante à la magie")
		fmt.Println("Vous avez perdu", p.gorgone.attaque, "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	} else if nbr_tour%2 != 0 {
		p.gorgone.resistance_magique = 20
		p.point_de_vie_actuel -= p.gorgone.attaque
		fmt.Println("La Gorgone vous a infligé", p.gorgone.attaque, "points de degats")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
}
