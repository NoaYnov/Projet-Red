package red
import "time"
import "fmt"

func (p *Personnage) PoisonPot() {
    fmt.Println(p.inventaire["potion de poison"])
    if p.inventaire["potion de poison"] >= 1 {
        p.RemoveInventory("potion de poison")
        interval := time.Second * 1
        for i := 0; i < 3; i++ {
            p.point_de_vie_actuel -= 10
            time.Sleep(interval)
        }
    }
}