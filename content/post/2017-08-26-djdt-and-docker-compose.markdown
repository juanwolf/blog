---
layout: post
title:  "Django debug toolbar not showing using docker compose"
date:   2017-08-26 10:00:00 +0000
categories: programming
---

# Django debug toolbar not showing using docker-compose

## Introduction

Hi all, after few minutes of frustration I finally got the django debug toolbar showing in my django apps running in docker managed by docker-compose. This issue happened to me quite a few times now and it really pisses me off to not know why this is happening. And that was a really simple issue...

![The title](https://theadultswimsquad.files.wordpress.com/2017/02/ep-379-3.jpg?w=679&h=381)


## The Issue

If you started to dockerize your django app, I am sure that you encountered some weird difficulties that you managed to fix after few google requests. There's some though that you might have fixed without any explanation.
That's the problem I find in every response I found about how to solve this problem.

I have to admit maybe not everyone is developing inside a docker environment but still I am surprised that I am the first one looking for a proper answer... (Thoughts: Maybe I just searched really badly actually). Anyway!

## The Quickfix

For the ones who really does give a shit about explanations, here's the fix:

```
# Run this command in the directory you would run docker-compose
docker network list | grep ${PWD##*/} | sed -r 's/^([0-9a-z]+).*$/\1/' | xargs docker network inspect  --format "{{ range .IPAM.Config }}{{ .Gateway }}{{ end }}"
# Add the IP to your ALLOWED_IPS and the djdt should now appeared :)
```

Ok that was the ultra short answer and please don't think I wrote that without thinking, it took me one hour to figure out this bloody line :thumbsup: (At least I learnt some stuff about Go Templates :p)

## The explanation

Docker-compose hides a lot of stuff. Final point.

See ya.

Seriously docker-compose hides quite a ton of stuff. What you might have not noticed was that your application is located in a specific network that docker-compose created for you. It will be named something like `my_project_default` or something like that. So your app has a proper network where it will receive a proper ip etc... However, how your localhost will be able to access to this subnet?? Thanks to the bridge function in docker, one interface will be created on your local machine proxing all the traffic to the specific network.
Example:

I am developing [Gringotts](https://github.com/juanwolf/gringotts) at the moment. I go inside the docker folder and run `docker-compose up`. This command will create a `docker_default` network (because I ran the command from the docker folder). We can check that by running `docker network ls`. Now we can run `ip addr` and we should see that a new interface appeared on the local machine. In my case it looks like that:

```
4: br-57a6b1bdcf1a: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:98:71:d9:10 brd ff:ff:ff:ff:ff:ff
    inet 172.19.0.1/16 scope global br-57a6b1bdcf1a
       valid_lft forever preferred_lft forever
    inet6 fe80::42:98ff:fe71:d910/64 scope link
       valid_lft forever preferred_lft forever
```

The good interface will be your network_id (showed with `docker network ls`) prefixed by "br_". And the IP displayed there is the one allocated to your local machine to access to the bridge network of your application.
So now I explained that, I could have done the short answer differently but anyway, the result is the same. So if you followed everything I explained, you understood that when you'll access to your app you'll not use the common localhost ip as 127.0.0.1 but the bridge one 172.19.0.1 instead. So that's why you could not see the django debug toolbar :)


Sur ce codez bien, Ciao!

