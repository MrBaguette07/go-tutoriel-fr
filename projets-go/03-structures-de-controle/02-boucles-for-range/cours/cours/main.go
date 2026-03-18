package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compter de 1 à 5
	for i := 1; i <= 5; i++ {
		fmt.Printf("Compteur : %d\n", i)
	}

	infini := 1

	// Équivalent d'un "while" dans d'autres langages
	for infini <= 3 {
		fmt.Printf("infini = %d\n", infini)
		infini++
	}

	i := 10

	// Pas d'initialisation, commence directement
	for ; i > 0; i-- {
		fmt.Printf("Décompte : %d\n", i)
	}
	fmt.Println("Décollage !")

	fruitsSlice()
	chercherElement()
	compterVoyelles()
	pyramideEtoile()
	utilisationMap()

	fruits := []string{"pomme", "banane"}
	for i, fruit := range fruits {
		i = 10         // N'affecte pas la boucle
		fruit = "kiwi" // N'affecte pas le slice
		fmt.Println(i, fruit)
	}

	println("Après la boucle for range :")
	for i, fruit := range fruits {
		fmt.Println(i, fruit) // Affiche les valeurs originales du slice
	}

	for i := 0; i < len(fruits); i++ {
		// Ici on peut modifier i si nécessaire
		fruits[i] = "orange" // Modifie le slice
		fmt.Println(i, fruits[i])
	}

	println("Après la boucle for range :")
	for i, fruit := range fruits {
		fmt.Println(i, fruit) // Affiche les valeurs originales du slice
	}

}

func fruitsSlice() {
	fruits := []string{"pomme", "banane", "orange", "kiwi"}

	for i, nombre := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", i, nombre)
	}

	nombre := 7

	fmt.Printf("Table de multiplication de %d :\n", nombre)
	for i := 1; i <= 10; i++ {
		resultat := nombre * i
		fmt.Printf("%d x %d = %d\n", nombre, i, resultat)
	}

	fmt.Println("Tables de multiplication de 1 à 5 :")

	for i := 1; i <= 5; i++ {
		fmt.Printf("\nTable de %d :\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}

func chercherElement() {
	noms := []string{"Alice", "Bob", "Charlie", "Diana"}
	recherche := "Charlie"
	trouve := false
	position := -1

	for index, nom := range noms {
		if nom == recherche {
			trouve = true
			position = index
			break
		}
	}

	if trouve {
		fmt.Printf("%s a été trouvé dans la liste à l'index %d.\n", recherche, position)
	} else {
		fmt.Printf("%s n'a pas été trouvé dans la liste.\n", recherche)
	}
}

func compterVoyelles() {
	phrase := "J'adore Seapalm ! C'est une super entreprise !"
	voyelles := "aeiouAEIOU"
	compteur := 0

	fmt.Printf("Phrase : %s\n", phrase)

	for _, caractere := range phrase {
		if strings.ContainsRune(voyelles, caractere) {
			compteur++
		}
	}
	fmt.Printf("Nombre de voyelles dans la phrase : %d\n", compteur)
}

func pyramideEtoile() {
	hauteur := 5

	fmt.Println("Pyramide d'étoiles :")
	for ligne := 1; ligne <= hauteur; ligne++ {
		// Imprimer les étoiles pour cette ligne
		for etoile := 1; etoile <= ligne; etoile++ {
			fmt.Print("* ")
		}
		fmt.Println() // Nouvelle ligne
	}
}

func utilisationMap() {
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	for nom, age := range ages {
		fmt.Printf("%s a %d ans.\n", nom, age)
	}
}
