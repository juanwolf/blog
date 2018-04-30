---
title: Introduction à Django - partie 2
date: 2015-11-26
tags: ["Django"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

(Re)Bonjour \! Si vous êtes là, c'est que vous êtes restés sur votre
faim pendant ce quickie de 15 minutes de django à Codeurs en Seine,
sinon, pour les personnes s'aventurant par hasard sur cet article, je
vous invite à passer votre chemin (le temps que la vidéo soit
disponible, s'il y en a). On va donc voir ce que j'ai ajouté à
l'application que nous avons codé pour obtenir le résultat de
[ces2015-django](https://github.com/juanwolf/ces2015-django)

## Introduction

Ce que nous allons faire durant cet article va être de décortiquer,
commit par commit quels ont été les notions vues et les principes de
django pour l'élaboration de chaque fonctionnalité / amélioration.

## Héritage des templates

Django permet de mettre en place un système d'héritage des templates
html. Cela nous permettra de créer par exemple une base commune à notre
site et ou créer des templates pour chaque éléments spécifiques de notre
application. Par exemple, nous pourrons nous en servir pour créer un
template spécifique pour le menu, un template spécifique pour un
article, etc... Cela va nous permettre de maintenir plus aisément chaque
partie de notre application.

Dans le commit
[205a43](https://github.com/juanwolf/ces2015-django/commit/205a43c29e6d3d09a351246e4f80b3001123fcc0),
on peut voir que j'ai supprimer une grande partie du code de home.html
pour le mettre dans blogengine/includes/base.html. Ceci va nous
permettre d'avoir une structure qui sera réutilisée par toutes les
différentes pages. Comment spécifier ce template est un template
parent, me demanderez-vous? Je vous répondrai que j'en ai pas la moindre
idée, par contre, pour toutes zones que vous aurez envie de redéfinir
dans les templates fils, vous devez inclure des block.

### Les blocks

Les blocs vous nous permettre de définir les zones à raffiner dans les
templates enfants du template base.html. Dans le template base.html,
j'ai simplement défini 2 blocs: content et footer. Ces deux blocs sont
définis
[ici](https://github.com/juanwolf/ces2015-django/commit/205a43c29e6d3d09a351246e4f80b3001123fcc0#diff-ac710ad52bd5f9914feeba7ef60d40b0R84).
On peut voir que la syntaxe est la suivante :

```django
    {% block nom_du_block %} {% endblock %}
```

Dans les templates fils, nous devrons seulement ajouter du contenu au
sein de ces deux balises afin de les raffiner.

### Extends et include

Passons maintenant aux nouveautés au sein du fichier home.html

On peut voir qu'une grande partie du fichier a été supprimée dû à la
création du template parent. En spécifiant

```django
{% extends blogengine/include/base.html %}
```

Nous précisons à django que ce template étends le template base.html, il
nous suffit alors dans le template home.html de redéfinir les blocks
content et footer avec le contenu adéquat.

Dans ce commit, nous n'avons malheureusement pas d'exemple d'include. Je
vais simplement vous expliquer le principe, contrairement au extends,
vous pouvez inclure des template à des endroits précis de vos templates.
Ceci peut être très pratique pour créer des templates par composant,
ici, nous aurions pu faire un template post.html que nous aurons inclus
dans la boucle for
([ici](https://github.com/juanwolf/ces2015-django/commit/205a43c29e6d3d09a351246e4f80b3001123fcc0#diff-2f7d65f01bb2d1dfb6cb3a1379b2ac1cR21)).
Cela aurait donné un truc du genre :

```django
{% for post in posts %}
    {% include blogengine/include/post.html %}
{% endfor %}
```

Et pour le fichier post.html

```django
{# @vars post : The instance of post to render #}
<div class=post-preview>
    <h2 class=post-title>{{ post.title }}</h2>
    <h3 class=post-subtitle>{{ post.content }}</h3>
    <p class=post-subtitle>{{ post.content }}</p>
    <p class=post-meta>{% trans Posté par :  %}{{ post.author }}</p>
</div>
```

## Les formulaires

Ok, là, on s'attaque à du lourd. Les formulaires en django sont assez
similaires aux models django.

Premièrement, on va créer notre formulaire. Encore une fois django nous
mets à disposition des classes pré construites afin de nous faciliter la
tâche. Ici, comme nous partons d'un model, il suffit d'étendre la classe
ModelForm de django.forms, et de préciser dans la classe Meta de quel
model il s'agit et quel champs nous voulons renseigner dans le
formulaire. Au cas où ce n'est pas clair :

```python
class PostForm(ModelForm):

    class Meta:
        model = Post
        fields = ['title', 'content', 'author']
```

On ne peut faire plus conçis \!

Ce formulaire va nous permettre deux choses, créer des règles de
validations champs par champs et va permettre d'afficher facilement dans
le template le formulaire associé. Django va se baser sur les champs
renseignés dans la liste fields, détecter quels sont leurs types au sein
du model, et créer l'HTML associé à ces champs. Exemple :

Nous avons trois champs : title, content et author. Title et Author sont
tous les deux des CharField avec un max\_length à 255. Django va alors
comprendre que les input liés à ces champs sont deux input avec un
size='255'. Cependant pour le champs content, django va détecter que
nous utilisons un TextField, ce qui implique que django va utiliser un
input de type textArea. Plutôt cool, non ?

### La validation

Créer des règles de validations d'un formulaire est assez facile en
django. Vous pouvez créer des fonctions pour la validation champs par
champs mais vous pouvez aussi ajouter une fonction pour la validation
globale de votre formulaire.

Pour créer une règle de validation pour un champs particulier, vous
devez créer une fonction préfixé de clean et suffixé de
le_nom_du_field. Exemple :

```python
    # Dans la classe PostForm
    def clean_content(self):
        # Le code nous permettant de vérifier que le champs est valide.
        # Ici on va vérifier que le contenu de l'article ne contient pas le mot épinard
        content = self.cleaned_data['content']
        if 'épinard' in content:
            raise ValidationError(JE DÉTESTE LES EPINARDS PUTAIN DE MERDE)

        return content
```

C'est un exemple, vous ne trouverez pas d'exemple comme celui-là sur le
github, j'ai laissé le comportement par défaut des formulaires qui est
de renvoyer un message d'erreur au cas où un des champs n'est pas
renseigné.

### Le Style

Ok, mais des formulaires générés côtés serveurs, c'est pas un peu la
merde pour les stylisés ? Un peu, mais on s'en sort bien. La ruse pour
injecter à nos inputs des styles particuliers est de modifier le widget
des champs du formulaire. Qu'est ce qu'un widget, me direz-vous ? En
fait, le widget est l'entité qui va s'occuper de générer l'HTML de notre
champs. Du coup supposons qu'on veuille appliquer une classe commune à
tous nos champs. On ferait alors:

```python
    def __init__(self, *args, **kwargs):
        super(PostForm, self).__init__(*args, **kwargs)
        for key, field in six.iteritems(self.fields):
            self.fields[key].widget.attrs['class'] = 'ma-super-classe-css'
            # On peut aussi ajouter n'importe quel type d'attribut tel que placeholder par exemple
            self.fields[key].widget.attrs['placeholder'] = 'placeholder'
```

Avec ce système, on peut alors créer des widget spécifiques pour chaque
élément non communs à django. Par exemple, vous utilisez un plugin
jquery pour un élément du formulaire -> création de widget !

### La vue

Comme toute page dans une application django, on va avoir besoin de
créer une vue pour notre formulaire. Django mets à disposition des vues
pré-existantes pour les opérations de CRUD sur un model. Ici, nous
voulons créer un objet en base, nous allons alors étendre la classe
CreateView (On a aussi des UpdateView et DeleteView). Un peu comme la
ListView qui permet de récuperer une liste d'instance d'un model
particulier, il nous suffit au sein de cette classe de précisier, le
model, le template et l'url en cas d'ajout réussi en base pour faire
fonctionner cette vue. Exemple :

```python
class CreatePost(CreateView):
    model = Post
    form_class = PostForm
    template_name = 'blogengine/create_post.html'
    success_url = reverse_lazy('homepage')

```

Vous vous demanderez peut être à quoi sert le reverse_lazy dans
l'attribut success_url ? Cette fonction nous permet de renseigner le
nom de la page à laquelle on veut rediriger l'utilisateur ayant soumis
un bon formulaire à partir du nom de l'url, ceci est plus maintenable
que de renseigner l'url en dur. (PETIT RAPPEL : Nous donnons toujours un
'name' à nos url afin de pouvoir facilement les réutiliser)

### L'URL

N'oublions pas de définir l'url de notre vue ! On reprends le principe
que nous avons vu pour la vue PostListView, dans
[blogengine/urls](https://github.com/juanwolf/ces2015-django/commit/4d68d728f049f955069daae04e7bf835e794e46c#diff-ed9ef2250d32df28fb81d6a3a6a3e887R9):

```python
    url(r'post/create', CreatePost.as_view(), name='create-post'),
```

Vous pouvez voir que sur le commit, j'ai fait une légère erreur à mettre
un / au début de l'url (À NE PAS REPRODUIRE CHEZ VOUS).

### Le template

Maintenant que nous avons tous le back de réaliser, passons au front.
Comme nous l'avons vu dans la vue, nous pouvons itérer sur les fields du
formulaire. Au sein du template, nous avons accès à une variable nommée
form. Cette variable a été injectée dans le contexte de la vue
CreatePost (comportement par défaut des Create, Delete, Update view). Au
sein du template, nous pouvons looper sur les champs comme ceci :

```django
{% for field in form %}
    {{ field }}
{% endfor %}
```

Ceci va simplement afficher le widget de chaque champs de notre
formulaire. Le problème est que si nous avons des erreurs dans notre
formulaire, elles ne seront pas affichées. Heureusement, que chaque
champs contient la liste d'erreurs qui lui est associé. On peut alors
afficher les erreurs comme ceci :


```django
{% for field in form %}
    {% if field.errors %}
        <ul>
            {% for error in field.errors %}
                <li>{{ error }}</li>
            {% endfor %}
        </ul>
    {% endif %}
    {{ field }}
{% endfor %}
```

## Déploiement

*Note: si jamais vous voulez tout simplement lancer l'application, je
vous invite à utiliser le script start.sh à la racine du projet plutôt
que de vous embêter avec ce qui suit.*

Pour le déploiement d'un projet django en production, vous avez
plusieurs choix entre uWSGI, supervisor ou avoir votre instance
elasticbeanstalk. Bon on va la faire à la cool et utiliser le plus
simple et le moins onéreux, uWSGI. L'avantage d'uWSGI, c'est qu'on peut
configurer le serveur avec un simple fichier, ou tout simplement lancer
une ligne de commande avec les arguments nécessaires (bon ça fait tout
de même une sacré ligne !).

On va décortiquer le fichier ligne par ligne :

```ini
[uwsgi]
plugins = python # On indique le plugin uwsgi utilisé (uwsgi peut lancer du php, perl, ruby..)
# Django-related settings
# Le chemin vers le projet django
chdir           = /home/juanwolf/ces2015-django/
# Le fichier wsgi.py (application mère .wsgi)
module          = ces2015.wsgi:application
# Le chemin vers le virtualenv du projet
home            = /home/juanwolf/ces2015-django/virtualenv/

# process-related settings
# master
master          = true
# Le nombre de processus utilisé
processes       = 10
# Emplacement de la socket unix qu'on va utiliser avec nginx
socket          = /home/juanwolf/ces2015-django/ces2015.sock
# Les permissions de la socket
chmod-socket    = 644

uid = juanwolf
gid = juanwolf
chown_socket = juanwolf
# On transforme les processus en démo en précisant l'emplacement du fichier de log
daemonize = /var/log/uwsgi/ces2015.log
# On nettoie l'environnement quand on éteint le processus
vacuum          = true
```

si on execute la commande :

```bash
uwsgi --ini ces2015.ini
```
Notre serveur devrait se lancer \! Par ailleurs, nous devons le lier à
nginx et utiliser la socket précédemment créée. Pour cela, on va faire
un petit peu de config nginx.

Créons premièrement le lien entre nginx et la socket. Dans votre
nginx.conf, ajoutez :

```nginx
upstream ces {
    server unix:///chemin/vers/la/socket/django.sock;
}
server {
    listen 80;
    server_name awesome.django.fr;

    location / {
        uwsgi_pass ces;
        include /etc/nginx/uwsgi_params;
    }
}
```

On relance nginx, et là, on a notre serveur nginx qui tourne et on peut
accéder à notre application toute fraichement installée !!!!

Faîtes attention à avoir désactivé le mode debug dans le settings.py,
[NE JAMAIS DEPLOYER UNE APPLI DJANGO EN DEBUG,
JAMAIS](https://docs.djangoproject.com/en/1.8/ref/settings/#debug).

