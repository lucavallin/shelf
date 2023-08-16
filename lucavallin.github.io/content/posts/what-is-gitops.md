---
showToc: true
TocOpen: false
hidemeta: false
comments: true
ShowCanonicalLink: false
disableShare: false
searchHidden: true
hideSummary: false
ShowReadingTime: true
ShowBreadCrumbs: true
ShowPostNavLinks: true
# ShowWordCount: true
ShowRssButtonInSectionTermList: true
UseHugoToc: true
editPost:
    URL: "https://github.com/lucavallin/lucavallin.github.io/tree/main/content"
    Text: "Suggest Changes"
    appendFilePath: true
# CanonicalLinkText
# disableAnchoredHeadings
# hideFooter
# ShowCodeCopyButtons
# ShareButtons
# robotsNoIndex

author: "Luca Cavallin"
date: "2021-12-02"
title: "What is GitOps?"
description: "DevOps practices such as CI/CD, version control, and collaboration have helped teams delivering higher quality software faster. Thanks to DevOps culture, development teams are able to deploy code into production several times per day, automating most of the processes.There is, however, still a need for teams to manually perform infrastructural activities, such as maintaining, managing, and provisioning the infrastructure: this takes a lot of time and effort! Automation is the key to relieve teams of such burden and that's what GitOps is for. In this post, I am going to explain more about GitOps and its pros and cons. Let's jump right in!"
canonicalURL: "https://binx.io/blog/2021/12/01/what-is-gitops/"
tags: ["gitops", "ci/cd", "devops"]
draft: false
cover:
    image: "/luca.png"
    alt: "Luca Cavallin | Software Engineer"
    caption: "Luca Cavallin | Software Engineer" # display caption under cover
    relative: false # when using page bundles set this to true
    hidden: false # only hide on current single page
# weight: 1
# aliases: ["/first"]
---
<img src="https://binx.io/wp-content/uploads/2021/11/gitops-900x407.png" alt="GitOps in a nutshell" width="900" height="407" class="size-medium wp-image-4311" /><br /><br /><br /><br />
GitOps is known as an *operational framework* that takes the best practices of DevOps, such as CI/CD, version control and collaboration, and applies them to infrastructure automation.
GitOps is a set of code-based practices that use Git, an open-source version control system, to manage infrastructure and application configurations. Git is  the single source of truth, and pull requests are used to verify and automatically manage/deploy infrastructure changes.
Just like developers use application source code, operation teams that practice GitOps use infrastructure as code (configurations files stored as code) to generate the same infrastructure environment during each deployment. In short, GitOps uses similar processes and tools used in software development to manage infrastructure, ensuring the automation that teams need.

### How does GitOps work?
Since GitOps involves Git as a version control system, it can be considered as an advancement in Infrastructure as Code (IaC). In GitOps, changes are triggered via pull requests that change state in the Git repository.

A GitOps workflow for updating or creating new feature is as follow:

- Initiate a pull request for the new feature in Git
- Review code and merge it to Git
- Git will automatically trigger CI and build pipelines, run tests, apply infrastructure changes and create -when in a container-based environment- a new image which is then uploaded to the registry
- Specifically for applications, deployment tools can be used to automatically update the running version on, for example, a Kubernetes cluster or another serverless product

The workflows of GitOps are meant to improve productivity and speed of development and deployment, along with ensuring that systems remain stable and reliable!

### Pros and Cons of GitOps
Following are the major pros and cons associated with GitOps:

**Pros**
- It enhances the DevOps team productivity, as they can instantly deploy new configurations of infrastructures. If the changes aren't performing as required, the Git history lets the team easily revert back to a stable state
- It ensures faster deployment, as the team does not have to switch between tools for deploying the application (application and infrastructure changes can be deployed using the same tools)
- It brings end-to-end standardization to the workflow
- It can scale to hundreds of Kubernetes clusters
- Since Git is used for storing the complete information of the deployed infrastructure, changes can easily be tracked and also it encourages a culture of sharing knowledge in teams
- With automation in infrastructure definition and testing, lots of manual work go away. The team becomes more productive, while downtime get reduced due to rollback/revert capability. GitOps can lower costs significantly!

**Cons**
- GitOps encourages collaboration, but it also sometimes becomes a time-consuming and tedious job. For example, the approval process in GitOps involves many stages, such as creating a merge request, approving the changes, and deploying the changes. Engineers that are used to doing manual and fast changes might feel the whole process time-consuming
- Collaborative culture also requires discipline from all the members in order to ensure commitment to the process. Moreover, teams are required to write and formalize everything down so that GitOps can work perfectly.

### Summary
GitOps is a powerful framework for managing modern infrastructure with a focus on the developer experience. It enables infrastructure management from the same version control system involved in application development, empowering teams to have a central collaborative environment with improved system reliability and stability.
