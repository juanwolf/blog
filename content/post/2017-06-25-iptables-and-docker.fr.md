
---
title: Iptables et Docker
date: 2017-06-25
tags: ["ops", "docker"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Hey\! Je jouais un peu avec mon serveur en production cet après-midi et
j'ai remarqué que chaque container docker ouvrait un port sur les
internets. Plutôt effrayant. Voici comment fixer ça.

![](/post_preview/20170625_201501_docker_security.png)

## Le fix

Première fois que je vous renvoie vers un autre article mais tout le
contenu est là: <https://fralef.me/docker-and-iptables.html>

## Le fix rapide

Si jamais vous êtes anglophobe, la solution est simple: Fixer votre
container à utiliser 127.0.0.1 plutôt que 0.0.0.0. Exemple:

    docker run --name nginx 127.0.0.1:9090:80 nginx

Du coup, vous pouvez accéder à ce container que depuis votre serveur sur
le port 9090.

## L'autre fix: Modifier le démon

L'autre solution est de désactiver l'intégration du démon docker avec
iptables, pour cela on doit ajouter une option au lancement de docker.

Premièrement on doit localiser où est le fichier de service pour docker,
si vous utilisez systemctl, vous pouvez le localiser en utilisant
systemctl cat docker.

Voici un exemple de fichier de service pour systemctl désactivant cette
intégration

    [Unit]
    Description=Docker Application Container Engine
    Documentation=https://docs.docker.com
    After=network-online.target docker.socket firewalld.service
    Wants=network-online.target
    Requires=docker.socket

    [Service]
    Type=notify
    # LA PARTIE IMPORTANTE EST JUSTE EN DESSOUS
    ExecStart=/usr/bin/dockerd -H fd:// --iptables=false
    ExecReload=/bin/kill -s HUP
    LimitNOFILE=1048576
    LimitNPROC=infinity
    LimitCORE=infinity
    TasksMax=infinity
    TimeoutStartSec=0
    Delegate=yes
    KillMode=process
    Restart=on-failure
    StartLimitBurst=3
    StartLimitInterval=60s

    [Install]
    WantedBy=multi-user.target

Si vous utilisez autre chose, je vous laisse vous débrouiller :p

Sur ce codez bien, ciao\!

