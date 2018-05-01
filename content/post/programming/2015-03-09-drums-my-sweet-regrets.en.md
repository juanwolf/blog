---
title: DRUMS - My sweet regrets
date: 2015-03-09
tags: ["DRUMS"]
categories: ["Programming"]

draft: false
author: "Jean-Loup Adde"
---

Hi everyone, today I will speak about a project which I loved doing
during the last six weeks, DRUMS. To give you a quick idea of what's
about this project, the tag page of DRUMS can help you to know more
about it !

![](/post_preview/20150322_142806_logo-alpha.png)

Now you've learned more about DRUMS, I will add the context to it. I
was in the project management team, and I had a title. I was the
integrator. My duties was to be sure everything will work together with
the specificty of our architecture (material and software). And with
that, I was responsable of 2 teams, because just to be integrator was
not enough (the little pleasure of titles).

Verse - Conception period
-------------------------

The conception period was during the end of october to the beginning of
january. This period was really hard because :

-   Hierarchy hard to maintain between students
-   Complex architecture because of incomplete technical skill (we're
    only students)
-   Discussion often not productive because of the number of student
    involved (33)
-   Some students were not involved at all in the project

I was one of the last point, and I regret that. Why did I ? Few
reasons :

-   Too much responsability
-   I felt I was inaudible
-   Inegal task distribution in my team
-   Hourly rate impossible to keep with the faculty

I was no the only one who didn't really participate to this
period and we felt it during the development period, we had a software
architecture document quite..., hum, how to say, useless, that's the
word.

### Architecture

The development period will begin with this architecture :

![first](https://docs.google.com/drawings/d/12KvZ83ab_6tWi9zW0ujbsKsmQC3vOgWdzmn-bgkZM6E/pub?w=1057&h=717)

However this architecture can make us ask ourself some question which
was solved only during the development period:

-   The data server will be able to support two triple store ?
-   The data server will be able to support a search platform ?
-   How the APIs will be easily accessible for the different teams ?

Some questions like that which deserve some response before the
development period.

### Working environment / Versioning

We used git for the versioning even if the project ownership adviced us
to SVN. "Because, the students had a formation for this tool (LOL).
The university give us 2 git repository, UI and webservices (back-end
and front-end If you prefer).

#### Front-End

For the first one, we used a system with 6 branches (which changed
after). A development branch for every team + the master branch to
deploy in production. The yellow team who was the integrator team UI,
the yellow branch was the integration one.

For this repository, We deployed continously on the development server
for every modification of the yellow branch, and a contineous deployment
to every modification in the master branch on the production server. The
yellow team was in charge of merge on the master when they were
delevering their work.

#### Back-End

For the back-end, it was a little bit more complicated. We used a
feature branching system, where every functionality will be in a proper
branch. And with that, every team had his own directory where they could
create their own ones for every project. For Dimitri Baëli
([@dbaeli](https://twitter.com/dbaeli)), it was an over
security, but it was really useful because the students was not used to
git so, having their own branch will secure every team to the others one
(we never know). And well obviously it was a security for me too, to win
time to not solve every problem of git. About the deployment as you
could see on the previous diagram, we had contineous integration. For
the back-end, every functionality was deployed on the devlopment server
(TOMCAT, my love). To put in production, I had a integration branch
where I developped some scripts which gave me to do the integration in
production through jenkins (merging, passing some variables to
production environment, building all wars and deploying every war to the
specific tomcat).

A little diagram to explay easily what I said in two paragphs :

![Schéma](https://docs.google.com/drawings/d/1lPynCo85-9cneoNmL0X50FqcBcnZJUX_pc3b-2NfABE/pub?w=963&h=845)

Chorus - Development period
---------------------------

This period has started and we still were modifying the software
architecture document (normal). The first week was ok, the team was
experimenting git (somes made few mistakes, nothing really important).
However a question went from the team who was working on audio player,
How we access to the data of the distant server ? Quick response : SFTP
!!! So this team tried to add some SFTP to their webservice but they
didn't succeed :'(. Thinking few times, I realized a lot of team will
need to access to files on the data server. So I created an API to solve
this problem.

The second week was stressful. The full project was blocked, well the
teams felt blocked because of the fact the API to communicate with the
database was late. Teams were losing motivation, and the one who was
developing the DB API was frustrated because of the unstopable remarks
from the other teams (one for all, as you can see). With that, the both
triple store on the data server was killing themself using too much
ressources. So we had to delete one :(. One last question was asked :
How we will share the APIs to the other team ? Two solutions :

-   Pull the APIs on their branch
-   Using a software who allow to share binaries (jar)

The first solution was scaring me. The students had only half a day of
formation about git, so they never really be confronted to conflict and
merge problem. The second option was the only one really valid.
Hopefully, we had a day to build a software factory (orchestretad by
[Dimitri Baëli](https://twitter.com/dbaeli)), we learned during
this module how important is a repository manager during a project (I
invite you to read this article if you're asking yourself what's for)
[http://maven.apache.org/repository-management.html](http://maven.apache.org/repository-management.html).
I added to the integration server this repository manager (Nexus) which
allowed us to share and deploy easily every version of the different
APIs (so simpler than do few merges for every update of the APIs).

Third week, everyone started to be used to. Some absences started to be
more and more frequent (even inside the project managing team). Some
little integration problem and a DB API finally usable :). The streaming
service was nearly finished when a problem came. Impossible to integrate
the database API inside the streaming service. Why ? This one was
developped by an other REST framework else than Spring data REST :'(
:'( :'(. Apparently someone of the management team allowed him to use
anything he wanted (who? We'll never know). So the guy had to develop
an other time his service using the good framework to integrate properly
the database API. A lost of time which be able to miss, if the
responsable of the team was watching properly this team (ME, LOL).

Fourth week. Tomcat started to be slow. He contained more than ten web
services and deployment was longer every build. Teams started to be
impatient. And even with the configuration I did for Tomcat to use more
RAM, nothing changed for Tomcat, the only solution was to kill it when
it started to be slow. An other problem appears, the web service to
manage users alternate between the state of stable and unstable and
that's blocked the UI team who could not test the functionnality they
were developping.

Fifth week, the last rush. Productivity multiplicated by twice. A lot of
web services went finished and delivered :). Thursday, the "last day
of development". Tomcat was getting slower and slower, the UI team
was rushing as they could for the integration of the services, even a
night was planned to prepare peacefully the presentation the next day.
If only... Tomcat died during the night. Destroyed by the weight of the
webservices which were deployed in it. Accablé et détruit par le poids
des web services qui lui ont été déployés :'( :'( :'( :'(. I left my
brother in arm at 4:30am. TOMCAT WAS DEAD. THE DAY BEFORE THE
PRESENTATION FUCK GOD NA SAKE (Sorry, how rough I am, but it was quite
frustrating). During the night, we thought to deploy the webservices in
the tomcat on the production environment. But it too was going to die.
Well, the night was a little descent into hell.

Last day, Friday 27 of February 2015. I tried to wake myself up at
8:30am (lol). I woke up at 10:30am. A text message to the chief to know
the current situation (we never know). PANIC. However all the staff of
the university were trying to figure that out. The system administrator
and the chief of the master. Even with these big names, tomcat kept
beyond. It was only 12:30pm when I could solve the problem. Well,
apparently it was me who destroyed the connection between tomcat of the
production environment and the database (see previous chapter). (For
them who knows, I add a connection inside the pg_hba.conf file of
postgreSQL badly formatted (oups).). Weirdly, everything went fine after
that. Integration with the speed of light, I throw all the webservices
in production, the UI team continued to integrate what they could with
the few time they had left before the demonstration. To make it harder,
between two integrations we had to present our work and how the project
works, it was a kind of big joke.

The demonstration, big suspense. EVERYTHING WORKED !! Only few
webservices was not integrated but everything worked :) :) :) :).

Solo
----

Now, let's analyse what's went wrong during the project :

-   Bad architecture vascillante
-   Motivation / implication
-   technical gap
-   6 weeks of development (but we can't do anything to change that)

Now, the question is : How we could avoid this last rush ?

### The architecture

We saw that the architecture we choosed in the case where a webservice
couldn't be deployed could block other team working with them (for us,
UI team). The architecture which could be used it's to add a tomcat
where every web services will be deployed only stable (or RELEASE) web
services and where the UI team will work on. We will have this
deployment architecture :

![Different](https://docs.google.com/drawings/d/14BNg_XQS6Nc2fdnreevxtc2C2zOgPmDz5RarvYOqJU4/pub?w=960&h=720)

But we will still have the problem of overabundance of the tomcat which
happens during the third cycle. Havind a service oriented architecture;
we could share properly the services between the both servers with a
tomcat for specified environment. Knowing that our 4 servers was all
full (data, integration, development, production), we could operate to
this way : Deployment of the web service on the integration tomcat, and
moving the web services stable to the development tomcat, and after
deleting it to the development server and putting it into the production
one when they were finished. We loose the copy style of the both server.
However we gain in performance. About web services repartitions, it
could be smart to create performance test to evaluate them and sharing
fairly them between the both servers. So we will have this new
architecture :


!["Deux](https://docs.google.com/drawings/d/1jenTFuboSP3rR4iNLs4cvDsYxc3git2XGzj5mbKp_VY/pub?w=960&h=720)

NB : We make sure the Apache server is configured to redirect every
request to the proper tomcat.

### Motivation

Don't forget we are in a universitary project, a hierarchy is hardly
respectable. Some students didn't hesitate to do what they wanted to do
during the development time (we only asked for 7h/day), and they were
saying it was because they work was done (...). A metric which could be
really useful to respond to this lazy people was the code coverage. I
didn't have time to install a Sonar for it, and I regret it. The
project manager team could say, a web service was not acceptable the
time they didn't do some percentage of code coverage.

For the motivation, we could install a little competition between
students with a notation based on failed/success build, web services
developed in time, code coverage, helping other team, etc... However,
the problem of this technic is to loose the students who are quite bad
or who are not competitive at all.

### Quality

We had some problem about the quality of some web services during this
project (ex : no indentation LOL). Develop a hook to stop students
commiting shitty code could be a good idea. But I didn't have time to
develop a module for my pre-commit hooks (you can see that here :
[http://blog.juanwolf.fr/2015/1/git-hooks/](http://blog.juanwolf.fr/2015/1/git-hooks/)).
But the problem of the hooks is they are installed in the client (so in
the machine of the developer), so the developer can avoid it :/. And
I'm sure for more than 70%, the students will be in this situation :

![Commit](http://www.commitstrip.com/wp-content/uploads/2015/03/Strip-Confession-650-finalenglish.jpg)

Setting up some code reviews on the code of every web services could be
really useful to prevent every drift. Some students needs to be
supervised if we want some quality in the code. So doing that for 1 or 2
times a week by team (to not discriminate a team to others) and we could
avoid the drift like using an other framework or setting some id hardly
in the code.

Outro
-----

I finally finished, and you're still here ? Impressive. I was devops /
integrator / system administrator during this project, other point could
be improved but I'm not the best one to talk about. Even some scaring
times, this project was a really good experience. 5little time of
advertising) This project was organised as part of the master 2 in
Software engineering at the university of Rouen. Don't hesitate to
react if at some point it's blurry or debatable! Ciao, et codez bien !

