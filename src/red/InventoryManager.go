package red
import "fmt"

func (p *Personnage) AccessInventory() {
    for k, v := range p.inventaire {
        fmt.Printf("|%v:%v| ", k, v)
        
    }
    fmt.Printf("\n")
}



func (p *Personnage) AddInventory(item string) {
    p.inventaire[item] += 1
     
}

func (p *Personnage) RemoveInventory(item string){
	p.inventaire[item] -= 1

}

func (p *Personnage) TestAddInventory(item string)bool{
    if len(p.inventaire) >= 15 || p.inventaire[item] >= 10 {
        return false
    }else{
        return true
    }
}