---
layout: post
title:  "juanwolf.fr is back!"
date:   2017-01-09 22:47:01 +0000
categories: other
---

# juanwolf.fr is BACK!

Hey! It has been a long time mate! I am happy to see you (again, maybe :))

A lot of things happened since I published an article and that is for a lot of reason.

## PERSONAL DISCLAIM OF WHY I DID NOT WORK DURING MY FREE TIME.

A lot of changes happened in my personal life. Changing of sex, Changing of city, changing completely of history, infiltrate the NSA, and I can't tell you more sadly (Universe topsecret defense).
No seriously, I mostly passed my time to look for a job in the fabulous city of London. Which I did (Houraaaaa (You can observe some autosatisfaction in the previous parenthesis)). So I've been a bit settling down.
And now I am back. I upgraded juanwolf.fr to the version 2.0.

## juanwolf.fr  2.0

My good old website, juanwolf.fr, was running on a old box on OVH (2 cores, 2GB of RAM) and started to limit the experiences I wanted to do on it. So It was time to move on!
If you realized the IP for juanwolf.fr changed.
So to prepare the migrations, I decided to make brand new the old applications running on this little server. For that: DOCKER. It took me some time to get where I wanted with it.
But now it's finally done. I moved as well of CI using gitlab-ci instead of jenkins (that I found so much easier for docker based builds but that will be a future subject).


Also, I wanted to move all the website into django to have a uniform ecosystem. So I made a new django app for the index, resume and about pages. But before that I blew my mind a bit
and tried to create a "django-resume" app to build a CV platform on top of it. This idea was not so bad until I realise the amount of work it would need. So instead I chose the quickest way.
Hard to realise that I don't have as much free time as before _sigh_

So to summarize:

* New Servers
* New django application for static pages (index, resume and about page) - [juanwolf.fr_static](https://github.com/juanwolf/juanwolf.fr_static)
* Remove old go proxy ["language-router"](https://github.com/juanwolf/language-router)
* Installation of RocketChat the time I am building a chat with django-channels
* Infrastructure provided with ansible - [playbooks](https://github.com/juanwolf/playbooks)

## E-Sport

I started to play to an awesome game called Rocket League (football + cars: more macho you die) quite intensively (Steam says that I played 180hours apparently (oups)).
Its gameplay is quite simple but really tricky which makes the challenge incredibly high but rewards you for the time that you spent (or trained I would say).
I think I will make an article in the future about it and the e-sport.

## Conclusion (or where to read if you're scared of 423 words)

So juanwolf.fr it's two brand new django applications running in Docker.
I also installed Rocket.Chat on chat.juanwolf.fr so if you have any questions do not hesitate to log into it.

When I will have finish the migration (so in a near future), you will find into this blog articles on:

* Docker (and/or Django + Docker)
* Ansible
* Gitlab CI
* E-Sport / Rocket League
* Doom

In a long time

* Django Channels (I want to do chat app with it, so it might take a while)


Time for me to sleep. Sur ce, codez-bien. Et bonne année.
