package red

import (
	"fmt"
)

func (p *Personnage) GobelinTurn(nbr_tour int) {

	fmt.Println("C'est le tour du gobelin")
	fmt.Println("Le goblin vous attaque")

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
