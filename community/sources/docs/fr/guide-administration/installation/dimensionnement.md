# Dimensionnement Canopsis

Voici un dimensionnement minimal recommandé pour Canopsis. 

Éléments    | CPU     | RAM  | Disque | IPv4 | Nom de domaine | Certificats TLS |
------------|---------|------|--------|------|----------------|-----------------|
Canopsis    |  8 vCPU | 10Go | 50Go   | Nécessaire    | Nécessaire              | Nécessaire               |
MongoDB     |  4 vCPU |  8Go | 50Go   | Nécessaire    | Optionnel     |                 |
TimescaleDB |  4 vCPU |  8Go | 50Go   | Nécessaire    | Nécessaire              |                |

Ce dimensionnement est à titre indicatif, il doit être adapté en fonction des environnements dans lesquels Canopsis est déployé.

