// package main qui calcul le salaire mensuel à partir du salaire annnuel.
package main

import "fmt"

var user string = "Alice"
var age int = 25
var salaireAnnuel float64 = 2500.50

// AfficherInformations affiche les informations de l'utilisateur.
func AfficherInformations() {
	// Affichage des informations de l'utilisateur
	fmt.Printf("Nom d'utilisateur: %s\nÂge: %d\nSalaire annuel: %.2f€\n", user, age, salaireAnnuel)
}

// CalculerSalaireMensuel calcule le salaire mensuel à partir du salaire annuel et du nombre de mois.
func CalculerSalaireMensuel(salaireAnnuel float64, mois float64) float64 {
	// Calcul du salaire mensuel en fonction du salaire annuel et du nombre de mois
	return salaireAnnuel / mois
}

func main() {
	AfficherInformations()
	salaireMensuel := CalculerSalaireMensuel(salaireAnnuel, 12)
	fmt.Printf("Salaire mensuel: %.2f\n", salaireMensuel)
}
