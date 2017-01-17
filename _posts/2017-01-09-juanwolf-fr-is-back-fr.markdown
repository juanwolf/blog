---
layout: post
title:  "juanwolf.fr est de retour !"
date:   2017-01-09 22:47:01 +0000
categories: other
---

# juanwolf.fr est de retour !

Hey! Ça fait un bail ! Content de vous (re)voir :)

Pas mal de choses se sont passées depuis que j'ai publié un article et ce n'est pas pour rien !

## EXCUSES PERSONNELLES DE POURQUOI JE N'AI PAS PU TRAVAILLER SUR MON TEMPS LIBRE.

Beaucoup de changement dans ma vie perso. J'ai changé de sex, d'identité, infiltré la NSA, et je ne peux pas vous dire plus pour le moment (Secret défense).
Non sérieusement, j'ai passé la plupart de mon temps à chercher un travail dans la fabuleuse ville de Londres.
Du coup,  je me suis un peu calmé niveau geekerie personnelle. Mais je suis de retour !
J'ai aussi upgradé juanwolf.fr à sa version 2.0

## juanwolf.fr  2.0

Mon bon vieux site internet, juanwolf.fr, tournait sur une vieille machine Kimsufi (2 coeur, 2GB de RAM) et commençait a me limiter dans ce que je voulais faire. Il était donc temps de tourner la page.

Je ne sais pas si vous avez remarqué mais l'adresse IP pour juanwolf.fr a changée (j'ai changé de serveur pour une version bien plus véner). Pour préparer la migration, j'ai décidé de passer toutes les applications sous docker. Ca m'a pris plutôt du temps pour tout faire mais le résultat y est. De plus, j'ai changé de serveur d'intégration continue, gitlab-ci au lieu de Jenkins (Je trouve gitlab-ci bien plus pratique pour tout ce qui est intégration docker, mais je vous en parlerai dans un prochain article)

De plus, je voulais supprimer des projets qui me semblaient superflu et garder une logique sur tout le site. J'ai désinstallé le routeur écrit en go, pour ajouter une app django. Plus de /en ou /fr.

Pour resumer:

* Nouveaux serveurs
* Nouvelle application django pour l'index, le CV et la page a propos - [juanwolf.fr_static](https://github.com/juanwolf/juanwolf.fr_static)
* Adieu bon vieux proxy - ["language-router"](https://github.com/juanwolf/language-router)
* Installation de RocketChat le temps que je construise une app avec django-channels
* Infrastructure provisionnée avec ansible - [playbooks](https://github.com/juanwolf/playbooks)

## E-Sport

J'ai commencé à jouer à Rocket League (football + voitures, plus con tu meures) assez fréquemment. (D'après Steam, j'aurai dépassé 180h en jeu, oups). Son gameplay est plutôt simple mais demande beaucoup de patience afin de maitriser complètement l'engin, je pense que j'écrirai un article a propos de ce jeu et de l'esport en general.

## Conclusion (ou ou lire si vous êtes presses)

Donc juanwolf.fr 2.0 c'est deux applications django tournant avec docker, une application de chat (Rocketchat) utilisant meteor [chat.juanwolf.fr](https://chat.juanwolf.fr)

Dans quelques temps, je vais recommencer a écrire des articles sur ces sujets :

* Docker (and/or Django + Docker)
* Ansible
* Gitlab CI
* E-Sport / Rocket League
* Doom

Dans un certains temps

* Django Channels (Apres que j'ai code le fameux chat donc ca risque de pas être tout de suite)


Bon allez, moi je suis parti. Sur ce, codez-bien. Et bonne année.
