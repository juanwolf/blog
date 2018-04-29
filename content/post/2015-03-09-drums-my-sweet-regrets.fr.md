
---
title: DRUMS - Mes tendres regrets
date: 2015-03-09
tags: ["DRUMS"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Salut tout le monde, aujourd'hui je vais vous parler d'un projet qui
m'a tenu à coeur ces six dernières semaines, DRUMS. Afin que vous ayez
un brève idée en quoi consiste ce projet, le tag contient une
description de ce super projet.

![](/post_preview/20150322_142806_logo-alpha.png)

Maintenant que vous en savez un peu plus sur DRUMS. Je vais vous y
ajouter le contexte. Faisant parti de l'équipe maître d'oeuvre du
projet, un titre m'a été attribué. Je suis intégrateur du projet, en
gros je m'occupe de l'interropérabilité des composants avec
l'architecture existante.  En plus de ça, je suis responsable de
2 équipes. Bref, y a de quoi péter une durite.

Couplet - Phase de conception
-----------------------------

La phase de conception s'est déroulée de fin octobre jusqu'à début
janvier. Cette période fut très complexe pour plusieurs raisons :

-   Pseudo hiérarchie entre étudiants difficile à respecter
-   Complexité architecturale dû à une expertise technique lacuneuse
    (nous sommes qu'étudiants)
-   Discussions souvent improductives dû aux nombres importants
    d'étudiants impliqués (33)
-   Implication de certains étudiants digne du néant

Pour ma part, j'ai fait parti de ce dernier point, et je le regrette
tout particulièrement. Pourquoi aies-je abandonner cette phase ?

-   Trop de responsabilités
-   L'impression d'être inaudible
-   Répartition des tâches inégales entre étudiants de la MOE
-   Taux horaires insoutenables en plus de la faculté

Je ne fus pas le seul à avoir laisser de côté cette phase et on l'a
vite senti lors de la phase de développement, ayant un DAL aussi utile
qu'un cheval culjat pendant une course hippique.

### Architecture

La phase de développement allait commencer avec cette architecture :

![](/media/django-summernote/2015-03-22/2a5f5fd7-18e8-43e0-b7d7-8a17638a0847.svg)\

Cependant cette architecture pose certaines questions qui n'ont été
résolues que rapidement en phase de développement :

-   Comment le serveur d'application accède aux serveurs de données ?
-   Le serveur de données supporterait t'il deux triple store ?
-   Le serveur de données allait-il supporter en plus une plateforme de
    recherche ?
-   Comment les API développées allaient-elles être accessibles
    facilement par les différentes équipes ?


Beaucoup de questions comme celles-ci qui auraient méritées des réponses
avant la phase de développement.

### Environnement de travail / Versionning

A ce niveau, on a utilisé git, malgré de forts conseils par la MOA
d'utiliser SVN. Soit disant, les étudiants avaient eu une formation sur
cet outil (LOL). L'université nous a fourni 2 dépôts git, ihm et
webservices (back-end et front-end si vous préférez).

#### Front-End

Pour le premier, nous sommes partis sur un système à 6 branches (qui fut
modifié par la suite). Une branche de dev pour chaque équipe + la
branche master servant à la mise en production. L'équipe jaune faisant
office d'équipe intégratrice IHM, leur branche était la branche
d'intégration parti IHM.

On a changé ce système à 6 branches pour un système à deux branches
durant le développement, où seul le master et la branche jaune allait
être utilisé (branche jaune qui sert de branche de développement pour
toutes les équipes). Le master serait déployé en production tant dis que
les modifications de la branche jaune serait déployé sur le serveur de
production.

#### Back-End

Au niveau du back-end se fut un poil plus complexe, on a adopté un
système de feature branching, où chaque branche acceuillait une nouvelle
fonctionnalité. En plus que chaque équipe et chaque projet étaient
développés dans un répertoire qui lui était propre, ce fut une "sur
sécurité" comme le disait Dimitri Baëli
([@dbaeli](https://twitter.com/dbaeli)) qui ne fut pas si
inutile vu le nombre de problèmes qu'a rencontré l'équipe IHM.
Multiplions ces problèmes par 5 et j'aurai perdu la tête pendant la
phase de dév à résoudre tout ça. Pour ce qui est déploiement, on avait
de l'intégration continue pour chaque fonctionnalités sur le serveur de
développement (TOMCAT, mon amour). Pour la mise en production, j'avais
une branche d'intégration qui contenait des scripts qui me permettaient
l'integration quasi automatique de tous les composants (merge de toutes
les branches, passage à l'environnement de production, construction de
tous les war, déploiement).

Un petit schéma, histoire de récapituler les deux précédents paragraphes:

![Schéma](https://docs.google.com/drawings/d/1lPynCo85-9cneoNmL0X50FqcBcnZJUX_pc3b-2NfABE/pub?w=963&h=845)

Refrain - Phase de développement
--------------------------------

La phase de développement avait commencé et nous étions toujours en
train de modifier le DAL (Il y a de quoi paniquer). La première semaine
se déroula correctement, les équipes se faisaient la main sur git malgré
quelques erreurs de certains qui furent facilement corrigées. Cependant
une question arriva de la part d'une équipe travaillant sur le lecteur
audio, Comment fait-on pour accéder au données du serveur distant ? La
réponse fut rapide : SFTP !!! Cette équipe a donc essayé d'ajouter le
SFTP à leur web services mais sans grand succès. Ce problème allait être
rencontré par beaucoup d'équipes voulant toucher au serveur de données,
la création d'une API était donc la meilleure chose à faire.

La deuxième semaine fut extrêmement stressante. Le projet entier fut
"bloqué", tout du moins les équipes se sentaient bloquées dû au
retard de l'API permettant la communication à la base de données. Les
équipes perdaient en motivation et l'équipe en charge de la BDD API se
sentait frustrée dû aux remarques incessantes des autres équipes (Une
promotion unie et soudée comme vous pouvez le voir). De plus, les
'triple store' sur le serveur de données s'entretuaient mutuellement,
il était impossible de faire cohabiter les deux TDB ensembles (RAM + CPU
insuffisant), il a donc fallu en supprimer un :'(. Et une ultime
question se posait, comment allons-nous partager les API aux différentes
équipes ? Deux solutions se proposaient :

-   Récupération des API sur leur branche respective
-   Utilisation d'une entité logicielle permettant le partage de jar

La première solution m'effrayait tout particulièrement. Les étudiants
n'ayant eu qu'une après midi de formation sur git et qu'ils n'ont
été confrontés que très rarement à des problèmes de merge, j'eus peur
d'être surchargé à superviser les merges de toutes les équipes. La
seconde solution fut donc la seule valable. Heureusement, ayant eu une
journée de mise en place d'une usine logicielle (orchestrée par Dimitri
Baëli ([@dbaeli](https://twitter.com/dbaeli))), nous avions
pris connaissance de l'importance d'un repository manager (Je vous
invite à lire cet article, si vous vous demandez à quoi cela sert (en
anglais)
[http://maven.apache.org/repository-management.html](http://maven.apache.org/repository-management.html).
J'ai donc ajouté au serveur d'intégration, un repository manager
(Nexus), qui a permis d'intégrer les API avec des dépendances maven
(bien plus simple que d'effectuer des merges à chaque mise à jour des
API).

Troisième semaine, tout le monde commençait à être rodé. Des absences
commençaient à être de plus en plus fréquentes. Des petits problèmes
d'intégration, et une API pour la base de données qui fut exlpoitable
:). Le service de streaming fut quasiment terminé qu'un problème
surgit. Impossible d'intégrer l'API de la BDD. Pourquoi ? Ce service
fut développé avec un framework REST autre que SPRING data REST :'(
:'( :'(. Apparemment un membre de la MOE fut consulté pour cette prise
de décision (Qui ? Mystère et boule de gomme). Il a donc dû être
développé de nouveau en utilisant le précédent framework. Une perte de
temps non négligeable, qui aurait-pu être évité, si le responsable de
l'équipe surveillait correctement ces équipes (MOI, LOL).

Quatrième semaine. Le tomcat commençait à peiner. Il contenait plus
d'une dizaine de web services et le déploiement devenait de plus en
plus long. Les équipes s'impatientaient de plus en plus. Malgré la
configuration de la RAM utilisée par tomcat, rien n'y faisait, la seule
solution fut de stopper et relancer le tomcat dès qu'il commençait à
suffoquer. Le service de gestion des utilisateurs vascillait d'êtat
stable à êtat instable, cela bloquait l'équipe IHM qui ne pouvait plus
tester ce qu'ils développaient.

Cinquième semaine, ultime rush. La productivité fut multipliée par 2
voir 3. Beaucoup de web services devenaient stables et fut livrés :).
Jeudi, "dernier jour de développement". Tomcat ramait de plus en
plus, l'équipe IHM rushait autant que possible l'intégration des
services au sein de l'IHM. Une nocturne était même prévue afin de
pouvoir préparer la soutenance du lendemain tranquillement. Si
seulement... Tomcat tomba durant la nuit, accablé et détruit par le
poids des web services qui lui ont été déployés :'( :'( :'( RIP
Tomcat. J'ai abandonné mon tendre compagnon à 4h30 du matin. Tomcat
était mort. LA VEILLE DE LA PRESENTATION (il y a de quoi être frustré).
Parallement, nous avions pensé à déployer les webservices sur le tomcat
de production. Lui aussi commençait à ramer, voir être inutilisable.
Bref, la nocture ne fut qu'une douce et lente descente en enfer.

Ultime jour, Vendredi 27 février 2015. J'essaya de me réveiller à 8h30
(lol). Je me leva à 10h30. Un sms au grand chef pour connaître l'actuel
situation (sait-on jamais). La même panique. Cependant le personnel de
la faculté fut mobilisé, l'amdinistrateur réseau et système de la
faculté, le dirigeant du master. Malgré ces grands noms de
l'université, tomcat n'en faisait qu'à sa tête et préférait rester
dans l'au-delà. Il fut seulement 12h30 quand j'ai pu déceler le
problème. Il s'avérait que j'avais détruit la connexion entre le
tomcat et la BDD en voulant travailler sur le tomcat de production la
veille (voir paragraphe précédent). (Pour ceux qui connaissent, j'ai
ajouté une une connexion dans le fichier pg_hba.conf de postgreSQL qui
était mal formattée (oups)). Et étrangement tout est reparti.
Intégration à la vitesse de la lumière, je balance les webservices en
production, l'équipe IHM continuant à développer avec le peu de temps
qu'il restait. Bien sûr entre deux intégrations, nous devions passer la
soutenance du projet... (Autant vous dire que c'était pas beau à
voir).


La démonstration, le grand suspens. TOUT FONCTIONNA !! Seuls quelques
web services n'ont pas été intégrés mais tout marchait :) :) :) :).

Solo
----

Maintenant analysons ce qui a pêché sur ce projet :

-   Architecture vascillante
-   Motivation / implication
-   Lacunes techniques
-   5 semaines de developpement seulement (mais ça on peut rien y faire)

Maintenant la question est, comment aurions-nous pu éviter ce rush de
dernière minute ?

### L'architecture

On a vu que l'architecture choisie dans le cas où un webservice ne
pouvait se deployer sur le tomcat, pouvait bloquer l'équipe IHM dans
leur développement. L'architecture qui aurait pu être choisie
connaissant ce problème aurait été d'ajouter un tomcat où serait
déployé seulement les versions stables (ou RELEASE) des webservices,
ainsi l'équipe IHM ne serait en aucun cas dépendant des autres équipes
et prier que les web services se déploient parfaitement. On aurait cette
architecture de déploiement :

![Différent](https://docs.google.com/drawings/d/14BNg_XQS6Nc2fdnreevxtc2C2zOgPmDz5RarvYOqJU4/pub?w=960&h=720)

Mais on aurait toujours ce soucis de surcharge des tomcats qui est
arrivé en cycle 3. Ayant une architecture orientée services, nous
aurions pu répartir les webservices sur deux machines différentes avec
un tomcat dédié pour chacune d'entre elle. Sachant que nos 4 serveurs
(données, intégrations, développement, production) étaient tous
surchargés, il aurait fallu opérer de cette manière : Déploiement du web
service sur le tomcat d'intégration, puis passage sur celui
développement s'il est stable. Supression des artéfacts présents en dev
et inté lors du passage en production. On perd ainsi le principe de
"copie" pour l'environnement de développement et de production.
Cependant on gagne en performance avec une meilleure répartition des web
services. Au niveau de la répartition des webservices, il aurait fallu
lancer des tests de performances qui aurait permi d'évaluer les
différents ws et permettre une répartition équitables des ressources
entre les deux tomcats. On aurait donc une nouvelle architecture de
cette forme :

![Deux](https://docs.google.com/drawings/d/1jenTFuboSP3rR4iNLs4cvDsYxc3git2XGzj5mbKp_VY/pub?w=960&h=720)

NB : On fait en sorte que le serveur Apache s'occupe des redirections
pour les webservices servant ainsi de proxy.

### Motivation

Etant dans un projet universitaire, une hiérarchie est difficilement
applicable. Certains élèves n'ont pas hésité à faire un peu ce que bon
leur semblait pendant les heures de développement, ce qui avait don
d'en énervé plus d'un, soit disant leur travail était terminé. Une
métrique qui aurait pu etre fort utile afin de répondre à ce dernier
argument très abstrait est le code coverage. N'ayant pas eu le temps de
l'installer, je le regrette un peu. L'équipe maître d'oeuvre aurait
pu utiliser cet argument pour prouver que le développement n'est pas
terminé tant que 'tant de pourcentage' de code coverage n'est pas
atteint.

Pour ce qui est de la motivation, nous aurions pu mettre en place une
pseudo compétition entre les équipes avec un système de notation sur les
builds ratés / réussis, web services développés à temps, code coverage,
aide aux autres équipes, etc\... Cependant le soucis de cette technique
est de perdre les gens qui n'ont pas d'esprit de compéttions.

### Qualité

On a eu quelques soucis d'ordre qualitatif durant le projet (Ex: pas
d'indentation lol). Le développement d'un hook git pour empêcher les
commits de code ne respectant pas un certains checkstyle, n'aurait pas
été un luxe. Le soucis des hooks, c'est qu'on aurait sûrement eu 70%
des étudiants dans cette situation (pour le pre-commit bien sûr) :

![Commic](http://www.commitstrip.com/wp-content/uploads/2015/03/Strip-Confession-650-final.jpg)

La mise en place de code review sur le code de chaque service, aurait
été très utile. Malgré l'expertise de certains, certains étudiants
nécessitent d'être encadré si on veut un minimum de qualité dans le
code. La mise en place de code review 1 ou 2 fois par semaine par équipe
aurait été très efficace. On aurait pu éviter les dérapages techniques
et quelques erreurs de conceptions.

Outro
-----

J'ai fininalement fini, et vous êtes toujours là ? Impressionnant. Je
vous rappelle que j'ai été devops / intégrateur / administrateur
système durant ce projet, certains autres points auraient pu être
ameliorés mais je ne suis pas le mieux placé pour en parler. Malgré les
frayeurs que nous a donné ce projet, ce fut une très bonne expérience.
(Petit moment de réclame) Je rappelle que ce projet a été organisé dans
le cadre du master 2 Génie de l'Informatique Logicielle à l'université
de Rouen. N'hésitez pas à réagir, si certains points vous semble flous
ou discutables ! Ciao, et codez bien !

