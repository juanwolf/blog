
---
title: Mitraillez votre web application avec Gatling !
date: 2015-10-21
tags: ["Jenkins", "scala"]
categories: ["Programmation"]

slug: mitraillez-votre-web-application-avec-gatling
aliases:
  - /programmation/mitraillez-votre-web-application-avec-gatling/

draft: false
author: "Jean-Loup Adde"
---

 Bonjour tout le monde, ça fait un bail ! Aujourd'hui on va
regarder le framework de tests de performances Gatling que j'ai pu
utiliser durant mon stage de fin de Master et vous montrer à quel point
il est facile à prendre en main !

![](/post_preview/20151021_200528_221px-Gatling_\(load_testing_tool\)_Logo.png)

Introduction
------------
### Pourquoi Gatling ?

1.  DSL pour l'écriture de scénarios
2.  Haute performance
3.  Rapport HTML généré après votre simulation
4.  Communauté ultra réactive
5.  Intégration avec Graphite en une ligne
6.  Et pour les gens pressés, il y a même un outil pour créer un
    scénario gatling à partir de votre utilisation de votre webapp
    depuis un navigateur


Détection des cas d'utilisations
---------------------------------

Afin d'optimiser vos tests de performances, vous aurez besoin de
détecter les cas d'utilisations de votre webapp, on appelera cela
des "Personas". Dans le cas de mon blog, j'en ai détecter
deux :

-   Les lecteurs ayant suivi un lien
-   Les gens qui sont venus par hasard

 Dans le premier cas, on aura une lecture longue de l'article
sans parcours du site et dans le deuxième cas, on aura un parcours
un peu random du site.

Je vais en profiter pour créer des objets regroupant les actions
définies dans les cas d'utilisations. Dans mon cas, je ne teste
que tous les différents liens disponibles sur le site, la page
d'accueil, une page tag, une page de catégorie et une page
d'article. Au sein du fichier Scenario.scala:

```scala
// Object regroupant les fonctions de navigations sur le site
object Browse {

    val goToIndex = exec(
        http("Go To Index")
        .get(indexPageUrl)
        .check(status.is(200))
    )

    val goToTag = exec(
        http("Go To Tag")
        .get(indexPageUrl + "/" + tag)
        .check(status.is(200))
    )

    val readArticle = exec(
        http("Go To Article")
        .get(indexPageUrl + "/" + article)
        .check(status.is(200))
    )

    val goToCategory = exec(
        http("Go To Category")
        .get(indexPageUrl + "/" + category)
        .check(status.is(200))
    )
}
```


Si par exemple on voulait tester les cas où un visiteur ajoute
des commentaires, on aurait ajouté un object 'Comment' contenant
toutes les fonctions possibles (add, edit, respond, etc...)



Représentation des personas
---------------------------

Maintenant que nous avons les fonctions pour simuler des utilisateurs
définissons nos personas !

Au sein du fichier Scenario.scala

```scala
class Scenario extends Simulation {
    // Variables de configurations (voir section feeders)
    val indexPageUrl = "http://blog.juanwolf.fr"
    val tag = "jenkins"
    val category = "video-games/"
    val article = "2015/3/continuous-integration-django-jenkins"

    object Browse { /*...*/ }

    val linkFollower = scenario('Link Follower').exec(
        Browse.readArticle,
        Browse.goToIndex
    )
    val randomers = scenario('Randomers').exec(
        Browse.goToIndex,
        Browse.goToTag,
        Browse.goToCategory,
        Browse.readArticle
    )

    // ....
}
```

Ici, nous n'avons pas d'authentification ou de différent types
d'utilisateurs donc la définition de personas reste assez simple.

Développement des scénarios
---------------------------

### Simulation

 Maintenant que nous avons nos personas, on peut passer à la phase
de création de la simuation. Nous allons avoir besoin de :

-   Fixer les objectifs de votre webapp (Temps de réponse,
    disponibilité, etc\...)
-   Définir quantativement les salves d'utilisateur que vous allez
    envoyer

Dans mon cas, je veux que mon site réponde en moins de 2 secondes
pour 100 utilisateurs simultanés histoire d'éprouver les
performances du blog. (Oui je n'ai pas un serveur très puissant
pour faire tourner ce site :/).

Au sein du fichier Scenario.scala

```scala
class Scenario extends Simulation {
  /*...*/

  setUp(
    linkFollower.inject(atOnceUsers(100))
    randomer.inject(atOnceUsers(100))
  )

}
```

Pour exécuter notre simulation, on va devoir utiliser le plugin maven de
gatling. Au sein de votre pom.xml, on va ajouter ces lignes:

```xml
<build>
    <plugins>
        <plugin>
            <groupId>io.gatling</groupId>
            <artifactId>gatling-maven-plugin</artifactId>
            <version>${gatling-plugin.version}</version>
            <executions>
                <execution>
                    <phase>test</phase>
                    <goals>
                        <goal>execute</goal>
                    </goals>
                    <configuration>
                        <!-- Default values -->
                        <dataFolder>src/main/resources</dataFolder>
                        <resultsFolder>target/gatling/results</resultsFolder>
                        <simulationsFolder>src/main/scala</simulationsFolder>
                        <simulationClass>fr.juanwolf.scenarios.Scenario</simulationClass>
                    </configuration>
                </execution>
            </executions>
        </plugin>
    </plugins>
</build>
```

Maintenant, nous pouvons exécuter notre simulation avec la directive
:

    mvn test

### Les sessions

Les sessions vont vous permettre de stocker des informations
durant l'exécution de votre scénario. Ces sessions sont
grossièrement une Map de String, Any. Vous pourrez injecter au sein de
vos sessions des données grâce aux feeders ou récuperer des
données avec l'EL (Expression Language) de gatling ou si aucun des
entités précédentes ne correspondent à vos besoins, vous pouvez
toujours intérragir directement aux sessions avec son api.

#### Les feeders

Les feeders ont été conçus pour vous permettre d'injecter des données à
partir d'un élément particulier. Vous pourrez alimenter vos scénarios à
partir de :

-   CSV feeders
-   JSON feeders
-   JDBC feeders
-   Sitemaps feeders

 Par exemple, dans mon cas, j'utiliserai le sitemap feeder afin
de parcourir tout mon site et effectué des tests sur chaque
article. Tiens, je vais ajouter ça à mon trello.

#### L 'EL de gatling

Grâce à l'EL de gatling (EL pour expression language), vous pourrez
accéder facilement à vos données au sein de la session. Ce langage vous
permettra d'éviter de manipuler la MAP de la session et même
d'utiliser l'API session ce qui est agréable.  Vous pouvez
accéder à vos données en spécifiant \${mavariableensesssion}. Dans le
cas où vous auriez stocké un json, objet, map, vous pouvez utiliser le
jsonpath associé pour accéder à cet élément. Exemple: Je stocke au sein
de la session ces variables :

```scala
record1: Map(
           "loc" -> "http://www.example.com/",
           "lastmod" -> "2005-01-01",
           "changefreq" -> "monthly",
           "priority" -> "0.8")

record2: Map(
           "loc" -> "http://www.example.com/catalog?item=12&desc=vacation_hawaii",
           "changefreq" -> "weekly")

record3: Map(
           "loc" -> "http://www.example.com/catalog?item=73&desc=vacation_new_zealand",
           "lastmod" -> "2004-12-23",
           "changefreq" -> "weekly")
```

Je peux accéder au deuxième lien grâce à l'EL:

    ${record2.loc}

Stylé non ?

PS: J'oublais cet EL n'est utilisé qu'au sein de string, donc pour
accéder à ces données vous devez les entourées de "". Donc pour
l'exemple précédent, je fais ma déclaration ainsi :

    var url: String = "${record2.loc}"

#### L'API session

Avec cet API, vous pourrez manipuler directement les attributs au sein
de la session grâce à des getter et setter. Exemple :

```scala
val scn = scenario("Test Sessions")
.exec(session => {
    session.set("key", "value").set("Bernard", "Lama")
    // On affiche le contenu de la session pour la clé "Bernard"
    println("BERNARD" + session("Bernard").as[String])
    session
})
```

**WARNING:** Faites attention aux getters, si vous ne spécifiez pas le
type de retour attendu, la session vous retournera un wrapper de type
SessionAttribute, donc n'oubliez pas le .as[LeType].

#### Les goals

Avoir des requêtes, c'est bien gentil mais comment vérifier que nos
attentes sont respectées ? Pour cela, vous devez aggréger vos scénarios
de "checks" et d'assertions.

Premièrement, vous pouvez spécifier quel est le code de retour HTTP
attendu. On a déjà rencontré ce type de verification au sein du fichier
User/WebSurfer. Reprenons un appel de ce fichier :

```scala
val goToIndex = exec(
    http("Go To Index")
    .get(indexPageUrl)
    .check(status.is(200))
)
```

C'est à ça que correspond la dernière partie. Nous allons vérifier que
le retour HTTP de l'appel est 200. Sinon on considère l'appel comme
échoué.

Et nous avons les assertions qui vont permettre de déterminer si un
appel est considéré comme échoué au niveau des performances de celui-ci.
Contrairement aux précédentes assertions, celles-ci se situent au niveau
de la configuration des scénarios.

Si on reprend le fichier Scenario.scala précédent, cela nous donnera :

```scala
setUp(
    websurfers.inject(atOnceUsers(150))
).assertions(
    // On veut que la globalité du blog réponde en moins d'une seconde et demie
    global.responseTime.max.lessThan(1500)

    // On ne veut qu'aucune requete n'échoue sur la globalité du site
    global.failedRequests.is(0)

    // On peut même spécifier des attentes pour des requetes spécifiques
    details("Go To Index").responseTime.max.lessThan(1000)
)
```

Les rapports
------------

Après chaque simulation, Gatling aura créé les résultats au sein
du dossier target. Ces résultats sont sous formes de fichier HTML.
Au sein de ces rapports, vous trouverez différentes sections. Une
globale et une détaillée pour chaque groupe de requètes.


![Rapport]/post_content/2015-10-23/2e7c0992-a166-4f10-9c0f-cdd21d92c205.png)


Sur ce rapport, on peut voir que la plupart des requêtes se sont mals
passées (temps supérieur à 1,2s et quasiment 20% ont échoué :'( :'()

Monitoring
----------

 Ce qui est aussi plutôt cool avec Gatling est sa facilité à être
couplé avec des outils de monitoring (surtout graphite). Il
suffit d'ajouter quelques lignes à votre configuration de gatling pour
que vos données soient publiées au sein de l'outil (développé en
django pour votre information)

Au niveau code, ça se présente ainsi :

-   Ajouter un gatling.conf dans vos ressources
-   Ajouter les lignes ci-dessous au sein du fichier créé précédemment

 Bam, juste ça et tout est envoyé sur graphite, cool non
? Vous pouvez même spécifier si vous voulez seuelement
envoyer les données globales (all), choisir le protocole,  change
la racine de publication des données et le taille du buffer

Tricks
------

### Intégration avec Jenkins

 Vous pouvez coupler vos tests de performances avec des serveurs
d'intégration continue tel que Jenkins. Personnellement, je m'en
sers pour lancer les tests de performances à partir d'un Cron mais
aussi pour afficher les tendances directement au sein de Jenkins
grâce au plugin jenkins de gatling. Ça ressemble à ça :
![Affichage](/post_content/2015-10-23/2023ad70-d2bb-4669-972d-d6a23f142899.png)

### Débugger

Pour débugger, c'est assez simple. Vous devez utiliser la
fonction exec de gatling et utiliser la variable session afin de
pouvoir débugger votre code. Exemple :

```scala

    val scn = scenario("Mon super scénario")
    .exec(http(""))
   .exec(session => {
        println(session)
        session // Pensez toujours à renvoyer la session à la fin de votre debug.
    })
```


Conclusion
----------



 Je pense que nous avons fait le tour. Bien sûr, je ne vous ai pas
tout montré ou avoir été beaucoup dans les détails.  Pour cela, je
vous invité à aller voir la documentation qui est très fournie. De plus,
j'ai conçu cet article autour de la campagne de tests de
performance sur ce site que vous pouvez trouver
[ici](https://gitlab.com/juanwolf/blog-stress-test). Si
vous avez des questions, n'hésitez pas à poster un commentaire, je me
ferai une joie de vous répondre (A moins que ce soit pour \#!/bin/bash
juanwolf , (MDR))


En espérant vous voir coder quelques scénarios, je vous laisse, j'ai de
nouveaux rollers à tester. Sur ce, codez bien.
