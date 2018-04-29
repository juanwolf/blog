
---
title: Comment créer un routeur de langue avec Go ?
date: 2014-06-18
tags: ["Go"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---


Si vous vous êtes posé comme moi la question "Comment pourrais-je
servir mes fichiers statiques correctement pour plusieurs langues ?",
alors cet article est exactement fait pour vous ! N'ayant trouvé aucun
équivalent répondant à mes attentes, je me suis donc attelé à la tâche
de le créer. Voyons comment j'ai procédé.

![](/old_blog/20150322_141925_gogogo1.jpg)

## Quelle arborescence de fichiers choisir ?

Ce fut la première question que je me suis posée. Une multitude de
solution s'offre à nous pour répondre à ce problème. De mon côté, j'ai
choisi de partir sur cette arborescence:  en, fr, stylesheets, js et
img. Les dossiers "en" et "fr" abriteront les fichiers html
traduit en fonction du code du dossier (on suppose que les fichiers
équivalents pour les deux langues portent le même nom), pour le reste je
pense que vous avez deviné par vous-même. 

## Quel type d'URL utiliser ?

Il est important de laisser l'utilisateur savoir où il est. Ayant une
architecture basique et des fichiers portant le même nom dans les
différents dossier, il est judicieux (dans mon cas) de laisser
l'arborescence apparaître comme elle est sur le serveur. Par ailleurs,
si vos fichiers ont un nom différent pour des versions équivalentes
(préférable), vous n'êtes pas obligé de montrer le dossier de langue
dans l'URL pour l'utilisateur, cependant le codage du routeur sera
plus ardu. L'idée principale est d'empêcher deux fichiers d'être
définis par la même URL.

## Quelle technologie utiliser ?

Afin de limiter le temps d'exécution entre les requêtes, nous nous
devons d'utiliser un langage permettant une parallélisation facile.
J'ai choisi [Go](http://golang.org/) car 

> "Il est non seulement facile de paralléliser, mais en plus une
> requète unique est elle même traitée très vite; le tout avec un
> langage moderne et facile à utiliser." -
> [Natsirtt](https://twitter.com/natsirtt)

Nous utiliserons cette technologie dans la suite du tutoriel (avec le
paquetage [gorilla mux](http://www.gorillatoolkit.org/pkg/mux)).
Cependant ce routeur est tout à fait réalisable avec une autre
technologie (node.js par exemple).

## Assez parlé, codons !

Euh... On va d'abord mettre en place notre environnement de
production. Assurez vous que votre variable d'environnement GOPATH
contient le dossier de votre projet. Petit rappel : 

```bash
    $ export GOPATH=/le/chemin/vers/votre/projet
```

Créez les dossiers bin, pkg, src au sein de votre projet. Allez dans src
et créez un dossier reprenant le nom de votre projet. Dans mon cas ce
sera 'language-router'.

Donc maintenant dans notre projet nous avons :

```
bin/
pkg/
src/
  language-router/
```

Rendez vous au sein du dossier src. Installons gorilla mux. (mux nous
propose des outils pour la création de routeur).

```bash
$ go get github.com/gorilla/mux
```

Vous devrez alors avoir cette arborescence pour src :

```
src/
    github.com/gorilla/mux/
    language-router/\
```

### MAINTENANT CODONS POUR DE VRAI COMME DES BONHOMMES OU FEMMES TRÈS MUSCLÉES !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

Commençons par créer le fichier language-router.go au sein de
src/language-router/

Pour s'assurer que notre installation est bonne, nous pouvons commencer
par coder un petit hello world comme ceci :

```go
package main

import (
 	"fmt"
)

func main() {
 	fmt.Printf("Hello world!")
}
```

Et si nous sommes aussi génial que nous pensons l'être, en exécutant :

```bash
    $ go run language-router.go
```

Le plus chaleureux message apparaît alors sur votre terminal :

```
Hello world!
```

Maintenant passons aux choses sérieuses.

Nous allons ajouter mux à nos paquetages (ou packages), et commencer à
créer notre propre routeur.

```go
package main

import (
 	"github.com/gorilla/mux"
 	"net/http"
 	"log"
)

func main() {
	router := mux.NewRouter()
        // On attribue le routeur à l'URL "/" (racine).
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
}
```

Lancez le programme (go run language-router). Lancez votre navigateur et
allez sur [localhost](localhost:8080). Vous devrez trouver
normalement une page 404. Ne vous inquiétez pas, tout va bien. Ce que
nous avons fait pour le moment et de créer le routeur et l'attribuer à
la racine du serveur. Cependant notre routeur ne fait strictement rien.
Changeons cela.

```go
package main

import (
 	"github.com/gorilla/mux"
 	"net/http"
 	"log"
)

const(
  // Notre chemin vers notre dossier contenant les fichiers statiques
 	ROOT_PATH = "/home/juanwolf/juanwolf.fr/fr"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  // Sert le fichier index.html
 	http.ServeFile(w, r, ROOT_PATH + "/index.html")
}

func main() {
 	router := mux.NewRouter()
  // Le routeur attribue à l'URL "/", la fonction rootHandler
 	router.HandleFunc("/", rootHandler)
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
}
```

Nous avons ajouté ici une fonction "rootHandler" qui va nous
renvoyer le même fichier (/home/juanwolf/juanwolf.fr/fr/index.html) et
cette fonction sera attribuée à l'url "/" (la racine). Nous avons
cité un peu plus haut que nous voulions une détection automatique du
langage préféré de l'utilisateur cependant pour le moment notre fichier
sera toujours la version française de index.html. IMPLEMENTONS ÇA
OHHHHHHHHHH YEAAAAAAAAAAAAAAH.

### Moi je veux bien, mais comment qu'on fait ?

Demandons à papy [W3C](http://www.w3.org/).

> [For a first contact, using the Accept-Language value to infer
> regional settings may be a good starting point, but be sure to allow
> them to change the language as needed and specify their cultural
> settings more exactly if necessary. Store the results in a database or
> a cookie for later visits.]{style=""font-family:" arial,=""
> 'lucida="" grande',="" helvetica,="" sans-serif;="" =""
> 21.600000381469727px;="" ="" justify;="" =""
> 14px;"=""}[ ]{style=""font-size:" 17.5px;="" ="" 1.42857143;"=""}

(Le lien vers l'article : [Accept-Language used for locale
setting](http://www.w3.org/International/questions/qa-accept-lang-locales#answer))

D'accord, faisons comme ça. Ravis de t'avoir revu l'ami ! 

Donc comme l'a dit papy, on va prendre en compte premièrement
l'attribut Accept-Language de la requête HTTP du client et par la suite
si l'utilisateur décide de changer de langue on stockera son choix dans
un cookie. Mais nous verrons cette dernière partie, un peu plus tard.
Gérons dès à présent les préférences de l'utilisateur :

```go
package main
import (
 	"github.com/gorilla/mux"
	"net/http"
 	"log"
 	"fmt"
 	"strings"
)

const(
 	ROOT_PATH = "/home/juanwolf/juanwolf.fr/"
 	LANG_DEFAULT = "en"
)

var languageMap = map[string]bool{
 	"en": true,
 	"fr": true,
}

func detectLanguageFromHTTPHeader(r *http.Request) string {
  // On récupère le Header de la requète
	header := r.Header
  // On récupère la valeur de "Accept-Language"
 	languagesRequest := header.Get("Accept-Language")
  // On coupe notre chaîne de caractères en fonction des ",".
 	languages := strings.Split(languagesRequest, ",")
 	// Pour chaque valeur dans notre tableau de langue
 	for _, language := range languages {
    // On supprime la qualité (q=*)
 		language_without_quality := strings.Split(language, ";")[0]
    // On supprime la région de la langue
		language_detected := strings.Split(language_without_quality, "-")[0]
    //Si la langue detectée fait partie des langues que nous disposons, on la renvoie.
 		if languageMap[language_detected] == true {
 			return language_detected
 		}
 	}
 	return LANG_DEFAULT
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  // On détecte la langue.
 	language := detectLanguageFromHTTPHeader(r)
  // On modifie le chemin de notre fichier en fonction de la langue
 	http.ServeFile(w, r, ROOT_PATH + language + "/index.html")
}

func main() {
 	router := mux.NewRouter()
 	router.HandleFunc("/", rootHandler)
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
 	}
}
```

Comme vous pouvez le voir sur le code précédent, nous avons modifier
notre variable ROOT_PATH afin qu'elle ne contienne plus la langue dans
son chemin. LANG_DEFAULT servira seulement si le langage préféré de
l'utilisateur n'est pas disponible au sein de languageMap. La variable
languageMap représente les langues que vous mettez à disposition de
l'utilisateur. Avant de nous lancer dans la décortication de la
fonction de détection du langage, regardons de quoi est constitué
'Accept-Language'.

Si on s'amuse à juste retourner la valeur d'Accept-Language, nous
retrouvons une chaîne de caractère de la forme
: en-GB,en;q=0.8,fr-FR;q=0.6,fr;q=0.4,es;q=0.2. Cet utilisateur a donc
pour langue préférée en-GB puis en puis fr-FR. Le q symbolise la
"qualité" de la langue, ce n'est autre qu'un poids pour la
langue. Si la première n'en a pas, c'est qu'elle est définie à 1 par
défaut.

Cela explique pourquoi nous avons autant spliter nos chaînes de
caractères. On réutilise donc la noble fonction précédemment créée dans
notre rootHandler. Et voilà !

### On devait faire apparaître la langue dans l'URL, non ?

En effet. Faisons ça !

```
package main
import (
 	"github.com/gorilla/mux"
	"net/http"
 	"log"
 	"fmt"
 	"strings"
)

const(
 	ROOT_PATH = "/home/juanwolf/juanwolf.fr/"
	LANG_DEFAULT = "en"
)

var languageMap = map[string]bool{
 	"en": true,
 	"fr": true,
}

func detectLanguageFromHTTPHeader(r *http.Request) string {
 	header := r.Header
 	languagesRequest := header.Get("Accept-Language")
 	languages := strings.Split(languagesRequest, ",")
 	fmt.Println(languages) 	for _, language := range languages {
 		language_without_quality := strings.Split(language, ";")[0]
 		language_detected := strings.Split(language_without_quality, "-")[0]
 		if languageMap[language_detected] == true {
 			return language_detected
 		}
 	}
 	return LANG_DEFAULT
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
 	language := detectLanguageFromHTTPHeader(r)
 	http.Redirect(w, r, r.URL.Path + language + "/", http.StatusFound)
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
  // On récupère la valeur de {lang} de l'URL
 	vars := mux.Vars(r)
 	langAsked := vars["lang"]
  // Si on a la langue, on sert le fichier.
 	if languageMap[langAsked] {
 		http.ServeFile(w, r, ROOT_PATH  + langAsked + "/index.html")
 	}
}

func main() {
 	router := mux.NewRouter()
 	router.HandleFunc("/", rootHandler)
  // On ajoute une URL qui prendra en compte la langue. ("/fr/ ou "/en/")
 	router.HandleFunc("/{lang}/", indexHandler)
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
}
```

Qu'est ce qui a changé ? Premièrement notre rootHandler ne sert plus de
fichiers mais renvoie vers l'URL constituée de la langue. Le routeur
dispose maintenant d'un nouveau Handler pour les URLs de langues
définies par indexHandler. 

### Gérons maintenant le cookie.

Pour cette partie nous supposons que votre site comporte un système de
choix de la langue tel qu'un select ou autre et que ce système met en
place un cookie contenant la langue choisie par l'utilisateur. Si votre
site met déjà à disposition un cookie contenant les préférences de
langue de l'utilisateur, vous pouvez passez la partie suivante.

### Atelier cuisine

Créons notre cookie du côtés des fichiers statiques. Nous allons avoir
maintenant besoin des fichiers javascript. On va donc ajouter une route
pour que notre routeur nous fournisse ces fichiers statiques.  

```go
func main() {
 	router := mux.NewRouter()
 	// Router section
 	router.HandleFunc("/", rootHandler)
 	// On sert les fichiers javascript
 	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
 		http.FileServer(http.Dir(ROOT_PATH + "js/"))))
 	router.HandleFunc("/{lang}/", indexHandler)
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
}
```

Notre fonction main() se transforme ainsi. Faites attention à bien
définir la route /js/ avant la route spécifique pour les langues sinon
vos fichiers ne seront jamais servis ! Ajoutons maintenant nos fonctions
à notre fichier javascript,on suppose que nous ferons la transition
entre les langues avec le select :

```html
    <select id="language-selection-select">     <option value="en">English </option>     <option value="fr">Français</option> </select>
```

Le JavaScript qui lui est joint.

```javascript
function saveLanguageChoosen($language) {
  document.cookie = "";
  var d = new Date();
  var exdays = 7;
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  var expires = "expires=" + d.toGMTString();
  console.log("language received = " + $language);
  document.cookie = "lang=" + $language +";expires=" + expires
               + ";domain=.juanwolf.fr;path=/";
}

$(document).ready(function() {
     $("#language-selection-select").change(function() {
         saveLanguageChoosen($(this).val());
     });
});
```

Maintenant nous avons notre propre cookie avec la langue voulue. On
pourrait ajouter à ce javascript une fonction de redirection lors du
changement de la langue. Si vous voulez que ce cookie soit disponible
sur tout votre site, le point est très important dans la définition du
domaine du cookie (domain=.domain.org)

REPASSONS AU GO
---------------

Nous devons maintenant lire le cookie s'il est présent sur
l'ordinateur de l'utilisateur.

```go
package main

import (
 	"github.com/gorilla/mux"
 	"net/http"
 	"log"
 	"strings"
)

const(
 	ROOT_PATH = "/home/juanwolf/juanwolf.fr/"
 	LANG_DEFAULT = "en"
         // Le nom du cookie
 	COOKIE_NAME = "lang"
         // La langue de l'attribut au sein du cookie.
 	COOKIE_LANG_ID = "lang"
)
// Les langues que nous mettons à disposition
var languageMap = map[string]bool{
 	"en": true,
 	"fr": true,
}

func readCookie(r *http.Request) string {
         // On lit le cookie.
 	cookie,err := r.Cookie(COOKIE_NAME);
         // S'il n'existe pas, on renvoie une chaine vide.
 	if (err != nil) {
 		return ""
	}
 	language := ""
 	cookieVal := strings.Split(cookie.String(), ";");
 	for i := 0; i < len(cookieVal); i++ {
		if strings.Contains(cookieVal[i], COOKIE_LANG_ID) {
 			langArray := strings.Split(cookieVal[i], "=")
 			language = langArray[1]
 		}
 	}
 	return language

}

func detectLanguageFromHTTPHeader(r *http.Request) string {
 	header := r.Header
 	languagesRequest := header.Get("Accept-Language")
 	languages := strings.Split(languagesRequest, ",")
	for _, language := range languages {
 		language_without_quality := strings.Split(language, ";")[0]
 		language_detected := strings.Split(language_without_quality, "-")[0]
 		if languageMap[language_detected] == true {
 			return language_detected
		}
 	}
	return LANG_DEFAULT
}

func detectLanguage(r *http.Request) string {
 	cookieResult := readCookie(r)
 	if cookieResult != "" {
 		return cookieResult
 	} else {
		return detectLanguageFromHTTPHeader(r)
 	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
 	language := detectLanguage(r)
 	http.Redirect(w, r, r.URL.Path + language + "/", http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
 	vars := mux.Vars(r)
	langAsked := vars["lang"]
	if languageMap[langAsked] {
		http.ServeFile(w, r, ROOT_PATH + "/" + langAsked + "/index.html")
 	}
}

func main() {
 	router := mux.NewRouter()
 	router.HandleFunc("/", rootHandler)
 	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
 		http.FileServer(http.Dir(ROOT_PATH + "js/"))))
 	router.PathPrefix("/stylesheets/").Handler(http.StripPrefix("/stylesheets/",
 			http.FileServer(http.Dir(ROOT_PATH + "stylesheets/"))))
 	router.HandleFunc("/{lang}/", indexHandler) 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err) 	}
}
```

Et voilà, nous avons donc notre routeur de langue flambant neuf ! Bien
sûr, il peut être soumis à des améliorations telles que la détection
automatique des langues disponibles sur le serveur, ou la définition
d'une 404 ou de sous domaine. Vous pourrez retrouver le lien du dépôt
de ce projet
[ici](https://bitbucket.org/juanwolf/language-router).

En attendant, codez bien !
