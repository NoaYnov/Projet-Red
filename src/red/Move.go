package red

import "fmt"

func (p *Personnage) SeDeplacer() { //Fonction qui permet de se déplacer dans le monde
	var direction int
	var choix int
	fmt.Println("Ou voulez vous aller ?")
	fmt.Println("1- Les marais de la mort")
	fmt.Println("2- les Montagnes désolées")
	fmt.Println("3- retourner au menu")
	fmt.Scanln(&direction)
	switch direction {
	case 1:
		fmt.Println("Vous vous trouvez dans les marais de la mort")
		fmt.Println("Que voulez vous faire?")
		fmt.Println("1. Parler à Fred")
		fmt.Println("2. Chercher une créature à combattre")
		fmt.Println("3-retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Print("\n")
			if p.TestIsInInventory("Ecaille de vers géant") == false {
				fmt.Println("Bien le bonjour aventurier ! Auriez-vous la gentillesses d'aider un vieux monsieur ?")
				fmt.Println("Depuis quelque temp un énorme vers creuse des trous de partout, si vous pourriez m'en débarraser vous en serait récompensé")
			} else {
				p.AddInventory("potion de mana")
				fmt.Println("Merci beaucoup aventurier ! Voici une potion de mana en récompense.")
			}
			p.SeDeplacer()
		case 2:
			fmt.Println("Vous combattez un vers géant")
			p.TrainingFightVersGeant()
		case 3:
			p.SeDeplacer()
		}
	case 2:
		fmt.Println("Vous vous trouvez dans les Montagnes désolées")
		fmt.Println("1. Parler à Darur")
		fmt.Println("2. chercher une créature à combattre")
		fmt.Println("3-retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Print("\n")
			if p.TestIsInInventory("Dent de troll") == false {
				fmt.Println("Salutation voyageur ! Hier mon chien Bernard a été tué par une saleté de Troll, j'aimerais lui rendre justice. Tuez-les et je vous récompenserais")
			} else {
				p.AddInventory("potion de vie")
				fmt.Println("Merci beaucoup aventurier ! Voici une potion de vie en récompense.")
			}
			p.SeDeplacer()
		case 2:
			fmt.Println("Le troll est supprimé pour cause de mise à jour, il resortira plus tard dans un dlc")
			p.SeDeplacer()
		case 3:
			p.SeDeplacer()

		}
	case 3:
		p.Menu()

	default:
		p.Menu()
	}
}
