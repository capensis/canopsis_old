# Installation de Canopsis avec Docker Compose

Cette procédure décrit l'installation de Canopsis avec Docker Compose.

!!! Warning
    La présente documentation décrit un déploiement où tous les services
    (*backends* et applicatifs) seront exécutés dans des conteneurs Docker, le
    tout sur une seule machine.

    Ce type d'installation permet de monter très facilement un environnement de
    test reproductible et jetable.

    Cependant, **cette architecture n'est pas adaptée pour de la production**.
    En effet, en production, il serait recommandé d'installer les services
    *backends* – bases de données en particulier – hors de Docker, puis de
    mettre les conteneurs applicatifs sur *plusieurs* hôtes Docker afin d'avoir
    une certaine haute disponibilité (tolérance de panne).

    Le déploiement multi-nœuds fait l'objet d'un accompagnement spécifique par
    [Capensis][capensis] dans le cadre d'une souscription à l'édition Pro de
    Canopsis.

[capensis]: https://www.capensis.fr/

## Prérequis

### Utilisation de Docker Compose

[Docker Compose](https://docs.docker.com/compose/) est l'orchestrateur Docker couramment supporté pour Canopsis.

!!! important
    Les configurations Docker Compose fournies par Canopsis sont maintenues et testées seulement avec Docker Compose. La compatibilité directe de ces configurations avec d'autres outils d'orchestration tels que Kubernetes, Docker Swarm, Consul, OpenShift, etc. n'est à ce jour pas assurée.

    Les clients disposant d'une souscription Canopsis Pro qui souhaitent déployer dans un cluster Kubernetes peuvent en revanche se tourner vers le [chart Helm][helm] maintenu par Capensis.

[helm]: ./installation-helm.md

### Prérequis de version du noyau Linux

Lors de l'utilisation de Docker, Canopsis nécessite **un noyau Linux 4.4 minimum sur votre système hôte**.

De nos jours, ce point est acquis à partir du moment où vous utilisez une version supportée d'une de ces distributions majeures orientées « serveur » : EL, Debian, Ubuntu LTS, SLES, OpenSUSE Leap.

En cas de doute, vérifiez votre version du noyau à l'aide de la commande suivante :

```sh
uname -r
```

Si la version affichée est inférieure à 4.4, vous devez soit utiliser une distribution GNU/Linux plus à jour, ou bien installer un noyau plus récent via des dépôts spécifiques, lorsque c'est possible.

!!! important
    L'utilisation de Docker Compose avec un noyau inférieur à 4.4 n'est pas prise en charge.

## Installation de Docker et Docker Compose

Vous devez tout d'abord [installer Docker](https://docs.docker.com/engine/install/), version 20.10 minimum (se référer à page [Prérequis des versions][prereq-versions]). Veuillez utiliser les dépôts officiels de Docker, et non pas ceux proposés par votre distribution.

Veillez à installer Docker Compose à cette occasion, comme indiqué dans les instructions de la documentation d'installation officielle du Docker Engine. Le paquet se nomme `docker-compose-plugin`.

!!! information
    Dans le passé, [Docker Compose][docker-compose] était distribué sous forme de programme séparé (binaire `docker-compose` à télécharger et installer). Ce mode d'utilisation est aujourd'hui abandonné par Docker.

    Depuis plusieurs versions majeures de Docker Engine, Compose est intégré à la ligne de commande `docker` avec le *plugin* Compose.  
    Concrètement, on utilise à présent `docker compose <subcommand>...` là où les commandes `docker-compose <subcommand>...` étaient auparavant employées.

    Ces évolutions de Compose ont aussi été accompagnées de nouveaux formats et de nouvelles possibilités pour l'orchestrateur (voir [Compose Specification][compose-spec]).

    Concernant Canopsis, depuis la version 22.10 les environnements de référence fournis se basent sur ces dernières versions. L'utilisation du *plugin* Compose intégré à Docker (commandes `docker compose ...`) est maintenant la norme.

## Lancement de Canopsis avec Docker Compose

Les images Docker officielles de Canopsis sont hébergées sur leur propre registre Docker, `docker.canopsis.net`.

### Récupération de l'environnement Docker Compose

Les environnements Docker Compose de référence pour Canopsis sont disponibles via
git :

=== "Canopsis Pro"

    Pour Canopsis Pro, les fichiers sont dans le
    [dépôt git canopsis-pro][canopsis-pro] dans la partie "Release"..

    Décompressez l'archive :

    ```
    tar -xvzf canopsis-pro-docker-compose-XX.XX.X.tar.gz
    ```

    Déplacez-vous ensuite dans le dossier contenant l'environnement :

    ```
    cd canopsis-pro-docker-compose-XX.XX.X
    ```

=== "Canopsis Community"

    Pour Canopsis Community, les fichiers sont dans le
    [dépôt git canopsis-community][canopsis-community] dans la partie "Release".

    Décompressez l'archive :

    ```
    tar -xvzf canopsis-community-docker-compose-XX.XX.X.tar.gz
    ```

    Déplacez-vous ensuite dans le dossier contenant l'environnement :

    ```
    cd canopsis-community-docker-compose-XX.XX.X
    ```

### Lancement de l'environnement

Récupérez les dernières images disponibles :

```sh
docker compose pull
```

Lancez ensuite la commande suivante, afin de démarrer un environnement Canopsis
complet :

```sh
docker compose up -d
```
## Vérification du bon fonctionnement

```sh
docker compose ps
```

Les services doivent être en état `Up`, `Up (healthy)` ou `Exit 0`. En fonction
des ressources de votre machine, il peut être nécessaire d'attendre quelques
minutes avant de voir l'ensemble des moteurs en état `Up`.

Vous pouvez ensuite procéder à votre [première connexion à l'interface Canopsis](premiere-connexion.md).

## Arrêt de l'environnement Docker Compose

```sh
docker compose down
```

## Rétention des logs

La mise en place d'une politique de rétention des logs nécessite la présence du logiciel `logrotate`.

Une fois que `logrotate` est installé sur votre machine, créer le fichier `/etc/logrotate.d/docker-container` suivant :

```
/var/lib/docker/containers/*/*.log {
  rotate 7
  daily
  compress
  minsize 100M
  notifempty
  missingok
  delaycompress
  copytruncate
}
```

Pour vérifier la validité de la configuration logrotate ajoutée, lancez la commande :

```sh
logrotate -dv /etc/logrotate.d/docker-container
```

Si vous souhaitez forcer une exécution manuelle de cette rotation sur-le-champ, vous pouvez éventuellement lancer la commande :

```sh
logrotate -fv /etc/logrotate.d/docker-container
```

[prereq-versions]: https://doc.canopsis.net/latest/guide-administration/installation/prerequis-des-versions/#prerequis-systemes
[compose-spec]: https://docs.docker.com/compose/compose-file/
[docker-compose]: https://docs.docker.com/compose/install/#install-compose
[canopsis-pro]: https://git.canopsis.net/canopsis/canopsis-pro/-/releases
[canopsis-community]: https://git.canopsis.net/canopsis/canopsis-community/-/releases