package red

import "fmt"

func (p *Personnage) Shop() {
	fmt.Println("Bienvenue dans le magasin")
	fmt.Println("Que voulez vous acheter ?")
	fmt.Println("1- Potion de soin")
	fmt.Println("2- Potion de poison")
	fmt.Println("3- spellBook : Boule de feu")
	fmt.Println("4- spellBook : Boule de glace")
	fmt.Println("5- spellBook : Boule de foudre")

	fmt.Println("6- Retour")
	var result int
	fmt.Scanln(&result)
	switch result {
	case 1:
		p.AddInventory("potion de vie")
		fmt.Println("Vous avez acheté une potion de vie")
		p.Menu()

	case 2:
		p.AddInventory("potion de poison")
		fmt.Println("Vous avez acheté une potion de poison")
		p.Menu()

	case 3:
		p.AddSkill("spellBook : Boule de feu")
		fmt.Println("Vous avez acheté un spellBook : Boule de feu")
		p.Menu()

	case 4:
		p.AddSkill("spellBook : Boule de glace")
		fmt.Println("Vous avez acheté un spellBook : Boule de glace")
		p.Menu()

	case 5:
		p.AddSkill("spellBook : Boule de foudre")
		fmt.Println("Vous avez acheté un spellBook : Boule de foudre")
		p.Menu()



	case 6:
		p.Menu()

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		p.Menu()

	
	
	
	
	}
}