package red

import "fmt"

func (p *Personnage) SeDeplacer() {
	var direction int
	var choix int
	fmt.Println("Ou voulez vous aller ?")
	fmt.Println("1- Les marais de la mort")
	fmt.Println("2- les Montagnes désolées")
	fmt.Println("3- les plaines arides")
	fmt.Println("4- retourner au menu")
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
			fmt.Println("Bien le bonjour aventurier ! Auriez-vous la gentillesses d'aider un vieux monsieur ?")
			fmt.Println("Depuis quelque temp un énorme vers creuse des trous de partout, si vous pourriez m'en débarraser vous en serait réconpensé")
			p.SeDeplacer()
		case 2:
			fmt.Println("Vous combattez un vers géant")
			p.TrainingFightVersGeant()
		case 3:
			p.SeDeplacer()
		}
	default:

	case 2:
		fmt.Println("Vous vous trouvez dans les Montagnes désolées")
		fmt.Println("1. Paarler à Darur")
		fmt.Println("2. chercher une créature à ccombattre")
		fmt.Println("3-retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("Salutation voyageur ! Hier mon chien Bernard a été tué par une saleté de Troll, j'aimerais lui rendre justice. Tuez-les et je vous récompenserais")
			p.SeDeplacer()
		case 2:
			fmt.Println("Vous combattez un troll")
			p.TrainingFightTroll()
		case 3:
			p.SeDeplacer()
		}
	case 3:
		fmt.Println("Vous vous trouvez dans les plaines arides")
		fmt.Println("1- Combattre la reine des Gorgones")
		fmt.Println("2-retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("Vous combattez la rein des Gorgones")
			p.TrainingFightGorgone()
		case 2:
			p.SeDeplacer()
		}
	case 4:
		p.Menu()
	}
}
