
---
title: Introduction à Terraform
date: 2018-01-05
tags: ["Go", "ops"]
categories: ["Programmation"]

draft: false
author: "Jean-Loup Adde"
---

Avec l'adoption massive du Cloud, différents outils sont nés pour
simplifier la vie des dev et des sys. admins. Ces outils ont créé une
nouvelle façon de gérer de l'infrastructure, vous avez surement entendu
parler de l'Infrastructure As Code aka IaC. Pour être franc, le "Cloud"
est une jolie métaphore pour parler de VMs que vous allez faire tourner
sur les serveurs d'une autre entreprise. Le gros plus sont les logiciels
qui ont été développé pour et autour ce qui a rendu la création
d'infrastructure beaucoup plus accessible. Et maintenant avec de simples
appels à certaines API vous pouvez provisionner / créer et détruire une
infra en un rien de temps! Ce serait sympa d'avoir quelques scripts pour
poppuler vos VPCs sur AWS mais aussi appeler Azure pour provisionner du
stockage, etc... À la fin vous risquerez de faire une curl overflow. Et
pourquoi pas créer un outil pour vous faciliter la vie avec tous ces
appels avec une syntaxe spécifique? Et bien c'est ce qu'est terraform et
nous allons voir dans cet article comment l'utiliser.

![](/post_preview/20180105_131244_terraform_logo.png)

>
>
> ATTENTION: Ce tutoriel utilisera AWS comme cloud provider. Chaque
> centimes que vous dépenserez est de votre responsabilité, si jamais
> vous avez utilisé votre "periode d'essaie" sur Amazon, vous allez
> devoir sortir votre portefeuille :p
>
>



Commençons doucement
--------------------



Ce que je trouve plutôt cool avec Terraform, c'est sa flexibilité. Vous
pouvez créer 100 fichiers, il va s'en occuper commme s'il n'était qu'un.
Mais commençons tranquilement.



Premièrement, installons terraform



### Installation



-   Mac: `brew install terraform`
-   Ubuntu: \`sudo apt-get install terraform
-   Windows: Le binaire est à disposition sur le site de terraform, et
    il faudra l'installer dans votre PATH



### Configurer votre account AWS



Vous devez aller dans la console AWS et créer un utilisateur spécifique.
Pour cela, allez dans la section IAM et créer un utilisateur genre
my\_terraform ou quelque chose comme ça



### Premier projet



Commençon par créer un fichier main.tf. Vous allez devoir créer un
utilisateur dans IAM afin de récupérer une clé d'accès et une clé
secrète.



Commençons from scratch:


```tf
mkdir terraform_lab
cd terraform_lab
git init
git remote add origin git://my_git_repo.git
```



```terraform
provider "aws" {
    access_key = "your_access_key"
    secret_key = "your_secret_key"
    region     = "eu-west-1" # You can change the region for the one your prefer
}

# Nice copy pasta from the doc (https://www.terraform.io/docs/providers/aws/r/instance.html)
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_vpc" "default" {
    cidr_block = "172.16.0.0/16"
}

resource "aws_internet_gateway" "default"{
    vpc_id = "${aws_vpc.default.id}"
}

resource "aws_subnet" "default" {
    vpc_id     = "${aws_vpc.default.id}"
    cidr_block = "172.16.0.0/16" # Just one big subnet covering the whole VPC. Of course do not use that in production.
}

resource "aws_security_group" "open_bar" {
    name = "open_bar"
    description = "Allow all connections inbound and outbound"
    vpc_id = "${aws_vpc.default.id}"
    ingress {
        from_port   = "0"
        to_port     = "0"
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port   = "0"
        to_port     = "0"
        cidr_blocks = ["0.0.0.0/0"]
        protocol    = "-1"
    }
}

resource "aws_instance" "simple_instance" {
    ami  = "${data.aws_ami.ubuntu.id}"
    instance_type = "t2.micro"
    subnet_id = "${aws_subnet.default.id}"
}
```

Vous pouvez voir assez simplement quels components nous avons créé, si
on compare ça à un script bash rempli de curl, c'est bien plus clair.

Maintenant voyons comment utiliser ce fichier et ce qu'il contient.

#### Plannification

Afin de faire quoi que ce soit regrettable, voyons ce que terraform va
créer. Pour cela:

`terraform plan`

Et vous devriez avoir un immense blob décrivant ce que vous allez créer.

Je vous conseil de sauvegarder tous vos "plan" avant de les appliquer.
De cette manière, vous êtes sûr que ce que le plan à plannifier va être
appliqué.

Pour cela, spécifiez un nom de fichier / chemin où vous voulez sauver
votre plan.

`terarform plan -out ./my_aws_plan`

Après avoir vérifié que tout allait bien, on peut appliquer notre plan.

`terraform apply ./my_aws_plans`

Et si vous vous rendez dans la console d'AWS, vous verrez un tout
nouveau VPC avec une instance.

Détruisons ce que nous venons de faire, le temps que je vous explique ce
que contient le main.tf.

Pour cela: `terraform destroy`.

Et voilà! On vient de détruire tous les composants que nous avions
précédemment créé, super simple hein?

### Euh What the fuck ?

#### Terraform plan

Plannifier est la meilleure forme pour tester, superviser et monitorer
ce que terraform va exécuter. Je vous invite à "plan" votre code assez
souvent, pour avoir un feedback assez rapide sur les erreurs de
syntaxes, etc... Attention cependant terraform ne va pas détecter les
erreurs logiques de votre code. Par exemple, si vous créez un subnet
avec ce CIDR block 172.0.0.1/32 et que vous définissez une instance ec2
avec une IP hors de ce sous-réseau du genre 192.168.1.1, terraform va
seulement détecter le problème une fois que vous allez appliquer vos
changements car l'API d'AWS va retourner une erreur.

Donc dans votre premier essaie, le plan vous montre que vous n'allez que
créer de nouvelles ressources (ce qui est logique vous me direz...).
Mais comment terraform va savoir ce qu'il doit détruire à ce qu'il va
devoir créer, etc...? Bref comment terraform s'y retrouve ? Si vous
regardez au sein de votre projet, vous devriez voir que terraform à créé
un nouveau fichier appelé "terraform.tfstate". Ce fichier va être
utilisé par terraform, un peu comme une base de données, il va stocker
toutes les infos à l'instant T de votre plan dans ce fichier, il est
donc SUPER important mais j'y reviendrai un peu plus tard.

#### Terraform apply

Cette action va, quant à elle, appliquer les changements que nous avons
fait. Si aucun fichier de "plan" est indiqué, terraform va mettre le
"state file" avant d'appliquer ces changements. Je vous conseille
vivement de sauvegarder vos plan et de les réutiliser pour les "apply"
au cas où vous travailler à deux sur la même infra. (Si vous comptez
travailler à plusieurs sur le même projet, je vous invite à utiliser un
"[backend](https://www.terraform.io/docs/backends/index.html)")

### La syntaxe

La syntaxe utilisée dans Terraform est la même que pour tous les outils
hashicorp. Ils utilisent le format ".hcl" qu'ils ont créé en se basant
sur le format JSON.

On peut retrouver deux types de définitions dans le format hcl:

-   `terraform_keyword name`
-   `terraform_keyword component_type component_id`

Le premier est plus souvent utilisé pour des fins de configuration tel
que variables, providers, etc... alors que le deuxième est plutôt
orienté définition de ressources.

**Attention**: Si jamais vous changez le component\_id entre deux apply,
terraform ne va pas détecter que vous avez seulement renommer la
ressource et détruira l'existante pour la remplacer avec une nouvelle.

#### La section "Provider"

Dans terraform vous devez toujours utiliser un provider, cette section
va permettre à terraform de se connecter à un cloud en particulier.

Dans notre exemple nous avons tout défini dans le fichier mais vous
pouvez aussi utiliser des variables d'environnement. Et l'avantage est
que vous pouvez même configurer terraform pour se connecter à plusieurs
"cloud provider".

#### Les sections "Resources"

Ce sera les sections que vous utiliserez le plus. Les ressources sont
les sections qui vont définir les composants que vous voulez créer,
modifier dans votre cloud. Il faut savoir qu'une fois la ressource
créée, vous pouvez utiliser certains attributs dans la suite de votre
project terraform. Par exemple une fois une instance ec2 est créée vous
pouvez récupérer son IP ou autre.

Une fois de plus, je vous invite à aller voir la doc de terraform (il y
a plus d'une centaines de ressources avec des variables tout aussi
différentes, donc la flemme très cher websurfeu r/se.

#### Interpolation

Si vous avez étudié un peu le code au dessus, vous avez du voir que nous
utilisons une variable pour définir le sous réseau de notre instance
ec2. Avec terraform, vous pouvez interpoler des variables en utilisant
`${}` afin d'avoir un peu de logique dans votre infra (il y a des loops,
des conditions et tout!!!).

### Raffiner notre projet

Avec le temps, votre projet va devenir beaucoup plus gros et confus.
Décomposons notre main.tf en plusieurs fichiers afin de nous y retrouver
un peu plus.

#### Décomposons le main.tf

Terraform vous permet de créer n'importe quels fichier vous voulez. Lors
d'un plan ou apply terraform va essayer de regrouper tous vos fichier en
un en évaluant un arbre de dépendance entre tous les fichiers.



Commençons par mettre le provider dans un nouveau fichier :


```tf
# providers.tf

provider "aws" {
    access_key = "your_access_key"
    secret_key = "your_secret_key"
    region     = "eu-west-1"
}
```



Vous pouvez commiter ce fichier en omettant les mots de passe (bien sûr
!) ou même utiliser les profiles configurables avec la cli d'aws.



Essayons :



`terraform apply`



**...**



Ok, vous ne lisez pas l'article en entier. Avant tout "apply" je vous
invite à faire un plan avant de détruire quoi que ce soit par erreur. De
toute façon même si vous êtes tombé dans la panneau cela n'a pas du
changer quoi que ce soit.



Continuons avec la configuration réseau. Nous allons mettre la création
du VPC et l'internet gateway.


```
# vpc.tf

resource "aws_vpc" "default" {
    cidr_block = "172.16.0.0/16"
}

resource "aws_internet_gateway" "default"{
    vpc_id = "${aws_vpc.default.id}"
}
```


On va faire de même avec les sous-réseaux, security groups et les
instances ec2.

```tf
# subnets.tf

resource "aws_subnet" "default" {
    vpc_id     = "${aws_vpc.default.id}"
    cidr_block = "172.16.0.0/16" # Just one big subnet covering the whole VPC. Of course do not use that in production.
}
```

```tf
# security_groups.tf

resource "aws_security_group" "open_bar" {
    name = "open_bar"
    description = "Allow all connections inbound and outbound"
    vpc_id = "${aws_vpc.default.id}"
    ingress {
        from_port   = "0"
        to_port     = "0"
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port   = "0"
        to_port     = "0"
        cidr_blocks = ["0.0.0.0/0"]
        protocol    = "-1"
    }
}
```

```tf
# instances.tf

# Oui je mets les ami avec les instances, pas besoin d'un fichier spécifique pour une data.

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "simple_instance" {
    ami  = "${data.aws_ami.ubuntu.id}"
    instance_type = "t2.micro"
    subnet_id = "${aws_subnet.default.id}"
}
```

Après avoir tout refactoré, vous aurez une structure un peu plus
convenable. Pour des projets de petite envergure, c'est grandement
suffisant. Cependant si votre projet commence à grossir encore plus,
vous allez besoin de gérer des dépendances un peu plus importantes et
vous allez vouloir moduler votre projet de manière plus optimales.

Pour cela nous avons les modules !

Modules
-------

Avec les modules vous pouvez séparer *logiquement* votre infrastructure
dans des modules ce qui va vous permettre d'éviter de vous répéter et
vous pouvez réutiliser vos modules dans différents projets, etc... Enfin
bref c'est super utile si vous voulez augmenter le nombre de composant
ou autre, seulement besoin d'incrémenter une variable et le tour est
joué.

### Architecture

Pour commencer avec les modules, vous devez créer une dossier "modules",
et créer un dossier spécifique au module en question. Dans le dernier
dossier, vous avez seulement besoin de créer trois fichiers :

-   variables.tf: Ce fichier contient toutes les variables paramétrant
    ce module. Par exemple, vous définierez une variable vpc\_id, ou des
    ids d'AMIs.

-   main.tf: Comme notre bon vieux main.tf, il va contenir toutes les
    ressources définissant notre module.

-   outputs.tf: Contient toutes les variables dont vous aurez besoin
    après la création de vos ressources. Par exemple, il est très
    fréquent d'output les IPs publiques après la création d'instances
    ec2.



Au cas où ce n'est pas clair:


```bash
mkdir -p modules/my_module
touch modules/my_module/{variables,main,outputs}.tf
```

### L'utiliser

Dans notre exemple, l'infrastructure est plutôt simple donc le module va
l'être tout aussi. On va juste créer notre "simple\_instance" au sein de
ce module. On va faire en sorte de pouvoir configurer dans quel subnet,
on peut installer cette instance.



Commençons par créer le module :


```bash
mkdir -p modules/my_cluster_of_instances
touch modules/my_cluster_of_instances/{main,variables,output}.tf
```

Mettons notre définition de la "simple\_instance" dans notre main.tf

```tf
# main.tf

resource "aws_instance" "simple_instance" {
    ami  = "${var.ami_id}"
    instance_type = "${var.instance_type}"
    subnet_id = "${var.subnet_ids}"
    count = "${cluster_size}"
}
```

Comme vous pouvez le voir plus haut, nous avons fait en sorte que toute
l'instance soit paramètrable. Nous allons devoir ajouter les variables
utilisées dans ce main.tf dans notre variables.tf. J'ai ajouté un
attribut "count" au cas où nous voulions augmenter le nombre d'instance
créée.

Configurons nos variables:

```tf
# variables.tf

variable "subnet_ids" {
    description = "The list of subnet id"
}

variable "cluster_size" {
    description "The number of instance you want"
    default = 1
}

variable "instance_type" {
    description = "The type of instance you want"
    default = "t2.micro"
}

variable "ami_id" {
    description = "The AMI to use on these instances"
}
```


Et maintenant, je veux connaître les IPs privées attribuées à mes
instances après l'exécution du module. Pour cela:


```tf
# output.tf

output "private_ips" {
    value = ["${aws_instance.simple_instance.*.private_ip}"]
}
```

Comme vous pouvez le voir j'ai utilisé une astérisque afin de référencer
**toutes** les instances créées.

Maintenant nous pouvons créer notre main.tf qui va utiliser ce module :

```tf
# main.tf

module "awesome_instance" {
    module_path = "modules/my_cluster_of_instances"
    ami_id = "${data.aws_ami.ubuntu.id}"
    subnet_id = "${aws_subnet.default.id}"
    instance_type = "t2.micro" # No need of this line as there's a default value
    cluster_size = 2 # Here we override the default value
}

# Et ici nous pouvons utiliser l'output tel une variable
# genre ${module.awesome_instance.private_ips}

output "private_ips_of_my_module" {
    value = ["${module.awesome_instance.private_ips}"]
}
```


Testons :

```bash
terraform get # Will create reference to our module
terraform plan # Should destroy what we had before
```


Malheureusement, terraform ne va comprendre la référence à notre
ancienne instance à juste bouger au sein d'un module. Donc encore une
fois, terraform va vouloir supprimer l'ancienne pour la remplacer avec
une nouvelle.

### Conclusion à propos de la structure d'un projet

Jusque là, la meilleure structure que j'ai rencontré est celle des
modules. Certes elle demande un peu plus d'expérience avec terraform.
Dans la plupart de mes projets, je différencie les environnements dans
deux dossiers et utilise un pour les modules à la racine, comme ceci:

1.  modules/
2.  env1/
3.  env2/

PLutôt sympa si vous avez pas mal de différences entre dev et prod par
exemple. Le mieux serait d'avoir exactement la même définition, genre un
main.tf commun mais seul les variables changent (encore plus dur niveau
implémentation)

Autres commandes
----------------

Pour finir cet article *super* *long*, je vais rapidement vous énoncer
quelques commandes que nous n'avons pas encore vu dans l'article.

### Taint

Dans terraform vous pouvez "taint" certaines ressources de votre projet,
cela va indiquer à terraform de la supprimer au prochain "apply". Plutôt
utile si vous avez des composants qui ont été marqué comme "not healthy"
par AWS.

En utilisant notre module, nous l'utiliserions comme ceci:

`terraform taint -module=my_cluster_of_instances simple_instance.0`

>
>
> **Attention**: Il n'y a pas encore de support pour les astérisques
> [this github
> issue](https://github.com/hashicorp/terraform/issues/)
>
>

### Graph

Si vous connaissez graphviz et que vous l'avez installé sur votre
machine, vous pouvez créer une graphe réprésentant votre infra.

### Import

Si jamais vous avez créé un peu d'infrastructure en utilisant
l'interface web, ajouter les définitions de ces éléments dans votre
projet ne vas pas être suffisant. Terraform ne va comprendre que vous
référez à cet élément. C'est pour cela que vous devez utiliser la
fonction "import". Elle va ajouter la définition de votre instance au
sein du tfstate.

Par exemple, si nous avions créé l'instance ec2 avec la console, nous
l'importerions comme ceci:

     terraform import aws_instance.simple_instance the_id_of_the_instance

Conclusion
----------

Il y a encore pleins de choses à dire sur terraform et j'ai déjà atteint
un nombre de mots conséquent... Je ne suis pas un expert non plus, donc
je vous invite à regarder sur plusieurs blog, retour d'expérience afin
de connaître la meilleure façon d'architecturer un projet terraform :)
Nous avons vu les bases de l'outil ainsi que sa puissance et sa facilité
d'utilisation. Le format utilisé est super intuitif comparé à du JSON.
Et terraform est assez maléable pour être accessible à tout niveau
d'utilisation. Je vous laisse terminer cette article en exécutant un
petit



    terraform destroy



Sur ce, codez bien! Ciao!




