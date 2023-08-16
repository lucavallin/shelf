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
date: "2021-05-02"
title: "How to Deploy ElasticSearch on GKE using Terraform and Helm"
description: "I was recently tasked with deploying ElasticSearch on **GKE using Terraform and Helm**, and doing so in most readable way possible. I wasn't very familiar with Helm before, so I did some research to find approach that would fulfill the requirements. In this post I will share with you the Terraform configuration I  used to achieve a successful deployment."
canonicalURL: "https://binx.io/blog/2021/05/02/how-to-deploy-elasticsearch-on-gke-using-terraform-and-helm/"
tags: ["google cloud", "gke", "terraform", "elasticsearch", "helm"]
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

Helm is, at its most basic, a templating engine to help you define, install, and upgrade applications running on Kubernetes. Using Helm, you can leverage its **Charts feature, which are simply Kubernetes YAML configuration files** (that can be further configured and extended) combined into a single package that can be used to deploy applications on a Kubernetes cluster. To be able to use Helm via Terraform, we need to **define the corresponding provider and pass the credentials needed to connect to the GKE cluster**.

```hcl
provider "helm" {
  kubernetes {
    token                  = data.google_client_config.client.access_token
    host                   = data.google_container_cluster.gke.endpoint
    cluster_ca_certificate = base64decode(data.google_container_cluster.gke.master_auth[0].cluster_ca_certificate)
  }
}
```

#### Terraform configuration
I am defining an `helm_release` resource with Terraform, which will deploy the ElasticSearch cluster when applied. Since I am using an Helm chart for the cluster, doing so is incredibly easy. All I had to do was tell Helm the name of the chart to use and where it is located (repository), along with the version of ElasticSearch that I would like to use.
With the `set` blocks instead, I can override the default values from the template: this makes it easy to select an appropriate storage class, amount of storage and in general any other piece of configuration that can be changed (you have to refer to the documentation of the chart itself to see which values can be overridden), **directly from Terraform**.

```hcl
resource "helm_release" "elasticsearch" {
  name       = "elasticsearch"
  repository = "https://helm.elastic.co"
  chart      = "elasticsearch"
  version    = "6.8.14"
  timeout    = 900

  set {
    name  = "volumeClaimTemplate.storageClassName"
    value = "elasticsearch-ssd"
  }

  set {
    name  = "volumeClaimTemplate.resources.requests.storage"
    value = "5Gi"
  }

  set {
    name  = "imageTag"
    value = "6.8.14"
  }
}
```
#### Creating a new storage class
I then had to provision a new storage class, which will be used by the ElasticSearch cluster to store data. The configuration below sets up the SSD (SSD is recommended for such purpose, since it's faster than a regular HDD) persistent disk that I referenced in the main configuration above.
```hcl
resource "kubernetes_storage_class" "elasticsearch_ssd" {
  metadata {
    name = "elasticsearch-ssd"
  }
  storage_provisioner = "kubernetes.io/gce-pd"
  reclaim_policy      = "Retain"
  parameters = {
    type = "pd-ssd"
  }
  allow_volume_expansion = true
}
```

#### Summary

In this blog post I have shown you how to deploy ElasticSearch on GKE using Terraform and Helm. The required configuration is simple and very readable, lowering the barrier to handling all of your infrastructure via Terraform, rather than, for example, using Cloud Marketplace, managed services, or other custom solutions.
