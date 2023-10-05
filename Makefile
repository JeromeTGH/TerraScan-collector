start:
        @echo "===> Tapez 'make update', pour mettre à jour le programme"

update:
        @echo "===> Suppression du fichier exécutable 'main'"
        @rm -f main
        @echo "===> Récupération des éventuelles modifs du programme, depuis le dépôt GIT ..."
        @git pull
        @echo "===> Génération d'un nouveau fichier exécutable"
        @go build ./main.go
        @echo "===> Mise à jour terminée !"
