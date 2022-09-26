package red
import "time"
import "fmt"

func (p *Personnage) SelfPoisonPot() {
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

func (p *Personnage) PoisonPot() {
    fmt.Println(p.inventaire["potion de poison"])
    if p.inventaire["potion de poison"] >= 1 {
        p.RemoveInventory("potion de poison")
        interval := time.Second * 1
        for i := 0; i < 3; i++ {
            p.goblin.point_de_vie_actuel -= 10
            time.Sleep(interval)
        }
        fmt.Println("Le gobelin a perdu 10 points de vie")
        fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
    }else{
        fmt.Println("Vous n'avez pas de potion de poison")
        p.Playerturn()}
}