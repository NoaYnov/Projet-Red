package red
import "fmt"

func (p *Personnage) AccessInventory() {
    for k, v := range p.inventaire {
        fmt.Printf("|%v:%v| ", k, v)
        
    }
    fmt.Printf("\n")
}



func (p *Personnage) AddInventory(item string) {
    if p.inventaire[item] <= 10 {
        p.inventaire[item] += 1
    } else {
        fmt.Println("Inventaire plein")
    }

}

func (p *Personnage) RemoveInventory(item string){
	p.inventaire[item] -= 1

}
