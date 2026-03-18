package main

import "fmt"

func main() {
	compteur := 5

	fmt.Printf("Compteur initial: %d\n", compteur)

	// Incrémenter (augmenter de 1)
	compteur++ // Équivalent à: compteur = compteur + 1
	fmt.Printf("Après compteur++: %d\n", compteur)

	// Décrémenter (diminuer de 1)
	compteur-- // Équivalent à: compteur = compteur - 1
	fmt.Printf("Après compteur--: %d\n", compteur)

	// Note: En Go, ++ et -- sont des instructions, pas des expressions
	// Donc ceci ne fonctionne PAS: result := compteur++
	// Il faut faire:
	compteur++
	result := compteur
	fmt.Printf("Résultat: %d\n", result)

	a := 3000000
	b := 200000

	if b != 0 && a/b > 2 {
		fmt.Println("Division OK")
	} else {
		fmt.Println("Division impossible ou résultat trop petit")
	}

	if b == 0 || a/b < 10 {
		fmt.Println("Condition remplie")
	}

	Exercice1()
	Exercice2()
	Exercice3()
	Exercice4()
}

func Exercice1() {
	// Calculateur simple

	a := 5
	b := 30000

	// Toutes les opérations arithmétiques de base
	addition := a + b
	soustraction := a - b
	multiplication := a * b
	division := a / b

	fmt.Printf("Addition: %d + %d = %d\n", a, b, addition)
	fmt.Printf("Soustraction: %d - %d = %d\n", a, b, soustraction)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, multiplication)
	fmt.Printf("Division: %d / %d = %d\n", a, b, division)
}

func Exercice2() {
	// Vérificateur d'âge
	age := 23

	if age >= 18 {
		fmt.Println("Vous êtes majeur.")
	} else if age >= 13 {
		fmt.Println("Vous êtes adolescent.")
	} else if age >= 65 {
		fmt.Println("Vous êtes senior.")
	} else {
		fmt.Println("Vous êtes mineur.")
	}

	// Déterminer les droits
	if age >= 18 {
		fmt.Println("Vous pouvez voter.")
	} else if age >= 15 {
		fmt.Println("Vous pouvez conduire.")
	} else if age >= 13 {
		fmt.Println("Vous pouvez regarder des films pour adolescents.")
	} else if age >= 16 {
		fmt.Println("Vous pouvez travailler à temps partiel.")
	} else {
		fmt.Println("Vous êtes trop jeune pour voter ou conduire.")
	}
}

func Exercice3() {
	// Calculateur d'imc avec recommandations
	poids := 70.0  // en kg
	taille := 1.75 // en m

	imc := poids / (taille * taille)
	fmt.Printf("Votre IMC est: %.2f\n", imc)

	if imc < 18.5 {
		fmt.Println("Vous êtes en sous-poids.")
	} else if imc < 25 {
		fmt.Println("Vous avez un poids normal.")
	} else if imc < 30 {
		fmt.Println("Vous êtes en surpoids.")
	} else {
		fmt.Println("Vous êtes en obésité.")
	}

	// Recommandations basées sur l'IMC selon l'âge
	age := 30
	if imc < 18.5 {
		if age < 18 {
			fmt.Println("Il est important de consulter un médecin pour évaluer votre croissance.")
		} else {
			fmt.Println("Il est recommandé de consulter un nutritionniste pour un plan alimentaire adapté.")
		}
	} else if imc < 25 {
		fmt.Println("Continuez à maintenir un mode de vie sain avec une alimentation équilibrée et de l'exercice régulier.")
	} else if imc < 30 {
		fmt.Println("Il est conseillé de consulter un professionnel de santé pour évaluer les risques et élaborer un plan de perte de poids.")
	} else {
		fmt.Println("Il est crucial de consulter un professionnel de santé pour une évaluation complète et un plan de traitement adapté.")
	}
}

func Exercice4() {
	// Système de tarification
	age := 70
	typeBillet := "simple"
	estSenior := age >= 65
	estEnfant := age < 12
	jour := "weekend"
	prixBase := 20.0
	statusAbonnement := "étudiant"

	prixDuBillet := prixBase

	if estEnfant {
		fmt.Printf("Prix du billet pour enfant: %.2f\n", prixBase*0.5)
		prixDuBillet = prixBase * 0.5
	} else if estSenior {
		fmt.Printf("Prix du billet pour senior: %.2f\n", prixBase*0.7)
		prixDuBillet = prixBase * 0.7
	} else {
		fmt.Printf("Prix du billet pour adulte: %.2f\n", prixBase)
	}

	if typeBillet == "aller-retour" {
		prixDuBillet *= 1.8
		fmt.Printf("Prix du billet aller-retour: %.2f\n", prixDuBillet)
	} else {
		fmt.Printf("Prix du billet simple: %.2f\n", prixDuBillet)
	}

	if jour == "weekend" {
		prixDuBillet *= 0.9
		fmt.Printf("Prix du billet avec réduction weekend: %.2f\n", prixDuBillet)
	} else {
		fmt.Printf("Prix du billet en semaine: %.2f\n", prixDuBillet)
	}

	if statusAbonnement == "étudiant" {
		prixDuBillet *= 0.8
		fmt.Printf("Prix du billet avec réduction étudiant: %.2f\n", prixDuBillet)
	}

	fmt.Printf("Prix final du billet: %.2f\n", prixDuBillet)
}

// Exercice 1: OK
// Exercice 2: OK
// Exercice 3: semi-ok
// Exercice 4: Semi-ok
