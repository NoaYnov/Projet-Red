package red

import (
	"fmt"
)

func (p *Personnage) PlayerTurn(nbr int) { //Permet de déclencher le tour du joueur avec toutes ces possibilitées d'actions
	if p.exit == true {
		p.exit = false
		return
	}
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
	fmt.Println("4- Fuir le combat")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		if nbr == 1 {
			p.goblin.point_de_vie_actuel -= p.force * (1 - p.goblin.resistance_physique/100)
			if p.goblin.point_de_vie_actuel <= 0 {
				p.goblin.point_de_vie_actuel = 0
				p.money += 15
				p.experience_actuel += 15
				fmt.Print("Vous avez gagné le combat")
				fmt.Println("Vous avez gagné 15 pièces d'or")
				fmt.Println("Vous avez gagné 15 points d'expérience")
				p.UpNiveau()
				p.exit = true
				p.Menu()

			} else {
				fmt.Println("Vous avez infligé", p.force*(1-p.goblin.resistance_physique/100), "point de dégats")
				fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
			}

		}
		if nbr == 2 {
			p.versgeant.point_de_vie_actuel -= p.force * (1 - p.versgeant.resistance_physique/100)
			if p.versgeant.point_de_vie_actuel <= 0 {
				p.versgeant.point_de_vie_actuel = 0
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

			} else {
				fmt.Println("Vous avez infligé", p.force*(1-p.versgeant.resistance_physique/100), "point de dégats")
				fmt.Println("PV du vers géant:", p.versgeant.point_de_vie_actuel, "/", p.versgeant.point_de_vie_max)
			}
		}
		if nbr == 3 {
			p.troll.point_de_vie_actuel -= p.force * (1 - p.troll.resistance_physique/100)
			if p.troll.point_de_vie_actuel <= 0 {
				p.troll.point_de_vie_actuel = 0
				p.money += 15
				p.experience_actuel += 15
				fmt.Println("Vous avez gagné le combat")
				fmt.Println("Vous avez gagné 15 pièces d'or")
				fmt.Println("Vous avez gagné 15 points d'expérience")
				if p.TestAddInventory("Dent de troll") == true {
					p.AddInventory("Dent de troll")
					fmt.Println("Vous récupérez une dent de troll")
				} else {
					fmt.Println("Vous n'avez pas de place dans votre inventaire")
				}
				p.UpNiveau()
				p.exit = true
				p.SeDeplacer()

			} else {
				fmt.Println("Vous avez infligé", p.force*(1-p.troll.resistance_physique/100), "point de dégats")
				fmt.Println("PV du troll:", p.troll.point_de_vie_actuel, "/", p.troll.point_de_vie_max)
			}
		}
		if nbr == 4 {
			p.gorgone.point_de_vie_actuel -= p.force * (1 - p.gorgone.resistance_physique/100)
			if p.gorgone.point_de_vie_actuel <= 0 {
				p.gorgone.point_de_vie_actuel = 0
				p.money += 15
				p.experience_actuel += 15
				fmt.Println("Vous avez gagné le combat")
				fmt.Println("Vous avez gagné 15 pièces d'or")
				fmt.Println("Vous avez gagné 15 points d'expérience")
				if p.TestAddInventory("médaillon de la victoire") == true {
					p.AddInventory("médaillon de la victoire")
					fmt.Println("Vous récupérez un médaillon de la victoire")
				} else {
					fmt.Println("Vous n'avez pas de place dans votre inventaire")
				}
				p.UpNiveau()
				p.exit = true
				p.SeDeplacer()
			} else {
				fmt.Println("Vous avez infligé", p.force*(1-p.gorgone.resistance_physique/100), "point de dégats")
				fmt.Println("PV de la gorgone:", p.gorgone.point_de_vie_actuel, "/", p.gorgone.point_de_vie_max)
			}
		}
	case 2:
		fmt.Println("Quel sort voulez vous utiliser?")
		fmt.Println("1- coup de poing / 10 mana")
		fmt.Println("2- Boule de Feu / 20 mana")
		fmt.Println("3- Boule de glace / 30 mana")
		fmt.Println("4- Boule de foudre / 50 mana")
		fmt.Println("5- Retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			if p.TestRemoveMana(10) == true {
				p.UseSkill("Coup de poing", nbr)
			} else {
				fmt.Println("Vous n'avez pas assez de mana")
			}
		case 2:
			if p.TestRemoveMana(20) == true {
				p.UseSkill("Boule de feu", nbr)
			} else {
				fmt.Println("Vous n'avez pas assez de mana")
			}
		case 3:
			if p.TestRemoveMana(30) == true {
				p.UseSkill("Boule de glace", nbr)
			} else {
				fmt.Println("Vous n'avez pas assez de mana")
			}
		case 4:
			if p.TestRemoveMana(50) == true {
				p.UseSkill("Boule de foudre", nbr)
			} else {
				fmt.Println("Vous n'avez pas assez de mana")
			}
		case 5:
			p.PlayerTurn(nbr)
			break
		default:
			fmt.Println("Vous n'avez pas choisi une option valide")
			p.PlayerTurn(nbr)
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
			p.TakePotPoison(nbr)
			fmt.Println("Vous avez utilisé une potion de poison sur l'ennemi")
		}
	case 4:
		p.money -= 2
		fmt.Println("Vous avez fuit le combat")
		fmt.Println("Vous avez perdu 2 pieces d'or")
		p.goblin.point_de_vie_actuel = 0
		p.versgeant.point_de_vie_actuel = 0
		p.troll.point_de_vie_actuel = 0
		p.gorgone.point_de_vie_actuel = 0
		p.exit = true
		p.Menu()
	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		p.PlayerTurn(nbr)
	}
}
