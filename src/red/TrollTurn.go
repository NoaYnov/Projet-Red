package red

import "fmt"

func (p *Personnage) TrollTurn(nbr_tour int) {

	fmt.Println("C'est le tour du troll")
	fmt.Println("Le troll vous attaque")

	if nbr_tour%3 == 0 {
		p.troll.resistance_magique *= 2
		p.point_de_vie_actuel -= p.troll.attaque
		fmt.Println("le troll devient 2 fois plus résistant à la magie")
		fmt.Println("Vous avez perdu", p.troll.attaque, "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
	if nbr_tour%3 != 0 {
		p.troll.resistance_magique = 20
		p.point_de_vie_actuel -= p.troll.attaque
		fmt.Println("Le troll vous a infligé", p.troll.attaque, "points de degats")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
}
