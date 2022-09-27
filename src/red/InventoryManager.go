package red

import "fmt"

func (p *Personnage) AccessInventory() {
	for k, v := range p.inventaire {
		if v > 0 {
			delete(p.inventaire, k)
		}
		fmt.Printf("|%v:%v| ", k, v)

	}
	fmt.Printf("\n")
}

func (p *Personnage) AddInventory(item string) {
	p.inventaire[item] += 1

}

func (p *Personnage) AddInventoryNbr(nbr int, item string) {
	p.inventaire[item] += nbr
}

func (p *Personnage) RemoveInventory(item string) {
	p.inventaire[item] -= 1

}

func (p *Personnage) RemoveInventoryNbr(nbr int, item string) {
	p.inventaire[item] -= nbr
}

func (p *Personnage) TestAddInventory(item string) bool {
	if len(p.inventaire) >= p.liminventaire || p.inventaire[item] >= p.limslot {
		return false
	} else {
		return true
	}
}

func (p *Personnage) TestRemoveInventory(nbr int, item string) bool {
	if p.inventaire[item] >= nbr {
		return true
	} else {
		return false
	}
}

func (p *Personnage) TestAddInventoryNbr(nbr int, item string) bool {
	if len(p.inventaire) >= p.liminventaire || p.inventaire[item] > p.limslot-nbr {
		return false
	} else {
		return true
	}
}
