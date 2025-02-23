# Mise à jour de Canopsis

Cette procédure décrit la mise à jour d'une instance mono-nœud de Canopsis.

Lisez l'ensemble de ce document avant de procéder à toute manipulation.

!!! important
    Les environnements n'ayant pas été installés en suivant l'une des [méthodes d'installation officielles de Canopsis](../installation/index.md#methodes-dinstallation-de-canopsis), notamment les environnements de type Docker Swarm ou en paquets multi-nœuds, ne sont pas pris en charge par cette procédure.

## Procédure de mise à jour

Vous devez tout d'abord lire **chacune** des [notes de version](../../index.md#notes-de-version) publiée entre votre version actuelle et celle que vous ciblez. [En savoir plus sur les numéros de version de Canopsis](numeros-version-canopsis.md).

Par exemple, si vous effectuez une mise à jour de Canopsis 24.04 à 25.04, vous devez :

*  consulter et appliquer toute procédure donnée dans les notes de version de Canopsis 24.04,
*  puis celles de Canopsis 24.10,
*  puis celles de Canopsis 25.04,
*  puis suivre le reste de cette procédure, selon votre méthode d'installation (paquets RPM, Docker Compose, charts Helm pour Kubernetes).

Si vous bénéficiez d'un développement spécifique (modules ou add-ons ayant été spécifiquement développés pour votre installation), assurez-vous de suivre toute procédure complémentaire vous ayant été communiquée.

!!! attention
    La mise à jour causera une **interruption de service** de Canopsis et des composants qui lui sont associés, durant son déroulement.

    Vous pouvez notamment utiliser la fonctionnalité de [diffusion de messages](../../guide-utilisation/menu-administration/diffusion-de-messages.md) afin de prévenir vos utilisateurs en amont.

### Mise à jour en installation par paquets RHEL

Les commandes suivantes doivent être réalisées avec l'utilisateur `root`.

Appliquez la mise à jour des paquets Canopsis :

```sh
dnf makecache
dnf --disablerepo="*" --enablerepo="canopsis*" update
```

Après avoir pris en compte toute éventuelle remarque des notes de version au sujet de paramètres dans les fichiers de type `canopsis.toml`, appliquez les changements de configuration en fonction de votre édition de Canopsis (Community ou Pro) :

!!! attention
    Dans le cas où des configurations spécifiques sont appliqués au travers d'un fichier `/opt/canopsis/etc/conf.d/canopsis-override.toml`. Il est nécessaire d'ajouter l'argument `-override /opt/canopsis/etc/conf.d/canopsis-override.toml` à la fin de votre commande `canopsis-reconfigure`.

=== "Canopsis Community"

    ```sh
    set -o allexport ; source /opt/canopsis/etc/go-engines-vars.conf
    /opt/canopsis/bin/canopsis-reconfigure -edition community -migrate-postgres=true -migrate-mongo=true -migrate-tech-postgres
    systemctl restart canopsis-engine-go@engine-action canopsis-engine-go@engine-axe canopsis-engine-go@engine-che.service canopsis-engine-go@engine-fifo.service canopsis-engine-go@engine-pbehavior.service canopsis-engine-go@engine-service.service canopsis-service@canopsis-api.service
    ```

=== "Canopsis Pro"

    ```sh
    set -o allexport ; source /opt/canopsis/etc/go-engines-vars.conf
    /opt/canopsis/bin/canopsis-reconfigure -edition pro -migrate-postgres=true -migrate-mongo=true -migrate-tech-postgres
    systemctl restart canopsis-engine-go@engine-action canopsis-engine-go@engine-axe canopsis-engine-go@engine-che.service canopsis-engine-go@engine-correlation.service canopsis-engine-go@engine-dynamic-infos.service canopsis-engine-go@engine-fifo.service canopsis-engine-go@engine-pbehavior.service canopsis-engine-go@engine-service.service canopsis-service@canopsis-api.service canopsis-engine-go@engine-remediation canopsis-engine-go@engine-webhook
    ```

Ne pas oublier d'appliquer toute éventuelle procédure supplémentaire décrite dans chacune des [notes de version](../../index.md#notes-de-version) qui vous concerne.

### Mise à jour en environnement Docker Compose

Après avoir suivi les notes de version, resynchronisez l'ensemble de vos fichiers Docker Compose avec les fichiers de référence correspondant à la version voulue :

* Canopsis Community :  
  <https://git.canopsis.net/canopsis/canopsis-community/-/releases>
* Canopsis Pro (autorisation nécessaire) :  
  <https://git.canopsis.net/canopsis/canopsis-pro/-/releases>

Si vous aviez surchargé des paramètres dans `canopsis-override.toml`, vous
souhaiterez certainement conserver votre propre fichier.

Si vous aviez d'autres personnalisations par rapport au modèle de configuration
Docker Compose des releases Canopsis, il vous appartient naturellement
d'étudier toute différence et de gérer la fusion des modifications.

Une fois les éventuelles différences étudiées et résolues, exécutez la commande suivante :

```sh
docker compose up -d
```

Ne pas oublier d'appliquer toute éventuelle procédure supplémentaire décrite dans chacune des [notes de version](../../index.md#notes-de-version) qui vous concerne, y commpris toute éventuelle remarque des notes de version au sujet du fichier `canopsis.toml`.

### Mise à jour en environnement K8S avec Helm

Après avoir suivi les notes de version, récupérez l'éventuelle nouvelle version
du chart depuis les dépôts, en ciblant la dernière version de la branche
majeure correspondant à la version majeure Canopsis souhaitée.

Si le chart n'a pas été mis à jour mais que vous souhaitez juste utiliser des
images de conteneurs Canopsis plus récentes (nouvelle version mineure),
surchargez simplement la valeur `image.tag` dans votre propre fichier de
*values*.

Dans tous les cas, vous terminerez en invoquant la commande `helm upgrade` pour
propager la mise à jour des configurations dans votre cluster Kubernetes.
Puis, après avoir attendu la bonne fin d'exécution du job `reconfigure`, vous
déclencherez le redémarrage de tous les éléments applicatifs (moteurs et
services Canopsis) via leurs `Deployments`.

Exemples :

```console
$ helm upgrade my-release charts/canopsis-pro -f my-values.yaml
$ kubectl rollout restart deploy -l app.kubernetes.io/instance=my-release
```

## Après la mise à jour

Durant le temps de coupure des services Canopsis, RabbitMQ se sera chargé de mettre en attente vos [évènements](../../guide-utilisation/vocabulaire/index.md#evenement). Ils seront alors « dépilés » et traités normalement par les moteurs Canopsis, dès leur redémarrage.

Cette accumulation d'évènements en attente peut, néanmoins, provoquer une latence des traitements, ou une augmentation de la consommation des ressources, en raison du rattrapage à effectuer. Cette incidence reste temporaire. Nous vous conseillons de [surveiller l'interface d'administration de RabbitMQ](../../guide-de-depannage/rabbitmq-webui/index.md) juste avant, durant et après la mise à jour, afin de mesurer l'état de « retour à la normale » de votre plateforme lors d'une période de maintenance de l'outil.

En revanche, tout appel fait aux API Canopsis durant cette période de maintenance n'aura pas été temporisé et devra donc être renouvelé s'il a échoué. Suivant les connecteurs ou configurations des sources envoyant ces évènements à destination de l'API HTTP de Canopsis, il peut donc exister un risque de perte d'évènements.

Une fois que le service est rétabli, vous pouvez vous connecter à nouveau sur l'interface Canopsis pour valider que tout fonctionne correctement.
