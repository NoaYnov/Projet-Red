package red

import "fmt"
func (p *Personnage)Menu(){
	
	personnage := p
	var result int
	var retour int

	fmt.Println("Choissisez une actions parmi: ")
	fmt.Println("1- Afficher les informations du personnage")
	fmt.Println("2- Accéder au contenu de l’inventaire")
	fmt.Println("3- Tableau de Skill")
	fmt.Println("4- Shop")
	fmt.Println("5- Quitter")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")

	fmt.Scanln(&result)
	switch result {
	case 1:
		p.Display()
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Println("Choissisez une actions parmi: retour ")
		fmt.Println("________________________________________________")
		fmt.Scanln(&retour)

		if retour == 3 {
			personnage.Menu()
		}

	case 2:
		fmt.Printf("\n")
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
			p.PoisonPot()
			fmt.Println(p.point_de_vie_actuel)
			p.Menu()

		case 3:
			personnage.Menu()

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
		fmt.Printf("\n")
		p.DisplaySkill()
		fmt.Println("1- Description des sorts")
		fmt.Println("2- Retour")
		fmt.Println("________________________________________________")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			// p.Description()
			fmt.Println("lil")
			personnage.Menu()
		
		case 2:
			personnage.Menu()

		}



	case 4:
		p.Shop()








	case 5:
		fmt.Println("Au revoir salope!!")
		break
	default:
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Println("Aucun menu associe")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Printf("\n")
		p.Menu()

	}
}

