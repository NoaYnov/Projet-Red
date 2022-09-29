package red

import (
	"fmt"
	"time"
)

func (p *Personnage) SelfPoisonPot() { // fonction qui permet de boire une potion de poison
	fmt.Println(p.inventaire["potion de poison"])
	if p.inventaire["potion de poison"] >= 1 {
		p.RemoveInventory("potion de poison")
		interval := time.Second * 1
		for i := 0; i < 3; i++ {
			p.point_de_vie_actuel -= 10
			time.Sleep(interval)
		}
	}
}

func (p *Personnage) TakePotPoison(nbr int) { // fonction qui permet d'utiliser une potion de poison sur un ennemi
	fmt.Println(p.inventaire["potion de poison"])
	if p.inventaire["potion de poison"] >= 1 {
		p.RemoveInventory("potion de poison")
		interval := time.Second * 1
		for i := 0; i < 3; i++ {
			if nbr == 1 {
				p.goblin.point_de_vie_actuel -= 10
				time.Sleep(interval)
			}
			fmt.Println("Le gobelin a perdu 10 points de vie")
			fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
			if nbr == 2 {
				p.versgeant.point_de_vie_actuel -= 10
				time.Sleep(interval)
			}
			fmt.Println("Le vers a perdu 10 points de vie")
			fmt.Println("PV du vers:", p.versgeant.point_de_vie_actuel, "/", p.versgeant.point_de_vie_max)
			if nbr == 3 {
				p.troll.point_de_vie_actuel -= 10
				time.Sleep(interval)
			}
			fmt.Println("Le troll a perdu 10 points de vie")
			fmt.Println("PV du troll:", p.troll.point_de_vie_actuel, "/", p.troll.point_de_vie_max)
			if nbr == 4 {
				p.gorgone.point_de_vie_actuel -= 10
				time.Sleep(interval)
			}
			fmt.Println("La gorgone a perdu 10 points de vie")
			fmt.Println("PV de la gorgone:", p.gorgone.point_de_vie_actuel, "/", p.gorgone.point_de_vie_max)
		}
	} else {
		fmt.Println("Vous n'avez pas de potion de poison")
		p.PlayerTurn(nbr)
	}
}
