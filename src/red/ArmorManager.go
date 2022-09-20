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

func (p *Personnage) AccessInventoryArmor() {
	for k, _:= range p.inventaireArmor {
		fmt.Println(k)
		
	}
	fmt.Printf("\n")
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

func (p *Personnage) DisplayHelmet()[]string {
	var result []string
	for k, v := range p.inventaireArmor {
		if v == "Helmet" {
			result = append(result, k)
			return result
		}
	}
	
}


