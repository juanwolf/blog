---
title: How to create an (awesome) 404 page in Django
date: 2014-10-19
tags: ["Django"]
categories: ["Programming"]

slug: how-to-create-404-page-django
aliases:
  - /programming/how-to-create-404-page-django

draft: false
author: "Jean-Loup Adde"
---

Hi everyone, today we will speak about the 404 page in Django. The
documentation is quite lazy about this subject so I wanted to make this
point a little bit more understandable. I have to admit, I look for
advices for a long time about it without really find what I needed.
Let's see what I have found during this research.

![](/post_preview/20150322_140627_django-logo-negative.png)

To create a 404 page, many options is possible :

  - Use the default Django's 404 page
  - Use a static template
  - Use a dynamic template

## Use the default Django's page

So... In this case, there's nothing to do, awesome (I don't even know
why you are here). But it's not the best :/.

## Use static html page

First, we need to configure Django to use the template system.

### Setup django to use templates

To use template in our project, we need to configure a little bit
django. We have to tell him where are our template folder. Nothing
really complicated, don't worry. Our project should have this kind of
archictecture:

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
```

Add in to the myproject folder (first one) a third folder named
templates. Now, let's go in the settings.py file in your myproject
folder (second one). Add this :

```python
    TEMPLATE_DIRS = [os.path.join(BASE_DIR, 'templates')]
```

Now, you're ready to create your templates. Little advice: Organize your
template by application, it's an easy way to have an organized project.
If you follow this advice, your project looks like this :

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

WARNING : To use a custom 404 page, be sure the DEBUG variable in your
settings.py file is equal to False, if not you'll always see the 404
debug page of Django

For this option, you can do :

  - Create a html page to the root of your templates folder with the
    name 404.html
  - Create a new view to render the static file.

### First option

Well, just add the html file to the templates folder root with the name
404.html, and it works \! (I know everything was in the title :'( )

### Second option

If the first option is not for you, you can still adapt how Django
reacts when there's a 404 error.

For that, you have to override the view of the handler404. Let's do this
peacefully:

1.  Create the view
2.  Indicate to Django the modification

#### Create a specific view for the 404 error

Here, we will be able to specify what Django will do when it has a 404
error IN ITS FACE. In our case, we will just return the static file,
nothing else. So, let's do this \! It's time to create the view in the
myapp views.py file. :

    def page_not_found_view(request):
         return render(request,'myapp/404.html')

WHAT YOUR 404 DOESN'T RENDER YET ?\! Calm down, let's do the second step
(it will be fine, don't stress out).

#### Indicate the modification of handler to Django

Now our view is ready to be used \! We have to tell to django to use it
when a 404 error is raised, so for that we have to change the
handler404.

For that, you have to add a line to your URL configuration file in the
myproject folder (it doesn't work if you do it in one of your urls.py
app's file). Add the line below:

    handler404 = 'myapp.views.page_not_found_view'

And that's it, you have now an awesome 404 page :).

## Use a dynamic template

This section a little bit more tricky (we will have to add 2 lines to
our view, that's a lot) is for people who need specific informations in
the database or something else more evolved to render their template.

Let's override the view:

    def page_not_found_view(request):
         return render(request,'myapp/404.html')

For this example, we will add all the posts available in our website.
Let's add our posts in the context variable.

    def page_not_found_view(request):
         context = RequestContext(request)
         context['posts'] = Post.objects.all()
         return render_to_response(myapp/404.html, context)

**UPDATE FOR DJANGO 1.11**:

    def page_not_found_view(request):
         context = RequestContext(request)
         context['posts'] = Post.objects.all()
         return render_to_response(myapp/404.html, context.flatten())

Now we want to render the data we had with the context variable. So we
create the template as below:

    <!DOCTYPE html>
     <html>
     <!-- Le Head and everything else -->
     <body>
         <h1>404 Error.</h1>
         <p>
         We saw that you tool the wrong path. That will be our little secret, don't worry ;)
         You can still have a look at these articles:
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

Tadam, you have your awesome 404 which shows all the articles titles
with a link. Awesome isn't it?
