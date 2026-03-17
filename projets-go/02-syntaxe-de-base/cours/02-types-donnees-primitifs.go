package main

import "fmt"

func main() {
	// Entiers signés
	var temperature int8 = -15       // Température en hiver
	var altitude int16 = 8848        // Mont Everest en mètres
	var population int32 = 2161000   // Population de Paris
	var timestamp int64 = 1640995200 // 1er janvier 2022

	// Entiers non signés
	var rouge uint8 = 255                 // Composant rouge RGB
	var port uint16 = 8080                // Port web
	var userId uint32 = 123456            // ID utilisateur
	var taillefichier uint64 = 1073741824 // 1 GB en octets

	// Type int général (recommandé pour l'usage courant)
	var age int = 25
	var compteur int = 0

	fmt.Printf("Température: %d°C\n", temperature)
	fmt.Printf("Altitude: %d m\n", altitude)
	fmt.Printf("Population: %d habitants\n", population)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Age: %d ans\n", age)
}
