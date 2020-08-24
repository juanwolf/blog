+++
title = "Introduction à org-mode"
author = ["Jean-Loup Adde"]
lastmod = 2020-08-24T08:55:20+01:00
tags = ["org", "intro", "emacs"]
draft = false
+++

Markdown est clairement le langage préféré des développeurs. En migrant sur Emacs, j'ai été surpris de constater que la documentation de nombreux plugins est écrite en "org". J'ai lu rapidement à propos et été surpris par ce que l'on peut accomplir avec cette technologie.

{{< figure src="/post_content/introduction-to-org-mode/org-mode-logo.png" alt="Le logo d'org-mode" caption="Figure 1: Le logo d'org-mode" >}}


## Pourquoi org-mode? {#pourquoi-org-mode}

Il se peut que vous vous demandiez, pourquoi créer un autre langage tel que markdown qui est devenu la référence. Il faut savoir que org a été créé dans les années 2000 afin de pouvoir utiliser les principes et méthodologies introduits par ["Get things done" de David Allen](https://gettingthingsdone.com/).

De plus, je ne pense pas qu'il y ait beaucoup d'intérêts à bouger vers org-mode si vous n'utilisez pas Emacs. Toute la puissance de la technologie repose sur son ecosystème et intégration à Emacs. Il faut voir org-mode comme un markdown sous testostérones.

Si vous êtes le genre de personne qui aime avoir un système tout-inclus, org-mode et Emacs vont être parfait pour vous. Sinon pas trop. Emacs va se transformer en usine à gaz et vous n'en sortirez jamais. Comme vous le sentez.


## Comment ? {#comment}

Il y a beaucoup de tutos sur les internets qui vont vous expliquer mieux que moi comment écrire en org. Voici un tutoriel qui m'a beaucoup aidé:

-   <https://karl-voit.at/2020/01/20/start-using-orgmode/>

Si je dois vous donner un conseil, commencez avec un seul fichier où vous regroupez toutes vos notes. Ça va vous permettre de découvrir l'écosystème d'org-mode. Rien que le besoin de créer des notes rapidement va vous introduire à org-capture (un systeme de template pour org-mode). Vouloir exécuter du code en local vous introduira à org-babel. Ajouter une sorte de gestion de projets vous incitera à utiliser les TODO et org-agenda.

Le langage en lui-même est équivalent à Markdown donc il ne devrait pas être trop compliqué à utiliser. Cependant son écosystème est tellement vaste qu'il peut être intimidant.


## L'écosystème {#l-écosystème}

Comme dit précédemment, la force d'org-mode réside dans les outils construit autour du langage. Par exemple, avec org-mode et emacs, je peux écrire des snippets dans ma documentation et les exécuter directement depuis celle-ci. Plus besoin de Makefile quand on peut tout exécuter depuis la documentation.

Au travail, j'ai tendance à travailler sur 15 trucs en même temps et il peut être difficile de justifier pourquoi les tâches des gros projets n'ont pas avancées. Pour éviter ça et être productif au max, j'utilise principalement la technique pomodoro pour ça. Coup de bol, org-pomodoro me permet de choisir une des tâches que j'ai créé dans mes fichiers org et lancer un timer de 25min. Je recois une notification une fois fini et je peux faire une pause. De plus, je peux même aggréger les données de mes pomodoros sur une tâche pour voir le temps total passé dessus.

"Ok t'es gentil avec org-mode mais j'ai déjà X, Y et Z pour gérer mes projets, ma doc, mes notes, etc..." Je comprends, je suis dans la même situation. Il existe des "providers" qui vous permettent de synchroniser vos fichiers org depuis des sources externes. Trello, exchange, Jira, Google Calendar et bien d'autres. Ainsi vous pouvez tout modifier à partir de fichiers "org" et avoir, ainsi, un mode textuel pour tous ces éléments externes.


## Les limitations {#les-limitations}

Forcémment, chaque technologie vient avec ces limitations. Comme énoncé précédemment, org-mode est principalement utilisable dans Emacs. Certaines personnes essayent de porter les extensions disponibles sur d'autres éditeurs mais je n'ai rien vu d'équivalents à celles d'Emacs.

Cette limitation est problématique pour l'adoption de la technologie. Dans mon cas, je ne peux pas partager mes fichiers orgs au travail. Personne n'utilise Emacs et donc ne verront pas l'utilité d'utiliser un format différent de Markdown.

De plus, chaque extension ajoutée à votre config ajoute de la complexité à votre setup. Toutes ces extensions doivent (pour la majorité) être configurée. Ce qui peut prendre pas mal de temps.

Pour finir, les intégrations sont principalement des projets personnel qui ont grande chance de ne pas être à jour ou être abandonné. Si vous êtes vraiment chaud, vous pouvez toujours utiliser votre talent inné de programmeur Elisp et supporter les projets dont vous avez besoin. Gros respect si vous faites ça.


## Conclusion {#conclusion}

Org-mode, c'est beau. Le langage en lui même ressemble à markdown mais ajoute bien plus de fonctionnalités une fois utilisé dans emacs avec son ecosystème. L'écosystème transforme un simple fichier en un playbook, un gestionnaire de projet, un gestionnaire d'habitudes, peut chiffrer vos fichiers et bien plus encore! Cependant, le fait que l'écosystème soit étroitement lié à Emacs fait que son adoption n'est pas si grande malgré toutes ces fonctionnalités. Si jamais la techno vous intéresse, n'hésitez pas à me pinger sur twitter et si le besoin se fait ressentir, j'écrierai peut-être un article sur mon usage d'org-mode. Sur ce, codez-bien!
