
---
title: How to dockerize an application?
date: 2017-02-26
tags: ["Go", "operations"]
categories: ["Programming"]

slug: how-to-dockerize-application
aliases:
  - /programming/how-to-dockerize-application

draft: false
author: "Jean-Loup Adde"
---

As you could have seen in the previous post, docker is quite handy to
make your application running EVERYWHERE. So now the problem you can
have is, how do you actually dockerize your application. I will give you
a guideline, feel free to contribute to it, I will surely be wrong or
missing something.

![](/post_preview/20170226_192322_docker-go.png)

## Define the requirements of your application

It can sound a bit a cheap advice, but before to start everything you
need to identify and isolate what your application needs. A database? A
task queue? etc... You can plan to have a container for each.

But also internally... Your application needs to run inside the
container so you can't only ship your code and thinking be done with it.
So for example let's say I wrote a Django application. But it will work
not by itself. I can install a uwsgi or running the default run command
inside the container... The default run is not production ready, so you
need to think about installing uwsgi... That's the kind of things that
you need to prepare before any dockerization...

## Define a setting policy

A really mind fuck issue when you play with containers. Do I need to
embed the configuration file? Should I use env variable? Yeap, we all
started to wonder this questions. And sadly there's no answer. Every
case as their bad and good points. Let's make a little list of it:

### Configuration as environement variable:

Instead of using a config file as usual you will need to setup one
variable for each settings you have to configure. It can be a big effort
at the first sight but can be really usefull on the long therm when you
don't need to do anything dodgy to change the configuration that's
inside your container.

#### Pros

  - Easily configurable

#### Cons

  - Behavior can change in function of the environment as you might not
    use same value for the env. variables

### Configuration embed in the container:

#### Pros

  - Static and sure that in dev or prod everything is working the same
  - Easy to scale

#### Cons

  - Nearly impossible to configure
  - Or if you want to configure it, you need to change the configuration
    before to build the container which can be dodgy sometimes even if
    you have some CI.

### Configuration mounted as a volume:

#### Pros

  - Work as most software and can avoid you to loose time.
  - Easy to scale vertically

#### Cons

  - You need the configuration file somewhere in the host (So you might
    some tools to do that)
  - Annoying for vertical scale as the configuration needs to leave on
    all the docker hosts.

### What to choose

Clearly no idea. I made 3 applications the last few months, they use the
3 ways and I am happy with none. I am tempted to say that the
configuration with env variable is a the best as you can easily change
this configuration and just depends on you to make sure that they are
filled the same way across environments or host. You can even use a tool
that will deploy your container on different hosts with the same env
variables. I use ansible in my case but there's lots, so feel free to
use the one you want :)

It depends also on your case and your needs. Need a quick and dirty
workaround? Config embed. No need to change the configuration? Config
embed. Big application that lives magically with a setting file -\>
setting file as volume. Feeling like a boyscoot and ready to break
things? -\> env variable.

## Install the minimum

When using docker, you knew that you would have to change some stuff.
That's where you will sweat a bit. Your application needs to be the
tiniest possible. Tinier and isolated is your app and more gain you
would have to use docker. Let's imagine you have a big app containing a
task queue and a webapp. And the first version is dockerized but both of
it are in the same container. Well, you can't scale (horizontally at
least) your task queue as you want or your webapp, you always need to
deploy both... So always have the strict minimum in your container.

To be clear, you will need to study every bit of your application to be
able to run every single piece of software independently to be able to
scale it properly.

## Define a retention policy

A big point that might force you to rebuild your container is the way
you will use docker volumes. It's important to think about it before to
continue in case your container gets bigger (and you'll loose everything
in it). Which it should not. So you need to detect any logging file, any
folder or file that can grow or get updated and define them as volume in
your dockerfile.

## Prepare some performance tests

Ok, let's imagine you finished with your dockerization but suddenly you
realize that your application is getting really slow. To make sure that
does not happen and be sure about it, is to create some performance
tests or using a tool showing you the quality of your app (New Relic for
example).

## Example

Would be a shame to not give you a little example before you get back to
your keyboard with a strange mood of dockerizing the world. I will add
more examples when I will have experienced more docker deployement. Feel
free to give me your example in the comments\!

### Django

As Django needs a settings file to survive, you might play first with a
volume for this setting file. As well you might have configure the
logging to log in a specific file... Then you need to add a volume in
your dockerfile pointing to this/these file(s). Would you like to serve
the statics on your application or with your proxy? If you choose to
serve the statics file with django, you can use whitenoise. If not you
need to define a new volume in your dockerfile pointing to your
STATIC\_ROOT AND call the python manage.py collectstatic. Your static
files should magically appears in your volume :)

Also, You will need a webserver to run your django application. I invite
you to use uwsgi which is extremely reliable.

And that's it, you have all the tips to build a nice django app
dockerized. You can have a look at the way that django-cookiecutter
generate its docker env, it's pretty cool.

## Conclusion

There's not much but that was the few points I would have liked to know
before to dockerize an application. Feel free to add in the comment your
guideline for your framework or your feedback of the article, it would
be helpful for everyone\! Sur ce, dockerizez bien\! Ciao\!

