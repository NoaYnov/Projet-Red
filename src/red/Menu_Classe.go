package red

import "fmt"

func (p Personnage) Menu_Classe() {
	var result int
	var retour int

	fmt.Println("Choissisez votre Classe parmi: ")
	fmt.Println("1- Berserk")
	fmt.Println("2- Mage")
	fmt.Println("3- Paladin")
	fmt.Println("4- Voleur")
	fmt.Println("5- Archer")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Scanln(&result)
	switch result {
	case 1:
		fmt.Println("Vous avez choisi les Berserk")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Berserk()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			classe = "Berserk"
			break
		case 2:
			p.Menu_Classe()
		}

	case 2:
		fmt.Println("Vous avez choisi les Mage")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Mage()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			classe = "Mage"
			break
		case 2:
			p.Menu_Classe()
		}
	case 3:
		fmt.Println("Vous avez choisi les Paladin")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Paladin()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			classe = "Paladin"
			break
		case 2:
			p.Menu_Classe()
		}
	case 4:
		fmt.Println("Vous avez choisi les Voleur")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Voleur()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			classe = "Voleur"
			break
		case 2:
			p.Menu_Classe()
		}
	case 5:
		fmt.Println("Vous avez choisi les Archer")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Archer()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			classe = "Archer"
			break

		case 2:
			p.Menu_Classe()
		}
	}
}
