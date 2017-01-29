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
docker run -d mongo
```

If you execute `docker ps`, you should get something like:

```

CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                          NAMESf2a68c669eeb        mongo               "/entrypoint.sh mo..."   6 seconds ago       Up 4 seconds        27017/tcp                      amazing_turing
```

### Linking a folder or a file from the host in the container

### Linking two containers together

## Creating our own container
## Orchestrating containers


