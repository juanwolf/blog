---
title: "What is a NAT Gateway?"
date: 2018-12-20T08:27:38+01:00
---

Looking into designing platform and new architectures can be quite intimidating, especially when it's your first time or you come from a software background.
Today I want to speak to you about the NAT Gateways, what is it? How it works? When to use it, etc... It's gonna be quite straight forward, _everyone_ use this component
but I would like to go even deeper and see what happens at the packet level.

## What does even NAT means??

So first things first, what is actually a NAT?? If we listen to papy wikipedia, he says:

> Network address translation (NAT) is a method of remapping one IP address space into another by modifying network address information in the IP header of packets while they are in transit across a traffic routing device.[1] The technique was originally used as a shortcut to avoid the need to readdress every host when a network was moved. It has become a popular and essential tool in conserving global address space in the face of IPv4 address exhaustion. One Internet-routable IP address of a NAT gateway can be used for an entire private network.
Sources: [Wikipedia](https://en.wikipedia.org/wiki/Network_address_translation)

Ok Grandpa, you kinda spoiled everyone and explained everything I wanted to cover in this article. Anyway... So as Grandpa said, it's pretty common to see an infrastructure using NAT Gateways, especially these days with the Cloud. In the cloud, there's one rule. Keep everything private (at least as much as possible).

Commonly when you start playing with the :cloud:, you'll create one instance, two, etc... but to access directly to these instances, you will setup those with a Public IP, it's easier. Then your infra gets a bit more mature (and you as well) and you realise that using Public IPs might turn against you if you get a bit too permissive with firewall rules (or security groups). So you decide to move everything into a private subnet (only accessible inside your virtual private cloud) and use the protagonist of this article. THE NAT GATEWAY.

We are quite lucky, cloud providers sells specific instances or services to take care of this for us. _Too handy_. But at the end, how is this actually working? Let's say you followed the instruction of your cloud providers to have a nat.

Then you have a platform like this:

![Schema of an architecture in the cloud with one vpc, one private subnet with four instances connected to the nat gateway which is inside a public subnet and relay the traffic to the internet](/post_content/2018-06-25/nat_vpc.svg)

As you can see in the diagram, every outbound connection goes through the NAT Gateway and goes back to the specific instance. My question is how? The instance has no public IP, so the server on the internet when he does receive a packet, sees only the NAT instance's IP. So how does a NAT gateway works?

## How does a NAT gateway work?

So as grandpa said, the NAT Gateway inject its own address inside each packet so the external server knows where to send back the packet (**Spoiler:** To the NAT Gateway).

But how does the NAT knows where to send the packet back?

![Nat packet translation](/post_content/2018-06-25/nat_anim.svg)

So what happen is:

Your request goes to the NAT Gateway, It will receive the packet and store the localisation of the sender and the destination in what we call a _NAT table_. The NAT will then change the _"from"_ header of your request to replace the source IP with the one of the NAT, send the packet to the proper destination, then your nat gets the response back, change the destination of the packet/request for the instance that's in your private subnet, and boom. Job done.

You can imagine there could be some clash if the red and the blue instance queries the exact same endpoint, the NAT will not know where to redirect the response.

So I was thinking to change the source port for a random one so we would know to which host redirect this request! I read a bit more on the internet that's actually what's happening (so pretty proud of myself there :smile:)

Let's summarize:

1. Your request/packet goes to the nat
2. The nat register the source IP, source port, destination IP, destination port and get a random port available and put all this info in the translation table
3. Modifies the source IP and the source port for the NAT public IP and the random port allocated to this request
4. Send the request/packet to destination
5. The destination server process the request and sends the request back to the NAT gateway (as it thinks it comes from it, lol, you fool)
6. The NAT receives the packet/request, check its translation table, if there's an entry for it, sends it in the private instance
7. The private instance receives the packet.

## Limitations

Using NATs allows us to keep our instances "_private_" and still give us external traffic. But there must be some limitation to it. Of course there is. During the second step of the translation (aka registration of the packet in the nat table), the NAT Gateway will allocate a random port for the response of the server. The problem is, let's say you have 20 servers running 20 services that sends 10 requests per second to external services. Sadly those external services gets unreachable, the connection will wait for a timeout of... let's say 15 seconds. In less than 10 seconds, the NAT gateway will report unhealthy or at least unable to process any more requests.

 Why is that?

You used all the [ephemeral ports](https://www.ncftp.com/ncftpd/doc/misc/ephemeral_ports.html) available on the NAT. :boom:

Let's do some math. We open 20 * 20 * 10 requests per seconds aka 4000requests/s. So each requests open an ephemeral port in the NAT. 4000 ports opened. Second second 4000 others, third second you have already 12000 ephemeral ports opened. It will vary of your operating system but according to [Wikipedia](https://en.wikipedia.org/wiki/Ephemeral_port#Range):

> Many Linux kernels use the port range 32768 to 61000. (The effective range is accessible via the /proc file system at node /proc/sys/net/ipv4/ip_local_port_range)

So roughly 28232 ports I would say (just guessing). That's our maximum of ephemeral ports that can be opened. Well congratulations, we exploded the limit and we'll have a lots of problems after that. Glad, it was just a role play and not an actual outage :sweat_smile:.

## Conclusion

Et voila! We saw how a NAT Gateway behave and its limitation. They are one of the core component used in an infra in the Cloud so I am sure knowing that will help you when dealing with SREs/DevOps. Sur ce codez bien!

### PS:
I tried to implement one [here](https://github.com/juanwolf/toran). I focused on getting a first version out before releasing this article but as the project were never ending (started in June) and I did not get the time (and the appetite at the end) to finish it properly, you can enjoy (if enjoyable) the draft I made of it.

