// package main qui créer une fonction qui modifie les propriétés d'une structure Compte (nom, solde) par référence.
package main

import "fmt"

// Compte représente un compte bancaire avec un nom et un solde.
type Compte struct {
	Nom   string
	Solde float64
}

// Fonction main pour tester la modification du compte
func main() {
	compte := Compte{Nom: "Alice", Solde: 1000.0}
	fmt.Printf("Compte avant modification: %+v\n", compte)
	modifierCompte(&compte, "Bob", 2000.0)
	fmt.Printf("Compte après modification: %+v\n", compte)
}

// Fonction qui modifie les propriétés d'une structure Compte (nom, solde) par référence.
func modifierCompte(c *Compte, nouveauNom string, nouveauSolde float64) {
	c.Nom = nouveauNom
	c.Solde = nouveauSolde
	fmt.Printf("Compte modifié: %+v\n", c)
}

// Exercice 2 : OK (20.0)
