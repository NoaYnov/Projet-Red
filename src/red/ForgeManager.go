package red
import "fmt"
func (p *Personnage) Forge() {
	var result0 int
	fmt.Println("Bienvenue dans la forge")
	fmt.Printf("Vous avez")
	p.AccessInventory()
	p.AccessMoney()
	fmt.Println("Que voulez vous fabriquer ?")
	fmt.Println("1- Chapeau de l’aventurier : 10 pièces , 1 tissu, 2 cuir")
	fmt.Println("2- Bottes de l’aventurier : 8 pièces , 2 tissu, 1 cuir")
	fmt.Println("3- Gants de l’aventurier : 3 pièces , 3 tissu, 1 cuir")
	fmt.Println("4- Pantalon de l’aventurier : 10 pièces , 4 tissu, 4 cuir")
	fmt.Println("5- Plastron de l’aventurier : 25 pièces , 8 tissu, 6 cuir")
	fmt.Println("6- Retour")
	fmt.Scanln(&result0)
	switch result0 {
	case 1:
		p.CraftForgeHelmet(10,"tissu",1,"cuir",2,"Chapeau_de_l’aventurier")
		p.Menu()
	case 2:
		p.CraftForgeBoots(8,"tissu",2,"cuir",1,"Bottes_de_l’aventurier")
		p.Menu()
	case 3:
		p.CraftForgeGloves(3,"tissu",3,"cuir",1,"Gants_de_l’aventurier")
		p.Menu()
	case 4:
		p.CraftForgePants(10,"tissu",4,"cuir",4,"Pantalon_de_l’aventurier")
		p.Menu()
	case 5:
		p.CraftForgeChestplate(25,"tissu",8,"cuir",6,"Plastron_de_l’aventurier")
		p.Menu()

	case 6:
		p.Menu()

	default:
		fmt.Println("Vous n'avez pas choisi une option valide")
		p.Menu()

	}
}


func (p *Personnage) CraftForgeHelmet(money int,ressource1 string, nbr1 int, ressource2 string,nbr2 int,item string){
	if p.TestLenAddArmor(item) == false {
		fmt.Println("Inventaire d'armure plein")
	} else if p.TestItemAddArmor(item) == false {
		fmt.Println("Vous avez déjà fabriqué cet objet")
	} else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	}else if p.TestRemoveInventory(nbr1,ressource1) == false {
		fmt.Println("Vous n'avez pas assez de",ressource1)
	}else if p.TestRemoveInventory(nbr2,ressource2) == false {
		fmt.Println("Vous n'avez pas assez de",ressource2)
	}else{  p.RemoveMoney(money)
			p.RemoveInventoryNbr(nbr1,ressource1)
			p.RemoveInventoryNbr(nbr2,ressource2)
			p.AddHelmet(item)
			fmt.Println("Vous avez fabriqué : ",item)}

}

func (p *Personnage) CraftForgeBoots(money int,ressource1 string, nbr1 int, ressource2 string,nbr2 int,item string){
	if p.TestLenAddArmor(item) == false {
		fmt.Println("Inventaire d'armure plein")
	} else if p.TestItemAddArmor(item) == false {
		fmt.Println("Vous avez déjà fabriqué cet objet")
	}else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	}else if p.TestRemoveInventory(nbr1,ressource1) == false {
		fmt.Println("Vous n'avez pas assez de",ressource1)
	}else if p.TestRemoveInventory(nbr2,ressource2) == false {
		fmt.Println("Vous n'avez pas assez de",ressource2)
	}else{  p.RemoveMoney(money)
			p.RemoveInventoryNbr(nbr1,ressource1)
			p.RemoveInventoryNbr(nbr2,ressource2)
			p.AddBoots(item)
			fmt.Println("Vous avez fabriqué : ",item)}

}

func (p *Personnage) CraftForgeGloves(money int,ressource1 string, nbr1 int, ressource2 string,nbr2 int,item string){
	if p.TestLenAddArmor(item) == false {
		fmt.Println("Inventaire d'armure plein")
	} else if p.TestItemAddArmor(item) == false {
		fmt.Println("Vous avez déjà fabriqué cet objet")
	} else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	}else if p.TestRemoveInventory(nbr1,ressource1) == false {
		fmt.Println("Vous n'avez pas assez de",ressource1)
	}else if p.TestRemoveInventory(nbr2,ressource2) == false {
		fmt.Println("Vous n'avez pas assez de",ressource2)
	}else{  p.RemoveMoney(money)
			p.RemoveInventoryNbr(nbr1,ressource1)
			p.RemoveInventoryNbr(nbr2,ressource2)
			p.AddGloves(item)
			fmt.Println("Vous avez fabriqué : ",item)}

}

func (p *Personnage) CraftForgePants(money int,ressource1 string, nbr1 int, ressource2 string,nbr2 int,item string){
	if p.TestLenAddArmor(item) == false {
		fmt.Println("Inventaire d'armure plein")
	} else if p.TestItemAddArmor(item) == false {
		fmt.Println("Vous avez déjà fabriqué cet objet")
	} else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	}else if p.TestRemoveInventory(nbr1,ressource1) == false {
		fmt.Println("Vous n'avez pas assez de",ressource1)
	}else if p.TestRemoveInventory(nbr2,ressource2) == false {
		fmt.Println("Vous n'avez pas assez de",ressource2)
	}else{  p.RemoveMoney(money)
			p.RemoveInventoryNbr(nbr1,ressource1)
			p.RemoveInventoryNbr(nbr2,ressource2)
			p.AddPants(item)
			fmt.Println("Vous avez fabriqué : ",item)}

}

func (p *Personnage) CraftForgeChestplate(money int,ressource1 string, nbr1 int, ressource2 string,nbr2 int,item string){
	if p.TestLenAddArmor(item) == false {
		fmt.Println("Inventaire d'armure plein")
	} else if p.TestItemAddArmor(item) == false {
		fmt.Println("Vous avez déjà fabriqué cet objet")
	} else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	}else if p.TestRemoveInventory(nbr1,ressource1) == false {
		fmt.Println("Vous n'avez pas assez de",ressource1)
	}else if p.TestRemoveInventory(nbr2,ressource2) == false {
		fmt.Println("Vous n'avez pas assez de",ressource2)
	}else{  p.RemoveMoney(money)
			p.RemoveInventoryNbr(nbr1,ressource1)
			p.RemoveInventoryNbr(nbr2,ressource2)
			p.AddChestplate(item)
			fmt.Println("Vous avez fabriqué : ",item)}

}