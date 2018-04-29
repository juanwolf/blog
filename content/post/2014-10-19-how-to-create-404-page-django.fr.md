---
title: Comment créer une (superbe) page 404 pour Django
date: 2014-10-19
tags: ["Django"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Salut tout le monde, aujourd'hui on traitera le sujet des 404 avec
Django. Je trouve la documentation très légère et je dois avouer que
j'ai un peu fouillé dans les entrailles du web pour trouver un semblant
de résultat intéressant. Regardons ce que j'ai pu trouvé au cours de mes
recherches.

![](/post_preview/20150322_140627_django-logo-negative.png)


Afin de créer une page 404 plusieurs options s'offrent à vous :

  - Utiliser la 404 par défaut de Django (beuuuuuuurk)
  - Utiliser un template statique
  - Utiliser un template dynamique.

## Utiliser la 404 par défaut de Django

Euh... donc dans ce cas là rien à faire, c'est génial. Sauf qu'à
l'affichage on a vu mieux.

## Utiliser un template statique

Nous devons avant tout configurer Django pour l'utilisation de
templates.

### Configuration pour l'utilisation des templates

Afin d'utiliser les templates dans notre projet, nous devons configurer
quelques peu django afin de lui indiquer où se trouve nos superbes
templates. Rien de bien complexe pour le moment, ne vous inquiétez pas.
Votre projet devrait avoir un projet avec cette architecture:

    myproject/
             myapp/
                 urls.py
                 views.py
                 models.py
                 tests.py
             myproject/
                 settings.py
                 urls.py
                 wsgi.py

Ajoutez au sein au dossier myproject/ (premier du nom) un troisième
dossier appelé templates. Rendez vous au sein du fichier settings.py du
dossier myproject (second du nom). Ajoutez-y :

    TEMPLATE_DIRS = [os.path.join(BASE_DIR, 'templates')]

vous voilà donc parer à créer vos templates. Je vous conseille
d'organiser vos templates par application, c'est un moyen simple de ne
pas avoir un projet désordonné. Vous aurez avoir le résultat suivant :

```
 myproject/
     myapp/
         urls.py
         views.py
         models.py
         tests.py
     myproject/
         settings.py
         urls.py
         wsgi.py
     templates/
         myapp/
```

ATTENTION : Afin de pouvoir utiliser une page 404 personnalisé, vous
devez définir la variable DEBUG à False au sein de votre fichier
settings.py, sinon une page 404 de débug Django s'affichera.

Pour cette option plusieurs choix s'offrent à vous :

  - Créer votre page à la racine du dossier template avec le nom
    404.html
  - Créer une pseudo-vue qui affichera votre page HTML.

### Première option

Ajoutez la page 404.html voulue à la racine de votre dossier templates
et le tour est joué !

### Seconde option

Si la première option ne convient pas à vos besoins, vous pouvez
toujours adapter le comportement de Django lors d'une erreur 404
(heureusement).

Pour cela vous devez redéfinir la vue du handler404 de Django (KES KI
DI). On va procéder en plusieurs étapes :

1.  Créer une vue
2.  Indiquer la modification du traitement d'une erreur 404.

#### Créer une vue spécifique à l'erreur 404

C'est ici que nous pourrons spécifier le comportement de django lors
d'une erreur 404. Dans notre cas nous ne voulons que renvoyer un
template prédéfini (page html). Pour cela rendez vous au sein du dossier
de votre application (myapp dans notre exemple) . Créons maintenant la
vue :

    def page_not_found_view(request):
         return render(request,'myapp/404.html')

Quoi votre 404 ne s'affiche pas encore ? Du calme, vous avez déjà oublié
la seconde étape.

#### Indiquer la modification du traitement d'une erreur 404

Maintenant que notre vue a été définie, nous pouvons indiquer à django
le traitement à suivre en cas d'erreur 404.

Vous devez ajouter une ligne au sein de votre configuration d'URL au
sein du dossier myproject (cela ne marche pas si le handler est défini
au sein d'une app). Ajoutez la ligne suivante :

    handler404 = 'myapp.views.page_not_found_view'

Et voilà, vous avez maintenant une magnifique page 404.

## Utiliser un template dynamique

Cette section un peu plus technique (2 lignes de codes en plus dans
notre vue, tout de même) s'adresse aux personnes qui ont besoin d'une
page 404 qui contient des informations stockées au sein de la BDD ou
tout autre traitement spécifique relatif à l'affichage de leur template.

Redéfinissons donc notre vue comme précédemment :

    def page_not_found_view(request):
         return render(request,'myapp/404.html')

Pour cet exemple, nous ajouterons tous les articles disponibles sur le
site. Ajoutons alors ces articles au sein de notre variable de contexte.

    def page_not_found_view(request):
         context = RequestContext(request)
         context['posts'] = Post.objects.all()
         return render_to_response(myapp/404.html, context)

**MISE A JOUR POUR DJANGO 1.11**:

    def page_not_found_view(request):
         context = RequestContext(request)
         context['posts'] = Post.objects.all()
         return render_to_response(myapp/404.html, context.flatten())

Maintenant nous pouvons manipuler tous les articles au sein de la
variable de contexte au sein de notre template tel que :

    <!DOCTYPE html>
     <html>
     <!-- Le Head et tout ce qui s'en suit -->
     <body>
         <h1>Erreur 404.</h1>
         <p>
         Nous vous avons vu prendre une mauvaise direction. Ça reste entre nous, ne vous inquiétez pas ;)
         Vous pouvez toujours visiter les articles suivants :
         </p>
         <ul>
             {% for post in posts %}
             <li>
                 <a href={{ post.get_absolute_url }}>
                     {{ post.title }}
                </a>
             </li>
             {% endfor %}
         </ul>
     </body>
     </html>

Et voilà vous avez votre superbe 404 qui s'affiche avec tous les titres
de vos articles. Génial n'est ce pas ?

Si vous avez encore d'autres techniques pour afficher des pages 404,
n'hésitez pas à lacher l'info dans les commentaires ! Codez bien,
ciao.

