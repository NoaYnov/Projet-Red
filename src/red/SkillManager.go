package red

import "fmt"

func (p *Personnage) AddSkill(skill string, degats int) {
	if p.TestAddSkill(skill) == true {
		p.skill[skill] = degats
	}
}


func (p *Personnage) DisplaySkill() {
	for k, v := range p.skill {
		println(k,":", v)
	}
}

func (p *Personnage) TestAddSkill(skill string) bool {
	for v, _ := range p.skill {
		if v == skill {
			return false
		}
	}
	return true
}



func (p *Personnage) TestRemoveSkill(skill string) bool {
	for v, _:= range p.skill {
		if v == skill {
			return true
		}
	}
	return false
}

func (p *Personnage) UseSkill(skill string) {
	if p.TestRemoveSkill(skill) == true {
		p.goblin.point_de_vie_actuel -= p.skill[skill]
		if p.goblin.point_de_vie_actuel <= 0 {
			p.goblin.point_de_vie_actuel = 0
		}
		fmt.Println("Vous avez utilisé ", skill)
		fmt.Println("Vous avez infligé", p.skill[skill], "points de degats")
		fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
	} else {
		fmt.Println("Vous n'avez pas ce sort")
		p.Playerturn()
	}
}
