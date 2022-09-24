package red

import "fmt"

func (p *Personnage) AccessAllArmor() {

	fmt.Println("Bienvenue dans le menu des équipements. Ils vous permettent d'améliorer vos statistiques de personnage")
	fmt.Println("Voici les armures que vous possédez :")
	fmt.Println(p.armor.helmet)
	fmt.Println(p.armor.boots)
	fmt.Println(p.armor.gloves)
	fmt.Println(p.armor.pants)
	fmt.Println(p.armor.chestplate)
}

func (p *Personnage) AddHelmet(res string) {
	p.armor.helmet = append(p.armor.helmet, res)
}

func (p *Personnage) AddBoots(res string) {
	p.armor.boots = append(p.armor.boots, res)
}

func (p *Personnage) AddGloves(res string) {
	p.armor.gloves = append(p.armor.gloves, res)
}

func (p *Personnage) AddPants(res string) {
	p.armor.pants = append(p.armor.pants, res)
}

func (p *Personnage) AddChestplate(res string) {
	p.armor.chestplate = append(p.armor.chestplate, res)
}

func (p *Personnage) TestLenAddArmor(item string) bool {
	if len(p.armor.helmet) >= 5 || len(p.armor.boots) >= 5 || len(p.armor.gloves) >= 5 || len(p.armor.pants) >= 5 || len(p.armor.chestplate) >= 5 {
		return false
	} else {
		return true
	}
}

func (p *Personnage) TestItemAddArmor(item string) bool {
	for _, v := range p.armor.helmet {
		if v == item {
			return false
		}
	}
	for _, v := range p.armor.boots {
		if v == item {
			return false
		}
	}
	for _, v := range p.armor.gloves {
		if v == item {
			return false
		}
	}
	for _, v := range p.armor.pants {
		if v == item {
			return false
		}
	}
	for _, v := range p.armor.chestplate {
		if v == item {
			return false
		}
	}
	return true
}
