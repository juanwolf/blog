---
title:  "La django debug toolbar n'apparaît pas avec docker-compose"
date:   2017-08-26
tags: ["programmation"]
---

# La django debug toolbar n'apparait pas avec docker-compose

## Introduction

Salut tout le monde, après quelques minutes de frustration, j'ai finalement reussi a avoir la toolbar s'afficher. Le probleme est apparu a chaque fois que j'utilisais docker-compose pour mon environemment de developement et ca commencait a me saouler serieusement de ne pas mettre la main sur la source du probleme. Et c'etait plutot simple a resoudre...

![The title](https://theadultswimsquad.files.wordpress.com/2017/02/ep-379-3.jpg?w=679&h=381)


## Le probleme

Si vous avez commence a dockerizer votre application django, je suis sur que vous avez forcement rencontre quelques problemes que vous avez facilement resolu avec plusieurs requetes sur google. Il y en a quelques uns par contre qui ont du vous demander un peu plus de recherche, et ca a ete mon cas aujourd'hui. J'ai ete un peu surpris que personne n'ait montre du doigt ce probleme mais bon =p.

Apres peut-être  que peu de personnes utilisent docker-compose, je ne sais pas :/ M'enfin bref!

## Le fix

Pour ceux qui veulent une réponse rapide:

```
# Utilisez cette commande dans le dossier où vous utilisez docker-compose d'habitude
docker network list | grep ${PWD##*/} | sed -r 's/^([0-9a-z]+).*$/\1/' | xargs docker network inspect  --format "{{ range .IPAM.Config }}{{ .Gateway }}{{ end }}"
# Ajoutez l'IP dans votre variable ALLOWED_IPS et la djdt devrait apparaitre :)
```

Ok c'etait pour la reponse rapide mais bon cette commande magique m'a pris tout de meme pas mal de temps.

## Pourquoi ?

Docker-compose cache la masse de trucs. Point final.

Des bisous.

Nan sérieusement docker-compose cache pas mal de trucs. Ce que vous n'avez peut-etre pas fait attention, c'est que docker-compose va creer un reseau specifique pour votre application. Il va avoir comme nom `mon_project_default` ou quelque chose comme ca. Donc votre app a un sous reseau qui lui est propre avec une IP specifique etc... Cependant, comment votre localhost va etre capable d'acceder a ce subnet ? Grace au mode reseau "bridge", cette utilisation de sous reseau est completement invisible car lorsque vous allez creer un reseau avec le mode bridge, docker va monter une nouvelle interface reseau sur votre machine qui va transferer vos requetes sur le bon sous reseau.

Exemple:

Je developpe actuellement [Gringotts](https://github.com/juanwolf/gringotts). Je vais dans le dossier `docker` et lance `docker-compose up`. Cette commande va creer un reseau nomme `docker_default` (car je l'ai lancee dans le dossier docker). On peut verifier qu'un reseau a ete cree en utilisant `docker network ls`. Maintenant on peut utiliser `ip addr` et on devrait voir une nouvelle interface est apparue sur la machine. Dans mon elle ressemble a ca:

```
4: br-57a6b1bdcf1a: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:98:71:d9:10 brd ff:ff:ff:ff:ff:ff
    inet 172.19.0.1/16 scope global br-57a6b1bdcf1a
       valid_lft forever preferred_lft forever
    inet6 fe80::42:98ff:fe71:d910/64 scope link
       valid_lft forever preferred_lft forever
```

La bonne interface va etre le reseau_id (que vous pouvez connaitre grace a `docker network ls`) prefixe par "br_". Et l'IP affichee est celle que votre machine va utiliser pour acceder au sous reseau de votre application.
Donc maintenant que je vous ai explique tout ca, je viens de remarquer que j'aurai pu faire la reponse rapide differemment mais bon ca marche aussi =p. Donc si vous avez tout compris, quand vous accedez a votre application vous n'allez pas avoir la fameuse IP 127.0.0.1 mais celle attribuee a l'interface creee qui va etre quelque chose comme 172.xx.0.1, c'est donc pour ca que la django debug toolbar ne s'affichait pas.


Sur ce codez bien, Ciao!

