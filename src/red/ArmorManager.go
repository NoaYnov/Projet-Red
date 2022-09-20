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
		fmt.Println(v)
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


func (p *Personnage) EquipHelmet(item string){
	var choix string
	fmt.Println("Quel item voulez vous equiper?")
	fmt.Scan(&choix)
	if p.helmet != choix{
		fmt.Println("Vous avez choisi de vous equiper de", choix)
		p.inventaireArmor[choix] = "Equipped"
	}else{
		fmt.Println("Cet item et deja equipé")
	}
}