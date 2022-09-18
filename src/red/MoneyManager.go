package red
import "fmt"
func (p *Personnage) AcessMoney() {
	println(p.money)

}



func (p *Personnage) AddMoney(amount int) {
	p.money += amount
}

func (p *Personnage) RemoveMoney(amount int) {
	p.money -= amount
}


func (p *Personnage) TestRemoveMoney(amount int)bool{
	if p.money >= amount {
		return true
	}else{
		return false
	}
}


func (p *Personnage) BuySkill(money int,skill string){
	if p.TestRemoveSkill(skill) == false {
		fmt.Println("Vous avez déjà ce skill")
	} else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	}else{p.RemoveMoney(money)
			p.AddSkill(skill)
			fmt.Println("Vous avez appris le",skill)}
}

func (p *Personnage) BuyItem(money int,item string){
	if p.TestAddInventory(item) == false {
		fmt.Println("Slot plein")
	} else if p.TestRemoveMoney(money) == false {
		fmt.Println("Vous n'avez pas assez d'argent")
	}else{p.RemoveMoney(money)
			p.AddInventory(item)
			fmt.Println("Vous avez acheté : ",item)}
}