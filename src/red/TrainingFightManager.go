package red

import "fmt"

func (p *Personnage) TrainingFightGobelin() {
	p.Gobelin()
	if p.initiative > p.goblin.initiative {
		p.CombatPlayerFirst(1)
	} else {
		p.CombatMobFirst(1)
	}
}

func (p *Personnage) TrainingFightVersGeant() {
	p.VersGeant()
	if p.initiative > p.versgeant.initiative {
		p.CombatPlayerFirst(2)
	} else {
		p.CombatMobFirst(2)
	}
}

func (p *Personnage) TrainingFightTroll() {
	p.Troll()
	if p.initiative > p.troll.initiative {
		p.CombatPlayerFirst(3)
	} else {
		p.CombatMobFirst(3)
	}
}

func (p *Personnage) TrainingFightGorgone() {
	p.Gorgone()
	if p.initiative > p.gorgone.initiative {
		p.CombatPlayerFirst(4)
	} else {
		p.CombatMobFirst(4)
	}
}

func (p *Personnage) CombatPlayerFirst(nbr int) {
	var nbr_tour int = 1

	for t := 1; t > 0; {
		if p.exit == true {
			p.exit = false
			break
		}
		if p.point_de_vie_actuel > 0 {
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
			} else {
				p.money += 15
				p.experience_actuel += 15
				fmt.Println("Vous avez gagné le combat")
				fmt.Println("Vous avez gagné 15 pièces d'or")
				fmt.Println("Vous avez gagné 15 points d'expérience")
				p.UpNiveau()
				p.exit = true
				p.Menu()
			}
		}
		if nbr == 2 {
			if p.versgeant.point_de_vie_actuel > 0 {
				p.VersGeantTurn(nbr_tour)
			} else {
				p.money += 15
				p.experience_actuel += 15
				fmt.Print("Vous avez gagné le combat")
				fmt.Println("Vous avez gagné 15 pièces d'or")
				fmt.Println("Vous avez gagné 15 points d'expérience")
				if p.TestAddInventory("Ecaille de vers géant") == true {
					p.AddInventory("Ecaille de vers géant")
					fmt.Println("Vous récupérez une écaille de vers géant")
				} else {
					fmt.Println("Vous n'avez pas de place dans votre inventaire")
				}
				p.UpNiveau()
				p.exit = true
				p.SeDeplacer()
			}
		}
		if nbr == 3 {
			if p.troll.point_de_vie_actuel > 0 {
				p.TrollTurn(nbr_tour)
			} else {
				p.money += 15
				p.experience_actuel += 15
				fmt.Print("Vous avez gagné le combat")
				fmt.Println("Vous avez gagné 15 pièces d'or")
				fmt.Println("Vous avez gagné 15 points d'expérience")
				p.UpNiveau()
				p.exit = true

				p.SeDeplacer()
			}
		}
		if nbr == 4 {
			if p.gorgone.point_de_vie_actuel > 0 {
				p.GorgoneTurn(nbr_tour)
			} else {
				p.money += 15
				p.experience_actuel += 15
				fmt.Print("Vous avez gagné le combat")
				fmt.Println("Vous avez gagné 15 pièces d'or")
				fmt.Println("Vous avez gagné 15 points d'expérience")
				p.UpNiveau()
				p.exit = true

				p.SeDeplacer()
			}
		}
	}
}

func (p *Personnage) CombatMobFirst(nbr int) {
	var nbr_tour int = 1
	for e := 0; e >= 0; {
		if p.point_de_vie_actuel > 0 {
			if nbr == 1 {
				p.GobelinTurn(nbr_tour)
			}
			if nbr == 2 {
				p.VersGeantTurn(nbr_tour)
			}
			if nbr == 3 {
				p.TrollTurn(nbr_tour)
			}
			if nbr == 4 {
				p.GorgoneTurn(nbr_tour)
			}
		}
		if p.point_de_vie_actuel <= 0 {
			p.Wasted()
			break
		}
		p.PlayerTurn(nbr)
		if p.exit == true {
			p.exit = false
			break
		}
	}
}
