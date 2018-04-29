
---
title: "[Archlinux] Comment fixer \"gcc: removing gcc-libs breaks dependency 'gcc-libs=7.1.1-2'\""
date: 2017-06-04
tags: ["ops", "archlinux"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Salut tout le monde, je viens d'avoir un petit problème en essayant de mettre à jour mon archlinux ce matin; Fixons ça en 2 minutes!

![](/post_preview/20170604_144529_archlinux-logo-1159446C2C-seeklogo.com.png)


Premièrement ce problème apparaitra seuelement si vous avez activé le support multilib d'archlinux ([https://wiki.archlinux.org/index.php/multilib](https://wiki.archlinux.org/index.php/multilib)). Pour fixer ce problème vous devez installer tous les paquets gcc en multilib ou aucun, mais pas les deux. Premièrement, trouvons quels paquets foutent le bordel. Utilisez :

```bash
sudo pacman -Qs gcc
```

Vous devriez avoir un output comme ça :

```
local/gcc 6.3.1-2 (base-devel)
  The GNU Compiler Collection - C and C++ frontends
local/gcc-libs-multilib 6.3.1-2
  Runtime libraries shipped by GCC for multilib
local/lib32-gcc-libs 6.3.1-2
  Runtime libraries shipped by GCC (32-bit)
local/libgsystem 2015.2+4+gd606be
  "Copylib" for system service modules using GLib with GCC
```

Dans mon cas on peut voir que gcc lui-même ne supporte pas multilib. Installons la version multilib.

```bash
pacman -S gcc-multilib
```

Et voilà! Problème résolu :) 2ez
