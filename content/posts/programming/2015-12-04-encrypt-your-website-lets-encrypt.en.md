
---
title: Encrypt your website with Let's encrypt!
date: 2015-12-04
tags: ["nginx"]
categories: ["Other"]

draft: false
author: "Jean-Loup Adde"
---

Hi all \! Today, we'll see how to change our http server to an https one
without paying one fucking cent thanks to let's encrypt which is now
officially in public beta. We will create our certificate and use them
with nginx \!

![](/post_preview/20151204_221310_letsencrypt.jpeg)

## Why HTTPS ?

First of all, we could ask why HTTPS ? That will change anything for
users ? Hmm, good question.

I heard a rumor that http website are not as well indexed in Google
(bullshit) , a rumor seems to say that firefox is going to blacklist all
website without HTTPS, etc... Let them talk and let focus on what
matters, understanding how it works.

As grand pa' wiki says, Https give the ability to web surfers to verify
the identity of a website thanks of the authentication certificate given
by a third authority. Well that's a bit rough, First what's a third
authority ?

Well well well M. Potter (Yeah I quote Harry Potter, a problem ??), an
authority is an organisation which will identify people (CAPTAIN
OBVIOUS). These organisation gives sort of id cards on the web. Well,
not only, a certificate does not help a websurfer to know where he is
but it's crypt datas from the server to the client, which is not to
forget.

## Let's Encrypt

Let's Encrypt, what's that? Well, it's a new third authority open source
which allows webmaster or (web makers) to pass their site in HTTPS
without asking a third authority (not free obviously) to provide them
certificate. So now, no one has an excuse to stay in HTTP. So let's
encrypt\!

## How to make it work on a nginx?

### Certificate generation

Firstly, we need to generate our certificate. For that, we will need to
use the letsencrypt client available on github, so let's clone the repo.

    git clone https://github.com/letsencrypt/letsencrypt.git

Now let's generate our
    certificate:

    ./letsencrypt-auto --cert-only --server https://acme-v01.api.letsencrypt.org/directory

You should have right now a blue screen in your terminal, it will ask
you some informations like which domain you want to encrypt, your
mail...

And.. that's it really, we need now to link the certificates to nginx
and it will done, a bit quick, I know :'( .

### Nginx configuration

Ok, we've done the hardiest step, now we just need to configure our
sweet nginx. For that, we will go in the directory of nginx (commonly:
/etc/nginx) and we'll modify the current configuration. You'll choose
the one your prefer, me I use directly the nginx.conf.

Normally your config should look like that:

    http {
        server {
            server_name blablablabla.com;
            listen 80;

            ///
        }
        # Etc...

    }

What we'll od it's to say to nginx that all http request will be
redirected to https.

```nginx
http {
    server {
        listen 80;
        server_name blablablabla.com;
        return 301 https://;
    }
}
```

We use a 301 code response to notify that our site is passed to HTTPS
(yeah you can be proud of that, BUT NOT YET WE'RE NOT DONE) (By the way:
301 = Permanently moved)

Now let's configure our virtual server:

```nginx
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
            # Your common things here
        }
    }
}
```

We restart the nginx

```bash
sudo systemctl restart nginx
```

ET VOILÀ, you know have a server running through HTTPS. Life is so easy.

