package red

import "fmt"

func (p *Personnage) Wasted() { //Fonction de mort du Joueur
	if p.point_de_vie_actuel <= 0 {
		fmt.Println("Vous êtes mort")
		fmt.Println("Vous retournez au menu principal avec la moitier de vos point de vie max")
		p.point_de_vie_max = p.point_de_vie_max / 2
		p.Menu()
	}
}
