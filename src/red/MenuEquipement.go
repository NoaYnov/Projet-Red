package red
import "fmt"

func (p *Personnage) MenuEquipement(){
	
	p.AccessAllArmor()
	
	var ts int
	fmt.Println("1- Equiper un casque")
	fmt.Println("2. Equiper des bottes")
	fmt.Println("3. Equiper des gants")
	fmt.Println("4. Equiper un pantalon")
	fmt.Println("5. Equiper un plastron")
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

func (p *Personnage) EquipBoots(){
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.boots)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "bottes":
		p.equipboots = p.armor.boots[0]
		fmt.Println("Vous avez équipé", p.equipboots)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

	case "Bottes_de_l’aventurier":
		p.equipboots = p.armor.boots[1]
		fmt.Println("Vous avez équipé", p.equipboots)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

		
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()

	}
}
func (p *Personnage) EquipHelmet(){
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.helmet)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "bonnet":
		p.equiphelmet = p.armor.helmet[0]
		fmt.Println("Vous avez équipé", p.equiphelmet)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

	case "Chapeau_de_l’aventurier":
		p.equiphelmet = p.armor.helmet[1]
		fmt.Println("Vous avez équipé", p.equiphelmet)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

		
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()

	}
}

func (p *Personnage) EquipGloves(){
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.gloves)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "gants":
		p.equipgloves = p.armor.gloves[0]
		fmt.Println("Vous avez équipé", p.equipgloves)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

	case "Gants_de_l’aventurier":
		p.equipgloves = p.armor.gloves[1]
		fmt.Println("Vous avez équipé", p.equipgloves)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

		
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()

	}
}

func (p *Personnage) EquipPants(){
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.pants)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "pantalon":
		p.equippants = p.armor.pants[0]
		fmt.Println("Vous avez équipé", p.equippants)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

	case "Pantalon_de_l’aventurier":
		p.equippants = p.armor.pants[1]
		fmt.Println("Vous avez équipé", p.equippants)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

		
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()

	}
}

func (p *Personnage) EquipChest(){
	fmt.Println("________________________________________________")
	fmt.Println(p.armor.chestplate)
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Retour")
	fmt.Println("Que voulez vous equiper ?")
	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	var ts2 string
	fmt.Scanln(&ts2)
	switch ts2 {
	case "retour":
		p.MenuEquipement()
	case "tunique":
		p.equipchestplate = p.armor.chestplate[0]
		fmt.Println("Vous avez équipé", p.equipchestplate)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

	case "Plastron_de_l’aventurier":
		p.equipchestplate = p.armor.chestplate[1]
		fmt.Println("Vous avez équipé", p.equipchestplate)
		fmt.Println("________________________________________________")
		fmt.Printf("\n")
		p.MenuEquipement()

		
	default:
		fmt.Println("________________________________________________")
		fmt.Println("Vous n'avez pas d'autre equipement à équiper")
		p.MenuEquipement()

	}
}