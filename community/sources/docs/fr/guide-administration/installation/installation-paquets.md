# Installation de paquets Canopsis sur Red Hat Enterprise Linux 8 et 9

!!! Information
    Si vous souhaitez réaliser une mise à jour, la procédure est décrite dans le [guide de mise à jour](../mise-a-jour/index.md).

Cette procédure décrit l'installation de Canopsis en mono-instance à l'aide de paquets RHEL 8 et 9. Les binaires sont compilés pour l'architecture x86-64.

!!! Warning
    L'installation de tous les services sur un seul nœud, décrite ici, permet
    de monter aisément un environnement de test.  
    Cependant, **ce type d'installation ne saurait être conseillé pour de la
    production**.

    Une plate-forme Canopsis de production repose généralement sur plusieurs
    nœuds, avec les mécanismes de haute disponibilité appropriés. Ce type de
    déploiement plus avancé est supporté uniquement dans le cadre
    d'une souscription à l'édition Pro de Canopsis.

    [Capensis][capensis] peut vous accompagner dans la mise en place de
    l'architecture Canopsis distribuée et des solutions de haute disponibilité.

[capensis]: https://www.capensis.fr/

Toutes les commandes dans cette procédure doivent être invoquées avec
l'utilisateur `root`.

## Prérequis

Assurez-vous d'avoir suivi les [prérequis réseau et de sécurité](../administration-avancee/configuration-parefeu-et-selinux.md), notamment concernant la désactivation de SELinux.

Pour vérifier l'état de SELinux :

```sh
sestatus
```

L'installation nécessite l'ajout de dépôts RPM tiers, ainsi qu'un accès HTTP et HTTPS pour le téléchargement de diverses dépendances. Plus de détails dans la [matrice des flux réseau](../matrice-des-flux-reseau/index.md).

!!! Information
    Notez que les versions de MongoDB, RabbitMQ, Redis et TimescaleDB dont l'installation est décrite ici sont les seules validées pour fonctionner avec Canopsis.

    Plus de détails sur les [prérequis des versions](prerequis-des-versions.md).

## Dépendances

### Mise à jour système

Assurez-vous que le système est à jour (l'installation sur RHEL 8 ou 9 suppose que le système est relié à des dépôts à jour de la distribution, et en particulier pas figé dans une ancienne version mineure 8.x ou 9.x.) :

```sh
dnf update
```

### Configuration système

Vous pouvez vérifier les limites de ressources systèmes avec la commande suivante :

```sh
ulimit -a
```

Pour appliquer la [configuration recommandée par le projet MongoDB](https://www.mongodb.com/docs/v7.0/reference/ulimit/), créez le fichier `/etc/security/limits.d/mongo.conf` :

```sh
cat << EOF > /etc/security/limits.d/mongo.conf
#<domain>      <type>  <item>         <value>
mongo           soft    fsize           unlimited
mongo           soft    cpu             unlimited
mongo           soft    as              unlimited
mongo           soft    memlock         unlimited
mongo           hard    nofile          64000
mongo           hald    nproc           64000
EOF
```

Désactivez la gestion des `Transparent Huge Pages (THP)` selon la [préconisation MongoDB](https://www.mongodb.com/docs/manual/tutorial/transparent-huge-pages/)

```sh
cat << EOF > /etc/systemd/system/disable-transparent-huge-pages.service
[Unit]
Description=Disable Transparent Huge Pages (THP)
DefaultDependencies=no
After=sysinit.target local-fs.target
Before=mongod.service
[Service]
Type=oneshot
ExecStart=/bin/sh -c 'echo never | tee /sys/kernel/mm/transparent_hugepage/enabled > /dev/null'
[Install]
WantedBy=basic.target
EOF
```

```sh
systemctl daemon-reload
systemctl enable --now disable-transparent-huge-pages
```

### Ajout des dépôts tiers

=== "RHEL 8"

    Ajout du dépôt pour PostgreSQL :

    ```sh
    dnf install https://download.postgresql.org/pub/repos/yum/reporpms/EL-8-x86_64/pgdg-redhat-repo-latest.noarch.rpm
    ```

    Ajout du dépôt pour MongoDB :

    ```sh
    cat << EOF > /etc/yum.repos.d/mongodb-org-7.0.repo
    [mongodb-org-7.0]
    name=MongoDB Repository
    baseurl=https://repo.mongodb.org/yum/redhat/\$releasever/mongodb-org/7.0/x86_64/
    gpgcheck=1
    enabled=1
    gpgkey=https://www.mongodb.org/static/pgp/server-7.0.asc
    EOF
    ```

    Ajout du dépôt pour RabbitMQ :

    ```sh
    ## primary RabbitMQ signing key
    rpm --import 'https://github.com/rabbitmq/signing-keys/releases/download/3.0/rabbitmq-release-signing-key.asc'
    ## modern Erlang repository
    rpm --import 'https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-erlang.E495BB49CC4BBE5B.key'
    ## RabbitMQ server repository
    rpm --import 'https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-server.9F4587F226208342.key'

    cat << EOF > /etc/yum.repos.d/rabbitmq.repo
    ##
    ## Zero dependency Erlang RPM
    ##

    [modern-erlang]
    name=modern-erlang-el8
    # Use a set of mirrors maintained by the RabbitMQ core team.
    # The mirrors have significantly higher bandwidth quotas.
    baseurl=https://yum1.novemberain.com/erlang/el/8/$basearch
            https://yum2.novemberain.com/erlang/el/8/$basearch
    repo_gpgcheck=1
    enabled=1
    gpgkey=https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-erlang.E495BB49CC4BBE5B.key
    gpgcheck=1
    sslverify=1
    sslcacert=/etc/pki/tls/certs/ca-bundle.crt
    metadata_expire=300
    pkg_gpgcheck=1
    autorefresh=1
    type=rpm-md

    ##
    ## RabbitMQ Server
    ##

    [rabbitmq-el8-noarch]
    name=rabbitmq-el8-noarch
    baseurl=https://yum2.novemberain.com/rabbitmq/el/8/noarch
            https://yum1.novemberain.com/rabbitmq/el/8/noarch
    repo_gpgcheck=1
    enabled=1
    # Cloudsmith's repository key and RabbitMQ package signing key
    gpgkey=https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-server.9F4587F226208342.key
        https://github.com/rabbitmq/signing-keys/releases/download/3.0/rabbitmq-release-signing-key.asc
    gpgcheck=1
    sslverify=1
    sslcacert=/etc/pki/tls/certs/ca-bundle.crt
    metadata_expire=300
    pkg_gpgcheck=1
    autorefresh=1
    type=rpm-md
    EOF
    ```

    Ajout du dépôt pour TimescaleDB :

    ```sh
    cat << EOF > /etc/yum.repos.d/timescale_timescaledb.repo
    [timescale_timescaledb]
    name=timescale_timescaledb
    baseurl=https://packagecloud.io/timescale/timescaledb/el/8/\$basearch
    repo_gpgcheck=1
    # TimescaleDB doesn’t sign all its packages
    gpgcheck=0
    enabled=1
    gpgkey=https://packagecloud.io/timescale/timescaledb/gpgkey
    sslverify=1
    sslcacert=/etc/pki/tls/certs/ca-bundle.crt
    metadata_expire=300
    EOF
    ```
=== "RHEL 9"

    Ajout du dépôt pour PostgreSQL :

    ```sh
    dnf install https://download.postgresql.org/pub/repos/yum/reporpms/EL-9-x86_64/pgdg-redhat-repo-latest.noarch.rpm
    ```

    Ajout du dépôt pour MongoDB :

    ```sh
    cat << EOF > /etc/yum.repos.d/mongodb-org-7.0.repo
    [mongodb-org-7.0]
    name=MongoDB Repository
    baseurl=https://repo.mongodb.org/yum/redhat/\$releasever/mongodb-org/7.0/x86_64/
    gpgcheck=1
    enabled=1
    gpgkey=https://www.mongodb.org/static/pgp/server-7.0.asc
    EOF
    ```

    Ajout du dépôt pour RabbitMQ :

    ```sh
    ## primary RabbitMQ signing key
    rpm --import 'https://github.com/rabbitmq/signing-keys/releases/download/3.0/rabbitmq-release-signing-key.asc'
    ## modern Erlang repository
    rpm --import 'https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-erlang.E495BB49CC4BBE5B.key'
    ## RabbitMQ server repository
    rpm --import 'https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-server.9F4587F226208342.key'

    cat << EOF > /etc/yum.repos.d/rabbitmq.repo
    ##
    ## Zero dependency Erlang RPM
    ##

    [modern-erlang]
    name=modern-erlang-el9
    # Use a set of mirrors maintained by the RabbitMQ core team.
    # The mirrors have significantly higher bandwidth quotas.
    baseurl=https://yum1.rabbitmq.com/erlang/el/9/$basearch
            https://yum2.rabbitmq.com/erlang/el/9/$basearch
    repo_gpgcheck=1
    enabled=1
    gpgkey=https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-erlang.E495BB49CC4BBE5B.key
    gpgcheck=1
    sslverify=1
    sslcacert=/etc/pki/tls/certs/ca-bundle.crt
    metadata_expire=300
    pkg_gpgcheck=1
    autorefresh=1
    type=rpm-md

    ##
    ## RabbitMQ Server
    ##

    [rabbitmq-el9-noarch]
    name=rabbitmq-el9-noarch
    baseurl=https://yum1.rabbitmq.com/rabbitmq/el/9/noarch
            https://yum2.rabbitmq.com/rabbitmq/el/9/noarch
    repo_gpgcheck=1
    enabled=1
    # Cloudsmith's repository key and RabbitMQ package signing key
    gpgkey=https://github.com/rabbitmq/signing-keys/releases/download/3.0/cloudsmith.rabbitmq-server.9F4587F226208342.key
        https://github.com/rabbitmq/signing-keys/releases/download/3.0/rabbitmq-release-signing-key.asc
    gpgcheck=1
    sslverify=1
    sslcacert=/etc/pki/tls/certs/ca-bundle.crt
    metadata_expire=300
    pkg_gpgcheck=1
    autorefresh=1
    type=rpm-md
    EOF
    ```

    Ajout du dépôt pour TimescaleDB :

    ```sh
    cat << EOF > /etc/yum.repos.d/timescale_timescaledb.repo
    [timescale_timescaledb]
    name=timescale_timescaledb
    baseurl=https://packagecloud.io/timescale/timescaledb/el/9/\$basearch
    repo_gpgcheck=1
    # TimescaleDB doesn’t sign all its packages
    gpgcheck=0
    enabled=1
    gpgkey=https://packagecloud.io/timescale/timescaledb/gpgkey
    sslverify=1
    sslcacert=/etc/pki/tls/certs/ca-bundle.crt
    metadata_expire=300
    EOF
    ```

### Configuration des dépôts

Exécuter la commande suivante et vérifier dans la sortie que les dépôts ajoutés
sont bien contactés (accepter les cléfs des différents dépôts lorsque demandé) :

```sh
dnf makecache
```

Désactiver le module PostgreSQL ([requis pour l'installation de TimescaleDB sur RHEL](https://docs.timescale.com/install/latest/self-hosted/installation-redhat/)) :

```sh
dnf module disable postgresql
```

=== "RHEL 8"

    Activer le module Nginx 1.20.* :

    ```sh
    dnf module disable php
    dnf module enable nginx:1.20
    ```

    Activer le module Redis 6.0.* :

    ```sh
    dnf module enable redis:6
    ```

=== "RHEL 9"

    Activer le module Nginx 1.24.* :

    ```sh
    dnf module disable php
    dnf module enable nginx:1.24
    ```

    Activer le module Redis 7.0.* :

    ```sh
    dnf module enable redis:7
    ```

### Installation

=== "RHEL 8"

    ```sh
    dnf install logrotate socat mongodb-org nginx redis timescaledb-2-postgresql-13-2.14.2 timescaledb-2-loader-postgresql-13-2.14.2
    dnf install erlang rabbitmq-server
    ```

    Pour éviter une mise à jour vers des versions non souhaitées de TimescaleDB
    ou RabbitMQ, vous devriez utiliser [*versionlock*][dnf-versionlock] :

    ```sh
    dnf install 'dnf-command(versionlock)'
    dnf versionlock add timescaledb-2-loader-postgresql-13 timescaledb-2-postgresql-13
    dnf versionlock add --raw 'rabbitmq-server-3.*'
    ```

    Les autres dépendances de Canopsis proviennent de canaux garantissant déjà le
    maintien dans la branche majeure souhaitée (exemples : MongoDB 7.0, Redis 6).

    [dnf-versionlock]: https://dnf-plugins-core.readthedocs.io/en/latest/versionlock.html

=== "RHEL 9"

    ```sh
    dnf install logrotate socat mongodb-org nginx redis timescaledb-2-postgresql-16 timescaledb-2-loader-postgresql-16
    dnf install erlang rabbitmq-server
    ```

    Pour éviter une mise à jour vers des versions non souhaitées de TimescaleDB
    ou RabbitMQ, vous devriez utiliser [*versionlock*][dnf-versionlock] :

    ```sh
    dnf install 'dnf-command(versionlock)'
    dnf versionlock add timescaledb-2-loader-postgresql-16 timescaledb-2-postgresql-16
    dnf versionlock add --raw 'rabbitmq-server-3.*'
    ```

    Les autres dépendances de Canopsis proviennent de canaux garantissant déjà le
    maintien dans la branche majeure souhaitée (exemples : MongoDB 7.0, Redis 6).

    [dnf-versionlock]: https://dnf-plugins-core.readthedocs.io/en/latest/versionlock.html

### Ouverture des ports

Pratiquer les ouvertures de ports nécessaires à l'accès au service.

Les commandes données couvrent le cas standard où le pare-feu système `firewalld` est utilisé, et servent surtout à rappeler les ports ou services à ouvrir. (cf. [matrice des flux réseau](../matrice-des-flux-reseau/index.md))

```sh
firewall-cmd --add-port=5672/tcp --add-port=15672/tcp --permanent
firewall-cmd --add-port=8080/tcp --permanent
firewall-cmd --add-port=27017/tcp --permanent
firewall-cmd --add-service=postgresql --permanent
firewall-cmd --add-service=redis --permanent
firewall-cmd --reload
```

### Configuration de MongoDB

Une clef d'authentification doit être créée pour sécuriser le futur *replicaset*
MongoDB. Pour cela, exécuter les commandes suivantes qui vont générer le
fichier `/etc/mongodb-keyfile` et lui attribuer les bonnes permissions et
appartenances :

```sh
openssl rand -base64 756 > /etc/mongodb-keyfile
chmod 400 /etc/mongodb-keyfile
chown mongod:root /etc/mongodb-keyfile
```

Ensuite, éditer le fichier `/etc/mongod.conf`. Il faut mettre les contenus
suivants dans les sections `security` et `replication` pour :

- protéger l'instance (authentification obligatoire)
- activer le mode *replicaset*

```yaml
security:
  authorization: enabled
  keyFile: /etc/mongodb-keyfile

replication:
  oplogSizeMB: 1024
  replSetName: rs0
```

!!! information
    Le nom de *replicaset* `rs0` est un exemple. Un autre nom peut être choisi,
    du moment que ce nom est répercuté plus tard dans l'URL MongoDB au sein du
    fichier `/opt/canopsis/etc/go-engines-vars.conf`.

On peut à présent activer et démarrer le service :

```sh
systemctl enable --now mongod.service
```

L'instance MongoDB étant démarrée, il reste à la configurer.

On se connecte dans un shell `mongosh` et on désactive la télémétrie :

```sh
mongosh
> disableTelemetry()
```

On initialise le *replicaset* :

```sh
> rs.initiate()
```

L'état du *replicaset* peut être vérifié avec la commande `rs.status()` :

```sh
> rs.status()
```

Au bout de quelques secondes, le prompt du shell `mongosh` doit faire apparaître
que le nœud est PRIMARY :

```sh
rs0:PRIMARY>
```

Lorsque c'est le cas, le *replicaset* est prêt. On poursuit avec la création
des utilisateurs MongoDB `root` puis `canopsis`, toujours dans le shell
`mongosh` :

```sh
> use admin
> db.createUser({user: "root", pwd: "UNMOTDEPASSEFORT", roles: [ { role: "root", db: "admin" }]})
> exit
```

On se reconnecte avec le shell `mongosh`, cette fois-ci en s'authentifiant en tant
que `root` MongoDB :

```sh
mongosh -u root -p UNMOTDEPASSEFORT
> use canopsis
> db.createUser({user: "cpsmongo", pwd: "canopsis", roles: [ { role: "dbOwner", db: "canopsis" }, { role: "clusterMonitor", db: "admin"}]})
> exit
```

Les manipulations d'installation dans MongoDB sont terminées.

### Configuration de TimescaleDB

Initialiser l'instance PostgreSQL puis initialiser TimescaleDB (cf. [documentation de l'outil de règlage](https://docs.timescale.com/timescaledb/latest/how-to-guides/configuration/timescaledb-tune/) de TimescaleDB) :

=== "RHEL 8"
    ```sh
    postgresql-13-setup initdb
    timescaledb-tune -yes --pg-config=/usr/pgsql-13/bin/pg_config
    echo "timescaledb.telemetry_level=off" >> /var/lib/pgsql/13/data/postgresql.conf
    ```

    Activer et démarrer le service :

    ```sh
    systemctl enable --now postgresql-13.service
    ```

    Se connecter à l'instance PostgreSQL avec l'identité du superuser `postgres` :

    ```sh
    sudo -u postgres psql
    ```

    Créer la base de données `canopsis` et l'utilisateur associé dans l'instance PostgreSQL :

    ```sql
    postgres=# CREATE database canopsis;
    postgres=# \c canopsis
    canopsis=# CREATE EXTENSION IF NOT EXISTS timescaledb;
    canopsis=# SET password_encryption = 'scram-sha-256';
    canopsis=# CREATE USER cpspostgres WITH PASSWORD 'canopsis';
    canopsis=# exit
    ```

=== "RHEL 9"
    ```sh
    postgresql-16-setup initdb
    timescaledb-tune -yes --pg-config=/usr/pgsql-16/bin/pg_config
    echo "timescaledb.telemetry_level=off" >> /var/lib/pgsql/13/data/postgresql.conf
    ```

    Activer et démarrer le service :

    ```sh
    systemctl enable --now postgresql-16.service
    ```

    Se connecter à l'instance PostgreSQL avec l'identité du superuser `postgres` :

    ```sh
    sudo -u postgres psql
    ```

    Créer la base de données `canopsis` et l'utilisateur associé dans l'instance PostgreSQL :

    ```sql
    postgres=# CREATE database canopsis;
    postgres=# \c canopsis
    canopsis=# CREATE EXTENSION IF NOT EXISTS timescaledb;
    canopsis=# SET password_encryption = 'scram-sha-256';
    canopsis=# CREATE USER cpspostgres WITH PASSWORD 'canopsis';
    canopsis=# GRANT ALL ON DATABASE canopsis TO cpspostgres;
    canopsis=# ALTER DATABASE canopsis OWNER TO cpspostgres;
    canopsis=# exit
    ```

### Configuration de RabbitMQ

Activer et démarrer le service :

```sh
systemctl enable --now rabbitmq-server.service
```

Activer le plugin qui apporte l'interface de management RabbitMQ, puis redémarrer le service :

```sh
rabbitmq-plugins enable rabbitmq_management
systemctl restart rabbitmq-server.service
```

Créer les objets RabbitMQ (vhost, utilisateur avec les bons droits) utiles à l'application Canopsis :

```sh
rabbitmqctl add_vhost canopsis
rabbitmqctl add_user cpsrabbit canopsis
rabbitmqctl set_user_tags cpsrabbit administrator
rabbitmqctl set_permissions --vhost canopsis cpsrabbit '.*' '.*' '.*'
```

### Démarrage de Redis

Ajouter un mot de passe ( ici `canopsis`)

=== "RHEL 8"

    ```sh
    sed -i 's/^# requirepass.*/requirepass canopsis/' /etc/redis.conf
    ```

=== "RHEL 9"

    ```sh
    sed -i 's/^# requirepass.*/requirepass canopsis/' /etc/redis/redis.conf
    ```

Activer et démarrer le service :

```sh
systemctl enable --now redis
```

## Installation de Canopsis Community ou Pro

Canopsis est disponible dans une édition « Community », open-source et gratuitement accessible à tous, et une édition « Pro », souscription commerciale ajoutant des fonctionnalités supplémentaires. Voyez [le site officiel de Canopsis](https://www.capensis.fr/canopsis/) pour en savoir plus.

Notez que l'édition Pro de Canopsis était auparavant connue sous le nom de « CAT » et que certains éléments peuvent encore la désigner sous ce nom.

Cliquez sur l'un des onglets « Community » ou « Pro » suivants, en fonction de l'édition choisie.

=== "Canopsis Community (édition open-source)"

    Ajout du dépôt de paquets Canopsis :

    === "RHEL 8"

        ```sh
        cat << EOF > /etc/yum.repos.d/canopsis.repo
        [canopsis]
        name = canopsis
        baseurl=https://nexus.canopsis.net/repository/canopsis/el8/community/
        gpgcheck=0
        enabled=1
        EOF
        ```

    === "RHEL 9"

        ```sh
        cat << EOF > /etc/yum.repos.d/canopsis.repo
        [canopsis]
        name = canopsis
        baseurl=https://nexus.canopsis.net/repository/canopsis/el9/community/
        gpgcheck=0
        enabled=1
        EOF
        ```

    Installation de l'édition open-source de Canopsis :

    ```sh
    dnf makecache
    dnf install canopsis
    ```

=== "Canopsis Pro (souscription commerciale)"

    !!! attention
        L'édition Pro nécessite une souscription commerciale ainsi qu'une authentification d'accès aux repos à renseigner dans `baseurl` du fichier `/etc/yum.repos.d/canopsis-pro.repo`.

    Ajout des dépôts de paquets Canopsis :

    === "RHEL 8"

        ```sh
        cat << EOF > /etc/yum.repos.d/canopsis.repo
        [canopsis]
        name = canopsis
        baseurl=https://nexus.canopsis.net/repository/canopsis/el8/community/
        gpgcheck=0
        enabled=1
        EOF

        cat << EOF > /etc/yum.repos.d/canopsis-pro.repo
        [canopsis-pro]
        name = canopsis-pro
        baseurl=https://user:password@nexus.canopsis.net/repository/canopsis-pro/el8/pro/
        gpgcheck=0
        enabled=1
        EOF
        ```

    === "RHEL 9"

        ```sh
        cat << EOF > /etc/yum.repos.d/canopsis.repo
        [canopsis]
        name = canopsis
        baseurl=https://nexus.canopsis.net/repository/canopsis/el9/community/
        gpgcheck=0
        enabled=1
        EOF

        cat << EOF > /etc/yum.repos.d/canopsis-pro.repo
        [canopsis-pro]
        name = canopsis-pro
        baseurl=https://user:password@nexus.canopsis.net/repository/canopsis-pro/el9/pro/
        gpgcheck=0
        enabled=1
        EOF
        ```

    Installation de Canopsis Pro :

    ```sh
    dnf makecache
    dnf install canopsis-pro
    ```

## Initialisation de Canopsis

Le fichier de configuration est `/opt/canopsis/etc/go-engines-vars.conf`, qui
est normalement dans l'état suivant :

```sh
CPS_MONGO_URL="mongodb://cpsmongo:canopsis@localhost:27017/canopsis?replicaSet=rs0"
CPS_AMQP_URL="amqp://cpsrabbit:canopsis@localhost:5672/canopsis"
CPS_POSTGRES_URL="postgresql://cpspostgres:canopsis@localhost:5432/canopsis"
CPS_REDIS_URL="redis://localhost:6379/0"
CPS_API_URL="http://localhost:8082"
#CPS_OLD_API_URL="http://localhost:8081"
CPS_POSTGRES_TECH_URL="postgresql://cpspostgres:canopsis@localhost:5432/canopsis_tech_metrics"
```

!!! Note
    On remarque dans `CPS_MONGO_URL` le nom de *replicaset*, ici `rs0`,
    qui doit être cohérent avec ce qui a été mis lors de la
    [configuration de MongoDB](#configuration-de-mongodb).

Il est maintenant temps d'initialiser l'instance Canopsis ;
cliquez sur l'un des onglets « Community » ou « Pro » suivants, en fonction de
l'édition choisie.

=== "Canopsis Community (édition open-source)"

    Provisionner Canopsis :

    ```sh
    set -o allexport; source /opt/canopsis/etc/go-engines-vars.conf; /opt/canopsis/bin/canopsis-reconfigure -migrate-postgres=true -edition community
    ```

    Activer et démarrer les services :

    ```sh
    systemctl enable --now canopsis-engine-go@engine-action.service \
                           canopsis-engine-go@engine-axe.service \
                           canopsis-engine-go@engine-che.service \
                           canopsis-engine-go@engine-fifo.service \
                           canopsis-engine-go@engine-pbehavior.service \
                           canopsis-service@canopsis-api.service \
                           canopsis.service
    ```

=== "Canopsis Pro (souscription commerciale)"

    Provisionner Canopsis :

    ```sh
    set -o allexport; source /opt/canopsis/etc/go-engines-vars.conf; /opt/canopsis/bin/canopsis-reconfigure -migrate-postgres=true -edition pro
    ```

    Activer et démarrer les services :

    ```sh
    systemctl enable --now canopsis-engine-go@engine-action.service \
                           canopsis-engine-go@engine-axe.service \
                           canopsis-engine-go@engine-che.service \
                           canopsis-engine-go@engine-correlation.service \
                           canopsis-engine-go@engine-dynamic-infos.service \
                           canopsis-engine-go@engine-fifo.service \
                           canopsis-engine-go@engine-pbehavior.service \
                           canopsis-engine-go@engine-remediation.service \
                           canopsis-engine-go@engine-webhook.service \
                           canopsis-service@canopsis-api.service \
                           canopsis-engine-python-snmp.service \
                           canopsis.service
    ```

Tester un envoi d'évènement :

```sh
curl -X POST -u root:root -H "Content-Type: application/json" -d '{
  "event_type": "check",
  "connector": "connector_test_creation_alarmes",
  "connector_name": "test",
  "component": "component_test_creation_alarmes",
  "resource": "resource_test_creation_alarmes",
  "source_type": "resource",
  "author": "QA_canopsis",
  "state": 2,
  "debug": true,
  "output": "Test création alarmes Canopsis"
}' 'http://localhost:8082/api/v4/event'
```

## Lancement de la Web UI de Canopsis

Installer le paquet :

!!! attention
    Le package `canopsis-webui` est disponible pour EL9 uniquement à partir de la version 24.04.2 !

```sh
dnf install canopsis-webui
```

Activation de https dans Canopsis:

Une configuration HTTPS est proposée avec Nginx, mais elle n'est cependant pas encore activée par défaut.  
Vous pouvez suivre la procédure suivante: [activation de https dans Canopsis](../administration-avancee/configuration-composants/reverse-proxy-nginx-https.md)

Activer et démarrer Nginx :

```sh
systemctl enable --now nginx.service
```

Une fois cette commande terminée, vous pouvez alors réaliser votre [première connexion à l'interface Canopsis](premiere-connexion.md).

## Mises à jour

Lorsque vous souhaitez réaliser une mise à jour, la procédure est décrite dans
le [Guide de mise à jour](../mise-a-jour/index.md).

Il est fortement conseillé d'utiliser à nouveau [*versionlock*][dnf-versionlock]
pour rester dans la version majeure de Canopsis que vous venez d'installer.

Ainsi, vous bénéficierez des mises à jour mineures. Mais vous ne risquez pas
une montée de version majeure non prévue vers une majeure suivante à l'occasion
d'une mise à jour de routine de l'ensemble des paquets système.

=== "Canopsis Community (édition open-source)"

    ```sh
    dnf versionlock add --raw 'canopsis-24.04.*'
    dnf versionlock add --raw 'canopsis-webui-24.04.*'
    ```

=== "Canopsis Pro (souscription commerciale)"

    ```sh
    dnf versionlock add --raw 'canopsis-pro-24.04.*'
    dnf versionlock add --raw 'canopsis-webui-24.04.*'
    ```
