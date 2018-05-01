
---
title: "[Archlinux] How to fix \"gcc: removing gcc-libs breaks dependency 'gcc-libs=7.1.1-2'\""
date: 2017-06-04
tags: ["operations", "archlinux"]
categories: ["Programming"]

draft: false
author: "Jean-Loup Adde"
---

Hi everyone, I was crossing this issue this morning so let's fix it\!

![](/post_preview/20170604_144529_archlinux-logo-1159446C2C-seeklogo.com.png)

First of all this issue is appearing because I enabled the multilib
support in my arch config
(<https://wiki.archlinux.org/index.php/multilib>). To fix this problem
you need to have all your gcc packages using multilib or none of them,
not both. First let's find out which packages you need to remove/install
with the multilib support. Run:

```bash
    sudo pacman -Qs gcc
```

You should have an output like:

```bash
local/gcc 6.3.1-2 (base-devel)
  The GNU Compiler Collection - C and C++ frontends
local/gcc-libs-multilib 6.3.1-2
  Runtime libraries shipped by GCC for multilib
local/lib32-gcc-libs 6.3.1-2
  Runtime libraries shipped by GCC (32-bit)
local/libgsystem 2015.2+4+gd606be
  Copylib for system service modules using GLib with GCC
```

In my case you can see that gcc itself is not multilib. Let's install
the multilib version of it.

```bash
pacman -S gcc-multilib
```

Et voilà\! Problem solved :) 2ez

