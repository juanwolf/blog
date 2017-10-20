---
title:  "Comment dockerizer une application ?"
date:   2017-02-22
tags: ["ops"]
---

# Comment dockerizer une application

## Introduction

Comme vous avez pu le voir dans mon article précédent, docker est l'outil parfait pour assurer que votre application fonctionne sur toutes les plateformes. Maintenant le problème que l'on pourrait se poser est, comment faire pour dockerizer une application. Je vais vous donner mon retour après avoir perdu quelques cheveux sur cette techno. N'hésitez pas à contribuer à ce guide, je vous en serais plus que reconnaissant !

## Définir les besoins de l'application

Avant de commencer quoi que ce soit, vous devez vous demander de quoi votre application à besoin. Une base de données ? Un gestionnaire de tâches ? Etc... vous pouvez envisager de créer un container pour chacun.

Vous devez aussi identifier au niveau de votre application les outils que vous neccésiterez. Votre application a besoin de tourner comme un service dans votre container. Par exemple, si vous dockerizé une application django, vous devrez utiliser uwsgi.


## Définir une police de rétention

J'ai perdu pas mal de cheveux sur ce point là. Devez vous laisser la configuration dans le container ou utiliser des variables d'environnement ? Malheureusement je n'ai pas de réponses à ces questions. Par ailleurs, elles ont toutes leurs avantages et inconvénients. En voici la liste :

### Configuration avec des variables d'environnement

Plutôt que d'utiliser un fichier de configuration, vous pouvez changer votre configuration afin qu'elle utilise des variables d'environnement. Cependant cette méthode peut vous demandez pas mal d'effort mais peut s'avérer utile sur le long terme.

#### Plus

* Facilement configurable
* Facilement "scalable" verticallement et horizontalement

#### Moins

* Comportement peut varier en fonction des environnements. (Peut-être éviter avec des outils d'orchestration)
* Peut demander beaucoup d'effort sur le refacto de la configuration

### Configuration contenue dans le container

#### Pour

* Figée et donc sûr qu'en prod ou dev tout fonctionne de la même façon
* Facilement "Scalable"

#### Cons

* Presque impossible à configurer une fois l'image construite
* Nécessite de l'intégration continue pour changer la configuration lors d'une modification du code

### Configuration montée comme un volume

#### Pour
* Fonctionne comme la plupart des applications (donc gain de temps de mise en place)
* Facilement "scalable" verticalement

#### Contre
* Le fichier de configuration doit se trouver sur le docker host
* Le fichier doit être présent et le même sur tous les docker hosts (Galère pour "scale" horizontal)

### UFC - Que Choisir ?

Franchement, aucune idée. J'ai créé 3 applications et elles utilisent les 3 différentes manières, et je suis pas super content du résultat. Je suis tenté de dire que les variables d'environnement est le meilleur choix. Vous pouvez toujours vous servir d'un fichier de configuration et repérer les éléments pouvant varier d'un environnement à un autre et les définir comme variables d'environnement. C'est le meilleur compromis que j'ai pu trouver. Afin que vous soyez sûr que vos variable d'environnements soient définies de la même manière dans votre environnement, je vous invite à utiliser un outil de déploiement ou d'orchestration. Personnellement j'utilise Ansible mais vous êtes libres de choisir celui qui vous convient.

Ce choix peut aussi varier en fonction de vos besoins. Besoin d'un résultat rapide mais crade ? Config contenue dans le container. Pas besoin de changer la configuration ? Configuration dans le container. Grosse application qui fonctionne comme de par magie avec un fichier de configuration -> Configuration montée comme un volume. Vous êtes en transe et pret à détruire des montagnes d'un seul coup de poing ? -> variable d'environnement.

## Installer le minimum

Quand vous utilisez docker, vous saviez que vous alliez changer quelques trucs. C'est là que vous allez suer un petit peu. Votre application a besoin d'être la plus petite possible. Plus votre application est petite et plus vous gagnerez à utiliser docker. Imaginons que votre application contienne une application web et un gestionnaire de tâche asynchrone et que la première version dockerizée contienne les deux éléments. Vous allez avoir du mal à répartir la charge entre ces deux éléments. Imaginons que la web app n'est quasiment jamais en surcharge mais que le gestionnaire de tâche l'est. Vous êtes obligé de redéployer un nouveau container avec la webapp qui sera jamais utilisée pour répondre à la charge du nombre de tâches... Donc pensez à découpler au maximum vos dépendences dans votre application.

## Définir une police de rétention

C'est le point qui m'a fait le plus redéployer des containers... Il est important de réfléchir à chaque élément de votre application qui pourait grossir dans votre container (car si vous avez besoin de détruire votre container vous perdrez toutes les données dans celui-ci). Clairement votre container ne doit pas grossir. Vous devez donc définir dans votre dockerfile tout élément qui pourrait modifier durant l'exécution de votre application.

## Exemple

Il serait dommage de vous laisser partir dockerizer la planète sans vous laisser un petit exemple.
J'essaierai d'ajouter des exemples dans le futur quand j'aurai expérimenté plus de déploiement de containers avec différentes techno.

### Django

Comme django a besoin d'un fichier de configuration pour fonctionner, vous pouvez envisager de commencer par monter votre configuration comme un volume, ce sera le plus simple et le plus rapide. De plus, vous avez sûrement configurer les logs pour logger dans un fichier. Pensez à définir un volume pour ce fichier. De même pour les fichiers statiques ou les médias. Il est nécessaire de définir un volume pour vos médias. Cependant pour vos statiques cela va dépendre si vous voulez les servir avec votre application ou avec un proxy. Pour la première solution je vous invite à regarder au niveau de whitenoise qui fait très bien le taf. Pour la seconde vous devez definir un volume pour votre STATIC_ROOT ET appeler collectstatic. Normalement vous devrez retrouver vos statics sur le docker host et pouvez les servir avec nginx ou apache. :)

N'oubliez pas que votre application a besoin de tourner comme un service. Vous devez donc penser à utliser uwsgi (ou autre) dans votre container. (Vous pouvez toujours utiliser la commande runserver pour votre environnement de dévelopemment, tant que ça reste QUE du développement)

Et c'est tout :). Je vous invite à regarder comment django-cookiecutter gère les environnements docker, c'est plutôt bien foutu.

## Conclusion

Ce fut bref mais ce sont les points que j'aurai aimé connaître avant de me lancer dans la dockerization d'applications. N'hésitez pas à contribuer à ce petit guide avec vos retours ou vos exemples ! Sur ce, dockerizez bien ! Ciao !
