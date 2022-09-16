package red

func (p *Personnage) AddSkill(skill string) {
	for _, v := range p.skill {
		if v == skill { 
			println("Vous avez déjà ce skill")
			break
	
		}
	}
	p.skill = append(p.skill, skill)
}

func (p *Personnage) RemoveSkill(skill string) {
	for i, v := range p.skill {
		if v == skill {
			p.skill = append(p.skill[:i], p.skill[i+1:]...)
		}
	}
}

func (p *Personnage) DisplaySkill() {
	if len(p.skill) == 0 {
		p.AddSkill("Coup de poing")
		

	}
	for _, v := range p.skill {
		println(v)
	}
println("\n")
}