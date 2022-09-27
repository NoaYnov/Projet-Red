package red

func (p *Personnage) Wasted() {
	if p.point_de_vie_actuel <= 0 {
		println("Vous es mort")
		p.point_de_vie_actuel = p.point_de_vie_max / 2
	}
}
