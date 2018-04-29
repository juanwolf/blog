
---
title: Why did I migrate all my projects to GitLab?
date: 2015-05-31
tags: ["git"]
categories: ["Programming"]

draft: false
author: "Jean-Loup Adde"
---

Since the popularization of git for IT projects, it became the most used
tool for your project. With this phenomen, new services has been created
as GitHub and Bitbucket (for the most known). The first one is focused
on the social aspect and the second one is based on his privatisation of
the projects (I'm only speaking about free offers). One way the most
secure to keep its code was to install a proper git server managed with
GitLab or GitBlit. However since few month, GitBlit launched a free
offer to host project with unlimited private projects. I gave it a shot,
and I fastly moved all my projects to this new platform, and that's this
point I'll try to explain during this article.

![](/post_preview/20150531_154323_b_1_q_0_p_1.jpg.png)

## GitLab, What is it?

From Grandpa Wikipedia, GitLab is :

> GitLab is a web-based Git repository manager with wiki and issue
> tracking features. GitLab is similar to GitHub, but GitLab has an open
> source version, unlike GitHub.

Thanks GrandPa' \! The problem of using GitLab was to own a proper
server to install it. However, everyone doesn't own one, and it's how
take advantage of the situation offering an git management already
hosted by themselves.

## Free offers available

I will only do a comparison between all the free online services to
manage your code.


|                                |   GitLab  | Bitbucket | GitHub      |
| ------------------------------ | --------- | --------- | ------------|
| Number Of Private Repositories | Unlimited | Unlimited | 0           |
| Max Number of Contributors     | Unlimited | 5         | Unlimited   |
| Contributor Stats              |  Yes      | No        | Yes         |
| SVN Support                    |  No       | Yes       | No          |
| Mercurial Support              |  No       | Yes       | No          |
| Code Search                    |  Yes      | No        | Yes         |
| Services integration           |  Good     | Low       | Really good |
| Milestone Management           |  Yes      | *         | *           |
| Continuous integration service |  Yes      | No        | No          |
| Branches protection            |  Yes      | No        | No          |

We can easily categorized the different services by tendancies. GitHub
is more oriented for social purpose to create contributions. Bitbucket
is focused on project management (Atlassian = JIRA = project management)
and GitLab is a kind of hybrid.

## Interface

The interface is sober and clear. We'll see all the main pages, and
you'll see it's really user friendly.

### Homepage

Classic:

![Home](/post_content/2015-05-31/e1d631a3-9fd8-4403-a3b1-87471d8a98dc.png)

The events linked to your projects, filterable by events (push, merge,
comment) and the list of your personnal projects.

### Profile page

It's a profile page more personal than one in GitHub, but anyway we keep
the contributions section to see your activity during the last month and
we can find all your personnal details such as linked profile, tweeter,
skype...

![Profile](/post_content/2015-05-31/ffb23551-9c55-4278-91ce-6c21fa83b617.png)

### Project Page

![GitLab](/post_content/2015-05-31/6ea6880e-03be-47e3-a5ec-e68dd8f18290.png)

Similar to the home page, it contains more sections in the menu, and we
can navigate easily to other pages to manage project such as the issues,
activity graph of the repo, activity, and the Mileston page.

### Milestone Page

The most interessing page in the project section. You can define some
milestone to manage your project. You can link issues to it so you can
see the progress of the project. An example here of the progress of this
article the 28th of May
:

![Milestone](/post_content/2015-05-31/abfd6071-0316-4269-85d5-1ecf8485fc22.png)

## Social

GitLab contains a little section called Explore where you can find all
the most popular projects in GitLab. But still nothing compare to the
social aspect to GitHub with his discover section.

So GitLab doesn't really bet on social purpose. However, the service is
available since few month, so we can expect some new features and maybe
a better social aspect for the new versions of GitLab.

## Conclusion

GitLa is one of the best hosting code platform at my sense. Because I'm
not a contributor, I'm more focused on a project management aspect than
a social one. Here GitLab does everything that I expect from a service
like that. But obviously I give the advice to create a project on GitHub
if you expect to work with a big community.

Personally, I migrate all my GitHub/Bitbucket projects on GitLab. Do you
?

