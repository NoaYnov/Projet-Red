package red

func (p *Personnage) TrainingFightGobelin() { // Fonction qui lance le combat contre le Gobelin en fonction de l'initiative
	p.Gobelin()
	p.CombatPlayerFirst(1)
	return

}

func (p *Personnage) TrainingFightVersGeant() { // Fonction qui lance le combat contre le Vers Géant en fonction de l'initiative
	p.VersGeant()
	p.CombatPlayerFirst(2)
	return

}

func (p *Personnage) TrainingFightTroll() { // Fonction qui lance le combat contre le Troll en fonction de l'initiative
	p.Troll()
	p.CombatPlayerFirst(3)
	return
}

func (p *Personnage) TrainingFightGorgone() { // Fonction qui lance le combat contre la Gorgone en fonction de l'initiative
	p.Gorgone()
	p.CombatMobFirst(4)
	return
}

func (p *Personnage) CombatPlayerFirst(nbr int) { // Combat tour par tour où le Joueur joue en premier
	var nbr_tour int = 1
	for t := 1; t > 0; {
		if p.exit == true {
			p.exit = false
			return
		}
		if p.point_de_vie_actuel > 0 {
			if nbr == 1 {
				if p.goblin.point_de_vie_actuel <= 0 {
					break
				}
				if p.versgeant.point_de_vie_actuel <= 0 {
					break
				}
				if p.troll.point_de_vie_actuel <= 0 {
					break
				}
				if p.gorgone.point_de_vie_actuel <= 0 {
					break
				}
			}
			p.PlayerTurn(nbr)
			if p.exit == true {
				p.exit = false
				break
			}
		} else {
			p.Wasted()
			break
		}
		if nbr == 1 {
			if p.goblin.point_de_vie_actuel > 0 {
				p.GobelinTurn(nbr_tour)
			}

		}

		if nbr == 2 {
			if p.versgeant.point_de_vie_actuel > 0 {
				p.VersGeantTurn(nbr_tour)
			}
		}
		if nbr == 3 {
			if p.troll.point_de_vie_actuel > 0 {
				p.TrollTurn(nbr_tour)
			}
		}
		if nbr == 4 {
			if p.gorgone.point_de_vie_actuel > 0 {
				p.GorgoneTurn(nbr_tour)
			}
		}

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
		if nbr2 == 3 {
			if p.troll.point_de_vie_actuel > 0 {
				p.TrollTurn(nbr_tour)
			} else {
				break
			}
			if p.exit == true {
				p.exit = false
				return

			}
		}
		if nbr2 == 4 {

			if p.gorgone.point_de_vie_actuel > 0 {
				p.GorgoneTurn(nbr_tour)
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
	}
}
