package red

import (
	"fmt"
)

func (p *Personnage) GobelinTurn(nbr_tour int) {

	if p.goblin.point_de_vie_actuel <= 0 {
		fmt.Println("Vous avez gagné le combat")
		fmt.Println("Vous avez gagné 15 pièces d'or")
		p.Menu()
	}
	fmt.Println("C'est le tour du gobelin")
	fmt.Println("Le goblin attaque")

	if nbr_tour%3 == 0 {
		p.point_de_vie_actuel -= p.goblin.attaque * 2
		fmt.Println("le gobelin vous infilge le double de ses degats")
		fmt.Println("Vous avez perdu", p.goblin.attaque*2, "points de vie")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}
	if nbr_tour%3 != 0 {
		p.point_de_vie_actuel -= p.goblin.attaque
		fmt.Println("Le gobelin vous a infligé", p.goblin.attaque, "points de degats")
		fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
	}

}
