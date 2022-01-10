+++
date = "2022-01-10T10:30:19+08:00"
title = "Version v0.0.3"
description = """
Monteur released its version `v0.0.3` with a bunch of new changes, notable
releases and supports, and upgrades. This section denotes all the important
notices related to this release.
"""
keywords = [
	"version",
	"v0.0.3",
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
url = "/en/versions/v0-0-3/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Version v0.0.3"

[thumbnails.1]
url = "/en/versions/v0-0-3/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Version v0.0.3"

[thumbnails.2]
url = "/en/versions/v0-0-3/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Version v0.0.3"


[menu.main]
parent = "Versions"
name = "v0.0.3"
pre = "🏅"
weight = 5
identifier = "versions-v0.0.3"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

This release primarily interate major feature improvements over the first
prototype to standardize its internal operations and external interactions.
There are non-backward compatible changes so please go through this document
clearly for what has changed.




## Downloadable Contents
Here are some of the downloadable archived pack available for manual deployment.

{{< release "v0-0-3" />}}




## Supported Variants
Here are the list of supported Monteur variants available for different
platforms and CPU architectures.

* `linux-amd64` - `.tar.gz`, `.deb`
* `linux-arm64` - `.tar.gz`, `.deb`
* `linux-arm` (v7) - `.tar.gz`, `.deb`
* `darwin-amd64` - `.tar.gz`
* `darwin-arm64` - `.tar.gz`
* `windows-amd64` - `.zip`
* `windows-arm64` - `.zip`




## Notable Changes
Here are some of the notable changes for this version.



### Backwards Compatible
Among the backward compatible changes are:

1. Fixed a critical Setup API bug where it failed to create MonteurFS main
   directories - [Issue 44](https://gitlab.com/zoralab/monteur/-/issues/44).
2. Add Version API for monteur app to easily obtain its version number -
   [Issue 43](https://gitlab.com/zoralab/monteur/-/issues/43).
3. Fixed monteur operations in legacy Arm environment -
   [Issue 39](https://gitlab.com/zoralab/monteur/-/issues/39).



### Non-Backwards Compatible
Among the non-backward compatible changes are:

**None**.




## What's Next
Here are the key development for the next iterations:

1. Native `.deb` package library support.
2. Native `Homebrew` support.




## Epilogue
That's all for this release. If you have any question, please feel free to
raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
