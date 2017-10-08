---
layout: post
title:  Introduction à Docker
date:   2017-01-22 23:47:01 +0000
categories: programmation
---

# Introduction à Docker

Wow, on attaque du lourd là. Docker. Docker, qu'est ce que c'est ? Cet outil va vous permettre de déployer des applications sans vous préoccuper de sur quelle machine, OS où vous voulez installer votre application elle marchera quoi qu'il en soit. Pour être bref, Docker va vous aider à encapsuler vos logiciels au sein de ultra légère VMs (qui ne n'en sont pas, mais on verra un peu plus tard dans l'article) qu'on appelle "containers".


## Summary

* Qu'est-ce que Docker ?
* Installation
* Qu'est ce qu'un container ?
  * Lancer un container
  * Exécuter une commande
  * Lié un dossier ou un ficher de l'host sur un container
  * Lié deux containers ensembles
* Créé notre propre docker image

## Installation

C'est assez simple d'installer docker sur archlinux :

`sudo pacman -S docker`

Sur Ubuntu, c'est un poil plus compliqué :

```
sudo apt-get install apt-transport-https \
  ca-certificates

curl -fsSL https://yum.dockerproject.org/gpg | \
sudo apt-key add - # We add the docker's official GPG key

sudo add-apt-repository \
"deb https://apt.dockerproject.org/repo/ ubuntu-$(lsb_release -cs) main"

sudo apt-get update
sudo apt-get -y install docker-engine
```

Maintenant vous pouvez utiliser des commandes comme root. Pour pouvoir utiliser les commandes docker avec votre utilisateur courant, vous devez l'ajouter dans le groupe docker
.
`usermod -a -G docker VotreUtilisateurr`

Maintenant vous devriez être capable d'utiliser la cli de docker. Je vous recommande vivement de faire cette petite manip', vous pourrez bénificier de l'autocomplétion sur zsh ou bash pour les noms de containers.

## Qu'est ce qu'un container?

Un container, comme je vous l'ai dit dans l'introduction est une ultra légère VM. Mais vous allez demander pourquoi je vous ai dit qu'en fait non... On peut prendre l'exemple d'une maison et d'un appartement (j'invente rien, je cite juste la doc officielle de Docker). La maison est construite de rien et va contenir beaucoup de choses, voir même plus que ce que vous aurez besoin pour vivre, de plus vous allez construire une maison non pas pour vous seuelement mais sûrement pour la partager avec votre famille, vos amis, vos animaux, etc... L'appartement quant à lui ne peut exister sans immeuble et ce que vous déciderez de placer dans l'appartement sera souvent le strict minimum à défaut de place (à moins que vous roulez sur l'or, chose que je vous souhaite). Dans cette petite métaphore la maison est la VM et l'appartement le container.

Maintenant on peut se demander, ok, c'est bien gentil d'installer des apparts partout mais il nous fait un immeuble, non ? Et bien, cet immeuble, c'est le fameux docker host ou engine (comme vous voulez). Le docker host va remplacer l'hypervisor et devenir le constructeur et l'hébergeur de vos containers. Je vous mets une image ci-dessous vous montrant la différence entre une architecture avec VMs et un archi avec containers.

![Image of differences of virtual machines and containers](https://imgur.com/MJHfm1c.jpg)

Comme on peut le voir plus haut, la VM contient bien des éléments qui ne sont pas forcément utile au fonctionnement de notre applicatio, le Guest OS.

Je ne sais pas si vous avez essayé de lancer une commande docker mais normalement vous devriez avoir une erreur du genre : `Is the docker host running?`, devinez quoi, bah non, il ne tourne pas. Pour faire tourner le docker engine ou host sur votre machine, executez cette commande :

`sudo systemctl start docker` or `sudo service docker start` depends which version of Ubuntu you use. And now, if you try `docker ps`, you should have:

`CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                     NAMES`

qui ne sert pas à grand chose mais on s'est débarassé du message d'erreur ce qui prouve que tout marche comme prévu !

### Lancer un container

Maintenant que nous avons notre docker engine / host qui tourne, construisons donc quelques appartements. Pour cet exemple nous allons utiliser une image super cool: jess/hollywood, si vous avez toujours voulu être un hacker comme dans les films hollywoodien, vous allez être comblé ! Je vous conseille de mettre votre terminal en plein écran et lancer la commande suivante :

`docker run -it jess/hollywood`

Stylé hein? Bon c'était juste pour le fun et merci à [@jessfraz](https://twitter.com/jessfraz) pour l'image.

Regardons ce que la commande précédente a fait. Si vous n'aviez pas localement une copie de l'image sur votre docker host, docker a effectué `docker pull jess/hollywood` pour la télécharger. Une fois le téléchargement effectué, docker va exécuter lancer le container et retourner la sortie du service directement dans notre terminal.
On aurait pu runner ce container en mode détaché mais il n'y aurait plus d'intéret d'utiliser ce container.

### Executer une commande.

Ok, commençons les choses sérieuses. On va installer une base de données sur notre docker host. Par contre ce coup-ci, on va y aller tranquilement.

```
docker pull mongo
docker run --name mongodb -d mongo
```

Si vous exécutez `docker ps`, vous devriez avoir un résultat comme celui-là:

```
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                          NAMES
 2a68c669eeb        mongo               "/entrypoint.sh mo..."   6 seconds ago       Up 4 seconds        27017/tcp                      mongodb
```

Comme vous pouvez le voir, on vient de créer un container nommé `mongodb` avec la docker image de mongo !

Maintenant, nous allons essayer de créer un admin dans la base données.

Pour cela :

```
docker exec -it mongodb mongo admin
```

Vous devriez avoir une sortie comme celle-ci :

```
connecting to: admin
>
```

Maintenant exécutons la commande :

`db.createUser({ user: 'etienne', pwd:'ALaTienne', roles: [{role: "userAdminAnyDatabase", db: "admin"}]})`

et normalement, vous devriez avoir ça :

```
Successfully added user: {
    "user" : "etienne",
    "roles" : [
        {
            "role" : "userAdminAnyDatabase",
            "db" : "admin"
        }
    ]
}
```


Pour résumer ce que nous venons de faire, On peut exécuter des commandes au sein d'un container, si le binaire est au sein du container, il nous suffit de lancer

`docker exec -it my_container my_command --m my_parameter`

**Conseil**: Si vous voulez vous rendre dans un container pour inspecter quoi que ce soit, vous pouvez utiliser cette commande

`docker exec -it my_container /bin/bash`.

### Lié un dossier ou un fichier depuis le docker host dans un container

Bon ok, on a notre base mongo et quelques données dedans. Cool, cool, cool... On est d'accord que des containers ça peut être facilement effacés ? Du coup que se passe t'il à ces données quand on détruit un container ?

Essayons, on verra après :

```
docker stop mongodb
docker rm mongodb
```

Lol, pas de panique, tout va bien, docker a juste réagit comme ça:

![Spongebob erasing all in his head](http://i.giphy.com/5YO4km322zuNy.gif)

Bon... On a tout perdu. On vient de faire une "Gitlab". Mais bon c'est pas grave, "ça arrive".

Gardez en tête que vos containers sont temporaires, un peu comme des vms, vous devriez être capables de les détruire et être capable de toujours avoir vos données et de récupérer l'état de votre vm comme si rien ne s'était passé.
Docker nous permet de stocker des fichiers ou dossiers sur notre docker host grâce à ce que l'on appelle des `docker volumes`.

Créons donc un volume docker lié à notre container afin de garder les données de notre mongodb.

Construisons de nouveau notre docker container mais cette fois-ci avec le volume :

`docker run --name mongodb -d -v /somehwere/youwant/to/store/the/data/on/the/host:/data/db mongo`

Comme vous pouvez le voir, nous avons ajouter l'option -v(pour volume, CAPTAIN OBVIOUS IN DA PLACE) avec ce pattern `chemin_du_dossier_sur_lhost:_chemin_dans_le_container`
Et si vous vous rendez dans le dossier que vous avez spécifié, vous pouvez voir que des données sont apparues dans ce dossier. Pas mal, hein ?

De cette façon, nous pouvons supprimer le container sans problème et lié de nouveau notre dossier pour récupérer un container dans le même êtat que celui que nous avons détruit.

Si nous ne voulions pas accéder aux données et simplement avoir un "espace de stockage" au sein de docker, on aurait pu créer un volume docker de cette façon :
`docker volume create --name mongodb_data`

et le lié de cette manière à notre container :

`docker run --name mongodb -d -v mongodb_data:/data/db mongo`

Vous pouvez aussi lister les volumes que vous avez sur votre docker host. Pour cela, utilisez :

`docker volume list`

Et bizarrement; vous devriez avoir plusieurs volumes... J'ai menti je m'excuse. Quand vous utilisez une image docker contenant des instructions "VOLUME" (voir la section Créer notre propre docker image), un volume non nommé va être créer par docker et contiendra les données du container. Dans l'image de mongodb ( que vous pouvez trouver [ici](https://github.com/docker-library/mongo/blob/master/3.4/Dockerfile), on peut voir qu'il y a une instruction VOLUME pour les dossiers /data/db, /data/configdb, ce qui veut dire que les volumes trouvés précédemment sont ceux de notre premier mongodb... Oups.


Si vous vouliez vraiment détruire les volumes d'un container, utilisez l'option -v : `docker rm -v my_container`.

### Lié deux container ensembles

OOOOOOh, un peu de réseau, the best part EVER. Il serait un peu étrange de laisser une base données sans avoir d'applications connectées à celle-ci. Malheureusement, nous n'allons pas créer d'application spécifique pour cet exempleDésolé.

![A poor guy crying in his shower](http://i.giphy.com/hmE2rlinFM7fi.gif)

**MAIS!** Nous allons utiliser le client mongodb (dans un container) pour se connecter à notre base de données.

Ajoutons notre utilisateur admin comme dans la deuxième section.

```
docker exec -it mongodb mongo admin
> db.createUser({ user: 'etienne', pwd:'ALaTienne', roles: [{role: "userAdminAnyDatabase", db: "admin"}]})
```

Maintenant créons notre client.

`docker run -it --rm --link mongodb:mongo mongo mongo -u etienne -p ALaTienne --authenticationDatabase admin mongo/apero`

Ok, le mongodb:mongo mongo mongo, est un peu chelou. Vous connaissez `docker run -it --rm`, on l'a déjà utilisé. -i -> interactive mode (garde STDIN ouvert), -t -> alloue un pseudo tty, et --rm supprime le container quand il est arrété.


On a ajouté le --link pour créer (suspense...) un lien entre le container que nous sommes en train de créer et celui nommé `mongodb`. `--link mongodb:mongo` veut dire `Créé un lien depuis le container mongodb avec l'alias mongo`. L'alias peut être considéré comme une entrée dans le fichier /etc/hosts liant l'IP du container au label.

Pour être sûr d'être compris, dans le container que nous créons, nous pouvons accéder à la base de donnée en utilisant le nom de domaine `mongo`.

J'ai expliqué `docker run -it --rm --link mongodb:mongo`, le prochain `mongo` est le nom de l'image docker. Et le reste est la commande que nous allons exécuter dans le container.

Retournons taffer un peu. Vous devirez avoir un curseur vous attendant, du genre :

`>`

Écrivons :

`> db.getName()`

et vous devriez avoir :

`apero`


Bon bah c'est l'heure de l'apéro apparemment : https://www.amazon.co.uk/d/Grocery/Ricard-Pastis-8712838324198-45-70cl/B0043A0B2U/ref=sr_1_1_a_it?ie=UTF8&qid=1485704879&sr=8-1&keywords=pastis . WOW..., bon on boira autre chose que du pastis à Londres.

Bref, on vient juste de lier deux container. *MAIS* Si vous avez regardé la documentation, vous avez sûrement vu que cette manière de faire est dépassée... Vu qu'on est des mecs au top de la technologie digitale, révolutionnant l'industrie et le monde dans une ambiance bien plus que familiale, faisons ça de manière classe!

Avec docker vous pouvez créer des réseaux pour vos applications et ils vous facilitent la liaison de vos containers ! Par exemple, si vous voulez créer encore une fois une base de données, tous les containers qui seront liés à cette base de données auront besoin d'être relancé pour pouvoir recréer le lien entre les containers. Avec les réseaux docker, vous pouvez en créer un, placer vos container dedans et comme de par magie, ils vont être capables d'intéragir entre eux. Vous devez toujours configurer votre application afin qu'elle utilise les noms de domaines correspondant aux noms de vos docker containers mais ça marche plutôt bien.

Créons en donc un ! Pour créer un réseau, utilisez cette commande :

`docker network create mongo_network`

Nous devons ajouter notre mongodb dans ce réseau. Détruisons notre container mongo et reconstruisons le.

```
docker stop mongodb
docker rm mongodb
docker run --network=mongo_network --name mongodb -d -v mongodb_data:/data/db mongo
```

Et pour le client...

`docker run -it --rm --network=mongo_network mongo mongo --host=mongodb -u etienne -p ALaTienne --authenticationDatabase admin mongo/apero`

Et tout devrait refonctionner comme précédemment ! Importante chose à savoir, vous ne pouvez plus utiliser les labels comme nous avions fait avec les liens. Donc faites attention aux noms des containers que vous attribuez.

On vient juste de terminer tous les points importants de la ligne de commande de docker. Félicitations (ça fait jamais de mal de se féciliter de temps en temps :D).Maintenant nous allons passer au niveau supérieur et créer notre propre image docker.

## Créer sa propre image

Commençons et terminons sur ce point.Quand vous voulez créer une nouvelle application avec votre framework préféré et que vous voulez utiliser docker, vous allez devoir créer un Dockerfile contenant les informations nécessaires pour installer votre application.

Construisons une simple application qui va retourner une string en sortie. On va construire étape par étape [cette application](https://github.com/juanwolf/hello-cli).

Premièrement, nous allons coder le cli.py. Le code :

```
#!/bin/env python

if __name__ == '__main__':
    print("Hello world!")
```

Maintenant nous pouvons ajouter le Dockerfile pour construire notre application.

La première étape qui est sûrement la plus importante est de choisir sur quelle image nous allons baser notre dockerfile. On peut commencer avec une image d'ubuntu ou directement utiliser une image optimisée pour python. La dernière image sera la plus facile à utiliser car nous n'aurons pas besoin d'installer de dépendance vu que tout sera contenu dans l'image python de base.

Créons donc notre `Dockerfile` là où vous avez créé le fichier `cli.py` :

```
# The docker image we base our one from
FROM python:3.6
# Information of the maintainer of this file
MAINTAINER "Who I am <HowTo@contact.me>"

# We copy the content of the directory to /opt/app in the container
COPY . /opt/app
# We change of directory to /opt/app
WORKDIR /opt/app

# We execute by default the cli.py file
ENTRYPOINT ["python", "cli.py"]
```

Si vous avez besoin de plus d'informations à propos de la syntaxe du dockerfile, la [documentation officielle](https://docs.docker.com/engine/reference/builder/) est votre amie.

Maintenant que nous avons notre Dockerfile, nous pouvons créer notre image. Pour cela :

`docker build -t hello_cli .`

Qui créera donc l'image docker basée sur le dockerfile trouvé dans le repertoire courant (d'où le `.`) avec le nom `hello_cli`. Si tout s'est bien passé, vous devriez avoir une sortie du genre:

```
Sending build context to Docker daemon 35.33 kB
Step 1 : FROM python:3.6
3.6: Pulling from library/python

5040bd298390: Pull complete
fce5728aad85: Pull complete
76610ec20bf5: Pull complete
52f3db4b5710: Pull complete
45b2a7e03e44: Pull complete
75ef15b2048b: Pull complete
e41da2f0bac3: Pull complete
Digest: sha256:cba517218b4342514e000557e6e9100018f980cda866420ff61bfa9628ced1dc
Status: Downloaded newer image for python:3.6
 ---> 775dae9b960e
Step 2 : MAINTAINER "blablabla"
 ---> Running in 67d47f331109
 ---> 8f3a0e87ad1d
Removing intermediate container 67d47f331109
Step 3 : COPY . /opt/code
 ---> e122d9d5f756
Removing intermediate container de3056b06428
Step 4 : WORKDIR /opt/code
 ---> Running in 5b72d5c6e2c2
 ---> 86224093a25a
Removing intermediate container 5b72d5c6e2c2
Step 5 : ENTRYPOINT python cli.py
 ---> Running in 80db6a18e17e
 ---> deee15fd090b
Removing intermediate container 80db6a18e17e
Successfully built deee15fd090
```

Comme vous pouvez le voir chaque commande dans le dockerfile est équivalent à une étape lorsque docker construit notre image.
Il est important de savoir que lorsque votre build échoue, docker va "cacher" toutes les étapes qui ont réussies pour que la prochaine construction de votre image soit plus rapide. De plus, docker va détecter si vous avez changé une étape et la re-exécutera ainsi que toutes les suivantes.

Pour vérifier que tout s'est bien passé, nous pouvons inspecter quelles sont les images que nous avons sur notre docker host. Lancez:

`docker images`

Vous devriez avoir une image `hello_cli` mais aussi l'image mongo que nous avons téléchargée dans le chapitre précédent. Maintenant nous pouvons
créer notre container et voir si tout fonctionne :

`docker run -it --rm hello_cli`

Et vous devriez voir :
`Hello world!`

AWESOME!

## Publier notre docker image

Si vous êtes très fier du travail que nous venons d'accomplir, nous pouvons pousser cette image au sein d'un registre docker (qui est une sorte de "base de données" d'image docker). Gitlab
intègre un registry par défaut mais vous avez aussi le docker hub. C'est comme vous voulez. Pour pousser une image, vous aurez besoin de vous logguer sur le registre depuis votre docker host avec cette commande `docker login myregistry.com`
et pousser votre docker image avec `docker push hello_cli` et vous serez maintenant capable de la télécharger sur tous vos docker hosts !

Le point important à se rappeler et que si vous utilisez un registre perso, vous devez vous authentifier avant de télécharger votre image.

## Conclusion

Et TADAM. On vient de voir à quel Docker est génial dans cette petite introduction. Bien sûr, ceci n'était qu'une introduction pour vous montrer comment Docker fonctionne mais je vous invite à jouer avec dès que vous créez une nouvelle application. Vous pouvez l'utiliser pour créer des environnements de développpement ou même l'utiliser en production et faire de l'intégration continue sans avoir peur d'une dépendance manquante ou quoi que ce soit. Pour ce dernier, je vous invite fortement à utiliser un "outil d'orchestration" pour gérer tous vos containers comme Nomad ou Ansible ou Kubernete, mais tu devras attendre un prochain article jeune padawan !

Sur ce, codez bien! Ciao!
