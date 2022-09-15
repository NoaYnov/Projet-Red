package red
import "fmt"
func (p Personnage) TakePotHeal(){
	fmt.Println(p.inventaire["potion de vie"])
	if p.inventaire["potion de vie"] >= 1{

		p.RemoveInventory("potion de vie")

		if p.points_de_vie_actuel < p.points_de_vie_max{
			p.points_de_vie_actuel += 25
			if p.points_de_vie_actuel > p.points_de_vie_max {
				p.points_de_vie_actuel = p.points_de_vie_max
		
				
			}
		} else{fmt.Println("Tu es deja Full vie")}
	}


	

}

