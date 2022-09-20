package red
import "fmt"
func (p *Personnage) MenuEquipement() {
	var choix int
	for {
		p.AccessInventoryArmor()
		p.AccessArmor()
		fmt.Println("1. Equip Helmet")
		fmt.Println("2. Equip Boots")
		fmt.Println("3. Equip Gloves")
		fmt.Println("4. Equip Pants")
		fmt.Println("5. Equip Chestplate")
		fmt.Println("11. Quit")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			p.EquipHelmet()
		// case 2:
		// 	p.EquipBoots()
		// case 3:
		// 	p.EquipGloves()
		// case 4:
		// 	p.EquipPants()
		// case 5:
		// 	p.EquipChestplate()
		// case 6:
		// 	p.UnequipHelmet()
		// case 7:
		// 	p.UnequipBoots()
		// case 8:
		// 	p.UnequipGloves()
		// case 9:
		// 	p.UnequipPants()
		// case 10:
		// 	p.UnequipChestplate()
		// case 11:
		// 	return
		}
	}
}