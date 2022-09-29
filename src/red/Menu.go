package red

import "fmt"

func (p *Personnage) Menu() { //Menu qui renvoi vers d'autres menus ou fonctions
	personnage := p
	var result int
	var retour int

	fmt.Println("Choissisez une actions parmi: ")
	fmt.Printf("\n")
	fmt.Println("1- Afficher les informations du personnage")
	fmt.Println("2- Accéder au contenu de l’inventaire")
	fmt.Println("3- Accéder au contenu de l’inventaire d’armure")
	fmt.Println("4- Tableau de Skill")
	fmt.Println("5- Magasin")
	fmt.Println("6- Forge")
	fmt.Println("7- Entrainement")
	fmt.Println("8- se déplacer")
	fmt.Println("9- Qui sont-ils ?")
	fmt.Println("10- Quitter")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Scanln(&result)
	switch result {
	case 1:
		p.Display()
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Println("Choissisez une actions parmi:")
		fmt.Println("1- retour")
		fmt.Println("________________________________________________")
		fmt.Scanln(&retour)
		if retour == 1 {
			personnage.Menu()
		}
	case 2:
		fmt.Printf("\n")
		fmt.Println(p.inventaire)
		p.AccessInventory()
		fmt.Printf("\n")
		fmt.Printf("\n")
		fmt.Printf("\n")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Println("Quel est votre choix ?")
		fmt.Println("1- Utiliser potion de soin")
		fmt.Println("2- Utiliser potion de poison")
		fmt.Println("3- Retour")
		fmt.Println("________________________________________________")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			p.TakePotHeal()
			fmt.Println(p.point_de_vie_actuel)
			p.Menu()
		case 2:
			p.SelfPoisonPot()
			fmt.Println(p.point_de_vie_actuel)
			p.Menu()
		case 3:
			p.Menu()
		default:
			fmt.Println("________________________________________________")
			fmt.Printf("\n")
			fmt.Println("Aucun choix associe")
			fmt.Println("________________________________________________")
			fmt.Printf("\n")
			fmt.Printf("\n")
			personnage.Menu()
		}
	case 3:
		p.MenuEquipement()
	case 4:
		fmt.Printf("\n")
		p.DisplaySkill()
		fmt.Println("1- Retour")
		fmt.Println("________________________________________________")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			personnage.Menu()
		}

	case 5:
		p.Shop()
	case 6:
		p.Forge()
	case 7:
		p.TrainingFightGobelin()
	case 8:
		p.SeDeplacer()
	case 9:
		fmt.Println("Les 2 artistes sont: ")
		fmt.Println("- ABBA ")
		fmt.Println("- Steven Spielberg")
	case 10:
		fmt.Println("ALT+F4")
		return
	default:
		fmt.Println("Aucun menu associe")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Printf("\n")
		p.Menu()
	}
}
