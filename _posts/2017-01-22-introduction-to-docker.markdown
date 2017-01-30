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

A container is as said multiple time above, a really light VMs, well I am sorry but I lied, a container is **not** a VM. So what's the difference then? We can take the example of a house and a flat (I did not make that up, it comes from the official docker's documentation). The house is built from nothing and contains everything you would need and even more. A flat, not really, It is a piece of a building, where (for most of us) contains only what to live. So the container is the flat, and the house is the VM.

Now we could wonder what is the famous "building" hosting all our flats then? It's the **famous** docker host. The docker host will replace the Hypervisor and be the builder of your container. So that's why it is not a VM, bye bye Hypervisor. An image will clarify the differences:

![Image of differences of virtual machines and containers](https://imgur.com/MJHfm1c.jpg)

So if take the reference of the flat and the house, we can see that the VMs contains more than needed to run just the application (the whole guest OS) compared to the container which contains only what the apps needs to work.

I don't know if you tried but if you run a docker command just after the installation you will have an error like `Is the docker host running?`, Well guess what, no it is not. To make the docker host running, let's execute this command:

`sudo systemctl start docker` or `sudo service docker start` depends which version of Ubuntu you use. And now, if you try `docker ps`, you should have:

`CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                     NAMES`

which does not say anything, but at least we know that our docker host is running!

### Running a container

Now that we have our docker host running, let's build some flats! For this example we will use a nice docker image: jess/hollywood, you will see it will be fun, especially if you always dreamed (or dreamt if you're American) to be a hacker.

`docker run -it jess/hollywood`

So normally you should have a fancy terminal with a lot of shit happening, forget about that, that was just for fun.

Now let's see what this command had actually done. If the jess/hollywood image did not existed on your docker host, the docker cli did a `docker pull jess/hollywood` first which retrieves the docker image on your docker host. Once the download is done, docker will execute run the container as a service but returning the stdout of the service on your stdout (really usefull for debug). You could run this container in detach mode and the container would run in the background (useful for production then).

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

Anyway, we just linked two docker container together and that's quite cool. We just finished to cover all the important features to know with docker. Congrats :D. Now we will get into a more advanced usage, creating our own docker image.

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
