package red

import "fmt"

func (p *Personnage) Menu_Race() { //Menu de choix de race
	var result int
	var retour int

	fmt.Println("Il existe aussi des classes qui possèdent des forces et les faiblesses contre différent type de créature")
	fmt.Println("Choissisez votre Race parmi: ")
	fmt.Println("1- Elfes")
	fmt.Println("2- Orcs")
	fmt.Println("3- Demons")
	fmt.Println("4- Humains")
	fmt.Println("5- Nain")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Scanln(&result)
	switch result {
	case 1:
		fmt.Println("Vous avez choisi la race des Elfes")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Elfe()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			race = "Elfes"
			break
		case 2:
			p.Menu_Race()
		default:
			fmt.Println("Vous n'avez pas choisi de Race")
			p.Menu_Race()
		}
	case 2:
		fmt.Println("Vous avez choisi la race des Orcs")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Orc()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			race = "Orcs"
			break
		case 2:
			p.Menu_Race()
		default:
			fmt.Println("Vous n'avez pas choisi de Race")
			p.Menu_Race()
		}
	case 3:
		fmt.Println("Vous avez choisi la race des Demons")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Demon()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			race = "Demon"
			break
		case 2:
			p.Menu_Race()
		default:
			fmt.Println("Vous n'avez pas choisi de Race")
			p.Menu_Race()
		}
	case 4:
		fmt.Println("Vous avez choisi la race des Humains")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Humain()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			race = "Humain"
			break
		case 2:
			p.Menu_Race()
		default:
			fmt.Println("Vous n'avez pas choisi de Race")
			p.Menu_Race()
		}
	case 5:
		fmt.Println("Vous avez choisi la Nains")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.Nain()
		p.Display()
		fmt.Println("1- confirmer")
		fmt.Println("2- retour")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		fmt.Scanln(&retour)
		switch retour {
		case 1:
			race = "Nain"
			break
		case 2:
			p.Menu_Race()
		default:
			fmt.Println("Vous n'avez pas choisi de Race")
			p.Menu_Race()
		}
	default:
		fmt.Println("Vous n'avez pas choisi de Race")
		p.Menu_Race()
	}
}
