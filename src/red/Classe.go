package red

func (p *Personnage) Berserk() {
	p.inventaire = map[string]int{"potion de vie": 3, "Hache": 1}
	p.classe = "Berserk"

}

func (p *Personnage) Paladin() {
	p.inventaire = map[string]int{"potion de vie": 3, "epee": 1}
	p.classe = "Paladin"
}

func (p *Personnage) Mage() {
	p.inventaire = map[string]int{"potion de vie": 3, "baton de magie": 1, "potion de mana": 2}
	p.classe = "Mage"
}

func (p *Personnage) Voleur() {
	p.inventaire = map[string]int{"potion de vie": 3, "double dague": 1}
	p.classe = "Voleur"
}

func (p *Personnage) Archer() {
	p.inventaire = map[string]int{"potion de vie": 3, "arc": 1}
	p.classe = "Archer"
}