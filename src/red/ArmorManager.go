package red

import "fmt"





func (p *Personnage) AddHelmet (item string) {
	p.inventaireArmor[item] = "Helmet"
}

func (p *Personnage) AddBoots (item string) {
	p.inventaireArmor[item] = "Boots"
}

func (p *Personnage) AddGloves (item string) {
	p.inventaireArmor[item] = "Gloves"
}

func (p *Personnage) AddPants (item string) {
	p.inventaireArmor[item] = "Pants"
}

func (p *Personnage) AddChestplate(item string) {
	p.inventaireArmor[item] = "Chestplate"
}

func (p *Personnage) DisplayArmor() {

	p.DisplayHelmet()
	p.DisplayBoots()
	p.DisplayGloves()
	p.DisplayPants()
	p.DisplayChestplate()
}

func (p *Personnage) TestLenAddArmor(item string)bool{
	if len(p.inventaireArmor) >= 25{
		return false
	}
	return true
}
func (p *Personnage) TestItemAddArmor(item string)bool{
	for v, _ := range p.inventaireArmor {
		if v == item {
			return false
		}
	}

	return true
}

func (p *Personnage) DisplayHelmet(){
	var result []string
	for k, v := range p.inventaireArmor {
		if v == "Helmet" {
			result = append(result, k)
			
		}
	}
	if len(result) != 0{
		fmt.Println(result)
	}else{
		fmt.Println("Vous n'avez pas de casque")
	}

}

func (p *Personnage) DisplayBoots(){
	var result []string
	for k, v := range p.inventaireArmor {
		if v == "Boots" {
			result = append(result, k)
			
		}
	}
	if len(result) != 0{
		fmt.Println(result)
	}else{
		fmt.Println("Pas de bottes")
	}
	

}

func (p *Personnage) DisplayGloves(){
	var result []string
	for k, v := range p.inventaireArmor {
		if v == "Gloves" {
			result = append(result, k)
			
		}
	}
	if len(result) != 0{
		fmt.Println(result)
	}else{
		fmt.Println("Pas de gants")
	}

}

func (p *Personnage) DisplayPants(){
	var result []string
	for k, v := range p.inventaireArmor {
		if v == "Pants" {
			result = append(result, k)
			
		}
	}
	if len(result) != 0{
		fmt.Println(result)
	}else{
		fmt.Println("Pas de pantalon")
	}
}

func (p *Personnage) DisplayChestplate(){
	var result []string
	for k, v := range p.inventaireArmor {
		if v == "Chestplate" {
			result = append(result,k)
			
		}
	}
	if len(result) != 0{
		fmt.Println(result)
	}else{
		fmt.Println("Pas de plastron")
	}
}


func (p *Personnage) EquipHelmet(choix int){

	switch choix {
	case 1:
		p.helmet = "Bonnet"
		fmt.Println("Vous avez choisi de vous equiper d'un bonnet")
	case 2:
		if p.TestItemAddArmor("Chapeau de l'aventurier") == false{
		p.helmet = "Chapeau de l'aventurier"
		fmt.Println("Vous avez choisi de vous equiper d'un chapeau de l'aventurier")
		}else{
			fmt.Println("Vous n'avez pas cet item")
		}
	
	}
}

func (p *Personnage) EquipBoots(choix int){
	
	switch choix {
	case 1:
		p.boots = "Bottes de cuir"
		fmt.Println("Vous avez choisi de vous equiper de bottes de cuir")
	case 2:
		if p.TestItemAddArmor("Bottes de l'aventurier") == false{
		p.boots = "Bottes de l'aventurier"
		fmt.Println("Vous avez choisi de vous equiper de bottes de l'aventurier")
		}else{
			fmt.Println("Vous n'avez pas cet item")
		}
	
	}
}

func (p *Personnage) EquipGloves(choix int){
	
	switch choix {
	case 1:
		p.gloves = "Gants de cuir"
		fmt.Println("Vous avez choisi de vous equiper de gants de cuir")
	case 2:
		if p.TestItemAddArmor("Gants de l'aventurier") == false{
		p.gloves = "Gants de l'aventurier"
		fmt.Println("Vous avez choisi de vous equiper de gants de l'aventurier")
		}else{
			fmt.Println("Vous n'avez pas cet item")
		}
	
	}
}

func (p *Personnage) EquipPants(choix int){
	
	switch choix {
	case 1:
		p.pants = "Pantalon de cuir"
		fmt.Println("Vous avez choisi de vous equiper de pantalon de cuir")
	case 2:
		if p.TestItemAddArmor("Pantalon de l'aventurier") == false{
		p.pants = "Pantalon de l'aventurier"
		fmt.Println("Vous avez choisi de vous equiper de pantalon de l'aventurier")
		}else{
			fmt.Println("Vous n'avez pas cet item")
		}
	
	}
}

func (p *Personnage) EquipChestplate(choix int){
	
	switch choix {
	case 1:
		p.chestplate = "Plastron de cuir"
		fmt.Println("Vous avez choisi de vous equiper de plastron de cuir")
	case 2:
		if p.TestItemAddArmor("Plastron de l'aventurier") == false{
		p.chestplate = "Plastron de l'aventurier"
		fmt.Println("Vous avez choisi de vous equiper de plastron de l'aventurier")
		}else{
			fmt.Println("Vous n'avez pas cet item")
		}
	
	}
}

