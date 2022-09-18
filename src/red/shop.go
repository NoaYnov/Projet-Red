package red

import "fmt"

func (p *Personnage) Shop() { 
	fmt.Println("Bienvenue dans le magasin")
	fmt.Println("Vous avez", p.money, "€")
	fmt.Println("Que voulez vous acheter ?")
	fmt.Println("1- Potion de soin : 3 pièces")
	fmt.Println("2- Potion de poison : 6 pièces")
	fmt.Println("3- spellBook : Boule de feu : 25 pièces")
	fmt.Println("4- spellBook : Boule de glace : 40 pièces")
	fmt.Println("5- spellBook : Boule de foudre : 150 pièces")

	fmt.Println("6- Retour")
	var result int
	fmt.Scanln(&result)
	switch result {
	case 1:
		p.BuyItem(3, "potion de vie")
		p.Menu()

	case 2:
		p.BuyItem(6,"potion de poison")
		p.Menu()

	case 3:
		p.BuySkill(25,"Boule de feu")
		p.Menu()

	case 4:
		p.BuySkill(40,"Boule de glace")
		p.Menu()

	case 5:
		p.BuySkill(150,"Boule de foudre")
		p.Menu()



	case 6:
		p.Menu()

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		p.Menu()

	
	
	
	
	}
}