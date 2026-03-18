# Makefile pour projet Go
.PHONY: fmt vet lint test doc clean

# Formater le code
fmt:
	go fmt ./...

# Vérifier avec go vet
vet:
	go vet ./...

# Linter (nécessite golint)
lint:
	golint ./...

# Exécuter tous les tests
test:
	go test -v ./...

# Générer la documentation
doc:
	godoc -http=:6060

# Vérification complète avant commit
check: fmt vet lint test
	@echo "Toutes les vérifications sont passées !"

# Nettoyer les fichiers générés
clean:
	go clean ./...
	rm -f *.out *.test