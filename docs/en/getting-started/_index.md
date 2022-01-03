+++
date = "2021-12-24T00:11:52+08:00"
title = "Getting Started"
description = """
Monteur can be quite intimidating to use if left alone in the dark. Hence, we
put up a getting started guide for you to begin with. This section guides you
on how to setup Monteur and using it for your repository.
"""
keywords = [
	"releases",
	"monteur",
	"configurations",
	"documentation",
]
draft = false
type = ""
# redirectURL=""
layout = "single"


[robots]
[robots.googleBot]
name = "googleBot"
content = ""


[modules]
extensions = [
	# Example: "sidebar",
]


[creators.holloway]
"@type" = "Person"
"name" = "'Holloway' Chew Kean Ho"


[thumbnails.0]
url = "/en/getting-started/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Getting Started with Monteur"

[thumbnails.1]
url = "/en/getting-started/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Getting Started with Monteur"

[thumbnails.2]
url = "/en/getting-started/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Getting Started with Monteur"


[menu.main]
parent = ""
name = "Getting Started"
pre = "🛫"
weight = 5
identifier = "getting-started"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## 1. Install Monteur
For starters, regardless on whether you're here to create a repository with
Montuer supports OR being instructed to setup Monteur for an existing supported
one, you need to install Monteur from our official download portals here
(new windows/tab shall open):

{{< link "/versions" "this" "" "_blank" "button" >}}
Download Monteur
{{< /link >}}




## 2. I'm Here To...
Depending on what you plan to do, this guide shall points you to the right
steps:



### A-Route) Work on Existing Repository Already with Monteur
That's easy, what you probably need to do is setup the localized repository
with the following command:

```bash {linenos=table,hl_lines=[],linenostart=1}
monteur setup
```

Once done, depending on your operating system (example here is in
Unix/Linux/MacOS operating system), you can source the localized
filesystem by:

```bash {linenos=table,hl_lines=[],linenostart=1}
source .monteurFS/config/main
```

and you're good to go!



### B-Route) Initialize My Repository with Monteur
To initialize your repository with Monteur, simply head over to our Monteur's
Init API section and follows the guide there properly:

{{< link "/internals/init-api" "this" "" "" "button" >}}
Monteur Init API
{{< /link >}}




## Epilogue
That's all for getting started with Monteur. If you have any queries, please
proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
