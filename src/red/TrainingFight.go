package red

import "fmt"

func (p *Personnage) TrainingFight() {
	p.Gobelin()
	var nbr_tour int = 1
	for i := 0; i >= 0; i++ {
		if p.point_de_vie_actuel > 0 {
			p.Playerturn()

		} else {
			p.Wasted()
			break
		}
		if p.goblin.point_de_vie_actuel > 0 {
			p.GobelinTurn(nbr_tour)
		} else {
			fmt.Println("Vous avez gagné le combat")
			fmt.Println("Vous avez gagné 15 pièces d'or")
			p.money += 15
			p.Menu()
			break
		}
		fmt.Println("Voulez vous continuer le combat ?")
		fmt.Println("1. Oui")
		fmt.Println("2. Non")
		var choix int
		fmt.Scanln(&choix)
		if choix == 2 {
			p.Menu()
			break
		} else {
			nbr_tour++
			fmt.Println("nous passons au tour:", nbr_tour)
		}
	}

}
