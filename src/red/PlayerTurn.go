package red

import (
	"fmt"
)

func (p *Personnage) Playerturn() {
	var choix int
	fmt.Println("C'est votre tour")
	fmt.Println("Que voulez vous faire?")
	fmt.Println("1-Attaquer")
	fmt.Println("2-Utiliser un sort")
	fmt.Println("3-Utiliser un objet")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		p.goblin.point_de_vie_actuel -= p.force
		if p.goblin.point_de_vie_actuel <= 0 {
			p.goblin.point_de_vie_actuel = 0
		}
		fmt.Println("Vous avez infligé", p.force, "points de degats")
		fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
	case 2:
		fmt.Println("Quel sort voulez vous utiliser?")
		fmt.Println("1-boule de Feu")
		fmt.Println("2-boude de glace")
		fmt.Println("3-boude de foudre")
		fmt.Scanln(choix)
		switch choix {
		case 1:
			// p.goblin.point_de_vie_actuel -= p.boule_de_feu
			if p.goblin.point_de_vie_actuel <= 0 {
				p.goblin.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé Boule de feu")
			fmt.Println("Vous avez infligé", p.force, "points de degats")
			fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
		case 2:
			// p.goblin.point_de_vie_actuel -= p.boule_de_glace
			if p.goblin.point_de_vie_actuel <= 0 {
				p.goblin.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé Boule de glace")
			fmt.Println("Vous avez infligé", p.force, "points de degats")
			fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
		case 3:
			// p.goblin.point_de_vie_actuel -= p.boule_de_foudre
			if p.goblin.point_de_vie_actuel <= 0 {
				p.goblin.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé Boule de foudre")
			fmt.Println("Vous avez infligé", p.force, "points de degats")
			fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
		}

	case 3:
		fmt.Println("Quel objet voulez vous utiliser?")
		fmt.Println("1-Potion de vie")
		fmt.Println("2-Potion de poison")
		fmt.Scanln(choix)
		switch choix {
		case 1:
			p.point_de_vie_actuel += 25
			fmt.Println("Vous avez utilisé une potion de soin")
			fmt.Println("Vous avez regagné 25 points de vie")
			fmt.Println("Vous avez", p.point_de_vie_actuel, "points de vie")
		case 2:
			fmt.Println("Vous avez utilisé une potion de poison")
		}

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		p.Playerturn()
	}

}
