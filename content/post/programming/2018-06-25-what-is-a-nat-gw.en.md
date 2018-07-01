---
title: "What is a NAT Gateway?"
date: 2018-06-25T08:27:38+01:00
draft: true
---

Looking into designing platform and new architectures can be quite intimidating, especially when it's your first time or you come from a software background.
Today I want to speak to you about the NAT Gateways, what is it? How it works? When to use it, etc... It's gonna be quite straight forward, _everyone_ use this component
but I would like to go even deeper and see what happens at the packet level.

## What is a NAT Gateway

So first things first, what is actually a nat gateway?? If we listen to papy wikipedia, he says:

> Network address translation (NAT) is a method of remapping one IP address space into another by modifying network address information in the IP header of packets while they are in transit across a traffic routing device.[1] The technique was originally used as a shortcut to avoid the need to readdress every host when a network was moved. It has become a popular and essential tool in conserving global address space in the face of IPv4 address exhaustion. One Internet-routable IP address of a NAT gateway can be used for an entire private network.
Sources: [Wikipedia](https://en.wikipedia.org/wiki/Network_address_translation)

Ok Grandpa, you kinda spoiled everyone and explained everything I wanted to cover in this article. Anyway... So as Grandpa said, it's pretty common to see an infrastructure using NAT Gateways, especially these days with the Cloud. In the cloud, there's one rule. Keep everything private. Then the NAT appears.

Commonly when you start playing with the :cloud:, you'll create one instance, two, etc... but to access directly to these instances, you will setup those with an Ephemeral IP (aka Public IP), it's easier. Then your infra gets a bit more mature (and you as well) and you realise that using Public IPs might turn against you if you get a bit too permissive with firewall rules (or security groups). So you decide to move everything into a private subnet (only accessible inside your virtual private cloud) and use the protagonist of this article. THE NAT  GATEWAY :princess:

We are quite lucky, cloud providers sells specific instances or services to take care of this for us. _too handy_. But at the end, how is this actually working? Let's say you followed the instruction of your cloud providers to have a nat.

Then you have a platform like this:

![Schema of an architecture in the cloud with one vpc, one private subnet with four instances connected to the nat gateway which is inside a public subnet and relay the traffic to the internet](/post_content/2018-06-25/nat_vpc.svg)

As you can see in the diagram, every outbound connection goes through the NAT Gateway and goes back to the specific instance. My question is how? The instance has no public IP, so the server on the internet when he does receive a packet, sees only the NAT instance's IP. So how does a NAT gateway works?

## How does a NAT gateway work?

So as grandpa spoiled us, the nat gateway literally receives every packet sent to it and inject inside of it the nat address ip so the external server knows where to send back the packet.

But how does the NAT knows where to send the packet back?

![Nat packet translation](/post_content/2018-06-25/nat_anim.svg)

So what happens is:

Your request goes to the NAT Gateway, It will receive the packet and store the localisation of the sender and the destination in what we call a _NAT table_. The nat will then change the from header of your request to replace the source IP with the one of the NAT, send the packet to the proper destination, then your nat gets the response back, change the destination of the packet/request for the instance that's in your private subnet, and boom. Job done.

You can imagine there could be some clash if the red and the blue instance queries the exact same endpoint, the NAT will not know where to redirect the response.

So I was thinking to change as well the source port for a random one so we would know to which host redirect this request! I read a bit more on the internet that's actually what's happening (so pretty proud of myself there :smile:)

Let's summarize:

1. Your request/packet goes to the nat
2. The nat register the source IP, source port, destination IP, destination port and get a random port available and put all this info in the translation table
3. Modifies the source IP and the source port for the NAT public IP and the random port allocated to this request
4. Send the request/packet to destination
5. The destination server process the request and sends the request back to the NAT gateway (as it thinks it comes from it, lol, you fool)
6. The NAT receives the packet/request, check its translation table, if there's an entry for it, sends it in the private instance
7. The private instance receives the packet.

## Let's implement one!

I was thinking that it could be a really cool exercise to implement it. As I need to lvl up my skills in go, you'll be the fruit of this experience :smile:

### 1. Defining the Translation table
### 2. Listening for every incoming packages
### 3. Make the magic happen


## Conclusion

Et voila! We saw how a NAT Gateway behave and its limitation. They are one of the core component used in an infra in the Cloud so I am sure knowing that will help you when dealing with SREs/DevOps.
