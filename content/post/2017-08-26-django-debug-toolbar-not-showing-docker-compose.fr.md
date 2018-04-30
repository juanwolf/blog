
---
title: La django debug toolbar n’apparaît pas avec docker-compose
date: 2017-08-26
tags: ["Django", "docker"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Salut tout le monde, après quelques minutes de frustration, la toolbar
est finalement apparue. Le problème est survenu à chaque fois que
j'utilisais docker-compose pour mon environnement de développement et ça
commençait à me saouler sérieusement de ne pas mettre la main sur la
source du problème. Et c’était plutôt simple à résoudre...

![](/post_preview/20170826_150352_docker-django.png)

## Le problème

Si vous avez commencé à dockerizer votre application django, je suis sûr
que vous avez forcément rencontré quelques problèmes que vous avez
facilement résolu avec plusieurs requêtes Google. Il y en a quelques uns
par contre qui ont du vous demander un peu plus de recherche, et ça a
été mon cas aujourd'hui. J'ai été un peu surpris que personne n'ait
montré du doigt ce problème mais bon.

Après peut-être que peu de personnes utilisent docker-compose, je ne
sais pas :/ M'enfin bref\!

## Le fix

Pour ceux qui veulent une réponse
    rapide:

```bash
# Utilisez cette commande dans le dossier où vous utilisez docker-compose d'habitude
docker network list | grep tools | sed -r 's/^([0-9a-z]+).*$/\1/' | xargs docker network inspect  --format {{ range .IPAM.Config }}{{ .Gateway }}{{ end }}
# Ajoutez l'IP dans votre variable ALLOWED_IPS et la djdt devrait apparaitre :)
```

Ok c’était pour la réponse rapide mais bon cette commande magique m'a
pris tout de même pas mal de temps.

## Pourquoi ?

Docker-compose cache la masse de trucs. Point final.

Des bisous.

Nan sérieusement docker-compose cache pas mal de trucs. Ce que vous
n'avez peut-être pas fait attention, c'est que docker-compose va créer
un réseau spécifique pour votre application. Il va avoir comme nom
`mon_project_default` ou quelque chose comme ça. Donc votre app a un
sous réseau qui lui est propre avec une IP spécifique etc... Cependant,
comment votre localhost va être capable d’accéder a ce subnet ? Grace au
mode réseau bridge, cette utilisation de sous réseau est complètement
invisible car lorsque vous allez créer un réseau avec le mode bridge,
docker va monter une nouvelle interface réseau sur votre machine et
celle ci va être utilisée pour requêter votre application.

Exemple:

Je développe actuellement
[Gringotts](https://github.com/juanwolf/gringotts). Je vais dans le
dossier `docker` et lance `docker-compose up`. Cette commande va créer
un réseau nomme `docker_default` (car je l'ai lancée dans le dossier
docker). On peut vérifier qu'un réseau a été créé en utilisant `docker
network ls`. Maintenant on peut utiliser `ip addr` et on devrait voir
une nouvelle interface est apparue sur la machine. Dans mon elle
ressemble a
    ça:

```bash
4: br-57a6b1bdcf1a: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:98:71:d9:10 brd ff:ff:ff:ff:ff:ff
    inet 172.19.0.1/16 scope global br-57a6b1bdcf1a
       valid_lft forever preferred_lft forever
    inet6 fe80::42:98ff:fe71:d910/64 scope link
       valid_lft forever preferred_lft forever
```

La bonne interface va être le reseau\_id (que vous pouvez connaître
grâce à `docker network ls`) préfixé par br\_. Et l'IP affichée est
celle que votre machine va utiliser pour accéder au sous réseau de votre
application. Donc maintenant que je vous ai explique tout ça, je viens
de remarquer que j'aurai pu faire la réponse rapide différemment mais
bon ça marche aussi =p. Donc si vous avez tout compris, quand vous
accédez à votre application vous n'allez pas avoir la fameuse IP
127.0.0.1 mais celle attribuée à l'interface créée qui va être quelque
chose comme 172.xx.0.1, c'est donc pour ça que la django debug toolbar
ne s'affichait pas.

Sur ce codez bien, Ciao\!

