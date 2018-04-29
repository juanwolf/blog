
---
title: Chiffrez votre site en HTTPS avec let's encrypt !
date: 2015-12-04
tags: ["nginx"]
categories: ["Autre"]

draft: false
author: "Jean-Loup Adde"
---

Salut tout le monde \! Aujourd'hui on va regarder comment passer notre
serveur en https assez facilement avec let's encrypt. On va générer nos
certificats et configurer notre petit nginx afin d'être comme les grands
en HTTPS

![](/post_preview/20151204_221310_letsencrypt.jpeg)

## Pourquoi HTTPS ?

Déjà la première question que l'on pourrait se poser est pourquoi HTTPS
? C'est vrai ça, pourquoi ? Cela change t'il quelque chose à notre
navigation ? Morf. Complexe question.

Une rumeur en SEO insinue qu'un site en HTTPS serait mieux référencé
(bullshit), une rumeur circule comme quoi firefox blacklisterait les
sites internets n'étant pas en HTTPS, etc... Laissons ces rumeurs
bullshit pour nous baser sur pourquoi et comment marche HTTPS.

Comme papy wiki le dit, HTTPS permet aux visisteurs de vérifier
l'identité du site web auquel il accède, grâce à un certificat
d'authentification émis par une autorité tierce, réputée fiable (et
faisant généralement partie de la liste blanche des navigateurs
internet). Bon, c'est un peu du charabiat au prime abord. Déjà qu'est ce
que l'autorité tierce ?

Et bien Jamy, (oui je fais des tournures à la C'est pas sorcier), une
autorité tierce est une organisation (lucrative ou non) qui va permettre
d'identifier des correspondants. Ce sont des organisations qui donnent
en quelques sortes des cartes d'identité sur le web. Enfin pas que, un
certificat ne sert pas qu'à indiquer à un surfeur du web si le site sur
lequel il est correspond bien à ce qu'il voulait visiter.

Le passage en HTTPS permet d'assurer aux visiteurs que leur connexion
avec les serveurs du site sont chiffrées et ça c'est pas rien.

## Let's Encrypt

Let's Encrypt, c'est quoi ? Et bien c'est une nouvelle autorité de
certification libre et open source et qui permette donc aux petits sites
comme celui sur lequel vous naviguez en ce moment de pouvoir être en
https, sans que vous n'ayez d'opérations à faire sur votre
navigateur.Car oui, les authorités de certifications ne certifie que
moyennant une certaine somme. :'( :'( MAIS maintnenant grâce à let's
encrypt plus personne sur le web n'a d'excuses de rester en HTTP.

## Comment mettre ça en place sur nginx ?

### La génération de certificats

Premièrement nous devons générer nos certificats. Pour cela, il va
falloir cloner le dépôt github de letsencrypt.

    git clone https://github.com/letsencrypt/letsencrypt.git

Maintenant, générons nos certificats
    :

    ./letsencrypt-auto --cert-only --server https://acme-v01.api.letsencrypt.org/directory

Vous aurez alors un écran bleu qui va remplacer votre console et vous
demandera quelques infos telles que les domaines à certifier, votre
email, etc...

Et voilà, nous avons nos certificats. Il ne reste plus qu'à utiliser ces
certificats avec notre nginx.

### La configuration de NGINX

Ok, on a fait le plus dur, maintenant, il ne nous reste plus qu'à
configurer notre cher et tendre nginx. Pour cela, on va aller dans le
répertoire de ce dernier (souvent /etc/nginx/) et on va modifier la
configuration actuelle. On a nginx.conf ou sites-enables/default. A vous
de choisir lequel vous utilisez pour votre configuration (moi
nginx.conf).

Normalement votre config actuelle doit ressembler à ça :

    http {
        server {
            server_name blablablabla.com;
            listen 80;

            ///
        }
        # Etc...

    }

Ce que nous allons faire et de dire à nginx que si on a uen requete en
HTTP, on le redirige vers HTTPS.

    http {
        server {
            listen 80;
            server_name blablablabla.com;
            return 301 https://;
        }
    }

On utilise 301 afin de notifier que notre site est passé en HTTPS (301 =
Permanently moved)

Maintenant on va configurer nginx pour écouter les entrées https

```
 http {
    server {
        listen 80;
        server_name blablablabla.com;
        return 301 https://;
    }

    server {
        listen 443;
        server_name blablablabla.com
        ssl on;
        ssl_certificate /etc/letsencrypt/live/blablablabla.com/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/blablablabla.com/privkey.pem;

        location / {
            # Vos trucs ici
        }
    }
}
```

On relance notre nginx :

    sudo systemctl restart nginx

ET TADAM, vous avez maintenant un serveur en HTTPS. Elle est pas belle
la vie ?

