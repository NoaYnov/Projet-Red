package red

import "fmt"
//le joueur doit choisir son age et son nom et son classe
var age int
var nom string
var race string
var classe string

func (p *Personnage) InitPerso(){
    p.money = 100


    p.helmet = "rien"
	p.chestplate = "rien"
	p.pants = "rien"
	p.boots = "rien"
	p.gloves = "rien"


    fmt.Println("Choisissez votre age")
    fmt.Scanln(&age)
    p.age = age









    fmt.Println("Choisissez votre nom")
    fmt.Scanln(&nom)
    p.nom = nom

    if p.age == 69 && p.nom == "SUS" {
        p.SUS()
    }else{



    p.Menu_Classe()
    p.classe = classe


    p.inventaireArmor = map[string]string{"Bonnet": "Helmet"}


    p.Menu_Race()
    p.race = race

    }














    fmt.Println("________________________________________________")
    fmt.Printf("\n")
    fmt.Println("Vous avez choisi de jouer en tant que", nom, "de", age, "ans et de classe", classe, "de race",race)
    p.Display()
    fmt.Println("________________________________________________")
    fmt.Printf("\n")

    fmt.Println("Inventaire")

    p.AccessInventory()
    fmt.Println("________________________________________________")
    p.Menu()


    
   
    


}

