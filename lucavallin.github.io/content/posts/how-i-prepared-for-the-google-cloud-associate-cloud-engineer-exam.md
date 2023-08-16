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
date: "2021-02-05"
title: "How I prepared for the Google Cloud Associate Cloud Engineer Exam"
description: "I recently passed the Google Cloud Associate Cloud Engineer certification exam. This post describes what the exam is about, how I prepared (with links to useful resources), and the strategies I used to answer the questions in under two minutes per question (skip down to Answering Strategies if that’s just what you want to know)."
canonicalURL: "https://binx.io/blog/2021/02/05/how-i-prepared-for-the-google-cloud-associate-cloud-engineer-exam/"
tags: ["google cloud", "certification", "cloud engineer"]
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

First, some background if you don’t know what this is about, and are trying to figure out if you should take the certification. The [Associate Cloud Engineer certification](https://cloud.google.com/certification/cloud-engineer "Associate Cloud Engineer certification ") is one of the most technical Google Cloud certifications: I think it is even harder than the Professional Cloud Architect exam I'm currently preparing for. Nevertheless, this is the first stepping stone into Google Cloud, making it an obvious choice to get started if you are a software engineer, like I am.

Google Cloud defines no prerequisites for this exam, but recommends a minimum of six months hands-on experience on the platform. The topics covered in the certification are:

- Projects, Billing, IAM
- Compute Engine, App Engine and Kubernetes Engine
- VPC and Networking
- Cloud Storage
- Databases (included are for example Cloud SQL, Memorystore, and Firestore)
- Operations Suite (Cloud Logging and other services). You might know this as Stackdriver.

The next step in the learning journey after the Associate Cloud Engineer exam is the [Professional Cloud Architect](https://cloud.google.com/certification/cloud-architect "Professional Cloud Architect") certification. From there, you can choose to specialize in Data and Machine Learning Engineering, Cloud Development and DevOps, or Security and Networking. (Check out the [certification paths here](https://cloud.google.com/certification#certification_paths "certification paths here")).

### Why Did I Choose Google Cloud?
Google Cloud strikes me as the most developer-friendly cloud provider. At the same time, it is also the smallest of the three main players but it is growing fast, meaning lots of room for new experts. I like the certification program  because it is role-based (instead of product-centered), preparing not just for this provider, but the position itself.

### About Me
I want to tell you a little bit about me, so you know how to interpret my advice. I have been a Software Engineer for the past 8+ years, doing mostly full-stack development. The ability to work on what “powers” the internet and an interest for what’s under the hood of software systems brought me to the cloud. Furthermore, the managed services provided by cloud providers make it easy to add all sorts of features into applications, from sending emails to advanced machine learning - these days nothing is stopping you from starting your next company from the comfort of your bedroom.

I grew up in Italy and moved to the Netherlands a few years ago, where I now live with my wife and cat. When I am not working on clouds and other atmospheric phenomena, I enjoy photography and cycling.

Feel free to connect with me on [LinkedIn](https://linkedin.com/in/lucavallin "LinkedIn")!

### What Does the Exam Look Like?
This is how a certification exam works: you have two hours to answer 50 multiple choice and multiple selection questions. That means you’ll have roughly two minutes per question, with 20 minutes left at the end to review your answers.

The exam can be taken remotely or in person, costs 125 USD (roughly 100 EUR), and can be retaken after a cool down period if you fail. The certification is valid for two years (you’ll have to take the full exam again to recertify.

### Exam Guide
Google publishes an [exam guide](https://cloud.google.com/certification/guides/cloud-engineer "exam guide") for every certification. It details the topics of the certification, providing details of what you need to know for each of the areas.

I used the exam guide to track my progress while studying. You can make a copy of it and use color coding to highlight the topic you feel confident about and those that need more of your attention.

### Learning resources
There are many books and online courses you can take to prepare for the exam, some more effective than others. For most people, a combination of written and audio-visual material works best. These are the resources that helped me best (and why):

- Official Google Cloud Certified Associate Cloud Engineer Study Guide ([book](https://www.wiley.com/en-us/Official+Google+Cloud+Certified+Associate+Cloud+Engineer+Study+Guide-p-9781119564416 "book"))
	- Most comprehensive resource
	- Questions after each chapter
	- Includes practice tests and flashcards
- ACloudGuru ([online course](https://acloudguru.com/course/google-certified-associate-cloud-engineer "online course"))
	- This is a good introduction (start with this course if you have no Google Cloud experience)
	- You also need the “[Kubernetes Deep Dive](https://acloudguru.com/course/kubernetes-deep-dive "Kubernetes Deep Dive")” course
	- Hands-on labs
	- This course does not cover everything you need to know

### Practice Tests
Practice tests are a key part of your preparation because they let you test your knowledge in a setting similar to the actual exam. If you finish the test, they provide you with a detailed explanation for each question, documenting the correct and wrong answers.

Take note of the topics that require attention and review the documentation accordingly. Once you consistently score at least 90% on practice tests, you are ready for the exam.

The Official Study Guide book provides review questions at the end of each chapter and an online portal with two practice tests. The ACloudGuru course also has a practice exam you can take, and you can find similar resources on Udemy.

### Answering Strategies
If you do the math, you’ll see that you only have two minutes to answer every question, which is not much, given that each of them is quite lengthy. Here are the strategies I used when answering the questions.

- Identify the core question
- Review it carefully since a single word can make the difference
- Eliminate answers that are wrong or in conflict with the question
- Choose the cheapest and most secure option
- Read the question again, keeping in mind the answer you chose

You can also mark answers during the test to come back to them later. While you don't have much time left for review at the end, this practice will save you from over thinking and losing valuable time.

### Taking the Remote Proctored Exam
If you decide to take the exam remotely, once you have arranged it, you will have to install a Sentinel tool provided by the testing authority, verify your identity and pass a number of checks.

**TIP**: The operator taking you through the process doesn’t talk, so you need to scroll the chat for their questions. For me, the chat didn’t auto-scroll, so it was a bit awkward at first.

### Summary
In this post, I shared my experience preparing and taking the Google Cloud Associate Engineer Certification exam. Here are the four most important things to take away from this:

- You can do this! If you focus your attention and take as many practice tests as possible, and carefully review your correct and incorrect answers, you’ll pass the exam.
- I think the Official Study Guide is hands down the best resource to use for preparation.
- It is very useful to have real-world experience and practice with gcloud CLI and Kubernetes.
- During the exam, attention to detail is important. Read every question carefully and read the question again after choosing your answer.
