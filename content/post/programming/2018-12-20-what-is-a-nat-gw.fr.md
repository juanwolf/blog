---
title: "Qu'est ce qu'une NAT Gateway?"
date: 2018-12-20T08:27:38+01:00
tags: ["ops", "cloud", "network"]
categories: ["Programmation"]

---

Créer des plateformes et des nouvelles architectures dans le Cloud peut-être intimidant. Surtout la première fois qu'on se lance dans le domaine de l'infrastructure. Certains composants apparaissent comme des boites noires ayant une place mystérieuse mais nécessaire dans l'infrastructure.

Aujourd'hui je veux vous parler des NAT Gateways. Qu'est ce qu'une NAT? Comment cela fonctionne, etc...

## NAT, Kezako ?

Commençons tranquillement. N.A.T. Que cela veut-il dire? Si l'on écoute papy Wikipedia:

> Network address translation (NAT) is a method of remapping one IP address space into another by modifying network address information in the IP header of packets while they are in transit across a traffic routing device.[1] The technique was originally used as a shortcut to avoid the need to readdress every host when a network was moved. It has become a popular and essential tool in conserving global address space in the face of IPv4 address exhaustion. One Internet-routable IP address of a NAT gateway can be used for an entire private network.
Sources: [Wikipedia](https://en.wikipedia.org/wiki/Network_address_translation)

Ok papy, tu as un peu spoilé tout le monde en expliquant tout d'un coup comme ça. Pragmatique, ce papy. Donc comme Wikipedia nous l'a énoncé, les NATs sont des composants très fréquents dans les infrastructures. Tout particulièrement dans des environnements où les IPv4 sont comptées au compte goute tel que les infrastructures partagées (ou cloud). De plus les NATs vont nous permettre de protéger nos composants de trafic externe en isolant nos VMs dans des sous-réseaux seulement accessible dans notre projet/Virtual Private Cloud/isolation créé rien que pour nous!

Généralement, quand vous commencez à toucher à du :cloud:, vous allez créer une instance, choisir l'os, définir une taille de disque et c'est tout. Votre VM va être créé avec une IP interne dans votre VPC/isolation/projet **et** une IP publique. Grâce à cette IP publique, vous allez pouvoir SSH et vous connecter au service que vous allez faire tourner sur cette machine. Cependant, vous permettez quiconque sur les internets d'y accéder aussi (si vous ne maitrisez pas les outils de sécurité mis à votre disposition). Votre infrastructure va murir (et vous aussi accessoirement) et vous allez vouloir bouger toutes vos instances dans des sous-réseaux privés. Mais comment accéder à des services externes si nos instances sont privées? C'est que rentre en jeu notre protagoniste, la NAT Gateway :tada:.

Nous sommes chanceux, les clouds sont des technologies maintenant très matures et les NATs peuvent être créer/provisionnées en quelques clics et managées par les clouds providers. Que demander de plus? Mais au final, comment une NAT fonctionne?

Disons que vous avez rempli le formulaire de création de la NAT gateway. Vous terminez avec une infrastructure telle que celle-ci:

![Schema of an architecture in the cloud with one vpc, one private subnet with four instances connected to the nat gateway which is inside a public subnet and relay the traffic to the internet](/post_content/2018-06-25/nat_vpc.svg)

Comme vous pouvez le voir sur le schéma, chaque connexion extérieure passe par la NAT gateway **et** revient à l'instance correspondante. Ma question est: Comment? :wNotre instance n'a clairement pas d'IP publique, comment le serveur externe puisse renvoyer un paquet sur cette instance. Cela n'a pas de sens. Qu'est ce donc que cette mascarade ? (Trop fier de cette blague)

## Comment la NAT Gateway fonctionne t-elle ?

Comme Wikipedia nous l'a dit précédemment, la NAT va modifier le paquet entrant en injectant son adresse IP et un port aléatoire afin que le serveur distant lui réponde à elle et seulement elle.

Mais comment la NAT va savoir où renvoyer le paquet quand elle le reçoit?

![Nat packet translation](/post_content/2018-06-25/nat_anim.svg)

Voilà ce qu'il se passe:

Votre paquet est envoyé à la NAT Gateway. La NAT va sauvegarder la configuration du paquet soit son IP d'origine, port d'origine, ip de destination, port de destination au sein de ce qu'on appelle une '_NAT table_'. Remplacer les information d'origine par son IP et un port qu'elle va allouer aléatoirement et envoyer le paquet. Quand elle reçoit le paquet en retour, elle va vérifier sur quel port, lire la NAT table, retrouver quelles étaient les informations du paquet au départ et va retourner le paquet a l'instance en question.

Etapes par étapes, cela nous donne:

1. Votre paquet est envoyé à la nat
2. La NAT enregistre l'IP et le port de la source, l'IP et le port de destination et alloue un port aléatoire dans la _NAT Table_
3. Elle va modifier l'IP et le port de la source afin que le serveur externe lui réponde
4. Elle Envoie le paquet à la destination
5. Le serveur externe reçoit le paquet et envoie la réponse à la NAT (il n'a aucune idée de la provenance original du paquet)
6. La NAT reçoit le paquet, vérifie sa _NAT Table_, retrouve l'instance originale avec le port qu'elle a ouvert spécifiquement pour ce paquet.
7. Change la destination du paquet pour l'IP et le port original de la requête, envoie le paquet
8. L'instance a reçu son paquet

## Limitations

Utiliser une NAT nous permet de garder nos instances privées même en ayant accès à des resources externes à notre VPC. Mais tout système vient avec ses limitations. La NAT doit avoir au moins une faiblesse ! Lors de la traduction de notre paquet, la NAT alloue un port aléatoire pour notre paquet. Le problème est que... Disons que vous avez 20 serveurs qui tournent 20 processus qui envoient 10 requêtes par secondes à des serveurs extérieurs et ces requêtes ont un timeout de... disons 15 secondes. Tous ces serveurs extérieurs sont indisponibles pendant une soirée. En moins de 10 secondes, votre NAT va paraitre comme instable ou défaillante et va être incapable de servir de nouvelles requêtes externes.

Pourquoi ça?

Vous venez d'utiliser tous les [ports éphémères](https://www.ncftp.com/ncftpd/doc/misc/ephemeral_ports.html) disponibles avec une seule NAT. :boom:

Faisons un peu de mathématiques. On crée 20 * 20 * 10 requêtes par seconde. Soit 4000 requête/sec. La NAT va allouer un port éphémère a chacune de ses requêtes. Première seconde: 4000 ports; Deuxième 8000 ports d'ouvert, etc...

Le nombre de ports éphémères disponibles va dépendre de votre OS mais d'après [Wikipedia](https://en.wikipedia.org/wiki/Ephemeral_port#Range):

> Many Linux kernels use the port range 32768 to 61000. (The effective range is accessible via the /proc file system at node /proc/sys/net/ipv4/ip_local_port_range)

Donc approximativement 28232 ports je dirai à vue de nez. Ceci est notre nombre de maximums de ports que l'on peut ouvrir sur une seule et même instance. Bravo, nous venons tout juste d'exploser les limites de notre NAT. Heureusement que ceci n'était qu'une mise en situation :sweat_smile:.

## Conclusion

Et voilà! Nous venons de voir comment fonctionne une NAT Gateway et ses contraintes. C'est un composant essentiel dans notre infrastructure dans le Cloud. Pas de Nat, pas de trafic externe :smile:. Sur ce codez bien, Ciao!

### Post Scriptum
J'ai essayé d'en implémenter une [ici](https://github.com/juanwolf/toran). Je voulais à tout prix terminer la première version de ce projet avant de sortir cet article. Malheureusement, je n'ai pas vu le bout de ce projet (commencé en Juin, tout de même) et je n'ai pas eu le temps (et le courage sur la fin) de le terminer. Qui sait? Peut être que ces lignes de code vous aideront un jour.
