package red

import "fmt"

//le joueur doit choisir son age et son nom et son classe
var age int
var nom string
var race string
var classe string

func (p *Personnage) InitPerso() {
	p.money = 100
	p.liminventaire = 15
	p.limslot = 10
	p.skill = map[string]int{"Coup de poing":10}

	p.armor.helmet = []string{"Bonnet"}
	p.armor.boots = []string{"Chaussures"}
	p.armor.gloves = []string{"Gants"}
	p.armor.pants = []string{"Pantalon"}
	p.armor.chestplate = []string{"Tunique"}

	fmt.Println("Bienvenue aventurier! Vous venez de commencer votre aventure dans la région de Tarate plus précisement dans la ville d'Andariel au sud du pays.")
	fmt.Println("Commençons par vous créer un personnage")
	fmt.Printf("\n")
	fmt.Println("Choisissez votre age")
	fmt.Scanln(&age)
	p.age = age

	fmt.Println("Choisissez votre nom")
	fmt.Scanln(&nom)
	p.nom = nom

	if p.age == 2004 && p.nom == "SUS" {
		p.SUS()
	} else {

		p.Menu_Classe()
		p.classe = classe

		p.Menu_Race()
		p.race = race

	}

	fmt.Println("________________________________________________")
	fmt.Printf("\n")
	fmt.Println("Vous avez choisi de jouer en tant que", nom, "de", age, "ans classe", classe, "et de race", race)
	p.Display()
	fmt.Println("________________________________________________")
	fmt.Printf("\n")

	fmt.Println("Voici la liste des objets de votre inventaire:")

	p.AccessInventory()
	fmt.Println("________________________________________________")
	p.Menu()

}
