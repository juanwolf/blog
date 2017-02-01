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

Ok let's do some serious stuff. Let's install a database on our docker host. Let's do it slowly this time.

```
docker pull mongo
docker run --name mongodb -d mongo
```

If you execute `docker ps`, you should get something like:

```
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                          NAMES
 2a68c669eeb        mongo               "/entrypoint.sh mo..."   6 seconds ago       Up 4 seconds        27017/tcp                      mongodb
```

As you can see we created a container from the mongo image with the name `mongodb` and it is running!

Now we gonna try to create an admin in the database.
For that:

```
docker exec -it mongodb mongo admin
```

You should get an output like:

```
connecting to: admin
>
```

Let's execute this command:

`db.createUser({ user: 'etienne', pwd:'ALaTienne', roles: [{role: "userAdminAnyDatabase", db: "admin"}]})`

and normally you should have the output :

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


So to summarize, what we just learned previously is that we can run commands inside a docker container,
if the binary exists (don't try to run vim in a container, it will never be installed) with the command
`docker exec -it my_container my_command --m my_parameter`

**Tip**: If you want to get inside a docker container to read logs or whatever, you can run
`docker exec -it my_container /bin/bash`.

### Linking a folder or a file from the host in the container

So ok, we have our little mongo database running and some data in it. Cool, cool cool... But we agree that containers
can be easily deleted, right? So what happens to the data when my container get destroyed???

Let's do it, we will think about it later:
```
docker stop mongodb
docker rm mongodb
```

Lol, don't worry, everything's fine, docker reacted just like that:

![Spongebob erasing all in his head](http://i.giphy.com/5YO4km322zuNy.gif)

Well, you just lost all your data.

Keep in mind that your containers are temporary, like vms, you should be able to destroy them and
be able to recover the same state when you create a new one. So here comes `docker volumes`.
The docker volumes will help you bind directories/files on the docker host. So now,
we will create a docker volume linked in our docker host to store the data of our mongodb.

Let's build again our docker container but this time with a volume:

`docker run --name mongodb -d -v /somehwere/youwant/to/store/the/data/on/the/host:/data/db mongo`

As you can see, we added the -v option (for volume, CAPTAIN OBVIOUS IN DA PLACE) with this pattern `location_of_the_volume_on_the_host:the_directory_to_get_in_the_container`
And as you can see, some data appeared in the host at the location you specfied. Awesome, isn't it?

With this practice we can do a `docker rm mongodb` without problem and link again our previous volume to have data in our database.

If we would not care at all of accessing to the data, we could have created a named docker volume like that:
`docker volume create --name mongodb_data`

and link it this way:

`docker run --name mongodb -d -v mongodb_data:/data/db mongo`

You can also list the volumes you have in your docker host. For that run:

`docker volume list`

And weirdly, you will have one or multiple entries in it. In fact I lied again (sorry). When you run containers from a docker image that contains `VOLUME` instructions in their dockerfile (please have a look at the Creating our own container section), an unamed volume will be created and will contain the folders it is specified in the dockerfile. In the mongodb dockerfile that you can find [here](https://github.com/docker-library/mongo/blob/master/3.4/Dockerfile), we can see their is an instruction VOLUME for the /data/db, /data/configdb folders, so that means that some your unamed volumes that you found executing `docker volume list` contains the data we inserted in the first section.

If we really wanted to destroy the volumes with the docker container in the example we could have ran this command `docker rm -v my_container`.

### Linking two containers together

OOOOOOh, some networking, the best part EVER. It would be weird to let a database on its own without any application connecting to it, don't you think? We will not create a specific application for this section, I am sorry.

![A poor guy crying in his shower](http://i.giphy.com/hmE2rlinFM7fi.gif)

**BUT!** We gonna use a mongodb client (inside a docker container) to connect to our database.

Let's add again the data in our mongodb.

```
docker exec -it mongodb mongo admin
> db.createUser({ user: 'etienne', pwd:'ALaTienne', roles: [{role: "userAdminAnyDatabase", db: "admin"}]})
```

Now let's create our client.

`docker run -it --rm --link mongodb:mongo mongo mongo -u etienne -p ALaTienne --authenticationDatabase admin mongo/apero`

Ok, the mongodb:mongo mongo mongo, looks pretty confusing. So `docker run -it --rm` We used it before. -i -> interactive mode (keeps STDIN open), -t -> Allocate a pseudo tty, and --rm Remove the container once the container stopped.


We added the --link to create (suspens...) a link between the container we are creating and the one running under the name `mongodb`. `--link mongodb:mongo` means `Create a link to the mongodb container with the alias mongo`. The alias can be considered as an entry inside the /etc/hosts file of the docker container containing the IP of the mongodb container with the "alias name" mongo.

Just to be understood, inside the container that we are creating, we can access to the database container with the hostname `mongo`.

So I explained `docker run -it --rm --link mongodb:mongo`, the next `mongo` is the name of the docker image we want. And the rest is the command that we are executing in the container (connecting the mongo client to the mongo database in the other container).


Let's get back to work. The time you read this section, you should have seen that in your terminal, you have a

`>`

Waiting for you to write something. Let's write:

`> db.getName()`

and you should get:

`apero`


Nice, it's now time for the apero: https://www.amazon.co.uk/d/Grocery/Ricard-Pastis-8712838324198-45-70cl/B0043A0B2U/ref=sr_1_1_a_it?ie=UTF8&qid=1485704879&sr=8-1&keywords=pastis . Fuck... hell, the pastis is quite expensive in this bloody island.

Anyway, we just linked two docker container together and that's quite cool. *BUT* If you read the documentation you saw this way of linking containers is called 'Legacy Links'. Yeah 'legacy'... We are cool kids right? Let's make it again but the right way!

With docker you can build networks for your application and it makes so much easier the way to deal with linking containers together! For example when you want to build again a database, all the linked container will need to be reloaded, and that sucks. With the network feature, you can create networks, put your containers in it and they will magically discover each other. You still need to configure your application to use the name of your container to connect to the other container but that works pretty well.

Let's do it! To create the network, run this command:

`docker network create mongo_network`

We need to add our mongodb to this network. So let's destroy our old container and rebuilt it.

```
docker stop mongodb
docker rm mongodb
docker run --network=mongo_network --name mongodb -d -v mongodb_data:/data/db mongo
```

And for the client...

`docker run -it --rm --network=mongo_network mongo mongo --host=mongodb -u etienne -p ALaTienne --authenticationDatabase admin mongo/apero`

And it should work like previously! A good thing to know is that you can't use "labels" as we did with the link. So be careful of how you name your containers, it will be easier to maintain.

We just finished to cover all the important features to know with docker. Congrats :D. Now we will get into a more advanced usage, creating our own docker image.

## Creating our own container

Let's start and finish on the advanced stuff. When you will want to build a new application with your lovely framework and that you want to use it with docker, you will have to create a specific Dockerfile, to build a docker image for your application.

Let's build a little application that print a string in the output. We gonna build step by step [this application](https://github.com/juanwolf/hello-cli).

First, We need to code our cli.py file. Here's the code:

```
#!/bin/env python

if __name__ == '__main__':
    print("Hello world!")
```

Now, we can add a Dockerfile to build  an image for this application.

We need to choose from which docker image we will base our one. We can start from an Ubuntu one, or directly using a python image. The last one will be easier, everything will be installied by default and we will only need to do
a python cli.py and the job is done :).

Let's create the file `Dockerfile` where our cli.py lives with this content:

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

If you want more information about the syntax and the command you can execute on this file, the [Dockerfile reference](https://docs.docker.com/engine/reference/builder/) is your friend.

Now that we have our Dockerfile, we need to build our image. For that:

`docker build -t hello_cli .`

Which will build the docker image with the Dockerfile found in the `.` directory with the name hello_cli. If everything went well, you should have an output like this one:

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

As you can see every command we made in the dockerfile equals to a step when building the image.
It's important to know that when a build failed docker will cache all the successfull steps to make
your next build quicker. Also, docker will detect if you changed a cached step and will make it again and all
the next ones.

To verify that everything went well, we can inspect which images we have on our docker host locally. Run:

`docker images`

You should have the hello_cli but also the mongo image we downloaded in the previous chapter. Now we can run
our docker image and see if everything works:

`docker run -it --rm hello_cli`

And you should see:
`Hello world!`

AWESOME!

## Publishing our docker image

If you are really proud of what we just accomplished, we can push this docker image to a registry (which is a "database" of docker image). Gitlab
comes with one integrated but there's also the docker hub. As you wish. To push an image, you will need to login to the registry first with `docker login myregistry.com`
and push your local image with `docker push hello_cli` and you are now able to download this image!

The important point to remember using a custom registry is to login before trying to pull an image.

## Conclusion

And TADAM. We had a quick view on all the awesomeness of Docker in this little introduction. Of course this is just to show you quickly how it works but I
invite you to play with it when you build a new application. You can use it as development environment or  even use it to deploy continuously applications without any
fear (well if you app has bugs, that's not docker's fault ;p). But for this last point it is heavily recommanded to use a software to orchestrate your containers such as
nomad or ansible or kubernete, but little padawan, you will need to wait for this topic in a future article!

Sur ce, codez bien! Ciao!
