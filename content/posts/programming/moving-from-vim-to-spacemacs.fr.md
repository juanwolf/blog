+++
title = "Migrer de Vim à Spacemacs"
author = ["Jean-Loup Adde"]
lastmod = 2020-04-19T13:50:45+01:00
tags = ["spacemacs", "vim"]
draft = false
+++

Cela va faire 5 ans que je suis dans l'industrie et j'ai joué avec pas mal d'éditeurs de code. Je dois vous l'avouer, c'est devenu un _meme_ autour de moi vu que j'adore expérimenter tous les éditeurs qui existent. VS Code, Vim, Emacs, IntelliJ, PyCharm, etc... C'est chronophage. Comme le titre vous a sûrement mis la puce à l'oreille, cet article sera à propos de ma transition de Vim vers Spacemacs. J'ai passé 4 ans sur Vim avant de bouger sur Spacemacs. J'ai eu une année d'égarement où j'ai utilisé PyCharm malgrès les louanges quotidienne de la bête technique qui travaillait à côté de moi.
Je suis sûr que vous avez déjà lu une multitude d'articles ayant comme sujet "_Pourquoi j'ai bougé de X à Y et voici pourquoi vous devriez faire de même._". Je ne suis pas d'accord avec la pratique et clairement, si vous n'avez pas les mêmes besoins que l'auteur, changer va-t-il en valoir la peine? Malheureusement, il n'y a pas de meilleurs éditeurs. **FIN**. Je vais vous donner quand même les raisons de pourquoi j'ai switché. Peut-être que vous vous y retrouverez.

{{< figure src="/post_content/moving-from-vim-to-emacs/spacemacs_logo.png" >}}


## Un peu de contexte {#un-peu-de-contexte}

Je suis passé de Développeur Web à Ingénieur Systèmes / SRE / Ingénieur plateforme (bref sys. admin avec un peu de dev). Pendant mon année de Développeur web, j'ai utilisé PyCharm comme mon IDE principal. Pour être franc, j'ai eu une super expérience. J'ai adoré le debuggueur intégré. Comme je ne travaillais que sur un seul projet django, un IDE était parfait. J'ai appris énormément sur Django grâce à l'IDE en creusant comment Django fonctionnait.
Cependant quand j'ai migré vers un travail pluridisciplinaire, c'est là que j'ai senti que PyCharm n'était pas l'outil pour mon taf de tous les jours. Clairement, c'est une usine à gaz et pour écrire de simples scripts ou changer du YAML, pas besoin d'un IDE. J'ai donc décidé de bouger sur Vim vu que je passais ma vie dans le terminal. La transition de PyCharm / IDEs de JetBrains vers Vim n'a pas été sans effort, je vous rassure tout de suite. De plus, configurer l'éditeur pour avoir un confort similaire aux IDEs m'a pris énormément de temps.

Syntax highlighting, Syntax checkers, Linters et auto complétion quand c'est possible.


## Pourquoi ? {#pourquoi}

Si vous fuyez quelque chose, vous devez avoir une raison. Mon problème avec Vim était que ma config était instable. Genre 15 outils différents pour faire tourner l'éditeur donc 15 outils à supporter et maintenir, un peu galère. De plus, pour un "simple éditeur", Vim prennait quasiment 40 secondes pour démarrer. Donc je cherchais une alternative. À ce moment là, VS Code faisait pas mal de bruit, "_la nouvelle ère de Microsoft_", j'ai donc essayé l'éditeur. J'ai eu du mal. Avec Vim, j'avais pris l'habitude de naviguer un peu partout avec des raccourcis clavier. On ressent très vite que ce n'est pas le cas avec VS Code. Enfin oui, on peut. Mais avec l'intégration Vim, on est limité à naviguer seulement dans le code avec les raccourcis vim, pas l'éditeur entier. De plus, je passais mon temps dans le terminal, j'avais mon setup avec tmux + vim, je voulais juste un expérience similaire. Donc un éditeur qui peut tourner dans le terminal. Je suis resté sur Vim jusqu'à ce qu'un collègue, un peu **insistant**, me montre Spacemacs. J'ai essayé et après avoir retourné sur Vim plusieurs fois, j'ai enfin passé à Spacemacs pour de bon. Clairement la migration avait des avantages et inconvénients:

**Pour:**

1.  Spacemacs est une configuration d'Emacs configuré par la communauté
2.  Spacemacs est très configurable
3.  Spacemacs peut être utilisé avec les raccourcis Vim :heart:
4.  L'éditeur utilise un langage de programmation (même si c'est du Lisp... _sigh_)
5.  Les raccourcis de Spacemacs sont mnémonic **et découvrable**

**Contre:**

1.  Apprendre comment Emacs fonctionne
2.  Apprendre Lisp :sweat_smile:
3.  On peut être vite dépassé par le nombre d'intégration et d'outils que Spacemacs contient


## Comment ? {#comment}

Si vous voulez essayer, c'est assez simple. Installez Emacs sur votre OS puis clonez [Spacemacs](https://github.com/syl20bnr/spacemacs) dans \`~/.emacs.d\`. Je vous invite à changer la branche utilisée dans .emacs.d pour "develop" comme cette branche est plus maintenue que master et contient bien plus de fonctionnalités.
Une fois emacs lancé, vous pouvez découvrir les raccourcis en pressant simplement "espace". Vous aurez une présentation de tous les raccourcis disponibles. p pour projet, w pour window, b pour buffer, etc... Les raccourcis s'apprennent très rapidement car la plupart ne sont que les initiales de l'action que vous voulez faire. Vous voulez voir l'arborescence de ficher de votre projet? "project -> tree": "Espace"+ "p"+ "t". Facile.


## Tout est pour le mieux donc ? {#tout-est-pour-le-mieux-donc}

Hum... Je passe toujours autant de temps sur ma configuration de Spacemacs. L'apprentissage de l'éditeur était un peu long mais j'ai gagné énormément en productivité. J'ai aussi appris Lisp et lu les "packages" de Spacemacs. Contribué à certains. Ça a été un sacré voyage mais je ne le regrette pas une seconde. De plus, en utilisant emacs en mode serveur, votre éditeur démarre en 2 secondes, c'est impressionant.
Donc oui, c'est clairement mieux et je ne retournerai pas sur Vim, c'est sûr. De plus, j'ai découvert org-mode et je l'utilise avec ce blog et au taf. Je ne pense pas qu'un seul éditeur n'aura l'intégration qu'Emacs a pour org-mode. J'écrierai un article sur org-mode comme je commence à pas mal m'en servir.

Il y a toujours des éléments instables dans ma configuration comme l'intégration avec le Language Server Protocol. Je peux vivre de sans mais l'auto-complétion reste sympa à avoir. Ça me force à lire la documentation ce qui n'est probablement pas plus mal. Mais bon, n'oublions pas que la configuration de Spacemacs est toujours en "development" et que ce genre d'instabilité est attendue.

Donc devez-vous bouger sur Spacemacs ? Probablement pas, je suis quasiment sûr que votre éditeur peut être configuré pour vos besoins. Mon conseil à 2 francs: Apprenez les raccourcis clavier plutôt que de naviguer à la souris. Sinon, rejoingnez moi dans un voyage de configuration et de weekends perdus à maximiser votre productivité, genre perdre un dimanche aprem pour gagner 2 minutes la semaine. Non ? Personne ? Je comprends pas... :joy:
