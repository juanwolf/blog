
---
title: Create easily a REST API with the django rest framework!
date: 2016-02-07
tags: ["Django", "django REST framework"]
categories: ["Programming"]

slug: create-rest-api-django-rest-framework
aliases:
  - /programming/create-rest-api-django-rest-framework

draft: false
author: "Jean-Loup Adde"
---

Hi! Are you near to create a new API, and you don't know which
framework to use? If this API will be added to a django project, the
django rest framework will be perfect to create this API with few lines
of code!

![](/post_preview/20160207_205447_logo.png)

What's a REST API?
-------------------
The REST architecture will allow us to think first of an API describing
the resources used in our django application.  The main point of
using a REST API is that we'll have a separation between the client and
the server. Then multiple clients can be developped consuming the
ressources given by the API we made !

Another advantage of using a REST architecture is that we can easily
request the API with simple HTTP request. GET requests will retrieve
elements, POST requests will modify them, DELETE requests will delete
them, etc... If you want to know more about the REST architecture I
invite you to read [this
article](https://www.ibm.com/developerworks/library/ws-restful/index.html).




### For what?



Well, if you want to share your data that will be used by other
platforms or if you want them to interact with other services.
For example, what's cool with all the APIs available is that you can
make them interact together, all you need to do is to create the bridge
between them. We can create a link between Trello and Twitter, we just
need to trigger modifications in one of them and modify the other. Here,
we'll just create a simple read only API to retrieve all data available
in this blog:



-   Posts
-   Categories
-   Tags



The django rest framework
-------------------------



The django rest framework will allow us to create easily and with few
lines of code, a REST API on top of our django application. So I
recommend you this framework if you have an existing django project. If
you're starting from scratch, don't worry, you just need to create a
models.py which is not really long or complex :).



### Prequel



For this little introduction, we'll start with a little django project
(this blog). As I said few lines before it's that you need a models.py
with few entities and you should be able to follow this introduction.

In my case, I have three ckasses in my models.py:

-   Post
-   Category
-   Tag

If you want a little diagram, they look like that:

!["Class](/post_content/2016-02-07/b5a4e317-bcb2-4cf1-8fc5-fbbbaa19a587.svg)

### Installation



To install it, you'll have two choices, install it with pip or from
github. I'll just show the first one because it's the most common and
the easiest one. It's time to activate your virtualenv , let's start!


```bash
pip install djangorestframework
pip install markdown
pip install django-filter
```

Or you can add them in your requirements.txt.

Now, let's create a new application to develop our API! Open a
terminal, go to the root of your django project and write:

```bash
python manage.py startapp api
```

If everything went ok, you should have a new folder called api,
containing a migrations folder, files like models.py, admin.py,
views.py, etc... You can delete the models.py, admin.py, and the
migrations folder.

### Viewsets

 Viewsets will act like views in django. They will allow us to
request the databasefor a specific request. We can compare them of
controllers in the MVC pattern.


In this viewsets, you'll have methods like:

-   get
-   post
-   list
-   create

But what's awesome, is that you have viewset for every taste! Generic,
ModelViewSet, ReadOnlyModelViewSet, etc...

In my case, I'll use only ReadOnlyModelViewSet to have a specific
viewset for my model and limiting actions to read actions only. (It
would be a shame if everyone could be able to post articles :/)

#### Let's code for god's sake!

In api/views.py


```python
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

### Serializers

Serializers will allow use to... suspens... serialize (yeah I know
that's weird) our models instancies to JSON and the opposite. We can
compare them to forms in django.

We'll add inside our serializer, fields that we'll need. We could then
write our serializer like that:


```python
# api/serializers.py

from rest_framework import serializers
from blogengine import models

class PostSerializer(serializers.Serializer):
    title = serializers.CharField(max_length=255)
    text = serializers.TextField()
    pub_date = serializers.DateField()
    # Abd the other fields...
```

That's what we could write if the ModelSerializer class did not exist!
We'll just need to specify which model we want to serialie and the
django rest framework will take care to create the fields automatically.
Pretty cool:


```python
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

And the class is not more readable :) (In my case, I put off the fields
argument to have all the attributes of the model).



#### Nested Serializers



For the case of this blog, we can see that a blog can be linked to
multiple tags and to a category. And we could like to nest this category
inside the post element in the JSON. To do so, we can add:

-   Create a serializer for the nested model and add it inside the model
    serializer.
-   Use a forgotten japanese secret



Let's start with the first point:



Create a nested serializer that will be added in our PostSerializer!

```python
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

```json
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


What's useful with a nested serializer is that you can add specific
rules for the serialization of the object such as adding specific fields
(some [SerializerMethodField](https://www.django-rest-framework.org/api-guide/fields/#serializermethodfield) for example).



And the last (and second) method waited by all. The method that will
revolutionize the world of API's creation. I'm speking of the famous
'depth' attribute.


```python
class PostSerializer(serializers.ModelSerializer):
    category = CategorySerializer()

    class Meta:
        model = models.Post
        depth = 1
```

This class attribute will allow us to specifcy the depth of
serialization of the entity we want. Here the 1 value will include the
category (with all its fields) inside our posts.

#### WARNING

Be careful! Here, we speak about serializations only. Nested element may
change your viewset to minimize the number of request in the database.
For example if we keep the previous viewset with the last serializer,
we'll have 1 request to retrieve 'N' posts and 'N' requests to
retrieve each category of each post. So I recommend you to read the doc
about
[prefetch_related](https://docs.djangoproject.com/en/1.9/ref/models/querysets/#prefetch-related)
and
[select_related](https://docs.djangoproject.com/en/1.9/ref/models/querysets/#select-related)
of django.

Example: In the previous example, we try to include categories in each
post JSON. As I said before we'll have N posts displayed but N + 1
requests done to the database. Why? To display our JSON, the django rest
framework will do:

```python
# 1
nested_obj = post.category
# 2
DefaultSerializer(nested_obj)
```


 Well, not really, but that's how it's globally working.
So django will execute a SQL request to retrieve the category element
from the post when the obj = post.category line will be executed.
To counter this problem, we'll need to warn django to load the nested
elements / foreign keys elements.


 For that, let's go to the api/views.py file.

```python
# api/views.py


class PostViewSet(ReadOnlyModelViewSet):
    """
    A simple viewset to retrieve all the posts of blog.juanwolf.fr
    """
    # queryset = models.Post.objects.all()  # La précédente queryset
    queryset = models.Post.objects.all().selected_related('category')
    serializer_class = serializers.PostSerializer
```

So with that, we'll have just one SQL request because django will take
care to request the database with a JOIN on the category table.

### Routers

 Routers will make the urls creation super easy without writting
them all!

 We'll just need to specify the viewset and the name that we
attribute to this viewset and the django rest framework will take care
of create the urls.

 If we create a router like that in our api/urls.py:


```python
# api/urls.py

from rest_framework import routers
from api import views

router = routers.SimpleRouter()
router.register(r'posts', views.PostViewSet)

urlpatterns = router.urls
```

Then we will have 2 urls created automatically.

-   '\^posts/\$' with the name 'post-list'
-   '\^posts/{pk}\$' with name 'post-detail'

So we'll have a lighter urls.py than a usual django urls.py.

Bonus: Let's add some doc to our API
-------------------------------------

To make client development easier for every developer, we need to
document correctly our API.

For that, we'll install swagger which an automatic generator of
documentation. It will show which urls are available to request but also
the type of return from the API.

For that, we'll start to add django_rest_swagger to our
requirements.txt or with pip, it's as you want.

After that, we'll add 'rest_framework_swagger' to our django
application list in our settings.py.



To finish, we need to include swagger urls to our urls.py


```python
#juanwolf_s_blog/urls.py
urlpatterns = [
    '',
    r'^api/docs/', include('rest_framework_swagger.urls'),
    r'^api/', include('api.urls')
]
```

Et voilà the result !
[https://blog.juanwolf.fr/api/docs/](https://blog.juanwolf.fr/api/docs/)

The End
-------

It's over for our little introduction of the django rest framework, I
hope you liked it! If you want to go learn more about it, I invite you
to go on [the website of the framework](https://www.django-rest-framework.org/) which is
really well documented! Sur ce, Codez bien !
