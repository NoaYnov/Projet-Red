package red

import (
	"fmt"
)

func (p *Personnage) Playerturn() {
	var choix int
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("__________________________________________________")
	fmt.Printf("\n")
	fmt.Println("C'est votre tour")
	fmt.Printf("\n")
	fmt.Println("__________________________________________________")



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
		fmt.Println("1-Coup de poing")
		fmt.Println("2-Boule de Feu")
		fmt.Println("3-Boule de glace")
		fmt.Println("4-Boule de foudre")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			p.UseSkill("Coup de poing")
		case 2:
			p.UseSkill("Boule de feu")
		case 3:
			p.UseSkill("Boule de glace")
		case 4:
			p.UseSkill("Boule de foudre")

		default:
			fmt.Println("Vous n'avez pas choisi une option valide")
			p.Playerturn()
		}
	

	case 3:
		fmt.Println("Quel objet voulez vous utiliser?")
		fmt.Println("1-Potion de vie")
		fmt.Println("2-Potion de poison")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			p.TakePotHeal()
		case 2:
			fmt.Println("Vous avez utilisé une potion de poison sur l'ennemi")
		}

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		p.Playerturn()
	}

}
