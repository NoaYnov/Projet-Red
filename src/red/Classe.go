package red

func (p *Personnage) Berserk() { // fonction qui permet de créer un personnage de classe Berserk
	p.inventaire = map[string]int{"potion de vie": 9, "Hache": 1}
	p.classe = "Berserk"

}

func (p *Personnage) Paladin() { // fonction qui permet de créer un personnage de classe Paladin
	p.inventaire = map[string]int{"potion de vie": 3, "lance": 1}
	p.classe = "Paladin"
}

func (p *Personnage) Mage() { // fonction qui permet de créer un personnage de classe Mage
	p.inventaire = map[string]int{"potion de vie": 3, "baton de magie": 1, "potion de mana": 2}
	p.classe = "Mage"
}

func (p *Personnage) Voleur() { // fonction qui permet de créer un personnage de classe Voleur
	p.inventaire = map[string]int{"potion de vie": 3, "double dague": 1}
	p.classe = "Voleur"
}

func (p *Personnage) Archer() { // fonction qui permet de créer un personnage de classe Archer
	p.inventaire = map[string]int{"potion de vie": 3, "arc": 1}
	p.classe = "Archer"
}
