---
title: "Why k8s Nodes Changes of Name When Using Aws"
date: 2018-01-31T21:58:23Z
draft: true
---

## Introduction

At work, I got a terrible issue with adding the AWS support to the OpenShift cluster. I really got confused as well as I was explicitly setting the nodename for this server. But whatever value I was putting in the config, it was never working, it always registered the node as the specific instance name that amazon would have given. But why??? Let's have a look.

## The homicide




