package red

import "fmt"

func (p *Personnage) TrainingFightGobelin() { // Fonction qui lance le combat contre le Gobelin en fonction de l'initiative
	p.mortgoblin = false
	p.Gobelin()
	p.CombatPlayerFirst(1)

}

func (p *Personnage) TrainingFightVersGeant() { // Fonction qui lance le combat contre le Vers Géant en fonction de l'initiative
	p.mortvers = false
	p.VersGeant()
	p.CombatPlayerFirst(2)

}

func (p *Personnage) CombatPlayerFirst(nbr int) { // Combat tour par tour où le Joueur joue en premier
	var nbr_tour int = 1
	for t := 1; t > 0; {
		if p.point_de_vie_actuel > 0 {
			p.PlayerTurn(nbr)
		} else {
			p.Wasted()
			break
		}
		if nbr == 1 {
			if p.mortgoblin == true {
				break
			}

			if p.goblin.point_de_vie_actuel > 0 {
				p.GobelinTurn(nbr_tour)
			}

		}
		if nbr == 2 {
			if p.mortvers == true {
				break
			}
			if p.versgeant.point_de_vie_actuel > 0 {
				p.VersGeantTurn(nbr_tour)
			}
		}
		fmt.Println("Nous passons au tour", nbr_tour+1)
	}
}

func (p *Personnage) CombatMobFirst(nbr2 int) { // Combat tour par tour où le Monstre joue en premier
	var nbr_tour int = 1
	for t := 1; t > 0; {
		if p.exit == true {
			p.exit = false
			return
		}
		if nbr2 == 1 {
			if p.goblin.point_de_vie_actuel > 0 {
				p.GobelinTurn(nbr_tour)
			} else {
				break
			}
		}
		if nbr2 == 2 {
			if p.versgeant.point_de_vie_actuel > 0 {
				p.VersGeantTurn(nbr_tour)
			} else {
				break
			}
			if p.exit == true {
				p.exit = false
				return

			}
		}

		if p.point_de_vie_actuel > 0 {
			p.PlayerTurn(nbr2)
		} else {
			p.Wasted()
			break
		}
		fmt.Println("Nous passons au tour", nbr_tour+1)

	}
}
