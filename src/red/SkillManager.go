package red

import "fmt"

func (p *Personnage) AddSkill(skill string, degats int) { //Ajoute un sort au personnage
	if p.TestAddSkill(skill) == true {
		p.skill[skill] = degats
	}
}

func (p *Personnage) DisplaySkill() { //Affiche les sorts du personnage
	for k, v := range p.skill {
		println(k, ":", v)
	}
}

func (p *Personnage) TestAddSkill(skill string) bool { //Test si le sort existe deja
	for v := range p.skill {
		if v == skill {
			return false
		}
	}
	return true
}

func (p *Personnage) TestRemoveSkill(skill string) bool { //Test si le sort est deja dans l'inventaire
	for v := range p.skill {
		if v == skill {
			return true
		}
	}
	return false
}

func (p *Personnage) UseSkill(skill string, nbr int) { //Utilise un sort en prevoyant le mana necessaire
	if nbr == 1 {
		if p.TestRemoveSkill(skill) == true {
			p.goblin.point_de_vie_actuel -= p.skill[skill] * (1 - p.goblin.resistance_magique/100)
			if p.goblin.point_de_vie_actuel <= 0 {
				p.goblin.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé ", skill)
			fmt.Println("Vous avez infligé", p.skill[skill]*(1-p.goblin.resistance_magique/100), "points de degats")
			fmt.Println("PV du gobelin:", p.goblin.point_de_vie_actuel, "/", p.goblin.point_de_vie_max)
		} else {
			fmt.Println("Vous n'avez pas ce sort")
			p.PlayerTurn(1)
		}
	}
	if nbr == 2 {
		if p.TestRemoveSkill(skill) == true {
			p.versgeant.point_de_vie_actuel -= p.skill[skill] * (1 - p.versgeant.resistance_magique/100)
			if p.versgeant.point_de_vie_actuel <= 0 {
				p.versgeant.point_de_vie_actuel = 0
			}
			fmt.Println("Vous avez utilisé ", skill)
			fmt.Println("Vous avez infligé", p.skill[skill]*(1-p.versgeant.resistance_magique/100), "points de degats")
			fmt.Println("PV du versgeant:", p.versgeant.point_de_vie_actuel, "/", p.versgeant.point_de_vie_max)
		} else {
			fmt.Println("Vous n'avez pas ce sort")
			p.PlayerTurn(2)
		}
	}
}
