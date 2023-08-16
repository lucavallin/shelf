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
date: "2021-06-11"
title: "How to use Terraform workspaces to manage environment-based configuration"
description: "I have recently worked on a 100%-Terraform based project where I made extensive use of Workspaces and modules to easily manage the infrastructure for different environments on Google Cloud. This blog post explains the structure I have found to work best for the purpose."
canonicalURL: "https://binx.io/blog/2021/06/11/how-to-use-terraform-workspaces-to-manage-environment-based-configuration/"
tags: ["google cloud", "terraform", "devops"]
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

Workspaces are separate instances of state data that can be used from the same working directory. You can use workspaces to manage multiple non-overlapping groups of resources with the same configuration.
To create and switch to a new workspace, after running `terraform init`, run:
```shell
terraform workspace create <name>
```
To switch to other workspaces, run instead:
```shell
terraform workspace select <name>
```

### Use the selected workspace in the Terraform files
The selected workspace is made available in your `.tf` files via the `terraform.workspace` variable (it's a string). I like to assign the value to a local variable called `environment` (since the name of the workspaces and that of the environments match).

```hcl
locals {
	environment = terraform.workspace
}
```

# Add the environment-based configuration to a new module
Now that you have a variable containing the environment (it's just a string) that you are operating on, you can create a module containing the environment-based configuration. I created a `vars` module inside the `modules` directory of my repository, which contains at least the following files:

- **`main.tf`**
This file will never changes, as it's only needed to aggregate the variables that will be exported.
```hcl
locals {
	  environments = {
		"development" : local.development,
		"acceptance" : local.acceptance,
		"production" : local.production
	  }
}
```
- **`outputs.tf`**
This file too, will never change. Here I am defining the output of the `vars` module so that it can be used from anywhere else in the Terraform repository.
The exported values are based on the selected workspace.
```hcl
output "env" {
	value = local.environments[var.environment]
}
```
- **`variables.tf`**
This file defines the variables required to initialize the module. The `outputs` of the module are based on the selected workspace (environment), which it needs to be aware of.
```hcl
variable "environment" {
	description = "The environment which to fetch the configuration for."
	type = string
}
```
- **`development.tf`** & **`acceptance.tf`** & **`production.tf`**
These files contain the actual values that differ by environment. For example, when setting up a GKE cluster, you might want to use cheap machines for your development node pool, and more performant ones in production. This can be done by defining a `node_pool_machine_type` value in each environment, like so:
```hcl
// in development.tf
locals {
	development = {
		node_pool_machine_type = "n2-standard-2"
	}
}
```
```hcl
// in acceptance.tf
locals {
	acceptance = {
		node_pool_machine_type = "n2-standard-4"
	}
}
```
```hcl
// in production.tf
locals {
	production = {
		node_pool_machine_type = "n2-standard-8"
	}
}
```

The `vars` module is now ready to be used from anywhere in the repository, for example in `main.tf` file. To access the configuration values, initialize the module like so:
```hcl
#
# Fetch variables based on Environment
#
module "vars" {
	source      = "./modules/vars"
	environment = local.environment
}
```

The correct configuration will be returned based on the Terraform Workspace (
environment name) being passed to it, and values can be accessed via ` module.vars.env.<variable-name>`. For example:
```hcl
node_pools = [
	{
		...
		machine_type = module.vars.env.node_pool_machine_type
		...
	}
]
```

### Summary
In this blog post I have shown you how you can use Terraform Workspaces to switch between different configurations based on the environment you are working on, while keeping the setup as clean and simple as possible. Are you interested in more articles about Terraform? Checkout [How to Deploy ElasticSearch on GKE using Terraform and Helm](/posts/how-to-deploy-elasticsearch-on-gke-using-terraform-and-helm/)!
