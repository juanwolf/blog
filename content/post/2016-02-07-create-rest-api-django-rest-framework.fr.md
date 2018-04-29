
---
title: Créez aisément une API rest grâce au django rest framework !
date: 2016-02-07
tags: ["Django", "django REST framework"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Salut ! Vous êtes sur le point de créer une nouvelle API. Vous
hésitez encore sur quel framework choisir ? Si cette API est un
ajout à votre projet django, le django rest framework sera parfait pour
créer en quelques lignes une API REST, documentée, accessible et
performante.

![](/post_preview/20160207_205447_logo.png)


Qu'est qu'une API REST ?
--------------------------

Le REST va permettre de nous orienter pûrement vers une API décrivant
les ressources utilisées dans notre application Django.  Le grand
avantage d'utiliser une API REST est que nous avons une séparation
entre Client et Serveur.  Ainsi de multiples Clients peuvent être
développés afin de consommer les ressources mises à disposition par
l'API.

Le grand avantage de REST est qu'il permet de requêter l'API avec de
simples requêtes HTTP. Les requêtes GET vont permettre de récupérer des
ressources, les requêtes POST d'en modifier, les requêtes DELETE d'en
supprimer, etc..

Si jamais vous voulez en savoir plus sur l'architecture REST, je vous
conseille [cet
article](http://www.ibm.com/developerworks/library/ws-restful/index.html)

### Pourquoi faire ?

Et bien, si jamais vous voulez mettre à disposition des données pour
quelles soient utilisées sur d'autres plateformes ou que vos données
intéragissent avec d'autres. Par exemple ce qui est cool avec
toutes les APIs maintenant disponibles sur le web est que par exemple si
vous voulez faire intéragir Trello avec twitter vous pouvez, Il vous
suffit de créer le pont entre les deux. Ici, nous créerons une simple
API en lecture seulement afin de récupérer toutes les données
disponibles sur le blog :

-   Articles
-   Catégories
-   Tags

Le django rest framework
------------------------

Le django rest framework va nous permettre de créer très facilement et
en très peu de lignes de code une API REST au dessus de notre
application django. Donc je vous conseille éperdument ce framework
si vous avez un projet django existant ! Si vous commencez from scratch,
il vous suffira de créer les modèles ce qui n'est pas non plus très
complexe :).

### Préquel

Pour ce petit tour d'intro, nous partirons d'un projet existant (ce
blog). Ce dont vous avez réelement besoin est d'avoir un models.py avec
quelques entités et ça devrait faire l'affaire ;)

Dans mon cas j'ai trois classes dans mon models.py:

-   Post (Article)
-   Category
-   Tag

Si jamais vous voulez un petit schéma, ils ont cette gueule là:
!["Diagramme](/post_content/2016-02-07/b5a4e317-bcb2-4cf1-8fc5-fbbbaa19a587.svg)

### Installation

Pour l'installation vous avez deux choix, installer depuis pip ou de
l'installer depuis github. On va se contenter de la première. Activer
votre virtualenv, nous allons commencer !

    pip install djangorestframework
    pip install markdown
    pip install django-filter

Ou sinon vous pouvez les ajouter dans votre fichier requirements.txt de
votre projet.

Maintenant créons une nouvelle application pour développer notre API !
Ouvrez un terminal, déplacez vous à la racine de votre projet django et
tapez:

    python manage.py startapp api

Si tout c'est bien passé, vous devriez avoir un dossier avec un dossier
migrations, et des fichiers models.py, views.py, apps.py, tests.py,
admin.py. Vous pouvez supprimer les fichiers models.py, admin.py et le
dossier migrations.

### Les viewsets

 Les viewsets vont jouer le rôle de vues comme dans django. Elles
vont permettre de requêter la base de données en fonction d'une requête
donnée, on peut les comparer à des controllers dans le pattern
MVC.

Dans ces viewsets vous aurez donc des méthodes telles que :

-   get
-   post
-   list
-   create

Mais ce qui est encore mieux avec les viewset, c'est que vous en avez
pour tous les gouts ! Generic, ModelViewSet, ReadOnlyModelViewset, etc..

Dans mon cas, j'utiliserai des ReadOnlyModelViewSet afin d'avoir un
viewset spécifique à mon modèle et limitant les actions sur cette entrée
de l'API à la lecture. (Ce serait dommage que tout le monde puisse
poster des articles :/)

#### Passons au code !

Dans api/views.py


```
# api/views.py
from blogengine import models
from api import serializers


class PostViewSet(ReadOnlyModelViewSet):
    """
    A simple viewset to retrieve all the posts of blog.juanwolf.fr
    """
    queryset = models.Post.objects.all()
    serializer_class = serializers.PostSerializer
    # On va créer le serializer juste après.
```

### Les serializers

Les serializers vont nous permettre de serialiser les instances en JSON
et transformer le JSON en instance python. On peut les comparer à des
formulaires en django.

Nous allons ajouter au sein de notre serializer, les champs que nous
avons besoin. On pourrait donc écrire un serializer de cette façon :

```python
# api/serializers.py

from rest_framework import serializers
from blogengine import models


class PostSerializer(serializers.Serializer):
    title = serializers.CharField(max_length=255)
    text = serializers.TextField()
    pub_date = serializers.DateField()
    # Et tous les autres champs ...
```

Voilà ce que l'on aurait pu écrire si les ModelSerializer n'existaient
pas ! On va juste avoir besoin de spécifier le modèle et le djrf
va s'occuper de créer les champs à la volée. Plutôt cool:

```
# api/serializers.py


from rest_framework import serializers
from blogengine import models


class PostSerializer(serializers.ModelSerializer):

    class Meta:
        model = models.Post
        # Vous pouvez ajouter un fields pour filtrer les
        # champs du modèle à sérialiser
        fields = ('title', 'text', 'pub_date')
```


Ce qui rend le code tout de suite plus lisible :) (Perso, j'ai enlevé
le fields afin d'avoir tous les attributs de mon modèles dans l'API.



#### Les Serializers imbriqués



Dans le cas du blog, on peut voir qu'un post peut être lié à plusieurs
tags et est lié à une catégorie. On pourrait vouloir imbriqué cette
catégorie et ces tags au sein de l'objet JSON en question. Pour cela
nous pouvons :

-   Créer un serializer pour l'entité imbriquée que l'on ajoutera dans
    le serializer de Post
-   Utiliser un secret ancestral japonais transmis de génération en
    génération dans la famille REST



Commençons par le premier point :



Créer un serializer qui sera imbriqué directement dans notre
PostSerializer !


```
class CategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = models.Category
        fields = ('id', 'name', 'description')


class PostSerializer(serializers.ModelSerializer):
    category = CategorySerializer()

    class Meta:
        model = models.Post
        fields = ('title', 'text', 'category')

```

```
{
    "id": 12,
    "pub_date": "2015-12-26T15:24:10Z",
    "title": "Sky Force Anniversary, an old style shoot'em up!",
    "text": "blablablba",
    "category": {
        'id': 2
        'name': 'Video Games',
        'description': "You'll find here all the posts [...]"
    }
}
```


On reste dans le basique mais on pourra par exemple ajouter des champs
spécifiques (des
[SerializerMethodField](http://www.django-rest-framework.org/api-guide/fields/#serializermethodfield)
par exemple).



Et la seconde et dernière méthode que vous attendez tous. La méthode qui
révolutionne le monde de la création de l'API depuis maintenant, pfiou,
longtemps. Je veux bien sûr parler de 'depth'.


```

class PostSerializer(serializers.ModelSerializer):
    category = CategorySerializer()

    class Meta:
        model = models.Post
        depth = 1

```

Cet attribut de classe vous permet de spécifier la profondeur de
sérialization que vous voulez pour ce modèle. Ici le définir à 1 va
permettre d'inclure la catégorie (avec tous ces champs) directement
dans les posts.

#### WARNING

Faites attention ! Ici, nous parlons seulement de sérialization. Un
nombre important de requête en base de données vont être exécutées afin
de récupérer ces éléments imbriqués. Je vous conseille fortement de lire
la documentation sur les
[prefetch_related](https://docs.djangoproject.com/en/1.9/ref/models/querysets/#prefetch-related)
et
[select_related](https://docs.djangoproject.com/en/1.9/ref/models/querysets/#select-related)
de django.

Exemple: Dans l'exemple précédent où nous essayons d'intégrer la
catégorie aux articles, nous allons avoir n requêtes faites en bases de
plus pour n articles au sein du JSON de retour. Pourquoi ? Pour
l'affichage de votre json, le django rest framework va faire:

```
# 1
nested_obj = post.category
# 2
DefaultSerializer(nested_obj)
```


 Bon j'improvise un peu le truc, mais c'est comme ça que ça se
passe en interne.  Ce qui implique que django va exécuter une
requete SQL quand la ligne obj = post.category va être exécutée.
Pour palier à ce problème nous devons dire à django de précharger
l'élément imbriqué  quand nous requétons la base de données pour
récupérer les articles.



 Pour cela, rendons nous dans le fichier api/views.py.


```
# api/views.py

class PostViewSet(ReadOnlyModelViewSet):
    """
    A simple viewset to retrieve all the posts of blog.juanwolf.fr
    """
    # queryset = models.Post.objects.all()  # La précédente queryset
    queryset = models.Post.objects.all().selected_related('category')
    serializer_class = serializers.PostSerializer
```

Ainsi nous n'aurons plus n requetes mais seulement 1 car django va
s'occuper de faire un JOIN SQL sur la table catégorie !

### Les routers

 Les routeurs vont nous permettre d'ajouter des urls aisément
sans avoir à les taper une à une.

Nous aurons seulement besoin de spécifier la viewset et le nom
qu'on lui attribue et le django rest framework va s'occuper de créer
les urls automatiquement.

Si l'on créé au sein de notre fichier api/urls.py un routeur tel
que:


```
api/urls.py

from rest_framework import routers
from api import views

router = routers.SimpleRouter()
router.register(r'posts', views.PostViewSet)

urlpatterns = router.urls
```

Nous allons avoir 2 urls créées automatiquement !


-   '\^posts/\$' avec le nom 'post-list'
-   '\^posts/{pk}\$' avec le nom 'post-detail'

On aura ainsi un fichier urls.py plus léger et plus facile à maintenir
que pour un projet django normal :)

Bonus: Ajoutons de la doc à notre API
-------------------------------------

Afin que tout développeur puisse consommer (aisément) notre API, il est
important d'y apporter de la documentation.

Pour cela, nous allons installer swagger qui est un générateur
automatique de documentation. Il va permettre de montrer les urls
disponibles pour les requêtes HTTP, mais aussi le type de retour de
chaque champs de notre JSON, etc, etc...

Pour cela, on va commencer par ajouter django_rest_swagger dans notre
requirements.txt ou avec pip c'est comme vous voulez.

En suite on va ajouter 'rest_framework_swagger' à notre liste
d'applications django dans le settings.py.

Et pour finir, nous allons inclure les urls swagger dans notre urls.py


```
#juanwolf_s_blog/urls.py
urlpatterns = [
    '',
    r'^api/docs/', include('rest_framework_swagger.urls')),
    r'^api/', include('api.urls'))
]
```

Et bim voilà le résultat !
[https://blog.juanwolf.fr/api/docs/](https://blog.juanwolf.fr/api/docs/)

The End
-------

C'est terminé pour notre petit tour d'horizon du django rest
framework, j'espère que ça vous a plu ! Si jamais vous voulez entrer
plus en détail sur le sujet, je vous invite à vous rendre sur [le site
du django rest framework](http://www.django-rest-framework.org/)
qui est très bien documenté ! Sur ce, Codez bien !

