package red
import "fmt"
func (p *Personnage) MenuEquipement() {
	var choix int
	var choixp int
	p.DisplayArmor()
	fmt.Println("1- Equiper un casque")
	fmt.Println("2. Equiper des bottes")
	fmt.Println("3. Equiper des gants")
	fmt.Println("4. Equiper un pantalon")
	fmt.Println("5. Equiper un plastron")
	fmt.Println("6- Retour")
	fmt.Scan(&choix)
	switch choix {
	case 1:

		p.DisplayHelmet()
		fmt.Println("________________________________________________")
		fmt.Println("vous avez choisi de vous equiper d'un casque")
		fmt.Println("Quel item voulez vous equiper?")
		fmt.Scan(&choixp)
		p.EquipHelmet(choixp)
		p.Menu()

	case 2:
		p.DisplayBoots()
		fmt.Println("________________________________________________")
		fmt.Println("vous avez choisi de vous equiper de bottes")
		fmt.Println("Quel item voulez vous equiper?")
		fmt.Scan(&choixp)
		p.EquipBoots(choixp)
		p.Menu()
	case 3:
		p.DisplayGloves()
		fmt.Println("________________________________________________")
		fmt.Println("vous avez choisi de vous equiper de gants")
		fmt.Println("Quel item voulez vous equiper?")
		fmt.Scan(&choixp)
		p.EquipGloves(choixp)
		p.Menu()
	case 4:
		p.DisplayPants()
		fmt.Println("________________________________________________")
		fmt.Println("vous avez choisi de vous equiper de pantalon")
		fmt.Println("Quel item voulez vous equiper?")
		fmt.Scan(&choixp)
		p.EquipPants(choixp)
		p.Menu()
	case 5:
		p.DisplayChestplate()
		fmt.Println("________________________________________________")
		fmt.Println("vous avez choisi de vous equiper d'un plastron")
		fmt.Println("Quel item voulez vous equiper?")
		fmt.Scan(&choixp)
		p.EquipChestplate(choixp)
		p.Menu()
	case 6:
		fmt.Println("________________________________________________")
		fmt.Println("Vous avez choisi de retourner au menu principal")
		p.Menu()

	default:
		fmt.Println("Aucun choix associe")
		p.Menu()
	}
	
}