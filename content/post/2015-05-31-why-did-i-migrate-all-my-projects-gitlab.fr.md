
---
title: Pourquoi ai-je migré tous mes projets sur Gitlab ?
date: 2015-05-31
tags: ["git"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Depuis la montée fulgurante de git au sein des projets informatiques, il
est devenu un des outils indispensables. Suivant ce phénomène,
l'apparition de services permettant d'héberger votre code ont vu le jour
tel que GitHub et Bitbucket pour les plus connus. Le premier joue sur
l'aspect communautaire et le second mise sur sa privatisation des
projets (pour les versions gratuites). Cependant les entreprises ont
plutôt eu tendance d'heberger eux-mêmes leur code avec GitLab ou
GitBlit. Cependant depuis quelques mois, GitLab a lancé une offre
gratuite d'hébergement avec un nombre de dépôt privé illimité. Ayant
tenté le coup, j'ai vite migré tous mes projets vers cette nouvelle
plateforme. Je vais vous expliquer pourquoi.

![](/post_preview/20150531_154323_b_1_q_0_p_1.jpg.png)

## GitLab, qu'est ce que c'est ?

D'après Wikipédia, GitLab c'est :

> GitLab CE est produit par GitLab B.V. avec un modèle de développement
> « open core ». Il permet :
>
>   - de gérer des dépôts Git ainsi que les utilisateurs et leurs droits
>     d'accès aux dépôts ;
>   - d'effectuer des examens de code et renforcer la collaboration avec
>     les demandes de fusion ;
>   - que chaque projet peut avoir un outil de ticket et un wiki.

L'inconvénient d'utilisation de GitLab était la nécessité de l'installer
(et de posséder un serveur). Hors, tout le monde n'en possède pas et
c'est ainsi que github et bitbucket ont tiré leur épingle du jeu en
offrant un service en ligne gratuit similaire.

## Les offres gratuites sur le marché


|                                |   GitLab  | Bitbucket | GitHub      |
| ------------------------------ | --------- | --------- | ------------|
| Nombre de dépôt privé	         | Unlimited | Unlimited | 0           |
| Nombre de contributeurs	       | Unlimited | 5         | Unlimited   |
| Statistiques sur les contributeurs |  Oui      | Non        | Oui         |
| Support SVN                    |  Non       | Oui       | Non          |
| Support Mercurial              |  Non       | Oui       | Non          |
| Recherche de Code	             |  Oui      | Non        | Oui         |
| Intégration de services	       |  Bonne    | Faible    | Très bonne  |
| Gestion de milestone (jalon)   |  Oui      | Non        | Oui          |
| Services d'intégration continue|  Oui      | Non        | Non          |
| Protections de Branches        |  Oui      | Non        | Non          |

## L'interface

L'interface est sobre et claire. On va décortiquer rapidement les
principales pages disponibles au sein de l'interface, vous verrez, on
prends vite ses marques.

### Page d'accueil

Tout ce qu'il y a de plus classique
:

![Page](/post_content/2015-05-31/e1d631a3-9fd8-4403-a3b1-87471d8a98dc.png)

Les événements liés à vos projets, filtrable par événement (push, merge,
comment) et la liste de vos projets perso non archivés.

### Page profil

Cest une page de profil plus personnelle que celle de github, on garde
tout de même les contributions / activité de la personne mais on
retrouve toutes les informations importantes de lutilisateur (réseaux
sociaux, projets perso).

![Page dun profil sur GitLab]()

### Page projet
![](/post_content/2015-05-31/6ea6880e-03be-47e3-a5ec-e68dd8f18290.png)

Un peu similaire à la page d'accueil, elle
contiendra cependant plus de sections au sein du menu telle que la
section milestone (que nous verrons tout de suite après), les issues,
les graphes d'activités des contributeurs, l'activité, etc...

### La page Milestone

Dans cette page vous pouvez définir
vos jalons / sprints au sein de GitLab et y ajouter des issues ce qui
vous permettra de voir l'avancement de votre projet. Exemple ici de
l'avancement de cet article au 28 Mai :

![](/post_content/2015-05-31/abfd6071-0316-4269-85d5-1ecf8485fc22.png)

## Le social

GitLab contient une minime section Explore où vous pourrez
retrouver tous les projets ayant le plus popularité au sein de GitLab.
Rien à voir en comparaison à l'aspect découverte de GitHub qui vous
permet par thème de découvrir les projets les plus populaires.

Donc GitLab comme Bitbucket ne mise pas essentiellement sur la
partie sociale de sa plateforme. Cependant n'oublions pas que le service
vient d'être lancé et sera sûrement amélioré au fil du temps.

## Conclusion

GitLab est une des meilleurs plateformes
d'hébergement de dépôt git à mon goût. N'étant pas un contributeur pour
les projets open source, j'attends d'une plateforme de m'offrir ce dont
j'ai besoin, une interface facile à utiliser se focalisant
principalement sur mes projets. Ici, GitLab répond tout à fait à mes
attentes. Gestion de jalon par Issue, collaborateur illimité, nombre de
dépôt privé illimité. Que demander, de plus ?

Il est clair que pour le moment la communauté de GitLab est assez faible et son côté
social peu développé, il est donc normal pour les nouveaux projets
open-source d'opter pour GitHub où les chances de contributions sont
nettement plus grandes.
