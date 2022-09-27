package red

import "fmt"

func (p *Personnage) AddSkill(skill string, degats int) {
	if p.TestAddSkill(skill) == true {
		p.skill[skill] = degats
	}
}

func (p *Personnage) DisplaySkill() {
	for k, v := range p.skill {
		println(k, ":", v)
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
	for v, _ := range p.skill {
		if v == skill {
			return true
		}
	}
	return false
}

func (p *Personnage) UseSkill(skill string, nbr int) {
	if nbr == 1 {
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
			p.PlayerTurn(1)
		}
	}
	if nbr == 2 {
		if p.TestRemoveSkill(skill) == true {
			p.versgeant.point_de_vie_actuel -= p.skill[skill]
			if p.versgeant.point_de_vie_actuel <= 0 {
				p.versgeant.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé ", skill)
			fmt.Println("Vous avez infligé", p.skill[skill], "points de degats")
			fmt.Println("PV du versgeant:", p.versgeant.point_de_vie_actuel, "/", p.versgeant.point_de_vie_max)
		} else {
			fmt.Println("Vous n'avez pas ce sort")
			p.PlayerTurn(2)
		}

	}
	if nbr == 3 {
		if p.TestRemoveSkill(skill) == true {
			p.troll.point_de_vie_actuel -= p.skill[skill]
			if p.troll.point_de_vie_actuel <= 0 {
				p.troll.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé ", skill)
			fmt.Println("Vous avez infligé", p.skill[skill], "points de degats")
			fmt.Println("PV du troll:", p.troll.point_de_vie_actuel, "/", p.troll.point_de_vie_max)
		} else {
			fmt.Println("Vous n'avez pas ce sort")
			p.PlayerTurn(3)
		}

	}
	if nbr == 4 {
		if p.TestRemoveSkill(skill) == true {
			p.gorgone.point_de_vie_actuel -= p.skill[skill]
			if p.gorgone.point_de_vie_actuel <= 0 {
				p.gorgone.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé ", skill)
			fmt.Println("Vous avez infligé", p.skill[skill], "points de degats")
			fmt.Println("PV de la gorgone:", p.gorgone.point_de_vie_actuel, "/", p.gorgone.point_de_vie_max)
		} else {
			fmt.Println("Vous n'avez pas ce sort")
			p.PlayerTurn(4)
		}

	}
}
