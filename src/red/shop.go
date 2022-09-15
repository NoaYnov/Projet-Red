package red

import "fmt"

func (p *Personnage) Shop() {
	fmt.Println("Bienvenue dans le magasin")
	fmt.Println("Que voulez vous acheter ?")
	fmt.Println("1- Potion de soin")
	fmt.Println("2- Retour")
	var result int
	fmt.Scanln(&result)
	switch result {
	case 1:
		p.AddInventory("Potion de soin")
		fmt.Println("Vous avez acheté une potion de soin")
		p.Menu()
	case 2:
		p.Menu()

	
	
	
	
	}
}