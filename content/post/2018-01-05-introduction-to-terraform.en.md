
---
title: Introduction to Terraform
date: 2018-01-05
tags: ["Go", "operations"]
categories: ["Programming"]

draft: false
author: "Jean-Loup Adde"
---

With the massive adoption of the "Cloud", different tools are borned
to simplificate dev / sys. admins lifes. These tools created a new way
of dealing with infra, the Infrastructure as Code aka IaC. To be honest,
the cloud is a nice word but let's face it, you just run vms on someone
else servers. The big plus to use those is that with a simple API call
you can provide, provision and destroy infrastructure at your glance. It
would be great to write some scripts to auto populate your VPCs on AWS
but as well calling azure to provision some storage etc... Well at the
end you might get "curl sick". Why not creating a tool that will deal
with all the API calls for you instead? Well that's what terraform is
and let's see how we can use this awesome tool.

![](/post_preview/20180105_131244_terraform_logo.png)


>
>
> Disclaimer: This tutorial will be using AWS as cloud provider. Any
> penny you'll spend on the platform is in your entire responsability
> and you're aware reading this article that you'll have to spend some
> if you're not eligible for the AWS free tier
>
>

Starting baby steps
-------------------

What I find amazing in terraform is the flexibility that the tools give
you. You can create 100 files, it will concatenate all of it,
understanding which part is dependant to the other and so on, but before
to have a funky project architecture let's start simple.



First of all, let's install terraform



### Installation



-   Mac: `brew install terraform`
-   Ubuntu: `sudo apt-get install terraform`
-   Windows: Download the binary on the terraform download page, and add
    it to your PATH.



### Setup your AWS account



You need to go in the aws console and setup a specific user to use
terraform. Go into the `IAM` panel and create a user. `my_terraform`



### First Project



Let's create a main.tf. If you're eligible for the AWS free tier feel
free to use that, if not, you'll need to take some cash out (sorry). You
need an IAM user as well to get an access key and a secret key.


We will start from scratch so:

```bash
mkdir terraform_lab
cd terraform_lab
git init
git remote add origin git://my_git_repo.git
```


```tf
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
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

resource "aws_instance" "simple_instance" {
    ami  = "${data.aws_ami.ubuntu.id}"
    instance_type = "t2.micro"
    subnet_id = "${aws_subnet.default.id}"
}
```

You see straight what components you're creating with terraform,
compared to a bash script full of curls, it's more clear right?

Now let's see how to use this file and dismistied it.

#### Planning

Before doing anything regrettable, let's see what terraform will do. For
that:

`terraform plan`

And you should have an output showing you every resource we gonna
create.

I advise you to save every plan you do before applying it. With this way
to proceed, you are sure that what the plan planned will be applied and
nothing more/less.

For that, just specify the filename/path where you want to save your
plan.

`terarform plan -out ./my_aws_plan`

Then after being happy with what the plan will do let's apply it.

`terraform apply ./my_aws_plans`

And if you go in the console you should see that you have a brand new
vpc, a subnet and an instance running.

Let's destroy all of these the time I explain you what we've just wrote
in this main.tf file.

For that: `terraform destroy`.

And tada, you come back to your original state, super easy isn't it?

### What just happened?

#### Terraform plan

Planning in terraform is the best way to test, supervise and monitor
what terraform will do. Everytime you want to test your terraform
project -\> plan, you'll see syntax /logical errors. PLanning does not
prevent on provider issues though. For example you create a subnet with
a range like 172.0.0.1/32 (this is just an example) and you define a ec2
instance inside the subnet with a private ip like 192.168.1.1, well
terraform will only get the error once you apply your change as AWS will
return an error when terraform will do the API call.

So in our first try, the terraform plan created only stuff, which is
normal as we started from scratch, but how do terraform will know how to
modify some infra from the previous run? If you have a look in your
project folder, you have a terraform.tfstate file. This is what
terraform use to know the state of the components you defined in your
main.tf in the cloud.

When you do a plan if this terraform.tfstate file exists, terraform will
pull the actual infrastructure configuration, update your tfstate file
and compare what's in the tfstate and what you defined in the main.tf
and will prompt you what will change once you run the apply.

#### Terraform apply

This action will actually do the API calls to your cloud provider, if no
plan\_file is specified terraform will update the state file **before**
the apply, and will update after. That's why you better using a plan all
the time in case your architecture gets modified between your plan and
your apply.

### The terraform syntax

The terraform syntax is a common syntax if you used other hashicorps
tools before. They use the .hcl format that they created which is JSON
superpowered with comments and many other stuff (Ok, I can't find
anything else right now, you got a point).

In Terraform, you have mostly two kind of definition:

-   `terraform_keyword name`
-   `terraform_keyword component_type component_id`

For the first syntax, the `<terraform_keyword>` can be :



-   provider : Will provide authentication to the cloud provider
-   variable : Create a variable for the scope of your terraform apply
-   configuration : Used to configure terraform, useful for backends
    configuration (we'll see that later)



Every things defined with terraform in your cloud provider will have the
same structure
`<terraform_keyword> <component_type> <component_id> { ... }`



The terraform keyword can vary between:



-   resource : Create a new component in the cloud using the provider
    configured
-   data : Will retrieve information about an existing component



The components type will all be different regarding which cloud provider
you're using. The terraform documentation is your friend in this case
(I'll not enumerate all of it as there's 100s)



And to finish the component\_id, can be everything you like, the time
you understand what it is. It used by terraform to identify the
different components created. For example, if we would have defined two
different kind of aws\_instance, it would have been great to use 2
different component\_id such as "front\_end" or "back\_end" etc...

**WARNING**: If for any reason you're changing this id, terraform will
consider it as a new resource and will remove the exisiting one instead
of modifying it.

#### Provider section

in terraform you'll always need a provider section, you can consider
this section as the authentication part of your terraform config. Here
we did setup all the configuration, but you can set it up with env
variables on the machine you're running terraform.

You can of course use multiple providers in case you write your
terraform files for azure, aws, vmware, etc...

#### Resources sections

This is where you define all the components you wants to **create** in
the cloud. They are provider specific and once they are created you can
access to different variables that this ressource will provide you. For
example, creating the ec2 instance, we reuse the id of the subnets we
created earlier.



Once again, everything is detailed for every resource in the terraform
documentation.



#### Interpolation

If you read what I wrote just above in the resources section, you
understood that the barbaric syntax used for defining the subnets of our
aws\_instance is a variable value. In terraform you can use strings with
\${} to give some logic to the attributes you define. You can reuse ids
of components you created (as we did), you can even use loops,
conditions, maps, and even some mathematics function.

### Raffined our project

With the time, your project might get bigger, by that I mean, your
main.tf file might contains thousands of line, which make it tough to
maintain. Let's start with an intuitive approach: cut our main.tf file
into pieces.

#### Cutting the main.tf

Terraform allows you to create whatever files you want, during a plan or
an apply it will try to concatenate all the `*.tf` files in the current
folder in one file. It will solve the dependencies between the different
file and should be able to recreate the main.tf file as it was before we
cut it in pieces. We just need to be **logical** on how we will do it.

Let's put the provider in a different file first. Let's create a
`providers.tf` file in the current folder containing just our aws
logging creds.


```tf
# providers.tf

provider "aws" {
    access_key = "your_access_key"
    secret_key = "your_secret_key"
    region     = "eu-west-1"
}
```

You can even commit this file with empty credentials, and just make sure
you'll never push it again.

Let's test it!

`terraform apply`

**...**

Ok you're not reading the whole article. I said to test we use **plan
not apply**. Well anyway if you fell in the trap, this should have not
change anything.

Let's continue with our network configuration. Let's put the VPC,
gateway in it.

```tf
# vpc.tf

resource "aws_vpc" "default" {
    cidr_block = "172.16.0.0/16"
}

resource "aws_internet_gateway" "default"{
    vpc_id = "${aws_vpc.default.id}"
}
```

Let's do the same with subnets, security\_groups and instances

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
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}
```


```tf
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



At the end of the refactoring you'll get a nicer structure. This
structure will last for few times until you infrastructure starts to get
bigger. With time you'll see some limit to it like in case of complex
dependencies or even about SOC (Separation of concerns)

So here comes modules.

Modules
-------

With modules you can seperate components of your infrastructure inside
*modules* allowing you to prevent any repetition in your infra
definition. Really handy when you want to scale up some components of
your current infra or when you want to refactor "all the masters".

### Architecture

To start with modules, you juxg need to create a modules folder at the
root of your project, and create a folder with the name of the module
you want to create. Inside the last folder you just need to create three
files and you will have created your first module! Those files are:



-   variables.tf: This file contains all the variables you would like to
    parametrize your module. For example you will pass some vpc\_id or
    some AMIs.

-   main.tf: Like our first main.tf it will contains all the resource
    definition of your module.

-   outputs.tf: Contains all the variables you will need after executing
    the modules. The most common use is to *output* the public ips of
    the future created instances.



### Using it



In our example, the infra is quite simple so this module will be as
well. We will just put our simple\_instance inside the module. But in a
way as we can configure on which subnet, which az it will be.



Let's start creating the module.

    mkdir -p modules/my_cluster_of_instances
		touch modules/my_cluster_of_instances/{main,variables,output}.tf

Let's put our definition of the "simple\_instance" in the main.tf


```tf
# main.tf

resource "aws_instance" "simple_instance" {
    ami  = "${var.ami_id}"
    instance_type = "${var.instance_type}"
    subnet_id = "${var.subnet_ids}"
    count = "${cluster_size}"
}
```

If you saw, I changed the variable to be something defined from the
variable.tf file. I added a count attribute in case we want to scale up
this module

Now let's configure the variables:

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



# output.tf

output "private_ips" {
    value = ["${aws_instance.simple_instance.*.private_ip}"]
}
```


The last one use a wildcard as we don't know how many instances will be
created in the module. It means "*Ouput* a list of the instances's
private ips"

So now let's create a main.tf at the root of our project calling this
module:


```tf
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


```bash
terraform get # Will create reference to our module
terraform plan # Should destroy what we had before
```

Sadly terraform will not understand that this module represent your old
instance, and will try to remove your old instance to put the two news.

### Conclusion about project structure

So far, the module structure is the best one I met at the moment. I
usually comes with few folders at the root of my projects:



1.  modules/
2.  env1/
3.  env2/



And using envx as a proper seperation between environments. It's really
handy when you have a lot of differences between environments.



Other commands with terraform
-----------------------------

To finish this *super* *long* article, I will briefly speaks about the
command we did not see in the article.

### Taint

In terraform you can taint some resources so they will get destroyed at
the next apply. Really useful if you want to start from scratch with
some components.

From our module example, we would use it like that:

`terraform taint -module=awesome_instance instance.0`

So you need to specify to which module you're tainting the ressource
from. And after that resource\_name.which\_one .

>
>
> **Warning**: There's no support for wildcard yet according to [this
> github issue](https://github.com/hashicorp/terraform/issues/)
>
>

### Graph

If you have graphviz installed in your laptop, you can create a graph of
the terraform resources you define.

### Import

If you created some resources in the UI, adding their definition in your
tf project will not be enough. You need to add it to your terraform
state using the import command.

For example if we did create the ec2 instance in the ui. We would add it
like that in the tfstate file.

     terraform import aws_instance.my_instance the_id_of_the_instance

Conclusion
----------

There's still much more to say about terraform but I'll stop here as I
reached nearly 2500 words... We have seen that Terraform is a great tool
to manage infrastructure with code. Their .tf format is a really good
thing compared to simple json definition. And terraform is quite
permissive so you can start with a simple project and finish with
hundreds of modules calling each others making the adopotion of the tool
quite easy. Anyway I leave you enjoy your `terraform destroy`

Sur ce, codez bien! Ciao!


