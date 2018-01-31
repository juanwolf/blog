---
title: "Why k8s Nodes Changes of Name When Using Aws"
date: 2018-01-31
tags: ["openshift", "kubernetes", "cloud", "aws"]
---

## Introduction

At work, I got a terrible issue with adding the AWS support to the OpenShift cluster. I really got confused as well as I was explicitly setting the nodename for this server. But whatever value I was putting in the config, it was never working, it always registered the node as the specific instance name that amazon would have given. But why??? Let's have a look.

## The homicide

A long time ago, a nice little instance in AWS called node00.k8s.internal were living its peaceful life. Deploying pods, serving services & passing good times with its masters, life were simple. But node00 started to feel sad. It was seeing all this nice ebs running around and it could not even use those! How unfortunate! Thinking about all the new services it could provide, node00 determined to fullfill at the best its duty, cross the line to the ebs, reach them and succeed to get one. node00 never been so happy! But the joy did not last long, it realised that the master were not speaking to it anymore! Node00 started to feel depress and started to cry. When it finished, it took a look at itself in the mirror and realised that since it crossed the barreer, it was not the same anymore! Node00.k8s.internal were gone but became ip-172-12-4-2.compute.internal instead. Realising that was the issue blocking it to speak to the master, Ex Node00 decided to proove to the master who it was. "I AM NODE00" it shouted to the master, the master did not respond, so it shouted, shouted and shouted again until it became tired of it. Ex node00 became so tired that it decided to give up and leave its life on its own, errant without hope, loosing all connection to the master. THE END.

Okay, let's talk with technical words now. So when I activated the AWS support my good old node node00.k8s.internal turned into the default aws instance name which is equal to its private dns name which looks like something like ip-172-12-4-2.compute.internal. And sadly for me, connection between all the components in openshift are super secure using certificates. Of course ip-172... were not part of the names allowed in the node certificate, so bye bye the node.
I tried everything to make the node takes its old name back. Setting up private dns zone, rename hostname of the ec2 instance, nothing worked and then I thought, fuck it, I'll dig into the code.

## The investigation

As OpenShift does not add much features on the cloud support I thought going directly inside kubernetes code to troubleshoot my issue. Let's start!

```
git clone https://github.com/kubernetes/kubernetes.git
```

You can go and take a coffee, realise you don't have any coffee left, going out and buy some, coming back, prepare it, drink it, that the download will not be finished. (It's actually not that bad when you have a 6MB/s connection)
There's more than 50000 commits so it'll take time (& space! nearly 2Go).

Anyway once that's done, let's look for some aws code:

```
cd kubernetes
grep -r -i "aws" .
```

Bloody heel, that's a lot. Let's reduce that to the go files.

```
grep -r --include="*.go" -i "aws" .
```

dsapoidspgjfdas[pg. Let's go a bit quicker:

```
grep -r --include="*.go" --exclude-dir="vendor" --exclude="*test*" --exclude-dir=test -i "aws" .
```

So we remove the tests file & folder, the vendor directory. And we end up with a nice start. There's a bunch of code for volumes but one file gets particularly my attention : `./pkg/cloudprovider/providers/aws/aws.go`, I have a feeling that this file is the one, the criminal.

Let's have a look into it. After a search on nodeName, we can find a nice structure:

```
type awsInstance struct {
    ec2 EC2
    // id in AWS
    awsID string

    // node name in k8s
    nodeName types.NodeName

    // availability zone the instance resides in
    availabilityZone string

    // ID of VPC the instance resides in
    vpcID string

    // ID of subnet the instance resides in
    subnetID string

    // instance type
    instanceType string
}
```

We are getting close! How this nodeName gets setted though...

```
// newAWSInstance creates a new awsInstance object
func newAWSInstance(ec2Service EC2, instance *ec2.Instance) *awsInstance {
    az := ""
    if instance.Placement != nil {
        az = aws.StringValue(instance.Placement.AvailabilityZone)
    }
    self := &awsInstance{
        ec2:              ec2Service,
        awsID:            aws.StringValue(instance.InstanceId),
        nodeName:         mapInstanceToNodeName(instance),
        availabilityZone: az,
        instanceType:     aws.StringValue(instance.InstanceType),
        vpcID:            aws.StringValue(instance.VpcId),
        subnetID:         aws.StringValue(instance.SubnetId),
    }

    return self
}

```

Ok, so the aws instance is getting its node name using the mapInstanceToNodeName function, so close, I am burning. Let's search for

```
func mapInstanceToNodeName
```

and ...


## The murderer

```
// mapInstanceToNodeName maps a EC2 instance to a k8s NodeName, by extracting the PrivateDNSName
func mapInstanceToNodeName(i *ec2.Instance) types.NodeName {
    return types.NodeName(aws.StringValue(i.PrivateDnsName))
}
```

BOOM. Here it is. This code is forcing the nodeName to be set as the PrivateDnsName of your ec2 instance. __sigh__ what a journey.

## Conclusion

We finally saw why kubernetes is changing suddenly the nodename of our node. You might think that's pointless, but in a cloud environment people might not think to setup private DNS zones and entries for their host, which AWS does it for you as well, so why bother? And basing your nodename on the hostname will not be enough in case someone rename it (like I did lol). So that's a nice prevention. But we were not expecting this *AT ALL* . So now you know, if you plan to use ebs for persistant storage make sure your nodes are names as their privateDNSName and you'll have no problem. If you want to distinguish your nodes, I highly invite you to setup labels on them such as kind or type, region, az etc... Anything that could be useful for you to pin a pod to a specific node. But anyway I hope you enjoyed this little investigation :)

Sur ce codez bien. Ciao !
