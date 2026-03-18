// package main qui créer une fonction saluer qui prend un nom et un moment de la journée ("matin", "après-midi", "soir") et affiche un message approprié.
package main

import "fmt"

func main() {
	var (
		nom, moment string
	)

	// Demander à l'utilisateur d'entrer son nom et le moment de la journée, puis appeler la fonction saluer pour afficher le message de salutation approprié.
	fmt.Print("Entrez votre nom : ")
	fmt.Scanln(&nom)
	fmt.Print("Entrez le moment de la journée (matin, après-midi, soir) : ")
	fmt.Scanln(&moment)
	saluer(nom, moment)
}

// saluer prend un nom et un moment de la journée en entrée et affiche un message de salutation approprié.
func saluer(nom string, moment string) string {
	switch moment {
	case "matin":
		fmt.Printf("Bonjour %s, bon matin !\n", nom)
	case "après-midi":
		fmt.Printf("Bonjour %s, bon après-midi !\n", nom)
	case "soir":
		fmt.Printf("Bonsoir %s, bonne soirée !\n", nom)
	default:
		fmt.Printf("Bonjour %s, bon moment de la journée !\n", nom)
	}
	return ""
}
