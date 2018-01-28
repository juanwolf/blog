---
title:  "Introduction à Terraform"
date:   2017-11-08
tags: ["terraform", "introduction"]
---

# Introduction à Terraform

## Intro

Avec l'adoption massive du Cloud, différents outils sont nés pour simplifier la vie des dev et des sys. admins.  Ces outils ont créé une nouvelle façon de gérer de l'infrastructure, vous avez surement entendu parler de l'Infrastructure As Code aka IaC. Pour être franc, le "Cloud" est une jolie métaphore pour parler de VMs que vous allez faire tourner sur les serveurs d'une autre entreprise. Le gros plus sont les logiciels qui ont été développé pour et autour ce qui a rendu la création d'infrastructure beaucoup plus accessible. Et maintenant avec de simples appels à certaines API vous pouvez provisionner / créer et détruire une infra en un rien de temps! Ce serait sympa d'avoir quelques scripts pour poppuler vos VPCs sur AWS mais aussi appeler Azure pour provisionner du stockage, etc... À la fin vous risquerez de faire une curl overflow. Et pourquoi pas créer un outil pour vous faciliter la vie avec tous ces appels avec une syntaxe spécifique? Et bien c'est ce qu'est terraform et nous allons voir dans cet article comment l'utiliser.

> ATTENTION: Ce tutoriel utilisera AWS comme cloud provider. Chaque centimes que vous dépenserez est de votre responsabilité, si jamais vous avez utilisé votre "periode d'essaie" sur Amazon, vous allez devoir sortir votre portefeuille :p

## Commençons doucement

Ce que je trouve plutôt cool avec Terraform, c'est sa flexibilité. Vous pouvez créer 100 fichiers, il va s'en occuper commme s'il n'était qu'un. Mais commençons tranquilement.

Premièrement, installons terraform

### Installation

* Mac: `brew install terraform`
* Ubuntu: `sudo apt-get install terraform
* Windows: Le binaire est à disposition sur le site de terraform, et il faudra l'installer dans votre PATH

### Configurer votre account AWS

Vous devez aller dans la console AWS et créer un utilisateur spécifique. Pour cela, allez dans la section IAM et créer un utilisateur genre my_terraform ou quelque chose comme ça

### Premier projet

Commençon par créer un fichier main.tf. Vous allez devoir créer un utilisateur dans IAM afin de récupérer une clé d'accès et une clé secrète.

Commençons from scratch:

```
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

Vous pouvez voir assez simplement quels components nous avons créé, si on compare ça à un script bash rempli de curl, c'est bien plus clair.

Maintenant voyons comment utiliser ce fichier et ce qu'il contient.

#### Plannification

Afin de faire quoi que ce soit regrettable, voyons ce que terraform va créer. Pour cela:

`terraform plan`

Et vous devriez avoir un immense blob décrivant ce que vous allez créer.

Je vous conseil de sauvegarder tous vos "plan" avant de les appliquer. De cette manière, vous êtes sûr que ce que le plan à plannifier va être appliqué.

Pour cela, spécifiez un nom de fichier / chemin où vous voulez sauver votre plan.

`terarform plan -out ./my_aws_plan`

Après avoir vérifié que tout allait bien, on peut appliquer notre plan.

`terraform apply ./my_aws_plans`

Et si vous vous rendez dans la console d'AWS, vous verrez un tout nouveau VPC avec une instance.

Détruisons ce que nous venons de faire, le temps que je vous explique ce que contient le main.tf.

Pour cela: `terraform destroy`.

Et voilà! On vient de détruire tous les composants que nous avions précédemment créé, super simple hein?

### Euh What the fuck ?

#### Terraform plan

Plannifier est la meilleure forme pour tester, superviser et monitorer ce que terraform va exécuter. Je vous invite à "plan" votre code assez souvent, pour avoir un feedback assez rapide sur les erreurs de syntaxes, etc... Attention cependant terraform ne va pas détecter les erreurs logiques de votre code. Par exemple, si vous créez un subnet avec ce CIDR block 172.0.0.1/32 et que vous définissez une instance ec2 avec une IP hors de ce sous-réseau du genre 192.168.1.1, terraform va seulement détecter le problème une fois que vous allez appliquer vos changements car l'API d'AWS va retourner une erreur.

Donc dans votre premier essaie, le plan vous montre que vous n'allez que créer de nouvelles ressources (ce qui est logique vous me direz...). Mais comment terraform va savoir ce qu'il doit détruire à ce qu'il va devoir créer, etc...? Bref comment terraform s'y retrouve ? Si vous regardez au sein de votre projet, vous devriez voir que terraform à créé un nouveau fichier appelé "terraform.tfstate". Ce fichier va être utilisé par terraform, un peu comme une base de données, il va stocker toutes les infos à l'instant T de votre plan dans ce fichier, il est donc SUPER important mais j'y reviendrai un peu plus tard.

#### Terraform apply

Cette action va, quant à elle, appliquer les changements que nous avons fait. Si aucun fichier de "plan" est indiqué, terraform va mettre le "state file" avant d'appliquer ces changements. Je vous conseille vivement de sauvegarder vos plan et de les réutiliser pour les "apply" au cas où vous travailler à deux sur la même infra.
(Si vous comptez travailler à plusieurs sur le même projet, je vous invite à utiliser un "[backend](https://www.terraform.io/docs/backends/index.html)")

### La syntaxe

La syntaxe utilisée dans Terraform est la même que pour tous les outils hashicorp. Ils utilisent le format ".hcl" qu'ils ont créé en se basant sur le format JSON.

On peut retrouver deux types de définitions dans le format hcl:

* `terraform_keyword name`
* `terraform_keyword component_type component_id`

Le premier est plus souvent utilisé pour des fins de configuration tel que variables, providers, etc... alors que le deuxième est plutôt orienté définition de ressources.

**Attention**: Si jamais vous changez le component_id entre deux apply, terraform ne va pas détecter que vous avez seulement renommer la ressource et détruira l'existante pour la remplacer avec une nouvelle.

#### La section "Provider"

Dans terraform vous devez toujours utiliser un provider, cette section va permettre à terraform de se connecter à un cloud en particulier.

Dans notre exemple nous avons tout défini dans le fichier mais vous pouvez aussi utiliser des variables d'environnement. Et l'avantage est que vous pouvez même configurer terraform pour se connecter à plusieurs "cloud provider".

#### Les sections "Resources"

Ce sera les sections que vous utiliserez le plus. Les ressources sont les sections qui vont définir les composants que vous voulez créer, modifier dans votre cloud. Il faut savoir qu'une fois la ressource créée, vous pouvez utiliser certains attributs dans la suite de votre project terraform. Par exemple une fois une instance ec2 est créée vous pouvez récupérer son IP ou autre.

Une fois de plus, je vous invite à aller voir la doc de terraform (il y a plus d'une centaines de ressources avec des variables tout aussi différentes, donc la flemme très cher websurfeu r/se.

#### Interpolation

Si vous avez étudié un peu le code au dessus, vous avez du voir que nous utilisons une variable pour définir le sous réseau de notre instance ec2.
Avec terraform, vous pouvez interpoler des variables en utilisant `${}` afin d'avoir un peu de logique dans votre infra (il y a des loops, des conditions et tout!!!).

### Raffiner notre projet

Avec le temps, votre projet va devenir beaucoup plus gros et confus. Décomposons notre main.tf en plusieurs fichiers afin de nous y retrouver un peu plus.

#### Décomposons le main.tf

Terraform vous permet de créer n'importe quels fichier vous voulez. Lors d'un plan ou apply terraform va essayer de regrouper tous vos fichier en un en évaluant un arbre de dépendance entre tous les fichiers.

Commençons par mettre le provider dans un nouveau fichier :

```
# providers.tf

provider "aws" {
    access_key = "your_access_key"
    secret_key = "your_secret_key"
    region     = "eu-west-1"
}
```

Vous pouvez commiter ce fichier en omettant les mots de passe (bien sûr !) ou même utiliser les profiles configurables avec la cli d'aws.

Essayons :

`terraform apply`

**...**

Ok, vous ne lisez pas l'article en entier. Avant tout "apply" je vous invite à faire un plan avant de détruire quoi que ce soit par erreur. De toute façon même si vous êtes tombé dans la panneau cela n'a pas du changer quoi que ce soit.

Continuons avec la configuration réseau. Nous allons mettre la création du VPC et l'internet gateway.

```
# vpc.tf

resource "aws_vpc" "default" {
    cidr_block = "172.16.0.0/16"
}

resource "aws_internet_gateway" "default"{
    vpc_id = "${aws_vpc.default.id}"
}

```

On va faire de même avec les sous-réseaux, security groups et les instances ec2.

```
# subnets.tf

resource "aws_subnet" "default" {
    vpc_id     = "${aws_vpc.default.id}"
    cidr_block = "172.16.0.0/16" # Just one big subnet covering the whole VPC. Of course do not use that in production.
}

```

```
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

```
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

Après avoir tout refactoré, vous aurez une structure un peu plus convenable. Pour des projets de petite envergure, c'est grandement suffisant. Cependant si votre projet commence à grossir encore plus, vous allez besoin de gérer des dépendances un peu plus importantes et vous allez vouloir moduler votre projet de manière plus optimales.

Pour cela nous avons les modules !

## Modules

With modules you can seperate components of your infrastructure inside _modules_ allowing you to prevent any repetition in your infra definition. Really handy when you want to scale up some components of your current infra or when you want to refactor "all the masters".


### Architecture

To start with modules, you juxg need to create a modules folder at the root of your project, and create a folder with the name of the module you want to create. Inside the last folder you just need to create three files and you will have created your first module! Those files are:

* variables.tf: This file contains all the variables you would like to parametrize your module. For example you will pass some vpc_id or some AMIs.

* main.tf: Like our first main.tf it will contains all the resource definition of your module.

* outputs.tf: Contains all the variables you will need after executing the modules. The most common use is to _output_ the public ips of the future created instances.

### Using it

In our example, the infra is quite simple so this module will be as well. We will just put our simple_instance inside the module. But in a way as we can configure on which subnet, which az it will be.

Let's start creating the module.

```
mkdir -p modules/my_cluster_of_instances
touch modules/my_cluster_of_instances/{main,variables,output}.tf
```

Let's put our definition of the "simple_instance" in the main.tf


```
# main.tf

resource "aws_instance" "simple_instance" {
    ami  = "${var.ami_id}"
    instance_type = "${var.instance_type}"
    subnet_id = "${var.subnet_ids}"
    count = "${cluster_size}"
}
```

If you saw, I changed the variable to be something defined from the variable.tf file.
I added a count attribute in case we want to scale up this module

Now let's configure the variables:

```
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

```
# output.tf

output "private_ips" {
    value = ["${aws_instance.simple_instance.*.private_ip}"]
}
```

The last one use a wildcard as we don't know how many instances will be created in the module. It means "_Ouput_ a list of the instances's private ips"

So now let's create a main.tf at the root of our project calling this module:

```
# main.tf

# Everything we had a bit earlier ... (vpcs, subnets until the instance resource)

module "awesome_instance" {
    module_path = "modules/my_cluster_of_instances"
    ami_id = "${data.aws_ami.ubuntu.id}"
    subnet_id = "${aws_subnet.default.id}"
    instance_type = "t2.micro" # No need of this line as there's a default value
    cluster_size = 2 # Here we override the default value
}

aws_security_group_rule "a_simple_sg_rule" {
    security_group_id = "${}"
    type = "ingress"
    from = 0
    to_port = 0
    protocol = -1
    cidr_block = ["${module.awesome_instance.private_ips}"] # Using here the output of our module
}
```

In a command line:

```
terraform get # Will create reference to our module
terraform plan # Should destroy what we had before
```

Sadly terraform will not understand that this module represent your old instance, and will try to remove your old instance to put the two news.

### Conclusion about project structure

So far, the module structure is the best one I met at the moment. I usually comes with few folders at the root of my projects:

1. modules/
2. env1/
3. env2/

And using envx as a proper seperation between environments. It's really handy when you have a lot of differences between environments.

## Other commands with terraform

To finish this _super_ *long* article, I will briefly speaks about the command we did not see in the article.

### Taint

In terraform you can taint some resources so they will get destroyed at the next apply. Really useful if you want to start from scratch with some components.

From our module example, we would use it like that:

`terraform taint -module=awesome_instance instance.0`

So you need to specify to which module you're tainting the ressource from. And after that resource_name.which_one .

> **Warning**: There's no support for wildcard yet according to [this github issue](https://github.com/hashicorp/terraform/issues/)


### Graph

If you have graphviz installed in your laptop, you can create a graph of the terraform resources you define.

### Import

If you created some resources in the UI, adding their definition in your tf project will not be enough. You need to add it to your terraform state using the import command.

For example if we did create the ec2 instance in the ui. We would add it like that in the tfstate file.

```
 terraform import aws_instance.my_instance the_id_of_the_instance
```

## Conclusion

 There's still much more to say about terraform but I'll stop here as I reached nearly 2500 words...
 We have seen that Terraform is a great tool to manage infrastructure with code. Their .tf format is a really good thing compared to simple json definition. And terraform is quite permissive so you can start with a simple project and finish with hundreds of modules calling each others making the adopotion of the tool quite easy. Anyway I leave you enjoy your `terraform destroy`

 Sur ce, codez bien! Ciao!
