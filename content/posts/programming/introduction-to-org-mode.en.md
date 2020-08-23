+++
title = "Getting started with org-mode"
author = ["Jean-Loup Adde"]
lastmod = 2020-08-23T19:49:14+01:00
tags = ["org", "intro", "emacs"]
draft = false
+++

Markdown is clearly programmers preferred (or at least most used) text format. By moving to Emacs, I got a bit curious on the documentation format used by the majority of the packages. This format is _org_. I read a bit about it and was blown away by what people were able to do with it.

{{< figure src="/post_content/introduction-to-org-mode/org-mode-logo.png" alt="org-mode's logo" caption="Figure 1: org-mode's logo" >}}


## Why org-mode? {#why-org-mode}

So you might wonder, why bother creating a new markup / text format. The thing is Org mode has been around since the 2000. It has been created to follow some productivity workflow called ["Get things done" by David Allen](https://gettingthingsdone.com/).
I don't think it would make much sense for someone to jump on org-mode without using Emacs. A bit of technology locking here unfortunately. Its main strength comes with the tooling around the language itself. Think about Markdown but with strong integration in your editor, generating calendars, todo lists, remainders and others.

So if you're more the kind of person that likes everything in one place, that's the way to go. If not, clearly not. Emacs will become a beast and you'll never get out of it. Your choice.


## How? {#how}

There's plenty of tutorial and guides to help you out. I don't think it's worth to repeat what's been done so here's some good links that helped me out:

-   <https://karl-voit.at/2020/01/20/start-using-orgmode/>

My main advice to get started is to just create a all.org in a notes git repo and just start writing stuff you'll need for later. How-to-s, Documentation, Meeting notes etc... Once you'll get the syntax, you'll naturally check what you can do with some org plugins and discover some powerful stuff.

Org-mode itself is just a markup language so it should not be too complicated to learn. However, its ecosystem is so massive, it can be tough to know where to start.


## The ecosystem {#the-ecosystem}

We said org-mode strength is the tooling made around its markup language. For instance, in org-mode and emacs, I can write some snippets and make them executable easily thanks to org-babel. That allows you to have some executable documentation (a bit like a juniper playbook). Which can be quite powerful if everyone is using emacs. No need of Makefile if you can run each steps directly from your README.org!

I love to track how I spend my time at work so I tend to use the pomodoro technic to catch that. Guess what, there's an org extension for that, org-pomodoro. I put all my stuff to do in a all.org, run org-pomodoro, choose which task I am working on and bim, I have a timer set in Emacs and will get an alert once the pomodoro is finished.
As well, it logs the time spent on each task so you can have aggregated data on the time spent for a specific item, pretty neat.

Now the main question is "That looks really promising but at work we use X and Y and Z to log work so I can't introduce a new tool to the team!". I am in the same position! However, you can collect data and convert them in org-mode with a lot of "external providers". Using Jira? org-jira is there to grab the tasks you have to do. Using Trello? org-trello is there. There's a lot more but you can grab data mostly anywhere and import them in your org setup.


## The Cons {#the-cons}

Obviously, every single technology comes with its caveats. As said before, org-mode is locked to Emacs unfortunately, so people having a org-mode workflow are quite locked to Emacs. And vice-versa, you want to play with it? Definitely better in the editor. There's some effort made to port the incredible Emacs ecosystem to other editors but it'll never be the real thing.

It's a bit problematic for the adoption of the technology. In my case, I can not use org-mode at work (or at least share my org stuff). No one is using emacs so no one will be able to parse those files / render them correctly and get any benefit from using that format.

Another point is that, you are adding layers / packages to your Emacs setup which means more configuration and testing which can be quite time consuming.

As well, the integrations are mainly single people side-projects and you can end up with outdated projects or not supporting your use case. If you're brave enough, it's your time to shine and implement what you want from those packages. Good luck with e-lisp though 😂.


## Conclusion {#conclusion}

Org-mode is a beautiful piece of software. Its markup is similar to markdown but includes way more functionalities when used in Emacs. The ecosystem of that language turns a simple file into a playbook, a project planner, an habit tracker, encrypted data and anything else you would think of! However, the fact that the ecosystem is tightly coupled to Emacs makes the adoption quite difficult and low even though its incredible capabilities. If you're interested about org-mode, don't hesitate and send me a tweet, if I receive a few, I might even write how I use org-mode.
