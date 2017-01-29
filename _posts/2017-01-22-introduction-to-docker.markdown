---
layout: post
title:  Introduction to Docker
date:   2017-01-22 23:47:01 +0000
categories: programming
---

# Introduction to Docker

Wow, big subject here. Docker. So what's docker? Docker will help you to deploy applications without worrying about the installation or the host you want to run it on. To be brief Docker will help you to encapsulate softwares inside lightweight virtual machines.

## Summary

* What's docker?
* Installation
* What is a container?
  * Running a container
  * Executing a command
  * Linking a folder/file on the host in the container
  * Linking two containers together

* Creating our own container
* Orchestrating containers together

## Installation

It's pretty straight forward. On archlinux:

`sudo pacman -S docker`

For ubuntu, it is a bit harder:

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

Now you can run docker commands as root but not as your current user, you need to add your user into the docker group.

`usermod -a -G docker YourLocaUser`

and you should be able to run any docker command. It is highly recommanded to add your user on the docker group. You will be able to use zsh plugins or others, so you will be able to have shell auto completion of the docker container.

## What is a container?

A container is as said multiple time above, a really light VMs, so what's the difference with a VM? We can take the example of a House and a flat. The House is built from nothing and contains everything you would need and even more. A flat, not really, It is a piece of a building, where (for most of us) contains only what to live. So the container is the flat, and the house is the VM.

Now we could wonder what is the famous "building" hosting all our flats then? It's the famous docker host. I don't know if you tried but if you try to run a docker command just after the installation you had an error like `Is the docker host running?`, Well guess what, no it is not. To make the docker host running, let's execute this command:

`sudo systemctl start dockerd` or `sudo service docker start` depends which version of Ubuntu you use. And now, if you try `docker ps`, you should have:

```CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                          NAMES```

which does not say anything, but at least we know that our docker host is running!

### Running a container

Now that we have our docker host running, let's build some flat! For this example we will use a nice docker image: jess/hollywood, you will see it will be fun, especially if you always dreamt to be a hacker.

`docker run -it jess/hollywood`

So normally you should have a fancy terminal with a lot of shit happening, forget about that, that was just for fun.

Now let's see what this command had actually done. If the jess/hollywood image did not existed on your docker host, the docker cli did a `docker pull jess/hollywood` first which retrieves the docker image on your docker host. Once the download is done, docker will execute run the container as a service but returning the stdout of the service on your stdout (really usefull for debug). You could run this container in detach mode and the container would run in the background (production then).

### Executing a command

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

*Tip*: If you want to get inside a docker container to read logs or whatever, you can run
`docker exec -it my_container /bin/bash`.

### Linking a folder or a file from the host in the container

So ok, we have our little mongo database running and some data in it. Cool, cool cool... But we agree that containers
can be easily removed right? So what happens to the data when my container get destroyed???

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
The docker volumes will help you bind directories/files on the docker host. So now that we know that,
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

### Linking two containers together

OOOOOOh, some networking, the best part EVER. It would be weird to let a database on its own without any application connecting to it, don't you think? We will not create a specific application for this section, I am sorry.

![A poor guy crying in his shower](http://i.giphy.com/hmE2rlinFM7fi.gif)

BUT! We gonna use a mongodb client (inside a docker container) to connect to our database.

Let's add again the data in our mongodb.

```
docker exec -it mongodb mongo admin
> db.createUser({ user: 'etienne', pwd:'ALaTienne', roles: [{role: "userAdminAnyDatabase", db: "admin"}]})
```

Now let's create our client.

`docker run -it --rm --link mongodb:mongo mongo mongo -u etienne -p ALaTienne --authenticationDatabase admin mongo/apero`

Ok, the mongodb:mongo mongo mongo, looks pretty confusing. So `docker run -it --rm` We used it before. -i -> interactive mode (keeps STDIN open), -t -> Allocate a pseudo tty, and --rm Remove the container once the container stopped.


We added the --link to create (suspens...) a link between the container we are creating and the one running under the name `mongodb`. `--link mongodb:mongo` means `Create a link to the mongodb container with the alias mongo`. The alias can be considered as an entry inside the file /etc/hosts of the docker container containing the IP of the mongodb container with the "alias name" mongo.

Just to be understood, inside the container that we are creating, we can access to the database container with the hostname `mongo`.

So I explained `docker run -it --rm --link mongodb:mongo`, the next `mongo` is the name of the docker image we want. And the rest is the command that we are executing in the container (connecting the mongo client to the mongo database in the other container).


Let's get back to work. The time you read this section, you should have seen that in your terminal, you have a

`>`

Waiting for you to write something. Let's write:

`> db.getName()`

and you should get:

`apero`


Nice, it's now time for the apero: https://www.amazon.co.uk/d/Grocery/Ricard-Pastis-8712838324198-45-70cl/B0043A0B2U/ref=sr_1_1_a_it?ie=UTF8&qid=1485704879&sr=8-1&keywords=pastis

Bloody hell, the pastis is quite expensive in this bloody island. Anyway, we just linked two docker container together and that's quite cool. We just finished to cover all the important features to know with docker. Congrats :D. Now we will get into a more advanced usage, creating our own docker image.


## Creating our own container
## Orchestrating containers


