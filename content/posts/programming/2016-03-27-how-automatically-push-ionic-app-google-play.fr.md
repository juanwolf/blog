
---
title: Intégration continue d'un projet ionic avec Jenkins
date: 2016-03-27
tags: ["Jenkins", "ionic"]
categories: ["Programmation"]

slug: integration-continue-ionic-avec-jenkins
aliases:
  - /programmation/integration-continue-ionic-avec-jenkins/

draft: false
author: "Jean-Loup Adde"
---

Salut tout le monde, aujourd'hui on va voir comment pousser
automatiquement chaque nouvelle mise à jour de notre application ionic
sur le Google Play.

![](/post_preview/20160327_124355_article-banner-ionic-ci-bis.png)

Configuration sur le serveur
----------------------------

Nous allons premièrement installer les dépendances sur le serveur sur
lequel vous avez jenkins.

### Installation de ionic

    sudo npm install -g ionic@beta

Ici j'utilise la version beta de ionic, vous pouvez très bien
l'enlever si vous utilisez la v1.

### Installation de cordova

    sudo npm install -g cordova

### Installation du JDK

Tout va dépendre de votre distribution, perso, je suis sous Archlinux,
du coup j'ai tout installer avec pacman.



    sudo pacman install oracle-jdk



Pour les utilisateurs d'Ubuntu, je vous invite à lire [cet
article](https://www.digitalocean.com/community/tutorials/how-to-install-java-on-ubuntu-with-apt-get%22).



### Installation du SDK Android



Vous devez installer la version 22 de l'API Android. Pour cela, lancez
l'Android SDK Manager, et choisissez l'API version 22.



Si vous avez un écran, vous pouvez utiliser l'Android SDK manager de
manière visuelle en lançant la commande :



    android sdk



### Génération du keystore



Afin de pouvoir signer notre application Android, nous allons devoir
créer un keystore que nous utiliserons à chaque construction de
l'application.



Pour cela exécutez la commande suivante:



    keytool -genkey -v -keystore lenomdevotrekeystore.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 10000



La commande vous demandera un mot de passe afin de protéger votre
keystore, gardez le bien en tête car nous en aurons besoin pour la
configuration du build jenkins.



Configuration de jenkins
------------------------

### Installation des plugins

Pour continuer, vous aurez besoin des plugins :

-   GooglePlayPluginPusher
-   EnvInject

Configuration de votre projet sur la Google Developers Console
--------------------------------------------------------------

Vous devez configurer la [Google Developers
Console](https://play.google.com/apps/publish/#ApiAccessPlace)
afin de créer des "credentials" pour votre jenkins afin qu'il
puisse pousser les modifications de votre APK sur le Google Play.



Dans Google Play Developers Console -\> Paramètres -\> Accès à l'API
-\> Comptes de services -\> Créer un compte de service -\> Afficher dans
Google Developers Console -\> Utiliser p12 -\> Télécharger
p12



Maintenant ajoutons cette clé à jenkins:



Dans Jenkins -\> Paramètres -\> Credentials -\> Ajouter Crédentials -\>
Ajouter Credentials Play Store -\> Balancer le p12



Création du job jenkins
-----------------------



### Définition des variables d'environnement



Dans la première section choisissez "This build is parametrized",
"Add Parameter", "String Parameter"



-   Name : STORE\_PASSWORD
-   Default value: Le mot de passe utilisé pour créer votre keystore.



Ajoutons une deuxième variable du même type:



-   Name: PATH\_TO\_KEYSTORE
-   Default value: Le chemin où vous avez créé votre keystore.



### Définition de la récupération du code source



#### 'Git Flow'

Afin de détecter les releases faites de notre application, nous allons
utiliser un pattern simple avec l'outil de versionning que vous
utilisez.

Premièrement, ne commitez jamais sur votre branche master, celle ci ne
sera utilisée que pour faire des merge et définir les releases de votre
application. Créez une branche dev afin de commitez tous les changements
fait de votre application (bref ce que vous utilisiez auparavant quoi).



Cas d'utilisation : Je commit plusieurs fois sur ma branche dev
jusqu'à temps que ma feature soit terminée, je veux envoyer cette
feature le Google Play. Je modifie la version de mon application dans le
fichier config.xml à la racine de mon projet. Je commit avec un message
annonçant la release: "release v0.0.4" (par exemple), je tag la
branche avec la version de ma release, je pousse le tout sur dev. Je
checkout sur master, je merge dev dans master. Je pousse master. Le
build jenkins détecte un changement, l'apk avec ma release est
construit et est déployé sur le Google Play.



#### Revenons à jenkins



Dans la section Source Code Management, choissisez votre logiciel de
versionning, ajoutez votre URL de votre dépôt (un truc du genre
git\@gitlab.com:juanwolf/blogapp.git) et définissez vos credentials
(username, password ou clé ssh)



Dans la section branches to build, ajoutez: "\*/master", Jenkins
va alors utiliser les sources que vous avez poussées sur master (ce que
l'on veut). Si jamais vous voulez un build spécifique pour votre
environnement de développement, faites de même et spécifiez la branche
\*/dev (je pense que vous aviez compris l'idée)

### Activation du build



Dans la section "Build Triggers", dans poll SCM, ajoutez



    */5 * * * *



Pour les néophites, ceci est un cron, on dit à jenkins de checker si les
sources ont été modifiées toutes les 5 minutes



Nous allons créer le script de construction de l'APK:



    \r\nnpm install # Super important (Installe toutes les dépendances de votre projet)\r\nionic build android --release; # Création de l'APK de release.\r\njarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 -keystore $PATH_TO_KEYSTORE platforms/android/build/outputs/apk/android-release-unsigned.apk alpha -storepass $STORE_PASSWORD # Signature de l'APK\r\n/opt/android-sdk/build-tools/23.0.2/zipalign -v 4 platforms/android/build/outputs/apk/android-release-unsigned.apk platforms/android/build/outputs/apk/$BUILD_ID.le.nom.de.votre.app.apk  # Création de l'APK à déployer.\r\n



### Publication de l'APK dans le Google Play



Ajoutez une "Post build action" à votre build "Upload Android
APK to Google Play".



Dans la section "Google Play Account", choisissez le crendentials
que vous avez créé avec le clé p12. Dans "APK files" ajoutez:



    platforms/android/build/outputs/apk/$BUILD_ID.le.nom.de.votre.app.apk



Dans "release track", mettez "production" (ou beta ou alpha)



Dans "Recent Changes", bon, je suis un peu bloqué sur cette
partie. Je cherche à modifier ce message depuis un changelog mais je
n'ai pas encore trouvé comment faire. Si jamais vous l'avez fait ou
avez une idée, je suis plus que preneur !!! Pour le moment, j'ai mis un
simple texte du genre "De nouvelles fonctionnalités ont été ajoutées,
veuillez vous rendre sur cette page pour en savoir plus".

Done !
------

Et voilà ! Nous avons notre job jenkins qui va déployer à chaque push
sur la branche master de notre application ionic. Vous pouvez bien sûr
ajouter des étapes à ce build jenkins tel que lancer les tests, générer
des rapports, etc... Sur ce, codez bien. Ciao !

