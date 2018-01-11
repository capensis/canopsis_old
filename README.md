# Makefile

## Idées

Des idées qui me semble intéressante à mettre en place dans le makefile.

 - [x] Cloner des repo privée (cat) [workaround](https://stackoverflow.com/a/27501176)
 utilisation de glide
 - [x] Makefile récursif pour factoriser un max de variables et de targets, et
 éviter un gros makefile de plusieurs centaines de ligne. Et pour offrir la
 possibité de recompiler un «sous-projet»
 - [x] initialiser le projet
   - [x] télécharger les dépendances
   ~~- [ ] cloner, si besoin, les repo privée/autres repo public de canopsis.~~
   ~~> Je ne sais pas si c'est une bonne idée sauf peut être pour cat. Il~~
   ~~> faudra que l'on en parle~~
 - [x] lancer les tests unitaires
 - [ ] lancer les tests  fonctionnels
 - [ ] faire des releases
 - [x] nettoyer le bordel (dépendance go, binaires, …)
 - [ ] afficher de l'aide (en cours)
 - [ ] avoir la possibilité de compiler tous les projets en une commande
 - [x] lancer les build docker
 - [ ] démarrer les images dockers en une seule fois. Docker-compose ?
 - [ ] exécuter des tests pour les images docker


## Problèmes à résoudre
 - comment installer canopsis
 - si dans le makefile on utilise ssh pour clone les repos, on va bloquer les
 utilisateur qui n'omt pas de clé ssh sur le repo. Ce que l'on peux toujours
 faire comme [ça](https://stackoverflow.com/a/22027731)
