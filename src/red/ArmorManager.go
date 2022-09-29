package red

import "fmt"

func (p *Personnage) AccessAllArmor() { // Fonction qui permet d'accéder à toutes les armures

	fmt.Println("Bienvenue dans le menu des équipements. Ils vous permettent d'améliorer vos statistiques de personnage")
	fmt.Println("Voici les armures que vous possédez :")
	fmt.Println("Voici votre inventaire de casque : ", p.armor.helmet)
	fmt.Println("Voici votre inventaire de bottes : ", p.armor.boots)
	fmt.Println("Voici votre inventaire de gants : ", p.armor.gloves)
	fmt.Println("Voici votre inventaire de jambières : ", p.armor.pants)
	fmt.Println("Voici votre inventaire de plastron : ", p.armor.chestplate)
}

func (p *Personnage) AddHelmet(res string) { // Fonction qui permet d'ajouter un casque à la liste des armures
	p.armor.helmet = append(p.armor.helmet, res)
}

func (p *Personnage) AddBoots(res string) { // Fonction qui permet d'ajouter des bottes à la liste des armures
	p.armor.boots = append(p.armor.boots, res)
}

func (p *Personnage) AddGloves(res string) { // Fonction qui permet d'ajouter des gants à la liste des armures
	p.armor.gloves = append(p.armor.gloves, res)
}

func (p *Personnage) AddPants(res string) { // Fonction qui permet d'ajouter des jambières à la liste des armures
	p.armor.pants = append(p.armor.pants, res)
}

func (p *Personnage) AddChestplate(res string) { // Fonction qui permet d'ajouter une plastron à la liste des armures
	p.armor.chestplate = append(p.armor.chestplate, res)
}

func (p *Personnage) TestLenAddArmor(item string) bool { // Fonction qui permet de vérifier si l'inventaire d'armure n'est pas plein
	if len(p.armor.helmet) >= 5 || len(p.armor.boots) >= 5 || len(p.armor.gloves) >= 5 || len(p.armor.pants) >= 5 || len(p.armor.chestplate) >= 5 {
		return false
	} else {
		return true
	}
}

func (p *Personnage) TestItemAddArmor(item string) bool { // Fonction qui permet de vérifier si l'armure n'est pas déjà dans l'inventaire d'armure
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

func (p *Personnage) DisplayEquipement() { // Fonction qui permet d'afficher les armures équipées
	fmt.Println("Voici vos équipements :")
	if len(p.equiphelmet) == 0 {
		fmt.Println("Vous avez rien équpé comme casque")
	} else {
		fmt.Println("Vous avez ", p.equiphelmet, " comme casque")
	}
	if len(p.equipboots) == 0 {
		fmt.Println("Vous avez rien équpé comme bottes")
	} else {
		fmt.Println("Vous avez ", p.equipboots, " comme bottes")
	}
	if len(p.equipgloves) == 0 {
		fmt.Println("Vous avez rien équpé comme gants")
	} else {
		fmt.Println("Vous avez ", p.equipgloves, " comme gants")
	}
	if len(p.equippants) == 0 {
		fmt.Println("Vous avez rien équpé comme jambières")
	} else {
		fmt.Println("Vous avez ", p.equippants, " comme jambières")
	}
	if len(p.equipchestplate) == 0 {
		fmt.Println("Vous avez rien équpé comme plastron")
	} else {
		fmt.Println("Vous avez ", p.equipchestplate, " comme plastron")
	}
}
