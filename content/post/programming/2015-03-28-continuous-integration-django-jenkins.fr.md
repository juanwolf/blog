
---
title: Intégration continue d'un projet django avec jenkins
date: 2015-03-28
tags: ["Django", "Jenkins"]
categories: ["Programmation"]

slug: integration-continue-django-jenkins
aliases:
  - /programmation/integration-continue-django-jenkins

draft: false
author: "Jean-Loup Adde"
---

Yo ! Aujourd'hui on va voir comment faire de l'intégration continue
pour un projet django. L'avantage de cette pratique, c'est que vous
n'aurez plus jamais peur de mettre en production et ça je peux vous dire
que c'est un vrai plaisir. Imaginez un monde dans lequel la feature
toute fraichement codée apparaît sur votre app django déployée sur un
serveur distant en quelques minutes. Incroyable hein? Voyons comment
mettre ça en place.

![](/post_preview/20150328_183333_jenkins-django.png)

## Jenkins

Jenkins est un serveur d'intégration continue. Il va nous permettre
d'automatiser différentes tâches comme par exemple, récupération des
sources toutes les 10 minutes, ou envoyer un mail aux gens qui commit
des changements qui ne passent pas les tests (mon préféré), etc... Nous,
dans notre cas, on veut utiliser jenkins pour nous permettre de déployer
automatiquement chaque nouvelle fonctionnalité sur notre serveur. Sur le
papier ça a l'air simple, en pratique, pas forcément.

Mais ne vous inquiétez pas cet article va tout vous décortiquez et vous
aurez installé un jenkins paramétré en moins de 2 (ans).

## INSTALLATION DE JENKINS

Pour ArchLinux :

    sudo pacman -S jenkins
    sudo systemctl start jenkins

Pour Ubuntu, je suppose un
    simple

    wget -q -O - https://jenkins-ci.org/debian/jenkins-ci.org.key | sudo apt-key add -
    sudo sh -c 'echo deb http://pkg.jenkins-ci.org/debian binary/ > /etc/apt/sources.list.d/jenkins.list'
    sudo apt-get update
    sudo apt-get install jenkins

Normalement, jenkins devrait tourner sur votre machine. Son port par
défaut est le 8090. Donc essayez de vous rendre sur
[http://localhost:8090/](http://localhost:8090). Vous devriez tomber sur
la page d'accueil de Jenkins. (Trop facile ce tuto)

## Installation de l'environnement virtuel

Je vous conseille fortement de créer un environnement virtuel pour votre
projet, ça vous permettra d'avoir une installation propre de python pour
votre projet, donc aucun conflit de librairie où je ne sais quoi, VOUS
N'AVEZ PLUS D'EXCUSE POUR L'ÉCHEC. Et dans ce tuto on en utilisera un,
donc autant en profiter \!

Pour l'installer, nous allons faire un :

    sudo pacman -S python-virtualenv

Ensuite nous pourrons créer un simple environnement virtuel pour python
avec la commande :

    virtualenv /le/path/ou/je/stocke/mon/env/

Et on l'utilisera avec la commande :

    source /le/path/ou/je/stocke/mon/env/bin/activate

Easy non ?

## Installation d'UWSGI

Pour lancer notre application Django, on va avoir besoin de UWSGI.

Pour ça, veuillez installer uwsgi et uwsgi-python (le module python
d'UWSGI) :

```
sudo pacman -S uwsgi
sudo pacman -S uwsgi-plugin-python

```

Pour Debian/Ubuntu :

    sudo apt-get install uwsgi
    sudo apt-get install uwsgi-plugin-python

N'oubliez pas d'adapter en fonction de votre version de python. Dans mon
cas, mon application est codée en python3. Sinon on aurait :

    uwsgi-plugin-python2

Cependant le soucis de cette installation est qu'elle changera avec
votre système si vous le mettez à jour. Afin de ne pas avoir ce genre de
problème, je vous invite à installer uwsgi via pip ou en l'ajoutant dans
votre requirements.txt au sein de votre environnement virtuel et vous
n'aurez plus besoin du plugin \!

    pip install uwsgi

## Lancement de votre application django

Là, soit vous attendez de configurer jenkins pour qu'il téléverse
(j'adore ce mot) pour la première fois les sources de votre application
sur votre serveur, soit vous en avez déjà une copie de celle-ci. Ici,
nous partons du pricipe qu'on a déjà les sources sur le serveur.

Pour lancer aisément votre application, on va devoir créer une .ini file
afin de configurer UWSGI. Prenons l'exemple pour une application appelée
mon-amour-pour-jean-travolta

    #mon-amour-pour-jean-travolta.ini file
    [uwsgi]
    # On spécifie le type d'appli
    plugins = python
    # Django-related settings
    # La racine du projet django (là où se situe le manage.py)
    chdir           = /chemin/vers/lapplication/mon-amour-pour-jean-travolta/
    # Le fichier wsgi
    module          = mon-amour-pour-jean-travolta.wsgi:application
    # Chemin absolu vers l'environnement virtuel
    home            = /chemin/vers/lapplication/mon-amour-pour-jean-travolta/venv-python3-blog/

    # process-related settings
    # master
    master          = true
    # Le nombre de worker pour votre application
    processes       = 10
    # La socket unix permettant la communication entre votre application django et votre serveur nginx ou apache
    socket          = /chemin/vers/lapplication/mon-amour-pour-jean-travolta/juanwolfs-blog.sock
    # Si jamais vous voulez exposer votre application directement décommenter la ligne du dessous
    # socket = 127.0.0.1:8888
    # Les permissions de la socket
    chmod-socket    = 644
    # l'utilisateur du processus
    uid = jean-travolta
    # le groupe du processus
    gid = jean-travolta
    # le propriètaire de la socket
    chown_socket = jean-travolta
    # Cree un service pour votre application (laissez cette valeur en commentaire le temps que vous n'êtes pas sur que tout fonctionne)
    # daemonize = /var/log/uwsgi/juanwolfs-blog.log
    # Nettoie l'environnement quand on quite uwsgi
    vacuum          = true

Si tout est bien paramétré, votre application django devrait se lancer
avec la commande :

    uwsgi --ini mon-amour-pour-jean-travolta.ini

Là, le petit soucis qu'on va avoir avec cette commande, c'est que nous
n'aurons aucun moyen de relancer l'application automatiquement et
facilement. Après plusieurs recherches infructueuses à ce sujet, j'ai
demandé l'avis à un expert. Voici son témoignage :

> [@Juan\_\_Wolf](https://twitter.com/Juan__Wolf) I tend to use
> touch-reload (<http://t.co/cGNueQDzad>), which allows you to specify a
> file that, when touched, restarts uwsgi.
>
> — Dominic Rodger (@dominicrodger)
> [March 17, 2015](https://twitter.com/dominicrodger/status/577809423420248064)

Merci encore [@dominicrodger](https://twitter.com/dominicrodger) pour le
coup de main \! Donc ajoutons l'option en question
    :

    uwsgi --ini mon-amour-pour-jean-travolta.ini --touch-reload mon-amour-pour-jean-travolta.ini
    # Et On lance cette commande pour recharger l'application
    touch mon-amour-pour-jean-travolta.ini

## Automatisation

Maintenant nous avons tous les outils nécessaires pour déployer notre
application django en un clic =). Il ne nous reste plus qu'à créer le
job qui executera tous les commandes que nous aurions executés si nous
avions déployés l'application à la main.

Le principe est simple : récupérer les sources, lancer l'environnement
virtuel, installer les dépendances de l'application, lancer les tests
(oui quand même), faire les changements au niveau de la BDD s'il y en a
(pour ça, j'utilise south), générer les fichiers de traductions,
relancer l'application, et sortir de l'environnement virtuel.

### Configuration de Jenkins

Pour récupérer les sources de votre application, elle devra être
disponible sur un dépôt de gestion de versions tel qu'un dépôt git ou
SVN (ou dropbox, lol je décone). Pour git, vous aurez besoin d'installer
un petit plugin. Allez dans la catégorie Administrer jenkins -\> Gestion
des plugins -\> Onglet Disponibles -\> GIT plugin. Si vous comptez
utiliser un autre serveur pour votre application vous aurez aussi besoin
du plugin Publish Over SSH.

### Création du job

Nous allons créer notre job, pour cela cliquez sur nouveau item. Pour
cet exemple je vais l'appeler
jean-travolta-l-integration-qui-envoie-du-pate, et faire un projet
free-style. Maintenant on va devoir configurer toutes les étapes que
nous avions préciser un peu plus haut. Commençons par la récupération
des sources.

#### Récupération des sources

Rendez-vous à la section gestion du code source. Plusieurs options
s'offrent à vous CVS, Subversion, Git (si vous avez installé le plugin).
Dans mon cas je vais selectionner Git et rentrer ces valeurs
:

![configuration](/post_content/2015-03-28/ea855e2f-12d3-4cc9-a6e6-cec855918cc0.png)


Notez que vous pouvez spécifier la branche à récupérer ce qui est très
pratique dans le cas où vous avez plusieurs environnement de
déploiement. Pour la dernière section, nous indiquons à jenkins de
scruter tout changement au niveau du dépôt git toutes les 2 minutes et
d'enclencher le build jenkins s'il y en a eu.

#### Construction

Maintenant, il ne nous reste plus qu'a faire tout le reste. On va donc
créer un petit script bash qui va nous executer tout ça. Le mien
ressemble à ça :

```
#!/bin/bash
virtualenv -q venv-python3-blog # Création de l'environnement virtuel s'il n'existe pas
source ./venv-python3-blog/bin/activate # Activation de l'environnement virtuel
pip install -r /requirements.txt # Installation des dépendances pour l'application
python /mon-amour-pour-jean-travolta/manage.py test blogengine # Lancement des test

python /mon-amour-pour-jean-travolta/manage.py schemamigration blogengine --auto --update # Chercher si changement en BDD (SOUTH)
python /mon-amour-pour-jean-travolta/manage.py migrate # Application des changements à la BDD
cd /mon-amour-pour-jean-travolta/;
python manage.py compilemessages # Génération des fichiers de traductions

touch /mon-amour-pour-jean-travolta.ini # On relance l'application

deactivate # On sort de l'environnement virtuel

```

Lancer le build, et normalement vous devriez avoir une sortie telle que
:

![Affichage](/post_content/2015-03-28/acc139bc-a4f3-4b53-82bf-3feea358f3a4.png)

Et voilà, on a notre plateforme d'intégration continue pour notre
application Django \! Cependant, concernant les tests, nous n'avons
aucun affichage mis à part le résultat de nos tests dans la console. Il
serait bien d'être en mesure d'avoir des indicateurs qualitatifs pour
notre application.

## BONUS - Création de rapports de tests

Là le but va être d'utiliser des outils qui vont nous évaluer la qualité
de notre code. Pour cela, vous aurez besoin de modifier un peu votre
projet et votre build jenkins. Premièrement, installez les plugins
JUnit, Cobertura ainsi que Violation au sein de Jenkins. Au niveau de
projet, vous aurez besoin d'ajouter à vos dépendances la librairie
django-jenkins et coverage qui va nous créer les rapports de tests \!
Donc dans votre requirements.txt ajoutez :

    django-jenkins==0.16.4
    coverage==3.7.1

Maintenant, on va modifier un peu votre build. Premièrement remplacez
dans votre script de build la directive de test afin qu'on utilise la
librairie django-jenkins :

    # python /mon-amour-pour-jean-travolta/manage.py test blogengine
    # Nouvelle
    python /mon-amour-pour-jean-travolta/manage.py jenkins --enable-coverage blogengine

Maintenant ajoutons l'affichage du rapport dans jenkins comme ceci
:

![Configuration](/post_content/2015-03-28/a35c6061-7200-47cb-ba02-86a816776103.png)

Et si on relance plusieurs fois le build, on aura alors cette affichage
:

![Affichage](/post_content/2015-03-28/47f11de3-d972-44b2-9a32-82982a7191fd.png)

Je suis donc en mesure de vous dire que mes 22 tests me permettent de
couvrir plus de 90% du code de mon application. ET OUAIS MESDAMES ET
MESSIEURS.

## Conclusion

Maintenant vous n'avez plus aucun prétexte pour ne plus développer sur
votre application django maintenant que vous n'avez plus rien à faire
pour la mettre en production ! Vous venez tout de même de mettre en
place une usine logicielle qui vous donne des indicateurs qualitatifs
sur votre application. Vous pourrez donc orienter facilement votre
développement et rejeter des builds si une certaine qualité n'est pas
atteinte. Si jamais vous connaissez d'autres plugins, d'autres métriques
/ graphes, ou tout bonnement une autre façon de faire que celle que je
viens de vous énoncer, laissez donc un commentaire. Sur ce, codez-bien
=) Ciao \!

