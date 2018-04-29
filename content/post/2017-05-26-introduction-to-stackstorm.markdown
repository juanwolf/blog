---
title:  Introduction to Stackstorm
date:   2017-05-26
tags: ["programming"]
draft: true
---

# Introduction to stackstorm

## Introduction

I am sure you've heard about github being able to deploy cool stuff from their chat platform and all this fancy kind of stuff. At my current job, we use stackstorm, a framework for developing remote execution tools easily and "refined".

## Chatops?

The last past years has been a big changes for entreprises, the raise of the devops practices, agility, telework and all this kind of stuff... A new trend at the moment on a devops/sysadmin point of view is using chat bots to deal with infrastructure and remote execution. Yeah you got it right, you'll use your chat platforms to deploy new applications, check the state of your app, or even provision new vms, amazing isn't it? That's what we call "Chatops".

## Why Stackstorm?

Their's tons of bot engine in the world. You might have heard about [hubot](https://github.com/hubot/hubot), the github bot engine written in coffeescript letting you write 'scripts' for custom actions.
It is quite simple to use. Just

```shell
yo hubot
```

and starts to write your scripts inside the "scripts" folder. At the first sight, you'll find that really handy but with time and the accumulation of your scripts, the project and the maintenance of your bot will get tough.

Stackstorm will had some complexity to your bot management, but organization as well. If you're just trying to make a simple bot. Go for hubot if not and you will use it for production purposes, I invite you to use this nice framework.

## Installation

For this little introcution you can use the [st2vagrant](https://github.com/stackstorm/st2vagrant) project. Of course you'll need a valid vagrant installation.
So with vagrant you'll need to run this commands:
```shell
git clone https://github.com/stackstorm/st2vagrant.git
cd st2vagrant
vagrant up
```

Let's check that everything worked:

```shell
st --help
```
Normally the help would display. That will be enough for our introduction.

## Architecture

Stackstorm is using the adapter pattern to abstract all the conneciton logic to the chat platform. It can be hubot, or any bot engine.
