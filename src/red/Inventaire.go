package red
import "fmt"

func (p *Personnage) AccessInventory() {
    for k, v := range p.inventaire {
        fmt.Printf("|%v:%v| ", k, v)
        
    }
    fmt.Printf("\n")
}



func (p *Personnage) AddInventory(item string) {
    
    if len(p.inventaire) == 15 {
         fmt.Println("Inventaire plein")
         
    } else {for k, _ := range p.inventaire {
        if p.inventaire[k] == 10 {
          fmt.Println("Slot plein")
          break
            
        } else {p.inventaire[item] += 1
        break}
        }
       
    }
     

    


}

func (p *Personnage) RemoveInventory(item string){
	p.inventaire[item] -= 1

}
