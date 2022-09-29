package red

import (
	"fmt"
)

func (p *Personnage) GobelinTurn(nbr_tour int) { // Fonction qui permet de faire le tour du gobelin lors d'un combat

	fmt.Println("C'est le tour du gobelin")
	fmt.Println("Le goblin vous attaque")

	if nbr_tour%3 == 0 {
		p.point_de_vie_actuel -= (p.goblin.attaque * 2) * (1 - p.resistance_physique/100)
		fmt.Println("le gobelin vous infilge le double de ses degats")
		fmt.Println("Vous avez perdu", p.goblin.attaque*2*(1-p.resistance_physique/100), "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
	if nbr_tour%3 != 0 {
		p.point_de_vie_actuel -= p.goblin.attaque * (1 - p.resistance_physique/100)
		fmt.Println("Le gobelin vous a infligé", p.goblin.attaque*(1-p.resistance_physique/100), "points de degats")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
}
