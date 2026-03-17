var nom string = "Alice"
nom2 = "Bob"
age := 30
var nom3 string

var nomGlobal string = "Variable globale"  // Accessible partout

func main() {
    fmt.Println(nomGlobal)  // Fonctionne
    afficherNom()
}

func afficherNom() {
    fmt.Println(nomGlobal)  // Fonctionne aussi
}

