start:
	@echo "TerraScan-collector make commands :"
	@echo " - make update         : pour mettre à jour le programme et (re)générer l'exécutable"
	@echo " - make build          : pour simplement (re)générer le fichier exécutable"
	@echo " - make clearlogfile   : pour vider le fichier log de suivi (activity.log)"
	@echo " - make viewlogfile    : pour éditer le fichier activity.log"

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

clearlogfile:
	@echo "===> Effacement du fichier activity.log"
	@rm -f ./logs/activity.log
	@echo "===> Création d'un nouveau fichier log"
	@touch ./logs/activity.log
	@echo "===> Fichier log purgé !"

viewlogfile:
	@echo "===> Ouverture du fichier activity.log"
	@nano ./logs/activity.log
	@echo "===> Fichier log fermé"