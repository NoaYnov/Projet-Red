package red

import "fmt"
//le joueur doit choisir son age et son nom et son classe
var age int
var nom string
var race string
var classe string

func (p *Personnage) InitPerso(){
    fmt.Println("Choisissez votre age")
    fmt.Scanln(&age)
    p.age = age









    fmt.Println("Choisissez votre nom")
    fmt.Scanln(&nom)
    p.nom = nom








    p.Menu_Race()
    p.race = race









    p.Menu_Classe()
    p.classe = classe









    fmt.Println("________________________________________________")
    fmt.Printf("\n")
    fmt.Println("Vous avez choisi de jouer en tant que", nom, "de", age, "ans et de classe", classe)
    fmt.Println("________________________________________________")


    
   
    


}

