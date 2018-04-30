
---
title: How to make a language router with Go?
date: 2014-06-18
tags: ["Go"]
categories: ["Programming"]

draft: false
author: "Jean-Loup Adde"
---

If you asked yourself the same question as me "How could I serve my static files for different languages ?", so this post is exactly for you ! Let's see how I did.

![](/post_preview/20150322_141925_gogogo1.jpg)

## Which file tree should we choose ?
It was the first question I asked myself. A lot of solution is possible to solve this problem. In my side, I choosed this file tree:  en, fr, stylesheets, js and img. The directories "en" and "fr" will have the html files translated (We suppose that the equivalent files has the same name),  and the other directories will contain all the css and js.

## Which kind of URL should we use ?
It's important that the user know where he is. As we have a basic architecture and files with the same name in the language folders, it's smart to let the file tree appears as it is on our server. But if you have different names for equivalent files, it's not important if the language appears in the URL, but the router will be harder to make. The idea is just to have different URL for every single html files.

## Which technology should we use ?
To limit time between requests, we have to use a language which give us the possibility to parallelize easily our router. I choose Go because

"It's also easy to parrellize it but one request it's traited fastly but with a modern language easy to use." - Natsirtt
and we will use this technology during all the tutorial (with the mux package). But this router is makeable with others technologies (as node.js).
STOP TO TALK, START TO CODE !

Hum... Well, calm down, we will first set up our production environment. Be sure that the environment variable GOPATH contain your project folder. Little reminder:

```bash
$ export GOPATH=/the/path/to/your/project
```

Now we have to create the directories bin, pkg, and src inside our project. Go (AHAHAHAHA) inside the src directory and create the folder with the same name as your project. IN my case it will be 'language-router'.

So now, in our project we have:

```
bin/
pkg/
src/
  language-router/
```

Let's go inside the src folder. We'll install the mux package (the package gorilla/mux implements a request router and dispatcher).

```bash
$ go get github.com/gorilla/mux
```
You should have this file tree for src :

```
src/
    github.com/gorilla/mux/
    language-router/
```

NOW LET'S GOTO CODE LIKE A FELLOW OR WOMEN WITH MUSCLES VERY DEVELOPED !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Let's begin to create the language-router.go file inside the src/language-router/

To make sure our environment is setted up, we can start to write a little hello world program as below :

```go
package main
import (
    "fmt"
)
func main() {
    fmt.Printf("Hello world!")
}
```
And if we are awesome as we think, executing :

```bash
$ go run language-router.go
```
The most beautiful message appears inside your terminal :

```
Hello world!
```

Now, let's start serious stuff.

We will add gorilla/mux at our packages, and start to create our own router.

```go
package main
import (
 	"github.com/gorilla/mux"
 	"net/http"
 	"log"
)
func main() {
    router := mux.NewRouter()
    http.Handle("/", router)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

Run the program (go run language-router). Start your browser and go to localhost. Normally, you'll find a 404 page. Don't worry, everything is fine. What we did for the moment is just to create the router and define it at the root of the server. However our router doesn't do anything. Let's change that !

```go
package main
import (
	"github.com/gorilla/mux"
 	"net/http"
 	"log"
)

const(
 	ROOT_PATH = "/home/juanwolf/juanwolf.fr/fr"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
 	http.ServeFile(w, r, ROOT_PATH + "/index.html")
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

We had a function "rootHandler" which return the  the same file (/home/juanwolf/juanwolf.fr/fr/index.html) and this function will be affected to the root URL ("/"). We said earlier that we wanted an automatic detection of user's preferred language but for the moment the file will be only on his french version :/ Nous avons cité un peu plus haut que nous voulions une détection automatique du langage préféré de l'utilisateur cependant pour le moment notre fichier sera toujours la version française de index.html. LET'S IMPLEMENT THAT OOOOOOOOOOOOOH YEEEEEEEEEEEEEEAAAAAAAAAAAAAAAAAAAAAH.

Ok why not, but how we do ?
Let's ask to granpa' W3C.

> For a first contact, using the Accept-Language value to infer regional settings may be a good starting point, but be sure to allow them to change the language as needed and specify their cultural settings more exactly if necessary. Store the results in a database or a cookie for later visits.
(The article link : Accept-Language used for locale setting)

So, as granpa' said, we'll take first the Accept Language inside the HTTP request and  if the user wants to change, we'll keep his inside a cookie. Now, we have to manage the Accept-Language value :

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
 	header := r.Header
 	languagesRequest := header.Get("Accept-Language")
 	languages := strings.Split(languagesRequest, ",")
 	fmt.Println(languages)
 	for _, language := range languages {
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

As you can see on this previous code, we modified our ROOT_PATH variable for she doesn't contains the language in it path. LANG_DEFAULT  will be usefull only if the preferred language is available in the languageMap variable. In this variable, we will put all the language available on our server. Before to start to see how was built the language detection function, let's see what contains the Accept-Language value.

If we play to return the Accept-Language value, we can see a string like that: en-GB,en;q=0.8,fr-FR;q=0.6,fr;q=0.4,es;q=0.2. This user has for preferred language en-GB next en next fr-FR. The "q" symbolize the language quality, It's just a value to make an order between the languages. If the first doesn't ave a q value, it's because she has the default value "1".

It explains why we had to split our strings. We use the function previously created in our rootHandler. Et voilà !

We had to show the language in the URL, right ?
Yeah you're right ! I already forgot. Let's do this!

```golang
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
 	fmt.Println(languages)
  for _, language := range languages {
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
 	vars := mux.Vars(r)
 	langAsked := vars["lang"]
 	if languageMap[langAsked] {
 		http.ServeFile(w, r, ROOT_PATH  + langAsked + "/index.html")
 	}
}

func main() {
 	router := mux.NewRouter()
 	router.HandleFunc("/", rootHandler)
 	router.HandleFunc("/{lang}/", indexHandler)
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
}
```

What changed ?Firstly our rootHandler doesn't serve files but redirect to the URL made with the language. The router has now a new handler for the languages URL defined by indexHandler.

### IT'S TIME TO MANAGE THE COOKIE.
For this party, we suppose that our site contains a system which give to the user the posibility to change the language and create a cookie containing it. If you're website has one, you can pass the section below.

### Time to cook
It's time to create our cookie. We'll now need to code it with javascript. We'll add a path to our router that provide javascript files.

```go
func main() {
 	router := mux.NewRouter()
 	// Router section
 	router.HandleFunc("/", rootHandler)
 	// Static js files
 	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
 		http.FileServer(http.Dir(ROOT_PATH + "js/"))))
 	router.HandleFunc("/{lang}/", indexHandler)
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
}
```
Be careful to define your path /js/ before the specified path for the languages because if you no your files will never be served ! Now we add our select for the language selection:

```javascript
select id="language-selection-select"
     option value="en">English /option    option value="fr">Français /option /select
The javascript associated:

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
Now we have our proper cookie with the language wanted. We could add to our javascript a redirection function when the user change the language. If you want that your cookie is available in all your website and all subdomain you have to add the '.' before your domain (domain=.domain.org).

GO BACK TO THE GO GO GO GO
Now we have to read the cookie if it's available on the client machine.

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
 	COOKIE_NAME = "lang"
 	COOKIE_LANG_ID = "lang"
)

var languageMap = map[string]bool{
 	"en": true,
 	"fr": true,
}
func readCookie(r *http.Request) string {
 	cookie,err := r.Cookie(COOKIE_NAME);
 	if (err != nil) {
 		return ""
	}
	language := "" 	cookieVal := strings.Split(cookie.String(), ";");
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
 	language := detectLanguage(r) 	http.Redirect(w, r, r.URL.Path + language + "/", http.StatusFound)
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
 	router.HandleFunc("/{lang}/", indexHandler)
 	http.Handle("/", router)
 	if err := http.ListenAndServe(":8080", nil); err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
}
```
Et voilà ! We have our own router ! Obviously, it can be improved like with a automatic detection of the languages inside the server or define a 404 page or subdomains. You can find the link to my repository here.

En attendant, codez bien ! xoxo.
