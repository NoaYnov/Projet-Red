package red

import (
	"fmt"
)

func (p *Personnage) DisplayMana() {
	fmt.Println("Mana : ", p.mana_actuel, "/", p.mana_max)
}

func (p *Personnage) TakePotMana() {
	fmt.Println(p.inventaire["potion de mana"])
	if p.inventaire["potion de mana"] == 0 {
		fmt.Println("Vous n'avez pas de potion de mana")
	} else {
		if p.inventaire["potion de mana"] >= 1 {
			if p.mana_actuel < p.mana_max {
				p.RemoveInventory("potion de mana")
				p.mana_actuel += 10
				if p.mana_actuel > p.mana_max {
					p.mana_actuel = p.mana_max

				}
			} else {
				fmt.Println("Tu es deja Full mana")
			}
		}
	}
}

func (p *Personnage) TestRemoveMana(nbr int) bool {
	if p.mana_actuel >= nbr {
		return true
	} else {
		return false
	}
}

func (p *Personnage) RemoveMana(nbr int) {
	p.mana_actuel -= nbr
}
