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

It's pretty straight forward. On ubuntu:

`sudo apt-get install docker`

## What is a container?

A container is as said multiple time above, a really light VMs, so what's the difference with a VM? We can take the example of a House and a flat. The House is built from nothing and contains everything you would need and even more. A flat, not really, It is a piece of a building, where (for most of us) contains only what to live. So the container is the flat, and the house is the VM.

Now we could wonder what is the famous "building" hosting all our flats then? It's the famous docker host. I don't know if you tried but if you try to run a docker command just after the installation you had an error like `Is the docker host running?`, Well guess what, no it is not. To make the docker host running, let's execute this command:

`sudo systemctl start dockerd` or `sudo service docker start` depends which version of Ubuntu you use. And now, if you try `docker ps`, you should have:

```CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                          NAMES```

which does not say anything, but at least we know that our docker host is running!

### Running a container
Now that we have our docker host running, let's build some flat!

### Executing a command


### Linking a folder or a file from the host in the container

### Linking two containers together

## Creating our own container
## Orchestrating containers


