package red

import "fmt"

func (p *Personnage) Shop() { 
	var NBR int
	var result7 int
	fmt.Println("Bienvenue dans le magasin")
	fmt.Println("Vous avez", p.money, "€")
	fmt.Println("Que voulez vous acheter ?")
	fmt.Println("1- Potion de soin : 3 pièces/unité")
	fmt.Println("2- Potion de poison : 6 pièces/unité")
	fmt.Println("3- spellBook : Boule de feu : 25 pièces")
	fmt.Println("4- spellBook : Boule de glace : 40 pièces")
	fmt.Println("5- spellBook : Boule de foudre : 150 pièces")
	fmt.Println("6- Tissus : 2 pièces/unité")
	fmt.Println("7- Cuir : 4 pièces/unité")
	fmt.Println("8- Retour")
	fmt.Scanln(&result7)
	switch result7 {
	case 1:
		fmt.Println("Combien de potion de vie voulez vous acheter ?")
		fmt.Scanln(&NBR)
		p.BuySeveralItem(3,NBR, "potion de vie")
		p.Menu()

	case 2:
		fmt.Println("Combien de potion de poison voulez vous acheter ?")
		fmt.Scanln(&NBR)
		p.BuySeveralItem(6,NBR,"potion de poison")
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
		fmt.Println("Combien de tissus voulez vous acheter ?")
		fmt.Scanln(&NBR)
		p.BuySeveralItem(2,NBR,"tissu")
		p.Menu()

	case 7:
		fmt.Println("Combien de cuir voulez vous acheter ?")
		fmt.Scanln(&NBR)
		p.BuySeveralItem(4,NBR,"cuir")
		p.Menu()

	case 8:
		p.Menu()

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		p.Menu()

	
	
	
	
	}
}