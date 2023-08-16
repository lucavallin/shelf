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
date: "2021-07-23"
title: "How to Deploy a Multi-cluster Service Mesh on GKE with Anthos"
description: "Anthos Service Mesh is a suite of tools that helps you monitor and manage a reliable service mesh on-premises or on Google Cloud. I recently tested it as an alternative to an unmanaged Istio installation and I was surprised at how much easier Anthos makes it to deploy a service mesh on Kubernetes clusters."
canonicalURL: "https://binx.io/blog/2021/07/23/how-to-deploy-a-multi-cluster-service-mesh-on-gke-with-anthos/"
tags: ["google cloud", "anthos", "gke"]
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

In this article, I am going to explain step-by-step how I deployed a multi-cluster, multi-region service mesh using Anthos Service Mesh. During my proof of concept I read the documentation at https://cloud.google.com/service-mesh/docs/install, but none of the guides covered exactly my requirements, which are:
- Multi-cluster, multi-region service mesh
- Google-managed Istio control plane (for added resiliency, and to minimize my effort)
- Google-managed CA certificates for Istio mTLS


### Deploy the GKE clusters
Deploy the two GKE clusters. I called them `asm-a` and `asm-b` (easier to remember) and deployed them in two different regions (`us-west2-a` and `us-central1-a`). Because Anthos Service Mesh requires nodes to have at least 4 vCPUs (and a few more requirements, see the complete list at): https://cloud.google.com/service-mesh/docs/scripted-install/asm-onboarding), use at least the `e2-standard-4` machines.

As preparation work, store the Google Cloud Project ID in an environment variable so that the remaining commands can be copied and pasted directly.
```shell
export PROJECT_ID=$(gcloud info --format='value(config.project)')
```

Then, to deploy the clusters, run:

```shell
gcloud container clusters create asm-a --zone us-west2-a --machine-type "e2-standard-4" --disk-size "100" --num-nodes "2" --workload-pool=${PROJECT_ID}.svc.id.goog --async

gcloud container clusters create asm-b --zone us-central1-a --machine-type "e2-standard-4" --disk-size "100" --num-nodes "2" --workload-pool=${PROJECT_ID}.svc.id.goog --async
```

The commands are also enabling Workload Identity, which you can read more about at: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity.

### Fetch the credentials to the clusters
Once the clusters have been created, fetch the credentials needed to connect to the them via `kubectl`. Use the following commands:
```shell
gcloud container clusters get-credentials asm-a --zone us-west2-a --project ${PROJECT_ID}
gcloud container clusters get-credentials asm-b --zone us-central1-a --project ${PROJECT_ID}
```

### Easily switch kubectl context with kubectx
`kubectx` makes it easy to switch between clusters and namespaces in kubectl (also known as context) by creating a memorable alias for them (in this case, `asma` and `asmb`). Learn more about the tool at: https://github.com/ahmetb/kubectx.
```shell
kubectx asma=gke_${PROJECT_ID}_us-west2-a_asm-a
kubectx asmb=gke_${PROJECT_ID}_us-central1-a_asm-b
```

### Set the Mesh ID label for the clusters
Set the `mesh_id` label on the clusters before installing Anthos Service Mesh, which is needed by Anthos to identify which clusters belong to which mesh. The `mesh_id` is always in the format `proj-<your-project-number>`, and the project number for the project can be found by running:
```shell
gcloud projects list
```

Use these commands to create the `mesh_id` label on both clusters (replace `<your-project-number>` with the project number found with the previous command:

```shell
export MESH_ID="proj-<your-project-number>"
gcloud container clusters update asm-a --region us-west2-a --project=${PROJECT_ID} --update-labels=mesh_id=${MESH_ID}

gcloud container clusters update asm-b --region us-central1-a --project=${PROJECT_ID} --update-labels=mesh_id=${MESH_ID}
```

### Enable StackDriver
Enable StackDriver on the clusters to be able to see logs, should anything go wrong during the setup!

```shell
gcloud container clusters update asm-a --region us-west2-a --project=${PROJECT_ID} --enable-stackdriver-kubernetes

gcloud container clusters update asm-b --region us-central1-a --project=${PROJECT_ID} --enable-stackdriver-kubernetes
```

### Create firewall rules for cross-region communication
The clusters live in different regions, therefore a new firewall rule must be created to allow communication between them and their pods. Bash frenzy incoming!

```shell
ASMA_POD_CIDR=$(gcloud container clusters describe asm-a --zone us-west2-a --format=json | jq -r '.clusterIpv4Cidr')
ASMB_POD_CIDR=$(gcloud container clusters describe asm-b --zone us-central1-a --format=json | jq -r '.clusterIpv4Cidr')
ASMA_PRIMARY_CIDR=$(gcloud compute networks subnets describe default --region=us-west2 --format=json | jq -r '.ipCidrRange')
ASMB_PRIMARY_CIDR=$(gcloud compute networks subnets describe default --region=us-central1 --format=json | jq -r '.ipCidrRange')
ALL_CLUSTER_CIDRS=$ASMA_POD_CIDR,$ASMB_POD_CIDR,$ASMA_PRIMARY_CIDR,$ASMB_PRIMARY_CIDR

gcloud compute firewall-rules create asm-multicluster-rule \
    --allow=tcp,udp,icmp,esp,ah,sctp \
    --direction=INGRESS \
    --priority=900 \
    --source-ranges="${ALL_CLUSTER_CIDRS}" \
    --target-tags="${ALL_CLUSTER_NETTAGS}" --quiet
```

### Install Anthos Service Mesh
First, install the required local tools as explained here: https://cloud.google.com/service-mesh/docs/scripted-install/asm-onboarding#installing_required_tools.

The `install_asm` tool will install Anthos Service Mesh on the clusters. Pass these options to fulfill the initial requirements:
- `--managed`: Google-managed Istio control plane
- `--ca mesh_ca`: Google-managed CA certificates for Istio mTLS
- `--enable_registration`: automatically registers the clusters with Anthos (it can also be done manually later)
- `--enable_all`: all Google APIs required by the installation will be enabled automatically by the script

```shell
./install_asm --project_id ${PROJECT_ID} --cluster_name asm-a --cluster_location us-west2-a --mode install --managed --ca mesh_ca --output_dir asma --enable_registration --enable_all

./install_asm --project_id ${PROJECT_ID} --cluster_name asm-b --cluster_location us-central1-a --mode install --managed --ca mesh_ca --output_dir asmb --enable_registration --enable_all
```

### Configure endpoint discovery between clusters
Endpoint discovery makes the clusters to communicate with each other, for example, it enables discovery of service endpoints between the clusters.

Install the required local tools as explained here: https://cloud.google.com/service-mesh/docs/downloading-istioctl, then run the following commands:
```shell
istioctl x create-remote-secret --context=asma --name=asm-a| kubectl apply -f - --context=asmb

istioctl x create-remote-secret --context=asmb --name=asm-b| kubectl apply -f - --context=asma
```

### Testing the service mesh
Anthos Service Mesh is now ready! Let's deploy a sample application to verify cross-cluster traffic and fail-overs.

### Create the namespace for the Hello World app
Create a new namespace on both clusters and enable automatic Istio sidecar injection for both of them. Since the Istio control plane is managed by Google, the `istio-injection- istio.io/rev=` label is set to `asm-managed`.

```shell
kubectl create --context=asma namespace sample

kubectl label --context=asma namespace sample istio-injection- istio.io/rev=asm-managed --overwrite

kubectl create --context=asmb namespace sample

kubectl label --context=asmb namespace sample istio-injection- istio.io/rev=asm-managed --overwrite
```

### Create the Hello World  service
Deploy the services for the` Hello World` app on both clusters with:

```shell
kubectl create --context=asma -f https://raw.githubusercontent.com/istio/istio/1.9.5/samples/helloworld/helloworld.yaml -l service=helloworld -n sample

kubectl create --context=asmb -f https://raw.githubusercontent.com/istio/istio/1.9.5/samples/helloworld/helloworld.yaml -l service=helloworld -n sample
```

### Create the Hello World deployment
Deploy the `Hello World` sample app, which provides an endpoint that will return the version number of the application (the version number is different in the two clusters) and an `Hello World` message to go with it.
```shell
kubectl create --context=asma -f https://raw.githubusercontent.com/istio/istio/1.9.5/samples/helloworld/helloworld.yaml -l version=v1 -n sample

kubectl create --context=asmb -f https://raw.githubusercontent.com/istio/istio/1.9.5/samples/helloworld/helloworld.yaml -l version=v2 -n sample
```

### Deploy the Sleep pod
The Sleep application simulates downtime. Let's use it to test the resilience of the service mesh! To deploy the Sleep application, use:
```shell
kubectl apply --context=asma -f https://raw.githubusercontent.com/istio/istio/1.9.5/samples/sleep/sleep.yaml -n sample

kubectl apply --context=asmb -f https://raw.githubusercontent.com/istio/istio/1.9.5/samples/sleep/sleep.yaml -n sample
```

### Verify cross-cluster traffic
To verify that cross-cluster load balancing works as expected (read as: can the service mesh actually survive regional failures?), call the` HelloWorld` service several times using the Sleep pod. To ensure load balancing is working properly, call the `HelloWorld` service from all clusters in your deployment.

```shell
kubectl exec --context=asma -n sample -c sleep "$(kubectl get pod --context=asma -n sample -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -- curl -sS helloworld.sample:5000/hello

kubectl exec --context=asmb -n sample -c sleep "$(kubectl get pod --context=asmb -n sample -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -- curl -sS helloworld.sample:5000/hello
```

Repeat this request several times and verify that the `HelloWorld` version should toggle between `v1` and `v2`. This means the request is relayed to the healthy cluster when the other one is not responding!


### Summary

In this article I have explained how I deployed Anthos Service Mesh on two GKE clusters in different regions with Google-managed Istio control plane and CA certificates. Anthos Service Mesh makes it simple to deploy a multi-cluster service mesh, because most of the complexity of Istio is now managed by Google.
