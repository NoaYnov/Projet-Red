package red
import "fmt"
func (p *Personnage) MenuEquipement() {
	var choix int
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
		fmt.Println(p.helmet, "est deja equipé")

	case 2:
		p.DisplayBoots()
	case 3:
		p.DisplayGloves()
	case 4:
		p.DisplayPants()
	case 5:
		p.DisplayChestplate()
	case 6:
		fmt.Println("________________________________________________")
		fmt.Println("Vous avez choisi de retourner au menu principal")
		p.Menu()

	default:
		fmt.Println("Aucun choix associe")
		p.Menu()
	}
	
}