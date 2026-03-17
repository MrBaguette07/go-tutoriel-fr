package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var petit int8 = 127 // Valeur maximale pour int8
	fmt.Printf("Avant: %d\n", petit)

	petit = petit + 127              // Débordement !
	fmt.Printf("Après: %d\n", petit) // Résultat: -2 (inattendu!)

	Exercice1()
	Exercice2()
	Exercice3()
	Exercice4()

}

func nombresDecimaux() {
	// float32 : précision simple
	var prix float32 = 19.99
	var taux float32 = 0.075

	// float64 : précision double (recommandé)
	var pi float64 = 3.141592653589793

	// Calculs
	total := prix * (1 + taux)
	circonference := 2 * pi * 10 // Rayon = 10

	fmt.Printf("Prix TTC: %.2f€\n", total)
	fmt.Printf("Circonférence: %.2f\n", circonference)

	// Notation scientifique
	var tresPetit float64 = 1.23e-10
	var tresGrand float64 = 6.02e23 // Nombre d'Avogadro

	fmt.Printf("Très petit: %e\n", tresPetit)
	fmt.Printf("Très grand: %e\n", tresGrand)
}

func nombresFlotants() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = a + b

	fmt.Printf("0.1 + 0.2 = %.17f\n", c) // Résultat: 0.30000000000000004

	// Pour les calculs monétaires, utilisez des entiers !
	var prixCentimes int = 10 + 20 // 30 centimes = 0.30€
	fmt.Printf("Prix exact: %.2f€\n", float64(prixCentimes)/100)
}

func chainesCaracteres() {
	// Déclarations simples
	var nom string = "Alice"
	var message string = "Bonjour tout le monde!"

	// Avec la syntaxe courte
	prenom := "Bob"
	ville := "Paris"

	// String vide
	var description string // Valeur par défaut: ""

	fmt.Printf("description: %s\n", description)

	// String multiligne
	var poeme string = `
	Les sanglots longs
	Des violons
	De l'automne
	`

	fmt.Printf("Nom: %s\n", nom)
	fmt.Printf("Prénom: %s\n", prenom)
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Ville: %s\n", ville)
	fmt.Printf("Poème:%s\n", poeme)
}

func OperationsString() {
	prenom := "Alice"
	nom := "Dupont"

	// Méthode 1 : avec +
	nomComplet := prenom + " " + nom

	// Méthode 2 : avec fmt.Sprintf
	presentation := fmt.Sprintf("Je m'appelle %s %s", prenom, nom)

	fmt.Println(nomComplet)   // Alice Dupont
	fmt.Println(presentation) // Je m'appelle Alice Dupont

	message := "Hello, 世界" // Mélange anglais/chinois

	fmt.Printf("Longueur en octets: %d\n", len(message))           // 13
	fmt.Printf("Nombre de caractères: %d\n", len([]rune(message))) // 9

	//[]rune(message) // Convertit la chaîne en une slice de runes (caractères Unicode)
	// rune est un type de données qui représente un point de code Unicode. En convertissant la chaîne en une slice de runes, on peut compter correctement le nombre de caractères, même si certains d'entre eux sont représentés par plusieurs octets (comme les caractères chinois).

	// Les valeurs "zéro" sont considérées comme false
	age := 0

	estMajeur := age >= 18
	aNom := nom != ""

	fmt.Printf("Est majeur: %t\n", estMajeur) // false
	fmt.Printf("A un nom: %t\n", aNom)        // true

	var lettre byte = 'A' // Caractère ASCII
	var valeur byte = 65  // Même chose (65 = code ASCII de 'A')

	fmt.Printf("Lettre: %c\n", lettre)              // A
	fmt.Printf("Valeur: %d\n", lettre)              // 65
	fmt.Printf("Identique: %t\n", lettre == valeur) // true

	var emoji rune = '😀'    // Emoji
	var chinois rune = '中'  // Caractère chinois
	var francais rune = 'é' // Caractère français avec accent

	fmt.Printf("Emoji: %c (code: %d)\n", emoji, emoji)
	fmt.Printf("Chinois: %c (code: %d)\n", chinois, chinois)
	fmt.Printf("Français: %c (code: %d)\n", francais, francais)

	// Parcourir une string caractère par caractère
	texte := "Héllo 世界 😀"
	for _, char := range texte {
		fmt.Printf("%c ", char)
	}
	fmt.Println()

}

func conversionEntier() {
	// Conversions entre entiers
	var a int = 42
	var b int64 = int64(a) // Conversion explicite requise
	var c int8 = int8(a)   // Attention au débordement !

	// Conversions entre entiers et flottants
	var entier int = 42
	var decimal float64 = float64(entier)
	var retourEntier int = int(decimal)

	// Conversions avec strings
	var nombre int = 123
	var texte string = fmt.Sprintf("%d", nombre) // int vers string

	fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
	fmt.Printf("Entier: %d, Décimal: %.1f\n", entier, decimal)
	fmt.Printf("Nombre en texte: %s\n", texte)
	fmt.Printf("Retour à l'entier: %d\n", retourEntier)
}

func conversionAvancee() {
	// String vers nombre
	texteNombre := "12999"
	nombre, err := strconv.Atoi(texteNombre)
	if err != nil {
		fmt.Println("Erreur de conversion:", err)
	} else {
		fmt.Printf("Nombre: %d\n", nombre)
	}

	// Nombre vers string
	valeur := 456
	texte := strconv.Itoa(valeur)
	fmt.Printf("Texte: %s\n", texte)

	// Float vers string
	pi := 3.14159
	piTexte := strconv.FormatFloat(pi, 'f', 2, 64)
	fmt.Printf("Pi: %s\n", piTexte) // 3.14
}

func profilUtilisateur() {
	var (
		nom         string  = "Alice"
		prenom      string  = "Dupont"
		age         int     = 25
		ville       string  = "Paris"
		estEtudiant bool    = true
		taille      float64 = 1.65
		noteMoyenne float64 = 15.5
		nbCours     uint8   = 6
	)

	fmt.Printf("Profil de l'utilisateur:\n")
	fmt.Printf("\tNom: %s\n", nom)
	fmt.Printf("\tPrénom: %s\n", prenom)
	fmt.Printf("\tAge: %d ans\n", age)
	fmt.Printf("\tVille: %s\n", ville)
	fmt.Printf("\tEst étudiant: %t\n", estEtudiant)
	fmt.Printf("\tTaille: %.2f m\n", taille)
	fmt.Printf("\tNote moyenne: %.1f\n", noteMoyenne)
	fmt.Printf("\tNombre de cours: %d\n", nbCours)

	tailleEnCm := taille * 100
	estMajeur := age >= 18

	fmt.Printf("Taille en cm: %.0f cm\n", tailleEnCm)
	fmt.Printf("Est majeur: %t\n", estMajeur)

	// NON
	var prix float64 = 0.1 + 0.2 // 0.30000000000000004
	fmt.Println(prix)
	// OUI
	var prixCentimes int = 10 + 20 // 30 centimes
	var prixEuros float64 = float64(prixCentimes) / 100
	fmt.Printf("%.2f\n", prixEuros)
}

func differentsTests() {
	var nom string = "Alice"

	nom = "B" + nom[1:]
	fmt.Println(nom)
}

func Exercice1() {
	// Calculateur d'IMC

	var (
		poids  float64 = 68.0 // en kg
		taille float64 = 1.75 // en m
	)

	imc := poids / (taille * taille)
	fmt.Printf("IMC: %.2f\n", imc)
}

func Exercice2() {
	// Profil complet

	// informations personnelles
	var (
		nom         string  = "Alice"
		prenom      string  = "Dupont"
		age         int     = 25
		ville       string  = "Paris"
		estEtudiant bool    = true
		taille      float64 = 1.65
	)

	// Les préférences de l'utilisateur
	var (
		couleurFavorite    string = "Bleu"
		langagePrefere     string = "Go"
		nombreHeuresCoding int    = 4
	)

	// Statistiques
	var (
		nbAmis      int     = 150
		nbPosts     int     = 45
		noteMoyenne float64 = 4.5
	)

	fmt.Printf("Profil de l'utilisateur:\n")
	fmt.Printf("\tNom: %s\n", nom)
	fmt.Printf("\tPrénom: %s\n", prenom)
	fmt.Printf("\tAge: %d ans\n", age)
	fmt.Printf("\tVille: %s\n", ville)
	fmt.Printf("\tEst étudiant: %t\n", estEtudiant)
	fmt.Printf("\tTaille: %.2f m\n", taille)
	fmt.Printf("\tCouleur favorite: %s\n", couleurFavorite)
	fmt.Printf("\tLangage préféré: %s\n", langagePrefere)
	fmt.Printf("\tNombre d'heures de coding par jour: %d\n", nombreHeuresCoding)
	fmt.Printf("\tNombre d'amis: %d\n", nbAmis)
	fmt.Printf("\tNombre de posts: %d\n", nbPosts)
	fmt.Printf("\tNote moyenne: %.1f\n", noteMoyenne)
}

func Exercice3() {
	// Convertisseur de température

	var celsius float64 = 25.0

	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15

	fmt.Printf("%.2f °C = %.2f °F\n", celsius, fahrenheit)
	fmt.Printf("%.2f °C = %.2f K\n", celsius, kelvin)
}

func Exercice4() {
	// Analyse de texte

	var texte string = "Go est un langage de programmation open source créé par Google."

	nbCaractere := len(texte)
	premierMot := texte[:2]
	contientCertainsMots := (len(texte) > 0) && (texte[0] == 'G' || texte[0] == 'g')

	contientGo := strings.Contains(texte, "Go") || strings.Contains(texte, "go")
	contientPython := strings.Contains(texte, "Python") || strings.Contains(texte, "python")

	fmt.Printf("Contient 'Go' ou 'go': %t\n", contientGo)
	fmt.Printf("Contient 'Python' ou 'python': %t\n", contientPython)
	fmt.Printf("Texte: %s\n", texte)
	fmt.Printf("Nombre de caractères: %d\n", nbCaractere)
	fmt.Printf("Premier mot: %s\n", premierMot)
	fmt.Printf("Contient 'Go' ou 'go' au début: %t\n", contientCertainsMots)
}

// Exercice 1 : OK
// Exercice 2 : OK
// Exercice 3 : OK
// Exercice 4 : Partiellement OK (il faut vérifier la présence de "Go" et "Python" dans le texte, pas seulement au début, je n'avais pas importé le package "strings" pour ça)
