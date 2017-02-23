---
layout: post
title:  "How to dockerize an application"
date:   2017-02-22 23:12:01 +0000
categories: other
---

# How to dockerize an application

## Introduction

As you could have seen in the previous post, docker is quite handy to make your application running EVERYWHERE. So now the problem you can have is, how do you actually "dockerize" your application. I will give you my guideline, feel free to contribute to it, I will surely be wrong or missing something.

## Define the requirements of the application

It can sound a bit cheap as advice, but before to start everything you need to identify and isolate what your application needs. A database? A task queue? etc... You can plan to have a container for each of your requirement.

But also internally... Your application needs to run inside the container so you can't only ship your code and thinking be done with it. So for example let's say I wrote a Django application. But it will not by itself. I can install a uwsgi or running the default run command inside the container... The default run is not production ready, so you need to think about installing uwsgi... That's the kind of things that you need to prepare before any dockerization..


## Define a setting policy

Biggest and hardest point ever when you start to play with containers. Do I need to embed the configuration file? Should I use env variable? Yeap, we all started to wonder this questions. And sadly there's no answer. Every case as their bad and good points. Let's make a little list of it:

### Configuration as environement variable:

Instead of using a config file as usual you will need to setup one variable for each settings you have to configure. It can be a big effort at the first sight but can be really usefull on the long therm when you don't need to do anything dodgy to change the configuration that's inside your container.

#### Pros

* Easily configurable

#### Cons

* Behavior can change in function of the environment as you might not use same value for the env. variables

### Configuration embed in the container:

#### Pros

* Static and sure that in dev or prod everything is working the same

#### Cons

* Nearly impossible to configure
* Or if you want to configure it, you need to change the configuration before to build the container which can be dodgy sometimes even if you have some CI.

### Configuration mounted as a volume:

#### Pros
* Work as most software and can avoid you to loose time.

#### Cons
* you need the configuration file somewhere in the host (So you might some tools to do that)
* Annoying for vertical scale as the configuration needs to leave on all the docker hosts.

### What to choose

Clearly no idea. I made 3 applications the last few month and they use the 3 ways, and I am happy with none. I am tempted to say that the configuration with env variable is a the best as you can easily change this configuration and just depends on you to make sure that they are filled the same way accross environment or host. You can even use a tool that will deploy your container on different hosts with the same env variables. I use ansible in my case but there's lots so feel free to use the one you want :)

Depends on your case and your needs. Need a quick and dirty workaround?  Config embed. No need to change the configuration? Config embed. Big application that lives magically with a setting file -> setting file as volume. Feeling like a boyscoot and ready to break things? -> env variable.

## Install the minimum

When using docker, you knew that would have to change some stuff. That's where you will need to sweat a bit. Your application needs to the tiniest possible. Tinier and isolated is your app and more gain you would have to use docker. Let's imagine a big app container a task queue and a webapp. Let's imagine the first version is dockerized but both are in the same container. Well, you can't scale (horizontally at least) as you want your application... So always have the strict minimum in your container.

## Define a retention policy

A big point that might force you to rebuild your container is the way you will use docker volumes. It's important to detect before to continue if your container might get bigger. Which it should not. So you need to detect any logging file, any folder or file that can get bigger and define them as volume in your dockerfile. It is easier to manage volumes that size inside a container.

