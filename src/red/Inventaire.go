package red
import "fmt"

func (p Personnage) AccessInventory() {
	fmt.Println("Inventaire:", p.inventaire)
}



func (p *Personnage) AddInventory(item string){
	p.inventaire[item] += 1


}

func (p *Personnage) RemoveInventory(item string){
	p.inventaire[item] -= 1

}