start:
	@echo "TerraScan-collector make commands :"
	@echo " - make update                  : pour mettre à jour le programme et (re)générer l'exécutable"
	@echo " - make build                   : pour simplement (re)générer le fichier exécutable"
	@echo " - make clearlogsfiles          : pour vider les fichiers log (activity.log et errors.log)"
	@echo " - make viewactitivylogfile     : pour éditer le fichier activity.log"
	@echo " - make viewerrorslogfile       : pour éditer le fichier errors.log"

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

clearlogsfiles:
	@echo "===> Effacement des fichiers activity.log et errors.log"
	@rm -f ./logs/activity.log
	@rm -f ./logs/errors.log
	@echo "===> Création de nouveaux fichiers log vides"
	@touch ./logs/activity.log
	@touch ./logs/errors.log
	@echo "===> Fichiers log purgés !"

viewactitivylogfile:
	@echo "===> Ouverture du fichier activity.log"
	@nano ./logs/activity.log
	@echo "===> Fichier activity.log fermé"

viewerrorslogfile:
	@echo "===> Ouverture du fichier errors.log"
	@nano ./logs/errors.log
	@echo "===> Fichier errors.log fermé"