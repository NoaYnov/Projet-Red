package red

import "fmt"

func (p *Personnage) UpNiveau() {
	if p.experience_actuel == p.experience_max {
		p.experience_actuel = 0
		p.experience_max *= 2
		p.niveau += 1
		p.force += 2
		p.point_de_vie_max += 5
		p.mana_max += 5
		fmt.Println("Vous avez gagné un niveau !")
		p.DisplayNiveau()
		p.DisplayExperience()

	} else if p.experience_actuel > p.experience_max {
		p.experience_actuel -= p.experience_max
		p.experience_max *= 2
		p.niveau += 1
		p.force += 2
		p.point_de_vie_max += 5
		p.mana_max += 5
		fmt.Println("Vous avez gagné un niveau !")
		p.DisplayNiveau()
		p.DisplayExperience()
	}
}

func (p *Personnage) DisplayNiveau() {
	fmt.Println("Vous êtes niveau", p.niveau)
}

func (p *Personnage) DisplayExperience() {
	fmt.Println("Vous avez", p.experience_actuel, "points d'expérience")
}
