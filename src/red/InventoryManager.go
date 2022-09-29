package red

import "fmt"

func (p *Personnage) AccessInventory() { //Affiche l'inventaire du personnage
	for k, v := range p.inventaire {
		fmt.Printf("|%v:%v| ", k, v)
	}
	fmt.Printf("\n")
}

func (p *Personnage) AddInventory(item string) { //Ajoute un item à l'inventaire du personnage
	p.inventaire[item] += 1
}

func (p *Personnage) AddInventoryNbr(nbr int, item string) { //Ajoute plusieurs items à l'inventaire du personnage
	p.inventaire[item] += nbr
}

func (p *Personnage) RemoveInventory(item string) { //Retire un item de l'inventaire du personnage
	p.inventaire[item] -= 1
}

func (p *Personnage) RemoveInventoryNbr(nbr int, item string) { //Retire plusieurs items de l'inventaire du personnage
	p.inventaire[item] -= nbr
}

func (p *Personnage) TestAddInventory(item string) bool { //Test si l'inventaire du personnage ou si ces slots sont plein
	if len(p.inventaire) >= p.liminventaire || p.inventaire[item] >= p.limslot {
		return false
	} else {
		return true
	}
}

func (p *Personnage) TestRemoveInventory(nbr int, item string) bool { //Test si le personnage a assez d'items pour les retirer
	if p.inventaire[item] >= nbr {
		return true
	} else {
		return false
	}
}

func (p *Personnage) TestAddInventoryNbr(nbr int, item string) bool { //Test si l'inventaire du personnage ou si ces slots sont plein pour plusieurs items
	if len(p.inventaire) >= p.liminventaire || p.inventaire[item] > p.limslot-nbr {
		return false
	} else {
		return true
	}
}

func (p *Personnage) TestIsInInventory(item string) bool { //Test si le personnage a un item dans son inventaire
	for k := range p.inventaire {
		if k == item {
			return true
		}
	}
	return false
}
