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
date: "2021-12-09"
title: "What is Google Cloud Deploy?"
description: "Teams are moving to the cloud to modernize their infrastructures, along with adopting DevOps practices to deliver faster, reliable, and quality software. They are making efforts to increase their deployment frequency, decrease the lead time for changes, lower change failure rate, and reduce the time to restore services after a failure.
One of the effective ways to achieve the above four metrics is to have an effective and robust continuous delivery pipeline."
canonicalURL: "https://binx.io/blog/2021/12/09/what-is-google-cloud-deploy/"
tags: ["google cloud", "cloud deploy", "devops"]
draft: false
cover:
    image: "/luca.png"
    alt: "Luca Cavallin | Software Engineer"
    caption: "Luca Cavallin | Software Engineer" # display caption under cover
    relative: false # when using page bundles set this to true
    hidden: true # only hide on current single page
# weight: 1
# aliases: ["/first"]
---
Developing container-based applications on Google Kubernetes Engine (GKE) can be challenging when there are dozens of pipelines across multiple environments. This is where Google Cloud Deploy comes into action!
[![Managed continuous delivery to GKE](https://storage.googleapis.com/gweb-cloudblog-publish/images/Cloud_Deploy_2.max-2800x2800.jpg "Managed continuous delivery to GKE")](https://storage.googleapis.com/gweb-cloudblog-publish/images/Cloud_Deploy_2.max-2800x2800.jpg "Managed continuous delivery to GKE")

### What is Google Cloud Deploy?
Google Cloud Deploy is a fully managed, scalable, and streamlined continuous delivery solution for GKE. It automates application delivery to multiple target environments via a defined promotion sequence.
With Google Cloud Deploy, you can construct reliable CI/CD pipelines that can automate build, deploy, and render jobs. A Cloud Deploy pipeline includes information, such as name (used for referring the pipeline), promotion sequence (used for providing the order of deployment to the targets), and targets (optional information). Moreover, it is easily integrable with popular tools, such as Gitlab CI, Jenkins, etc. It is accessible via CLI and API and also brings Skaffold (a command line tool that facilitates continuous development for Kubernetes-native applications) to your pipelines, thereby enhancing the reliability of pipelines.
<br /><br />
[![Solving for continuous delivery challenges](https://storage.googleapis.com/gweb-cloudblog-publish/original_images/cloud-deploy-pp-blog-post-3.gif "Solving for continuous delivery challenges")](http://https://storage.googleapis.com/gweb-cloudblog-publish/original_images/cloud-deploy-pp-blog-post-3.gif "Solving for continuous delivery challenges")

### Use Cases of Google Cloud Deploy
Some of the main use cases of Google Cloud Deploy are as follow:

- **Streamlined Continuous Delivery**: Cloud Deploy is highly useful to have simple and effective continuous delivery to Google Kubernetes Engine (GKE). You can create releases and advance them via environments, like test, production, and staging. In addition, its API, CLI, or web console can be used to have a one-step simple releases' rollback and promotion.
- **Integrated Solution**: Cloud Deploy is a tightly integrated GKE deployment platform. It comes pre-integrated to Cloud Audit Logs, Cloud Logging, and IAM. This way, it is effective to attain traceability with Cloud Audit Logs, monitor release events with Cloud Logging, and lockdown release progressions via IAM.
- **Scalable and Fully-Managed**: Cloud Deploy is a fully-managed service, which implies no expensive infrastructure to set up and maintain for the GKE CD pipeline. It can scale CD processes seamlessly and ensure their management through simple declarative configuration. It also assists to have a centralized view of all the pipelines!
<br /><br />
Learn more about it from this introduction video:
<br /><br />
<iframe width="560" height="315" src="https://www.youtube.com/embed/Il8FlhR9jKM" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

### A deployment example
Let's understand how Google Cloud Deploy works by learning how to register a delivery pipeline. In order to register the delivery pipeline, you have to run `gcloud beta deploy apply` for every pipeline configuration file. So, the command to register delivery pipeline with Cloud Deploy including its targets (the specific cluster and its configuration into which to deploy the application) is as follows (the `--region` and `--project` flags are optional):

```shell
gcloud beta deploy apply --file=PIPELINE_CONFIG \
--region=LOCATION \
--project=PROJECT
```

The structure of a `clouddeploy.yaml` file looks like this:
```yaml
   apiVersion: deploy.cloud.google.com/v1beta1
    kind: DeliveryPipeline
    metadata:
     name:
     annotations:
     labels:
    description:
    serialPipeline:
     stages:
     - targetId:
       profiles: []
     - targetId:
       profiles: []
     ---

     apiVersion: deploy.cloud.google.com/v1beta1
     kind: Target
     metadata:
      name:
      annotations:
      labels:
     description:
     requireApproval:
     gke:
      cluster: projects/[project_name]/locations/[location]/clusters/[cluster_name]

     executionConfigs:
     - privatePool:
         workerPool:
         serviceAccount:
         artifactStorage:
       usages:
       - [RENDER | DEPLOY]
     - defaultPool:
         serviceAccount:
         artifactStorage:
       usages:
       - [RENDER | DEPLOY]

     ---
```

To learn more about how to use Google Cloud Deploy, you can visit its [official guide](https://cloud.google.com/deploy "official guide").

### Summary
When it comes to streamlining the continuous delivery in Google Kubernetes Engine (GKE), Google Cloud Deploy presents an ideal choice for organizations. Being a fully-managed, easily scalable, and integrable solution, it makes releases and deployments an efficient and streamlined process.
