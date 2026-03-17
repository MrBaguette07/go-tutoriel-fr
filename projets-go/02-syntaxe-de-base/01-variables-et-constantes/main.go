package main

import "fmt"

var (
	nom    string  = "Alice"
	age    int     = 25
	taille float64 = 1.75
	actif  bool    = true
)

const (
	Dimanche = iota * 2
	Lundi
	Mardi
	Mercredi
	Jeudi
	Vendredi
	Samedi

	NomApplication  = "MonApp"
	Version         = "1.0.0"
	MaxUtilisateurs = 100
)

func main() {

	var nom string = "Alice"
	fmt.Println("Nom initial:", nom)

	nom = "Bob"
	fmt.Println("nouveau nom:", nom)

	age := 30
	fmt.Println("Age initial:", age)

	age = 26
	fmt.Println("Nouvel age:", age)

	fmt.Println("Dimanche, lundi, mardi, mercredi, jeudi, vendredi, samedi: ", Dimanche, Lundi, Mardi, Mercredi, Jeudi, Vendredi, Samedi)

	calculatrice()
	infosUser()

	println("--- Exercice 1 ---")
	Exercice1()

	println("--- Exercice 2 ---")
	Exercice2()

	println("--- Exercice 3 ---")
	Exercice3()

	println("--- Exercice 4 ---")
	Exercice4()
}

func calculatrice() {
	a := 10
	b := 5

	addition := a + b
	soustraction := a - b
	multiplication := a * b
	division := a / b

	fmt.Printf("Addition: %d + %d = %d\n", a, b, addition)
	fmt.Printf("Soustraction: %d - %d = %d\n", a, b, soustraction)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, multiplication)
	fmt.Printf("Division: %d / %d = %d\n", a, b, division)
}

func infosUser() {
	// Informations utilisateur
	nom := "Alice"
	prenom := "Dupont"
	age := 25
	estPremium := true

	// Variables calculées
	nomComplet := prenom + " " + nom
	peutUtiliserApp := age >= 18

	// Affichage
	fmt.Printf("=== %s v%s ===\n", NomApplication, Version)
	fmt.Printf("Utilisateur: %s\n", nomComplet)
	fmt.Printf("Age: %d ans\n", age)
	fmt.Printf("Statut Premium: %t\n", estPremium)
	fmt.Printf("Peut utiliser l'app: %t\n", peutUtiliserApp)
	fmt.Printf("Utilisateurs max: %d\n", MaxUtilisateurs)
}

func gestionStock() {

	var stockPommes int = 50
	var stockOranges int = 30
	var stockBananes int = 20

	ventesPommes := 12
	ventesOranges := 8
	ventesBananes := 15

	stockPommes -= ventesPommes
	stockOranges -= ventesOranges
	stockBananes -= ventesBananes

	stockTotal := stockPommes + stockOranges + stockBananes

	// Affichage du rapport
	fmt.Println("=== Rapport de Stock ===")
	fmt.Printf("Pommes restantes: %d\n", stockPommes)
	fmt.Printf("Oranges restantes: %d\n", stockOranges)
	fmt.Printf("Bananes restantes: %d\n", stockBananes)
	fmt.Printf("Stock total: %d\n", stockTotal)
}

func Exercice1() {
	var nom string = "Alice"
	var prenom string = "Dupont"
	var age int = 25
	var ville string = "Paris"

	fmt.Printf("Mon nom est %s\n", nom)
	fmt.Printf("Mon prénom est %s\n", prenom)
	fmt.Printf("J'ai %d ans\n", age)
	fmt.Printf("Je vis à %s\n", ville)
}

func Exercice2() {
	const anneeActuelle = 2026
	var dateNaissance int = 2003

	age := anneeActuelle - dateNaissance
	fmt.Printf("Age de la personne: %d ans\n", age)
}

func Exercice3() {
	var tempCelsius float64 = 25.0

	tempFahrenheit := (tempCelsius * 9 / 5) + 32

	fmt.Printf("%.2f °C = %.2f °F\n", tempCelsius, tempFahrenheit)
}

func Exercice4() {
	var (
		mensuel float64 = 5000.0
	)

	var (
		chips   float64 = 1500.0
		loyer   float64 = 2000.0
		courses float64 = 800.0
	)

	depensesMensuel := chips + loyer + courses
	soldeMensuel := mensuel - depensesMensuel

	fmt.Printf("Revenus mensuels: %.2f\n", mensuel)
	fmt.Printf("Dépenses mensuelles: %.2f\n", depensesMensuel)
	fmt.Printf("Solde mensuel: %.2f\n", soldeMensuel)
}
