package red

import "fmt"

func SeDeplacer() {
	var direction int
	var choix int
	fmt.Println("Ou voulez vous aller ?")
	fmt.Println("1- Les marais de la mort")
	fmt.Println("2- les Montagnes désolées")
	fmt.Println("3- les plaines arides")
	fmt.Scanln(&direction)
	switch direction {
	case 1:
		fmt.Println("Vous vous trouvez dans les marais de la mort")
		fmt.Println("Que voulez vous faire?")
		fmt.Println("1. Se battre")
		fmt.Println("2. Parler à Fred")
		fmt.Println("3-retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("Vous combattez un vers géant")
		case 2:
			fmt.Print("Vous parlez à Fred")
		case 3:
			SeDeplacer()
		}

	case 2:
		fmt.Println("Vous vous trouvez dans les Montagnes désolées")
		fmt.Println("1. Se battre")
		fmt.Println("2. Parler à Bruno")
		fmt.Println("3-retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("Vous combattez un troll")
		case 2:
			fmt.Print("Vous parlez à Bruno")
		case 3:
			SeDeplacer()
		}

	case 3:
		fmt.Println("Vous vous trouvez dans les plaines arides")
		fmt.Println("1. Se battre")
		fmt.Println("2. Parler à Iris")
		fmt.Println("3-retour")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("Vous combattez un ")
		case 2:
			fmt.Print("Vous parlez à Fred")
		case 3:
			SeDeplacer()
		}
	}
}
