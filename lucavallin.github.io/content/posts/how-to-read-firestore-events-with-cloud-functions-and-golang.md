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
date: "2021-07-27"
title: "How to Read Firestore Events with Cloud Functions and Golang"
description: "So you want to know 'How to Read Firestore Events with Cloud Functions and Golang' ? You're in the right place! I recently worked on a side project called 'Syn' (https://github.com/lucavallin/syn - Old Norse for vision!) which aims at visually monitoring the environment using the Raspberry Pi, Google Cloud and React Native."
canonicalURL: "https://binx.io/blog/2021/07/27/how-to-read-firestore-events-with-cloud-functions-and-golang/"
tags: ["google cloud", "firestore", "raspberry pi", "iot", "golang", "react"]
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
The Raspberry Pi uses a tool called `motion`, which takes pictures (with the Pi camera) when movement is detected. The pictures are then uploaded to Cloud Storage, and a Cloud Function listens for events that are triggered when a new object is uploaded to the bucket. When a new picture is uploaded, the function tries to label the image using Vision API and stores the result in Firestore. Firestore triggers an event itself, which I am listening for in a different Cloud Function, which uses IFTTT to notify me when movement has been detected and what Vision API has found in the image. The goal of this blog post is to explain how to parse Firestore events, which are delivered to the function in a rather confusing format!

### Shape and colour of the Event struct
The Cloud Function processing new events (= new uploads from the Raspberry Pi, when movement is detected), after labelling the picture with Vision API, creates a new Event object and stores it in the `Events` collection in Firestore.
The Event struct looks like this:
```go
type Event struct {
	URI     string    `json:"uri" firestore:"uri"`
	Created time.Time `json:"created" firestore:"created"`
	Labels  []Label   `json:"labels" firestore:"labels"`
}
```

The struct also references an array of Labels. Label is a struct defined as:

```go
type Label struct {
	Description string  `json:"description" firestore:"description"`
	Score       float32 `json:"score" firestore:"score"`
}
```

This is the result once the information has been persisted to Firestore:

<img src="https://binx.io/wp-content/uploads/2021/07/Screenshot-2021-07-27-at-10.22.55-900x494.png" alt="Syn - Firestore" width="900" height="494" class="size-medium wp-image-3722" />

### Create a function to listen to Firestore events
Another function, called `Notify`, listens for events from Firestore (and then notifies the user via IFTTT), which are triggered when new data is added into the database. I have used Terraform to setup the function:

```hcl
resource "google_cloudfunctions_function" "notify" {
  project               = data.google_project.this.project_id
  region                = "europe-west1"
  name                  = "Notify"
  description           = "Notifies of newly labeled uploads"
  service_account_email = google_service_account.functions.email
  runtime               = "go113"
  ingress_settings      = "ALLOW_INTERNAL_ONLY"
  available_memory_mb   = 128

  entry_point = "Notify"

  source_repository {
    url = "https://source.developers.google.com/projects/${data.google_project.this.project_id}/repos/syn/moveable-aliases/master/paths/functions"
  }

  event_trigger {
    event_type = "providers/cloud.firestore/eventTypes/document.create"
    resource   = "Events/{ids}"
  }

  environment_variables = {
    "IFTTT_WEBHOOK_URL" : var.ifttt_webhook_url
  }
}
```

The `event_trigger` block defines the event that the function should listen for. In this case, I am listening for `providers/cloud.firestore/eventTypes/document.create` events in the `Events` collection.

### What does a "raw" Firestore event look like?
Using `fmt.Printf("%+v", event)`, we can see that the Firestore event object looks like this:
```json
{OldValue:{CreateTime:0001-01-01 00:00:00 +0000 UTC Fields:{Created:{TimestampValue:0001-01-01 00:00:00 +0000 UTC} File:{MapValue:{Fields:{Bucket:{StringValue:} Name:{StringValue:}}}} Labels:{ArrayValue:{Values:[]}}} Name: UpdateTime:0001-01-01 00:00:00 +0000 UTC} Value:{CreateTime:2021-07-27 09:22:03.654255 +0000 UTC Fields:{Created:{TimestampValue:2021-07-27 09:22:01.4 +0000 UTC} File:{MapValue:{Fields:{Bucket:{StringValue:} Name:{StringValue:}}}} Labels:{ArrayValue:{Values:[{MapValue:{Fields:{Description:{StringValue:cat} Score:{DoubleValue:0.8764283061027527}}}} {MapValue:{Fields:{Description:{StringValue:carnivore} Score:{DoubleValue:0.8687784671783447}}}} {MapValue:{Fields:{Description:{StringValue:asphalt} Score:{DoubleValue:0.8434737920761108}}}} {MapValue:{Fields:{Description:{StringValue:felidae} Score:{DoubleValue:0.8221824765205383}}}} {MapValue:{Fields:{Description:{StringValue:road surface} Score:{DoubleValue:0.807261049747467}}}}]}}} Name:projects/cvln-syn/databases/(default)/documents/Events/tVhYbIZBQypHtHzDUabq UpdateTime:2021-07-27 09:22:03.654255 +0000 UTC} UpdateMask:{FieldPaths:[]}}
```
...which is extremely confusing! I was expecting the event to look exactly like the data I previously stored in the database, but for some reason, this is what a Firebase event looks like. Luckily, the `"JSON to Go Struct"` IntelliJ IDEA plugin helps making sense of the above:

```go
type FirestoreUpload struct {
	Created struct {
		TimestampValue time.Time `json:"timestampValue"`
	} `json:"created"`
	File struct {
		MapValue struct {
			Fields struct {
				Bucket struct {
					StringValue string `json:"stringValue"`
				} `json:"bucket"`
				Name struct {
					StringValue string `json:"stringValue"`
				} `json:"name"`
			} `json:"fields"`
		} `json:"mapValue"`
	} `json:"file"`
	Labels struct {
		ArrayValue struct {
			Values []struct {
				MapValue struct {
					Fields struct {
						Description struct {
							StringValue string `json:"stringValue"`
						} `json:"description"`
						Score struct {
							DoubleValue float64 `json:"doubleValue"`
						} `json:"score"`
					} `json:"fields"`
				} `json:"mapValue"`
			} `json:"values"`
		} `json:"arrayValue"`
	} `json:"labels"`
}
```

While still confusing, at least now I can split up the struct so I can reference the types correctly elsewhere in the application should I need to, for example, loop through the labels.

### Cleaning up the event structure

The `FirestoreUpload` can be split up in order to have named fields rather than anonymous structs. This is useful to be able to reference the correct fields and types elsewhere in the application, for example when looping through the labels.

```go
package events

import (
	"github.com/thoas/go-funk"
	"time"
)

//FirestoreEvent is the payload of a Firestore event
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields
type FirestoreValue struct {
	CreateTime time.Time       `json:"createTime"`
	Fields     FirestoreUpload `json:"fields"`
	Name       string          `json:"name"`
	UpdateTime time.Time       `json:"updateTime"`
}

// FirestoreUpload represents a Firebase event of a new record in the Upload collection
type FirestoreUpload struct {
	Created Created `json:"created"`
	File    File    `json:"file"`
	Labels  Labels  `json:"labels"`
}

type Created struct {
	TimestampValue time.Time `json:"timestampValue"`
}

type File struct {
	MapValue FileMapValue `json:"mapValue"`
}

type FileMapValue struct {
	Fields FileFields `json:"fields"`
}

type FileFields struct {
	Bucket StringValue `json:"bucket"`
	Name   StringValue `json:"name"`
}

type Labels struct {
	ArrayValue LabelArrayValue `json:"arrayValue"`
}

type LabelArrayValue struct {
	Values []LabelValues `json:"values"`
}

type LabelValues struct {
	MapValue LabelsMapValue `json:"mapValue"`
}

type LabelsMapValue struct {
	Fields LabelFields `json:"fields"`
}

type LabelFields struct {
	Description StringValue `json:"description"`
	Score       DoubleValue `json:"score"`
}

type StringValue struct {
	StringValue string `json:"stringValue"`
}

type DoubleValue struct {
	DoubleValue float64 `json:"doubleValue"`
}

// GetUploadLabels returns the labels of the image as an array of strings
func (e FirestoreEvent) GetUploadLabels() []string {
	return funk.Map(e.Value.Fields.Labels.ArrayValue.Values, func(l LabelValues) string {
		return l.MapValue.Fields.Description.StringValue
	}).([]string)
}
```

The `GetUploadLabels()` function is an example of how the `FirestoreUpload` event object should be accessed. Here I am also using the [go-funk](https://github.com/thoas/go-funk "go-funk") package, which adds some extra functional capabilities to Go (but the performance isn't as good as a "native" loop).

### Summary

In this article I explained how to read Firestore events from Cloud Functions listening for them. The examples are written in Golang, but different languages will need to parse the messages in a similar way. Although not handy, this is the current format of Firestore events! Luckily, once you know how to read them, the rest is simple!
