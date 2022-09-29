package red

import "fmt"

func (p *Personnage) MenuEquipement() { // Menu des equipements
	p.AccessAllArmor()
	fmt.Println("________________________________________________")
	p.DisplayEquipement()
	var ts int
	fmt.Println("1- Equiper un casque")
	fmt.Println("2. Equiper des bottes")
	fmt.Println("3. Equiper des gants")
	fmt.Println("4. Equiper un pantalon")
	fmt.Println("5. Equiper un plastron")
	fmt.Printf("\n")
	fmt.Println("6- Retour")
	fmt.Scanln(&ts)
	switch ts {
	case 1:
		p.EquipHelmet()
	case 2:
		p.EquipBoots()
	case 3:
		p.EquipGloves()
	case 4:
		p.EquipPants()
	case 5:
		p.EquipChest()
	case 6:
		p.Menu()
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Aucun choix associe")
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()
	}
}

func (p *Personnage) EquipBoots() { // Fonction d'equipement de bottes
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.boots)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("écrivez le nom de l'objet")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "Chaussures":
		p.equipboots = p.armor.boots[0]
		fmt.Println("Vous avez équipé", p.equipboots)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()
	case "Bottes":
		p.equipboots = p.armor.boots[1]
		fmt.Println("Vous avez équipé", p.equipboots)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.TestEquipement()
		p.MenuEquipement()
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()
	}
	p.TestEquipement()
}

func (p *Personnage) EquipHelmet() { // Fonction d'equipement de casque
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.helmet)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("écrivez le nom de l'objet")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "Bonnet":
		p.equiphelmet = p.armor.helmet[0]
		fmt.Println("Vous avez équipé", p.equiphelmet)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()
	case "Casque":
		p.equiphelmet = p.armor.helmet[1]
		fmt.Println("Vous avez équipé", p.equiphelmet)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.TestEquipement()
		p.MenuEquipement()
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()
	}
	p.TestEquipement()
}

func (p *Personnage) EquipGloves() { // Fonction d'equipement de gants
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.gloves)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("écrivez le nom de l'objet")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "Gants":
		p.equipgloves = p.armor.gloves[0]
		fmt.Println("Vous avez équipé", p.equipgloves)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()
	case "Gantlets":
		p.equipgloves = p.armor.gloves[1]
		fmt.Println("Vous avez équipé", p.equipgloves)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.TestEquipement()
		p.MenuEquipement()
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()
	}
	p.TestEquipement()
}

func (p *Personnage) EquipPants() { // Fonction d'equipement de jambieres
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.pants)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("écrivez le nom de l'objet")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "Pantalon":
		p.equippants = p.armor.pants[0]
		fmt.Println("Vous avez équipé", p.equippants)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()
	case "Jambieres":
		p.equippants = p.armor.pants[1]
		fmt.Println("Vous avez équipé", p.equippants)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.TestEquipement()
		p.MenuEquipement()
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()
	}
	p.TestEquipement()
}

func (p *Personnage) EquipChest() { // Fonction d'equipement de plastron
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.chestplate)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("écrivez le nom de l'objet")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "Tunique":
		p.equipchestplate = p.armor.chestplate[0]
		fmt.Println("Vous avez équipé", p.equipchestplate)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()
	case "Plastron":
		p.equipchestplate = p.armor.chestplate[1]
		fmt.Println("Vous avez équipé", p.equipchestplate)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.TestEquipement()
		p.MenuEquipement()
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()
	}
}

func (p *Personnage) TestEquipement() { //Test si une piece d'armure est équipé pour appliqué ces effets
	if p.equipboots == "Bottes" {
		p.resistance_physique += 2
	}
	if p.equipchestplate == "Plastron" {
		p.point_de_vie_max += 8
	}
	if p.equiphelmet == "Casque" {
		p.point_de_vie_max += 3
	}
	if p.equippants == "Jambieres" {
		p.point_de_vie_max += 6
	}
	if p.equipgloves == "Gantlets" {
		p.force += 1
	}
}
