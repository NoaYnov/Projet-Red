package red

import "fmt"

func (p *Personnage) TakePotHeal() { // Fonction qui permet de prendre une potion de vie
	fmt.Println(p.inventaire["potion de vie"])
	if p.inventaire["potion de vie"] == 0 {
		fmt.Println("Vous n'avez pas de potion de vie")
	} else {
		if p.inventaire["potion de vie"] >= 1 {
			if p.point_de_vie_actuel < p.point_de_vie_max {
				p.RemoveInventory("potion de vie")
				p.point_de_vie_actuel += 25
				if p.point_de_vie_actuel > p.point_de_vie_max {
					p.point_de_vie_actuel = p.point_de_vie_max
				} else {
					fmt.Println("Tu es deja Full vie")
				}
			}
		}
	}
}
