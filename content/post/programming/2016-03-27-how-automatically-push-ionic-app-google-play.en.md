
---
title: How to automatically push your Ionic app changes to Google Play
date: 2016-03-27
tags: ["Jenkins", "ionic"]
categories: ["Programming"]

slug: how-automatically-push-ionic-app-google-play
aliases:
  - /programming/ow-automatically-push-ionic-app-google-play/

draft: false
author: "Jean-Loup Adde"
---

Hi everyone, today we'll see how to automatize every new functionality
that we'll build in our little ionic app to the Google Play.

![](/post_preview/20160327_124355_article-banner-ionic-ci-bis.png)


Server configuration
--------------------



Firstly, we'll install all the dependencies on the server where's
Jenkins.



### Ionic installation



    sudo npm install -g ionic@beta



Here I use the beta version of ionic (to use angular 2 =D), so you can
remove the \@beta if you're on a v1 project.



### Cordova installation



    sudo npm install -g cordova



### JDK installation



Here, it will depend on your linux distribution. I'm using Archlinux,
so a simple command with pacman and everything is installed.



    sudo pacman install oracle-jdk



For Ubuntu users, I invite you to read [this
article](https://www.digitalocean.com/community/tutorials/how-to-install-java-on-ubuntu-with-apt-get%22).



### Android SDK installation



You have to install the 22 version of the android API (why? I don't
have a god damn clue). For that, launch the Android SDK manager and
choose to install the 22 API version.



If you're not using your server with only a terminal, you can use the
android sdk manager visually using this command:



    android sdk



### Keystore generation



To sign our android application, we'll need to create a keystore that
we'll use to sign our application.



To create it, run this command:



    keytool -genkey -v -keystore thenameofyourkeystore.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 10000



The command should ask you a password to protect your keystore, keep it
in mind because we'll need it to configure the next jenkins build.



Jenkins configuration
---------------------



### Plugins installation



To continue, you'll need this plugins:



-   GooglePlayPluginPusher
-   EnvInject



Project configuration on the Google Developers Console
------------------------------------------------------



You have to configure the [Google Developers
Console](https://play.google.com/apps/publish/#ApiAccessPlace)
to create specific credentials for jenkins. These credentials will allow
jenkins to push remotely your APK to the Google Play.



In the Google Play Console Developer -\> Settings -\> API Credentials
-\> Services Account -\> Create a service account -\> Display in the
Google Developers Console -\> Use p12 -\> Download p12

Now let's add our p12 key to Jenkins:

In Jenkins -\> Settings -\> Credentials -\> Add Credentials -\> Add Play
Store Credentials -\> Put your p12 key

Jenkins job creation
--------------------



### Environmnent variable definition



In the first section, choose "This build is parametrized", "Add
Parameter", "String Parameter"



-   Name : STORE\_PASSWORD
-   Default value: The password use to create your keystore.



Let's add a second variable of the same type:



-   Name: PATH\_TO\_KEYSTORE
-   Default value: The path you saved your keystore.



### Source Code Management



#### 'Git Flow'



To detect each release made for our application, we'll need to use a
simple pattern with the versionning tool you're using.



Firstly, never commit on the master branch. This branch will be only
used to merge our work from other branches and will be used to define
releases of the application. Create a new branch called dev that we'll
use to commit each modification of the code application.



If it's not clear, let's make a use case: I commit multiple times on
my "dev" branch until my new feature is available. I want to send
this new features straight after coding it to users (so pushing it to
the Google Play store). I modify the version inside the config.xml file
in my project root incrementing the version number. I commit with a
comment like "Release v0.0.4". I tag my branch with the release
version number. I push everything on the dev remote branch. I checkout
to the master local branch. I merge dev into master. I push master.
Jenkins detects a change on the master branch, it polls the code, build
the apk, sign it, and push it the Google Play.



#### Let's go back to Jenkins



In the Source Code Management, choose your versionning software, add the
project url repository (something like
git\@gitlab.com:juanwolf/blogapp.git) and define the credentials to
access to it (username, password or ssh key)



In the "branches to build" section, add: "\*/master".
Jenkins will then use the sources pushed into the remote branch called
master. If you want a specific build for development environment, make
it exactly the same but precise the \*/dev branch. (I think you add
already understood =) )

### Build Activiation

In the section "Build Triggers", in Poll SCM, add

    */5 * * * *

This is a cron, We say to Jenkins to check if there's
changes every 5 minutes.

### Building the release APK

Let's add our APK building script


```shell
npm install # Really important (it installs all your project dependencies)
ionic build android --release; # APK release creation
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 -keystore $PATH_TO_KEYSTORE platforms/android/build/outputs/apk/android-release-unsigned.apk alpha -storepass $STORE_PASSWORD # Signing the APK
/opt/android-sdk/build-tools/23.0.2/zipalign -v 4 platforms/android/build/outputs/apk/android-release-unsigned.apk platforms/android/build/outputs/apk/$BUILD_ID.blog.juanwolf.fr.apk  # Creating the relased and signed APK
```

### Pushing the APK to the Google Play

Add a Post build action to your build "Upload Android APK to Google
Play".

In Google Play Account, select the credentials you created with your p12
key. Inside APK files add the:

    platforms/android/build/outputs/apk/$BUILD_ID.blog.juanwolf.fr.apk

In the release track, put production (or beta or alpha, it depends what
you want to deploy)

In Recent Changes, well I'm a bit stuck with this part, I would love to
generate this message from a file (from a changelog for example) but I
don't know how to do that, so if you did I will love that you tell how
the hell you've made that!!!. For the moment, I added a static message
like New content added please check this file to see what's new.

Done !
------

Et voilà ! Our jenkins job will deploy every changes on the master
branch directly on the Google Play. You can also add step into the
building script to run tests or to create reports on the app, etc...
Sur ce, codez bien. Ciao !

