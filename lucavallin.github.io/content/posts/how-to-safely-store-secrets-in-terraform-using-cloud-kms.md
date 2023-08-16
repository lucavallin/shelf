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
date: "2021-07-04"
title: "How to safely store secrets in Terraform using Cloud KMS"
description: "I have recently worked on a project where I needed to configure a Helm release with secrets hard-coded in Terraform. With Cloud KMS, I could encrypt the secrets so that they could safely be committed to git. In this article, I am going to show you how the process works."
canonicalURL: "https://binx.io/blog/2021/07/04/how-to-safely-store-secrets-in-terraform-using-cloud-kms/"
tags: ["google cloud", "cloud kms", "terraform", "security"]
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
Since my project is already up and running, all I had to do was to create a Cloud KMS `keyring` and `crypto key` that will be used for encrypting and decrypting secrets. This can be done via Terraform with a few resources:

```hcl
data "google_project" "secrets" {
  project_id = "secrets"
}

resource "google_kms_key_ring" "this" {
  project  = data.google_project.secrets.project_id
  name     = "default"
  location = "europe-west4"
}

resource "google_kms_crypto_key" "this" {
  name     = "default"
  key_ring = google_kms_key_ring.this.self_link
}

resource "google_project_iam_member" "this" {
  project = data.google_project.secrets.project_id
  role    = "roles/cloudkms.cryptoKeyEncrypter"
  member  = "group:encrypter@example.com"
}
```

In this example, I am also assigning the `cloudkms.cryptoKeyEncrypter` role to an imaginary `encrypter@example.com` group so that members can encrypt new secrets.

### Encrypt the secrets
I then used this command to encrypt a secret for use in Terraform via Cloud KMS:

```bash
echo -n <your-secret> | gcloud kms encrypt --project <project-name> --location <region> --keyring default --key default --plaintext-file - --ciphertext-file - | base64
```

To encrypt a whole file instead, you can run:
```bash
cat <path-to-your-secret> | gcloud kms encrypt --project <project-name> --location <region> --keyring default --key default --plaintext-file - --ciphertext-file - | base64
```
- `<your-secret>` is the string or file you want to encrypt
- `<project-name>` is the project whose KMS keys should be used to encrypt the secret.
- `<region>` is the region in which the KMS keyring is configured.

If you are working on macOS, you can append `| pbcopy` to the command so the resulting output will be added automatically to the clipboard.

Finally, an encrypted string looks like this:
```
YmlueGlzYXdlc29tZWJpbnhpc2F3ZXNvbWViaW54aXNhd2Vzb21lYmlueGlzYXdlc29tZWJpbnhpc2F3ZXNvbWViaW54aXNhd2Vzb21lYmlueGlzYXdlc29tZQ
```

### Configure the Cloud KMS data source
The encrypted string is not stored anywhere. What I did instead, is to use the Cloud KMS data source to decrypt it on-the-fly.
```hcl
data "google_kms_secret" "secret_key" {
  crypto_key = data.google_kms_crypto_key.this.self_link
  ciphertext = "YmlueGlzYXdlc29tZWJpbnhpc2F3ZXNvbWViaW54aXNhd2Vzb21lYmlueGlzYXdlc29tZWJpbnhpc2F3ZXNvbWViaW54aXNhd2Vzb21lYmlueGlzYXdlc29tZQ"
}
```

The decrypted value can be referenced with:

```hcl
data.google_kms_secret.secret_key.plaintext
```

Terraform will decrypt the string automagically and replace it with the actual value where it is referenced.

### Summary

In this short post I have explained how I used Cloud KMS to encrypt strings that can be safely committed to git and used in Terraform. If you have multiple environments that you need to support, remember that KMS keys are different for each project, therefore you will need to encrypt the same string again, once for each project. If you need some input on how to organize a multi-project Terraform repository, have a look at my recent article [How to use Terraform workspaces to manage environment-based configuration
](/posts/how-to-use-terraform-workspaces-to-manage-environment-based-configuration/).
