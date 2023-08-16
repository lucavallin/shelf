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
# # ShowWordCount: true
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
date: "2021-07-26"
title: "A comparison of Google Cloud Pub/Sub and NATS"
description: "This article presents a comparison of Cloud Pub/Sub and NATS as message brokers for the distributed applications. We are going to focus on the differences, advantages and disadvantages of both systems."
canonicalURL: "https://binx.io/blog/2021/07/26/a-comparison-of-google-cloud-pub-sub-and-nats/"
tags: ["google cloud", "cloud pubsub", "nats.io"]
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

### Cloud Pub/Sub

Cloud Pub/Sub provides messaging and ingestion features for event-driven systems and streaming analytics. The highlights of the tool can be summarized as follows:

- Scalable, in-order message delivery with pull and push modes
- Auto-scaling and auto-provisioning with support from zero to hundreds of GB/second
- Independent quota and billing for publishers and subscribers
- Global message routing to simplify multi-region systems

Furthermore, Cloud Pub/Sub provides the following benefits over non-Google-managed systems:

- Synchronous, cross-zone message replication and per-message receipt tracking ensures reliable delivery at any scale
- Auto-scaling and auto-provisioning with no partitions eliminates planning and ensures workloads are production ready from day one
- Filtering, dead-letter delivery, and exponential backoff without sacrificing scale help simplify your applications
- Native Dataflow integration enables reliable, expressive, exactly-once processing and integration of event streams in Java, Python, and SQL.
- Optional per-key ordering simplifies stateful application logic without sacrificing horizontal scale—no partitions required.
- Pub/Sub Lite aims to be the lowest cost option for high-volume  event ingestion. - Pub/Sub Lite offers zonal storage and puts you in control of capacity management.

Some use cases of Cloud Pub/Sub include:

- Google’s stream analytics makes data more organized, useful, and accessible from the instant it’s generated. Built on Pub/Sub along with Dataflow and BigQuery, their streaming solution provisions the resources needed to ingest, process, and analyze fluctuating volumes of real-time data for real-time business insights. This abstracted provisioning reduces complexity and makes stream analytics accessible to both data analysts and data engineers.
- Pub/Sub works as a messaging middleware for traditional service integration or a simple communication medium for modern microservices. Push subscriptions deliver events to serverless webhooks on Cloud Functions, App Engine, Cloud Run, or custom environments on Google Kubernetes Engine or Compute Engine. Low-latency pull delivery is available when exposing webhooks is not an option or for efficient handling of higher throughput streams.

### Features

Cloud Pub/Sub offers the following features:

- **At-least-once delivery**: Synchronous, cross-zone message replication and per-message receipt tracking ensures at-least-once delivery at any scale.
- **Open**: Open APIs and client libraries in seven languages support cross-cloud and hybrid deployments.
- **Exactly-once processing**: Dataflow supports reliable, expressive, exactly-once processing of Pub/Sub streams.
- **No provisioning, auto-everything**: Pub/Sub does not have shards or partitions. Just set your quota, publish, and consume.
- **Compliance and security**: Pub/Sub is a HIPAA-compliant service, offering fine-grained access controls and end-to-end encryption.
- **Google Cloud–native integrations**: Take advantage of integrations with multiple services, such as Cloud Storage and Gmail update events and Cloud Functions for serverless event-driven computing.
- **Third-party and OSS integrations**: Pub/Sub provides third-party integrations with  Splunk and Datadog for logs along with Striim and Informatica for data integration. Additionally, OSS integrations are available through Confluent Cloud for Apache Kafka and Knative Eventing for Kubernetes-based serverless workloads.
- **Seek and replay**: Rewind your backlog to any point in time or a snapshot, giving the ability to reprocess the messages. Fast forward to discard outdated data.
- **Dead letter topics**: Dead letter topics allow for messages unable to be processed by subscriber applications to be put aside for offline examination and debugging so that other messages can be processed without delay.
- **Filtering**: Pub/Sub can filter messages based upon attributes in order to reduce delivery volumes to subscribers.

### Pricing

Cloud Pub/Sub is free up to 10GB/month of traffic, and above this threshold, a flat-rate of $40.00/TB/month applies.

### Summary of Cloud Pub/Sub

Cloud Pub/Sub is the default choice for cloud-native applications running on Google Cloud. Overall, the pros and cons of the tool can be summarized with the following points:


**Main advantages**
- Google-managed. There is no complex setup or configuration needed to use it.
- Integrations. Cloud Pub/Sub integrates seamlessly with other Google Cloud services, for example Kubernetes Engine.
- Secure. End-to-end encryption enabled by default and built-in HIPAA compliance.

**Main disadvantages**
- Limited choice of patterns. Cloud Pub/Sub implements the Publisher-Subscriber pattern, however, it is possible to set up a one-to-many Pub/Sub system as described at https://cloud.google.com/pubsub/docs/building-pubsub-messaging-system.


Refer to https://cloud.google.com/pubsub/docs/overview for further information.

### NATS

NATS is a message broker that enables applications to securely communicate across any combination of cloud vendors, on-premise, edge, web and mobile, and devices. NATS consists of a family of open source products that are tightly integrated but can be deployed easily and independently. NATS facilitates building distributed applications and it provides Client APIs are in over 40 languages and frameworks including Go, Java, JavaScript/TypeScript, Python, Ruby, Rust, C#, C, and NGINX.  Furthermore, real time data streaming, highly resilient data storage and flexible data retrieval are supported through JetStream , which is built into the NATS server.

 The highlights of the tool can be summarized as follows:

- With flexible deployments models using clusters, superclusters, and leaf nodes, optimize communications for your unique deployment. The NATS Adaptive Edge Architecture allows for a perfect fit for unique needs to connect devices, edge, cloud or hybrid deployments.
- With true multi-tenancy, securely isolate and share your data to fully meet your business needs, mitigating risk and achieving faster time to value. Security is bifurcated from topology, so you can connect anywhere in a deployment and NATS will do the right thing.
- With the ability to process millions of messages a second per server, you’ll find unparalleled efficiency with NATS. Save money by minimizing cloud costs with reduced compute and network usage for streams, services, and eventing.
- NATS self-heals and can scale up, down, or handle topology changes anytime with zero downtime to your system. Clients require zero awareness of NATS topology allowing you future proof your system to meet your needs of today and tomorrow.


Some use cases of NATS include:

- Cloud Messaging
 - Services (microservices, service mesh)
 - Event/Data Streaming (observability, analytics, ML/AI)
- Command and Control
 - IoT and Edge
 - Telemetry / Sensor Data / Command and Control
- Augmenting or Replacing Legacy Messaging Systems

## Features

NATS offers the following features:

- **Language and Platform Coverage**: Core NATS: 48 known client types, 11 supported by maintainers, 18 contributed by the community. NATS Streaming: 7 client types supported by maintainers, 4 contributed by the community. NATS servers can be compiled on architectures supported by Golang. NATS provides binary distributions.
- **Built-in Patterns**: Streams and Services through built-in publish/subscribe, request/reply, and load-balanced queue subscriber patterns. Dynamic request permissioning and request subject obfuscation is supported.
- **Delivery Guarantees**: At most once, at least once, and exactly once is available in JetStream.
- **Multi-tenancy and Sharing**: NATS supports true multi-tenancy and decentralized security through accounts and defining shared streams and services.
- **AuthN**: NATS supports TLS, NATS credentials, NKEYS (NATS ED25519 keys), username and password, or simple token.
- **AuthZ**: Account limits including number of connections, message size, number of imports and exports. User-level publish and subscribe permissions, connection restrictions, CIDR address restrictions, and time of day restrictions.
- **Message Retention and Persistence**: Supports memory, file, and database persistence. Messages can be replayed by time, count, or sequence number, and durable subscriptions are supported. With NATS streaming, scripts can archive old log segments to cold storage.
- **High Availability and Fault Tolerance**:  Core NATS supports full mesh clustering with self-healing features to provide high availability to clients. NATS streaming has warm failover backup servers with two modes (FT and full clustering). JetStream supports horizontal scalability with built-in mirroring.
- **Deployment**: The NATS network element (server) is a small static binary that can be deployed anywhere from large instances in the cloud to resource constrained devices like a Raspberry PI. NATS supports the Adaptive Edge architecture which allows for large, flexible deployments. Single servers, leaf nodes, clusters, and superclusters (cluster of clusters) can be combined in any fashion for an extremely flexible deployment amenable to cloud, on-premise, edge and IoT. Clients are unaware of topology and can connect to any NATS server in a deployment.
- **Monitoring**: NATS supports exporting monitoring data to Prometheus and has Grafana dashboards to monitor and configure alerts. There are also development monitoring tools such as nats-top. Robust side car deployment or a simple connect-and-view model with NATS surveyor is supported.
- **Management**: NATS separates operations from security. User and Account management in a deployment may be decentralized and managed through a CLI. Server (network element) configuration is separated from security with a command line and configuration file which can be reloaded with changes at runtime.
- **Integrations**: NATS supports WebSockets, a Kafka bridge, an IBM MQ Bridge, a Redis Connector, Apache Spark, Apache Flink, CoreOS, Elastic, Elasticsearch, Prometheus, Telegraf, Logrus, Fluent Bit, Fluentd, OpenFAAS, HTTP, and MQTT, and more.

### Pricing

There are no fees involved with deploying NATS, however, the costs of the instances running the system and related maintenance (and related time-cost) must be taken into account. The final cost depends on the number and type of instances chosen to run NATS.

### Summary of NATS

NATS is a CNCF-regnonized message broker.
Overall, the pros and cons of the tool can be summarized with the following points:

**Main advantages**

- It supports more patterns. Streams and Services through built-in publish/subscribe, request/reply, and load-balanced queue subscriber patterns. Dynamic request permissioning and request subject obfuscation is supported.

**Main disadvantages**

- User-managed. While NATS can be deployed as a Google Cloud Marketplace solution, more complex scenarios like multi-regional clusters require an extensive amount of user-supplied configuration, both for NATS itself and related resources (for example, firewall rules). Using the Helm-charts provided by NATS to run it on Kubernetes however, facilitates many aspects of the process (see https://docs.nats.io/nats-on-kubernetes/nats-kubernetes)


Refer to https://docs.nats.io/ for further information.

### Conclusion

Cloud Pub/Sub and NATS are both excellent, battle-tested message brokers. Whether you pick one or the other, it's often up to your requirements and preferences. Personally, I would always recommend Cloud Pub/Sub where the requirements allow for it, because of a high degree of integration with other Google Cloud products and because, being managed by Google, Cloud Pub/Sub frees engineers from the complex and time consuming process of setting up and maintaining a third-party solution.
