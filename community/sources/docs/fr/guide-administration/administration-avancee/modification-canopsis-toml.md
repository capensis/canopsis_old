# Modification du fichier de configuration toml `canopsis.toml`

## Description

Le fichier `canopsis.toml` regroupe la plupart des réglages fondamentaux des différents moteurs et services de Canopsis.

!!! note
    Les réglages d'exploitation « du quotidien » se situent plutôt dans l'interface web de Canopsis.

    D'autres réglages propres à certains moteurs se font au travers de :

    - leurs [variables d'environnement](variables-environnement.md) ;
    - leurs options de lancement, que l'on retrouve dans l'aide intégrée au binaire de chaque moteur (`-help`) ;
    - valeurs définies par défaut dans un fichier de configuration `canopsis.toml`, surchargeables (objet de cette page).

## Emplacement

L'emplacement du fichier de configuration diffère entre les différents types d'environnement d'installation proposés par Canopsis.

| Type d'environnement                    | Emplacement du fichier                                     |
|-----------------------------------------|------------------------------------------------------------|
| Paquets RPM                             | `/opt/canopsis/etc/canopsis.toml`                          |
| Docker Compose/K8S (Canopsis Pro)       | `/canopsis-pro.toml` dans le conteneur `reconfigure`       |
| Docker Compose/K8S (Canopsis Community) | `/canopsis-community.toml` dans le conteneur `reconfigure` |

!!! attention
    La modification de valeurs ne doit pas se faire dans le fichier d'origine.

    On utilise plutôt un fichier de surcharge (`canopsis-override.toml`), comme
    indiqué dans la suite de cette documentation.

## Modification et maintenance du fichier

Le fichier `canopsis-override.toml` permet de surcharger la configuration par défaut.
Ce fichier contiendra seulement les paramètres qui diffèrent de la configuration par défaut.

=== "Paquets RPM"

    Le fichier est situé au chemin suivant : `/opt/canopsis/etc/conf.d/canopsis-override.toml`.

    Lors de la mise à jour de Canopsis, vos modifications dans ce fichier seront préservées par le gestionnaire de paquets `dnf`.

=== "Docker Compose"

    Le fichier est situé dans le conteneur `reconfigure` au chemin suivant : `/opt/canopsis/etc/conf.d/canopsis-override.toml`.

    Montez-y votre fichier personnalisé à l'aide d'un volume.

    Ceci est déjà en place dans les environnements de référence Docker Compose :

    ```yaml
    services:
      # ...

      reconfigure:
        # ...
        volumes:
          - ./files/canopsis/reconfigure/conf.d/:/opt/canopsis/etc/conf.d/
    ```

    Un exemple de fichier avec des paramètres commentés est présent dans la
    configuration Docker Compose livrée.

=== "K8S avec le chart Helm (Pro)"

    Vous référer au chart Helm et sa documentation.

    Méthode 1 (applicable dans tous les cas) :

    - mettez votre fichier `canopsis-override.toml` dans un `ConfigMap`
    - indiquez le nom du `ConfigMap` via la *value* `reconfigure.tomlConfigMap`

    Méthode 2 (applicable si vous avez fait un `helm pull` et maintenez votre
    propre copie locale du chart) :

    - placez votre fichier dans le répertoire local `files/reconfigure/` prévu
    à cet effet

## Prise en compte des modifications

Après toute modification de valeur dans les fichiers `.toml` de Canopsis,
la commande `canopsis-reconfigure` doit être relancée, puis les services et
moteurs Canopsis doivent être redémarrés.

=== "Paquets RPM"

    Exécuter les commandes suivantes :

    ```bash
    set -o allexport ; source /opt/canopsis/etc/go-engines-vars.conf
    /opt/canopsis/bin/canopsis-reconfigure
    systemctl restart canopsis.service
    ```

=== "Docker Compose"

    Exécuter les commandes suivantes :

    ```sh
    docker compose restart reconfigure
    docker compose restart
    ```

=== "K8S avec le chart Helm (Pro)"

    Après modification du contenu du fichier local ou du `ConfigMap` de votre
    cru, il faut redéployer le chart avec la commande `helm upgrade`.

    Le job `reconfigure` sera alors recréé et lancé (nom de la forme
    `release-name-reconfigure-x-y-z` basé sur votre nom de release et la
    version de Canopsis en cours).

    Attendez la bonne fin d'exécution du job `reconfigure`, puis redémarrez
    tous les `Deployments` correspondant aux moteurs et services Canopsis, par
    exemple avec ce type de commande :

    ```bash
    kubectl rollout restart deploy -l app.kubernetes.io/instance=$RELEASE_NAME
    ```

## Description des options

### Section [Canopsis.global]

| Attribut                             | Exemple de valeur          | Description                          |
| :----------------------------------- | :------------------------- | :----------------------------------- |
| PrefetchCount                        | 10000                      |
| PrefetchSize                         | 0                          |
| ReconnectTimeoutMilliseconds         | 8                          | Délai de reconnexion auprès des services tiers (redis, mongodb, rabbitmq, ...)  |
| ReconnectRetries                     | 3                          | Nombre de tentative de reconnexion aux services tiers |
| MaxExternalResponseSize              | 10485760 # 10Mb            | Taille maximale d'une réponse d'API tierce |
| BuildEntityInfosDictionary           | true, false                | Activation de la génération des infos d'entités dans un dictionnaire servant à l'auto complétation coté WebUI | 
| BuildDynamicInfosDictionary          | true, false                | Activation de la génération des infos d'alarmes dans un dictionnaire servant à l'auto complétation coté WebUI | 
| InfosDictionaryLimit                 | 1100                       | Nombre maximum de paires clé/valeur à stocker dans les dictionnaires mentionnés ci-dessus |
| EventsCountTriggerDefaultThreshold   | 10                         | Valeur par défaut du nombre d'événements utilisé par le trigger |


### Section [Canopsis.file]

| Attribut       | Exemple de valeur                        | Description                          |
| :------------- | :--------------------------------------- | :----------------------------------- |
| Dir            | "/opt/canopsis/var/lib"                  | Emplacement des fichiers uploadés. Utilisé pour le module de [remédiation](../../guide-utilisation/remediation/index.md), des icônes, des exports csv, et dans le cadre de la multi-instanciation du service `api` |
| UploadMaxSize  | 314572800 # 300Mb                        | Taille maximale d'un fichier à uploader (en octet) |
| SnmpMib        | ["/usr/share/snmp/mibs"]                 | Emplacement des fichiers MIB qui seront utilisés par le module SNMP |
| IconMaxSize    | 10240 # 10Kb                             | Taille max des fichiers d'icônes |


### Section [Canopsis.alarm]

| Attribut                          | Exemple de valeur     | Description                          |
| :-------------------------------- | :---------------------| :----------------------------------- |
| StealthyInterval                  |                       | Encore utilisé ?          |
| :warning: Obsolète :warning: EnableLastEventDate               | true,false            | Active la mise à jour du champ `last_event_date` d'une alarme à chaque événement :warning: Depuis Canopsis 23.10, la date de dernier changement est nécessairement calculée :warning:  | 
| CancelAutosolveDelay              | "1h"                  | Délai de résolution effective d'une alarme après annulation depuis l'interface graphiqe |
| DisplayNameScheme                 | "{{ rand_string 3 }}-{{ rand_string 3 }}-{{ rand_string 3 }}" | Schéma utilisé pour générer le champ `display_name` d'une alarme |
| OutputLength                      | 255                   | Nombre maximum de caractères du champ `output` avant troncage | 
| LongOutputLength                  | 1024                  | Nombre maximum de caractères du champ `long_output` avant troncage | 
| DisableActionSnoozeDelayOnPbh     | true,false            | Si `vrai` alors le délai du snooze n'est pas ajouté à un comportement périodique |
| TimeToKeepResolvedAlarms          | "720h"                | Délai durant lequel les alarmes résolues sont conservées dans la collection principale des alarmes |
| AllowDoubleAck                    | true,false            | Permet d'acquitter plusieurs fois une alarme |
| ActivateAlarmAfterAutoRemediation | true,false            | Permet de décaler l'activation d'une alarme après l'exécution de la remédiation automatique |
| EnableArraySortingInEntityInfos   | true,false            | Active ou désactive le tri dans les listes utilisées dans les attributs d'événements. Par exemple, si un événement contient `info1=["item2", "item1"]` et que l'option est activée alors info1 vaudra en sortie `info1=["item1", "item2"]` |
| CropStepsNumber                   | 0                     | Nombre de steps `stateinc` ou `statedec` continus avant de les "compresser" en un step "statecounter".<br>Cerla permet de diminuer drastiquement le nombre de steps d'un alarme |

### Section [Canopsis.timezone]

| Attribut | Exemple de valeur | Description                           |
| :------- | :-----------------| :------------------------------------ |
| Timezone | "Europe/Paris"    | Timezone générale du produit Canopsis |


### Section [Canopsis.data_storage]

| Attribut      | Exemple de valeur | Description                           |
| :------------ | :-----------------| :------------------------------------ |
| TimeToExecute | "Sunday,23"       | Jour et heure d'exécution de la politique de rotation des données définie dans le module `Data Storage` | 


### Section [Canopsis.import_ctx]

| Attribut            | Exemple de valeur     | Description                           |
| :------------------ | :---------------------| :------------------------------------ |
| ThdWarnMinPerImport | "30m"                 | Durée d'import au délà de laquelle une alarme mineure sera générée |
| ThdCritMinPerImport | "60m"                 | Durée d'import au délà de laquelle une alarme critique sera générée |
| FilePattern         | "/tmp/import_s.json"  | Pattern de nommage des fichiers temporaires d'import  |

### Section [Canopsis.api]

| Attribut                 | Exemple de valeur  | Description                           |
| :----------------------- | :------------------| :------------------------------------ |
| TokenSigningMethod       | "HS256"            | Méthode de signature d'un token d'authentification |
| BulkMaxSize              | 1000               | Taille maximum d'un batch (api endpoint) de changement en données en base |
| ExportMongoClientTimeout | "1m"               | Durée maximum d'un export au format CSV |
| AuthorScheme             | ["$username"]      | Permet de définir la manière de représenter l'auteur d'une action dans Canopsis. Ex : `["$username", " ", "$firstname", " ", "$lastname", " ", "$email", " ", "$_id"] ` |
| MetricsCacheExpiration   | "24h"              | Durée de validité du cache des API liées aux métriques |

### Section [Canopsis.logger]

| Attribut            | Exemple de valeur  | Description                                             |
| :------------------ | :------------------| :------------------------------------------------------ |
| Writer              | "stdout"           | Canal de sortie du logger. **`stdout`** ou **`stderr`** |

### Sous-section [Canopsis.logger.console_writer]
Toute modification dans cette section nécessite un redémarrage de Canopsis

| Attribut            | Exemple de valeur                           | Description                                             |
| :------------------ | :-------------------------------------------| :------------------------------------------------------ |
| Enabled             | true                                        | Active ou désactive le mode [ConsoleWriter](https://github.com/rs/zerolog#pretty-logging). Si désactivé alors les messages sont loggués en JSON. |
| NoColor             | true                                        | Active ou désactive les couleurs dans les logs |
| TimeFormat          | "2006-01-02T15:04:05Z07:00"                 | Format des dates des messages de logs au format [GO](../../guide-utilisation/templates-go/index.md) |
| PartsOrder          | ["time", "level", "caller", "message"]      | Ordre des parties des messages de logs parmi "time", "level", "message", "caller", "error" |

### Section [Canopsis.metrics]

| Attribut               | Exemple de valeur  | Description                           |
| :--------------------- | :------------------| :------------------------------------ |
| Enabled                | false|true         | Activation / Désactivation des métriques |
| FlushInterval          | "10s"              | Délai d'écriture des métriques dans la base de données |
| SliInterval            | "1h"               | Les longs intervalles de SLI sont découpés en plus petits intervalles définis par cet attribut. <br />Une valeur faible augmente la précision des métriques mais nécessite plus d'espace disque. <br />Une valeur élevée diminue la précision des métriques mais nécessaite moins d'espace disque. <br /> "1h" est la valeur recommandée dans la mesure où l'intervalle le plus petit gérée par l'interface graphique correspond à 1 heure (Ne peut pas être > "1h" |
| UserSessionGapInterval | "1h"               | Précision des temps de session utilisateur. Prenons un utilisateur, actif sur la l'interface graphique de 9:00 à 18:00, avec UserSessionGapInterval=1h. <br />La table des métriques user ressemblera à :<br />- 9:00   3600 <br />- 10:00 3600<br />- 11:00 3600<br />- ...<br />- 17:00 3600<br /> |
| AllowedPerfDataUnits   | ["%","°C","s","us","ms","B","KB","MB","TB","c"]  | Liste blanche des unités de métriques autorisées dans un événement |


### Section [Canopsis.tech_metrics]

| Attribut            | Exemple de valeur  | Description                           |
| :------------------ | :------------------| :------------------------------------ |
| Enabled             | false|true         | Active ou non la collecte des [métriques techniques](../../guide-de-depannage/metriques-techniques/index.md) |
| DumpKeepInterval    | "1h"               | Détermine le temps durant lequel les dumps seront disponibles avant leur suppression                    |


### Section [Canopsis.template.vars]

| Attribut                | Exemple de valeur  | Description                           |
| :---------------------- | :------------------| :------------------------------------ |
| system_env_var_prefixes | ["ENV_"]           | Les variables d'environnement peuvent être utilisées dans des [templates Go](../../guide-utilisation/templates-go/index.md) sous la forme `{{ .Env.System.ENV_var }}` ou dans l'interface graphique en [Handlebars](../../guide-utilisation/cas-d-usage/template_handlebars.md) sous la forme `{{ env.System.ENV_var }}`.<br />Seules les variables dont le prefixe est mentionné dans ce paramètre seront lues. |
| var1                    | "valeur1"          | Ces variables peuvent être utilisées dans des [templates Go](../../guide-utilisation/templates-go/index.md) sous la forme `{{ .Env.var }}` ou dans l'interface graphique en [Handlebars](../../guide-utilisation/cas-d-usage/template_handlebars.md) sous la forme `{{ env.var1 }}` |


### Section [Remediation]

| Attribute | Example | Description
| ------ | ------ | ------ |
| http_timeout | "1m" | Timeout de connexion au serveur distant |
| launch_job_retries_amount | 3 | Nombre de tentatives d'exécution du job sur le serveur distant |
| launch_job_retries_interval | "5s" | Intervalle de temps entre 2 tentative d'exécution d'un job |
| wait_job_complete_retries_amount | 12 | Nombre par défaut de tentatives de récupération du statut d'un job |
| wait_job_complete_retries_interval | 5s | Intervalle par défaut entre 2 tentatives de récupération du statut d'un job |
| pause_manual_instruction_interval | 15s | Délai d'inactivité de l'utilisateur après lequel une consigne manuelle est mise en pause |

**Exemples**

1. Rundeck est défaillant. Le moteur `remediation` essaie de se connecter à Rundeck. Après le délai `http_timeout`, la requête est considérée en échec.
1. Le moteur `remédiation` émet une requête vers Rundeck pour déclencher un job. Rundeck renvoie une erreur 500. Le moteur tente de déclencher le job `launch_job_retries_amount` fois toutes les `launch_job_retries_interval`.  
1. Le moteur `remediation` récupère le statut d'un job Rundeck. Rundeck renvoie un statut **running**. Le moteur répète cette requête `wait_job_complete_retries_amount` fois toutes les `wait_job_complete_retries_interval`.
1. Un utilisateur exécute une consigne manuelle, il ferme son navigateur. Après `pause_manual_instruction_interval`, la consigne est mise en pause

### Section [HealthCheck]

| Attribut            | Exemple de valeur  | Description                           |
| :------------------ | :------------------| :------------------------------------ |
| update_interval     | "10s"              | Intervalle de mise à jour des informations de HealthCheck | 

## Références

Pour référence, vous pouvez aussi consulter les fichiers d'origine complets,
avec les valeurs par défaut et des exemples commentés (en anglais) dans les
sources de Canopsis :

- [`canopsis-community.toml`][canopsis-community-toml]
- [`canopsis-pro.toml`][canopsis-pro-toml]

Prenez soin de sélectionner la branche `release-YY.MM` en fonction de la
version majeure de Canopsis qui vous concerne, la branche `develop`
représentant la prochaine version de Canopsis.

[canopsis-community-toml]: https://git.canopsis.net/canopsis/canopsis-community/-/blob/develop/community/go-engines-community/cmd/canopsis-reconfigure/canopsis-community.toml
[canopsis-pro-toml]: https://git.canopsis.net/canopsis/canopsis-community/-/blob/develop/community/go-engines-community/cmd/canopsis-reconfigure/canopsis-pro.toml
