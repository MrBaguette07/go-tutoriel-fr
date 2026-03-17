package main
import "fmt"

var cacahuete string = "coco"
cacahuete := "coco"


type Personne struct {
	Nom string
	Age int
} // Le type struct Personne qui a deux champs : Nom et Age

func (p Personne) String() string {
	return "Bonjour, je suis " p.Nom " et j'ai " p.Age " ans."
} // Une fonction qui retourne une string en fonction du paramètre de type Personne

type Parleur interface {
	Saluer() string
} // Un type d'interface



// La je voudrais ouvrir un fichier .md par exemple et lire son contenu
fichier, err := os.Open("monfichier.md")
if err != nil {
	return err
}

defer fichier.Close()


go Saluer() // ça exécute en parallèle

ch := make(chan string) // création d'un canal pour communiquer entre les goroutines

// Make c'est pour créer un canal, la dans ce cas précis on créer un canal de type string, on peut aussi faire make(chan int) pour un canal d'entiers par exemple
// Mais est-ce que chan est obligatoire et que veut dire chan ?
// chan est un mot-clé en Go qui est utilisé pour déclarer un canal. Un canal est une structure de données qui permet la communication entre les goroutines. Les goroutines peuvent envoyer et recevoir des valeurs à travers les canaux, ce qui facilite la synchronisation et la coordination entre elles. En utilisant chan, on peut créer des canaux de différents types, comme chan string pour un canal de chaînes de caractères ou chan int pour un canal d'entiers.
// Donc chan est obligatoire pour déclarer un canal, et il indique que la variable qui suit est un canal de communication entre les goroutines.



go func() {
	ch <- "Hello from the goroutine!" // envoie un message dans le canal
}() 

message := <-ch // reçoit le message du canal
fmt.Println(message) // affiche le message reçu

// Que veut dire fmt ? 
// fmt est un package de la bibliothèque standard de Go qui fournit des fonctions pour formater et afficher des données. Il est couramment utilisé pour afficher des messages à la console, formater des chaînes de caractères, et effectuer des opérations d'entrée/sortie. Par exemple, fmt.Println() est une fonction qui affiche une ligne de texte à la console, tandis que fmt.Sprintf() permet de formater une chaîne de caractères en fonction de certains paramètres.
// Doit on importer le package fmt pour pouvoir l'utiliser ?
// Oui, pour utiliser le package fmt, il est nécessaire de l'importer dans votre code Go. Vous pouvez le faire en ajoutant la ligne suivante au début de votre fichier Go :
// Pourquoi j'ai une erreur en haut de mon code ? avec package ?
// L'erreur que vous avez en haut de votre code est probablement due au fait que vous n'avez pas déclaré le package auquel appartient votre code. En Go, chaque fichier doit appartenir à un package, et la première ligne de votre fichier doit déclarer ce package. Par exemple, si vous voulez que votre code fasse partie du package "main", vous devez ajouter la ligne suivante au début de votre fichier :

