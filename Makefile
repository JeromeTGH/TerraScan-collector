start:
	@echo "===> Tapez 'make update' pour mettre à jour le programme et (re)générer l'exécutable"
	@echo "        ou 'make build' pour simplement (re)générer le fichier exécutable"

build:
	@echo "===> Suppression du fichier exécutable 'main', si existant"
	@rm -f main
	@echo "===> Génération d'un nouveau fichier exécutable"
	@go build ./main.go
	@echo "===> Génération terminée !"

update:
	@echo "===> Suppression du fichier exécutable 'main', si existant"
	@rm -f main
	@echo "===> Récupération des éventuelles modifs du programme, depuis le dépôt GIT ..."
	@git pull
	@echo "===> Génération d'un nouveau fichier exécutable"
	@go build ./main.go
	@echo "===> Mise à jour terminée !"
