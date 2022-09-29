package red

import (
	"fmt"
)

func (p *Personnage) DisplayMana() { //Affiche la mana du perso
	fmt.Println("Mana : ", p.mana_actuel, "/", p.mana_max)
}

func (p *Personnage) TakePotMana() { //Utilise une potion de mana si le perso en a une et si sa mana est inferieur a sa mana max
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

func (p *Personnage) TestRemoveMana(nbr int) bool { //Test si le perso a assez de mana pour utiliser une attaque
	if p.mana_actuel >= nbr {
		return true
	} else {
		return false
	}
}

func (p *Personnage) RemoveMana(nbr int) { //Enleve de la mana au perso(pas la fonction la plus utile)
	p.mana_actuel -= nbr
}
