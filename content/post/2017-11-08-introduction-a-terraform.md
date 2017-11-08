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

### First Project

Let's create a main.tf. If you're eligible for the AWS free tier feel free to use that, if not, you'll need to take some cash out (sorry). You need an IAM user as well to get an access key and a secret key.

We will start from scratch so:

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

You see straight what components you're creating with terraform, compared to a bash script full of curls, it's more clear right?

Now let's see how to use this file and dismistied it.

#### Planning

Before doing anything regrettable, let's see what terraform will do. For that:

`terraform plan`

And you should have an output showing you every resource we gonna create.

I advise you to save every plan you do before applying it. With this way to proceed, you are sure that what the plan planned will be applied and nothing more/less.

For that, just specify the filename/path where you want to save your plan.

`terarform plan -out ./my_aws_plan`

Then after being happy with what the plan will do let's apply it.

`terraform apply ./my_aws_plans`

And if you go in the console you should see that you have a brand new vpc, a subnet and an instance running.

Let's destroy all of these the time I explain you what we've just wrote in this main.tf file.

For that: `terraform destroy`.

And tada, you come back to your original state, super easy isn't it?

### What just happened?

#### Terraform plan

Planning in terraform is the best way to test, supervise and monitor what terraform will do. Everytime you want to test your terraform project -> plan, you'll see syntax /logical errors. PLanning does not prevent on provider issues though. For example you create a subnet with a range like 172.0.0.1/32 (this is just an example) and you define a ec2 instance inside the subnet with a private ip like 192.168.1.1, well terraform will only get the error once you apply your change as AWS will return an error when terraform will do the API call.

So in our first try, the terraform plan created only stuff, which is normal as we started from scratch, but how do terraform will know how to modify some infra from the previous run? If you have a look in your project folder, you have a terraform.tfstate file. This is what terraform use to know the state of the components you defined in your main.tf in the cloud.

When you do a plan if this terraform.tfstate file exists, terraform will pull the actual infrastructure configuration, update your tfstate file and compare what's in the tfstate and what you defined in the main.tf and will prompt you what will change once you run the apply.

#### Terraform apply

This action will actually do the API calls to your cloud provider, if no plan_file is specified terraform will update the state file **before** the apply, and will update after. That's why you better using a plan all the time in case your architecture gets modified between your plan and your apply.

### The terraform syntax

The terraform syntax is a common syntax if you used other hashicorps tools before. They use the .hcl format that they created which is JSON superpowered with comments and many other stuff (Ok, I can't find anything else right now, you got a point).

To configure our terraform with different providers, variables, or even configure the backends terraform will use to store your state file.

You can find as well those but are to configure terraform and have this syntax `<terraform_keyword> <name> { ... }` :

* provider : Will provide authentication to the cloud provider
* variable : Create a variable for the scope of your terraform apply
* configuration : Used to configure terraform, useful for backends configuration (we'll see that later)

Every things defined with terraform in your cloud provider will have the same structure `<terraform_keyword> <component_type> <component_id> { ... }`

The terraform keyword can vary between:

* resource : Create a new component in the cloud using the provider configured
* data : Will retrieve information about an existing component

The components type will all be different regarding which cloud provider you're using. The terraform documentation is your friend in this case (I'll not enumerate all of it as there's 100s)

And to finish the component_id, can be everything you like, the time you understand what it is. It used by terraform to identify the different components created.
For example, if we would have defined two different kind of aws_instance, it would have been great to use 2 different component_id such as "front_end" or "back_end" etc...

**WARNING**: If for any reason you're changing this id, terraform will consider it as a new resource and will remove the exisiting one instead of modifying it.

#### Provider section

in terraform you'll always need a provider section, you can consider this section as the authentication part of your terraform config.
Here we did setup all the configuration, but you can set it up with env variables on the machine you're running terraform.

You can of course use multiple providers in case you write your terraform files for azure, aws, vmware, etc...

#### Resources sections

This is where you define all the components you wants to **create** in the cloud. They are provider specific and once they are created you can access to different variables that this ressource will provide you. For example, creating the ec2 instance, we reuse the id of the subnets we created earlier.

Once again, everything is detailed for every resource in the terraform documentation.

#### Interpolation

If you read what I wrote just above in the resources section, you understood that the barbaric syntax used for defining the subnets of our aws_instance is a variable value.
In terraform you can use strings with ${} to give some logic to the attributes you define. You can reuse ids of components you created (as we did), you can even use loops, conditions, maps, and even some mathematics function.

### Raffined our project

With the time, your project might get bigger, by that I mean, your main.tf file might contains thousands of line, which make it tough to maintain. Let's start with an intuitive approach: cut our main.tf file into pieces.

#### Cutting the main.tf

Terraform allows you to create whatever files you want, during a plan or an apply it will try to concatenate all the *.tf files in the current folder in one file. It will solve the dependencies between the different file and should be able to recreate the main.tf file as it was before we cut it in pieces. We just need to be **logical** on how we will do it.

Let's put the provider in a different file first. Let's create a `providers.tf` file in the current folder containing just our aws logging creds.

```
# providers.tf

provider "aws" {
    access_key = "your_access_key"
    secret_key = "your_secret_key"
    region     = "eu-west-1"
}
```

You can even commit this file with empty credentials, and just make sure you'll never push it again.

Let's test it!

`terraform apply`

**...**

Ok you're not reading the whole article. I said to test we use **plan not apply**. Well anyway if you fell in the trap, this should have not change anything.

Let's continue with our network configuration. Let's put the VPC, gateway in it.

```
# vpc.tf

resource "aws_vpc" "default" {
    cidr_block = "172.16.0.0/16"
}

resource "aws_internet_gateway" "default"{
    vpc_id = "${aws_vpc.default.id}"
}

```

Let's do the same with subnets, security_groups and instances

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

# Yeah I put the ami with instances. No need elsewhere.
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

At the end of the refactoring you'll get a nicer structure. This structure will last for few times until you infrastructure starts to get bigger. With time you'll see some limit to it like in case of complex dependencies or even about SOC (Separation of concerns)

So here comes modules.

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