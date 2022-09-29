package red

import "fmt"

func (p *Personnage) AccessMoney() { // Affiche l'argent du joueur
	println(p.money)
}

func (p *Personnage) AddMoney(amount int) { // Permet d'ajouter de l'argent au joueur
	p.money += amount
}

func (p *Personnage) RemoveMoney(amount int) { // Permet de retirer de l'argent au joueur
	p.money -= amount
}

func (p *Personnage) TestRemoveMoney(amount int) bool { //Test si le joueur a assez d'argent
	if p.money >= amount {
		return true
	} else {
		return false
	}
}

func (p *Personnage) BuySkill(money int, skill string, degats int) { //Permet d'acheter un skill tout en verifiant qu'il y est assez d'argent
	if p.TestAddSkill(skill) == false {
		fmt.Println("Vous avez déjà ce skill")
	} else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	} else {
		p.RemoveMoney(money)
		p.AddSkill(skill, degats)
		fmt.Println("Vous avez appris le", skill)
	}
}

func (p *Personnage) TestRemoveMoneyNbr(amount int, nbr int) bool { // Permet de tester si le joueur a assez d'argent pour acheter plusieurs items
	if p.money >= amount*nbr {
		return true
	} else {
		return false
	}
}

func (p *Personnage) RemoveMoneyNbr(amount int, nbr int) { // Permet de retirer de l'argent de de plusieurs achats au joueur
	p.money -= amount * nbr
}

func (p *Personnage) BuySeveralItem(money int, nbr int, item string) { // Permet d'acheter plusieurs items tout en verifiant qu'il y est assez d'argent
	if p.TestAddInventoryNbr(nbr, item) == false {
		fmt.Println("Slot plein")
	} else if p.TestRemoveMoneyNbr(nbr, money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	} else {
		p.RemoveMoneyNbr(nbr, money)
		p.AddInventoryNbr(nbr, item)
		fmt.Println("Vous avez acheté : ", nbr, item)
	}
}
