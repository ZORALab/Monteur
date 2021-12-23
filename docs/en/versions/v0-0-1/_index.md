+++
date = "2021-12-24T00:51:19+08:00"
title = "Version v0.0.1"
description = """
Monteur released its version `v0.0.1` with a bunch of new changes, notable
releases and supports, and upgrades. This section denotes all the important
notices related to this release.
"""
keywords = [
	"version",
	"v0.0.1",
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
url = "/en/versions/v0-0-1/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Version v0.0.1"

[thumbnails.1]
url = "/en/versions/v0-0-1/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Version v0.0.1"

[thumbnails.2]
url = "/en/versions/v0-0-1/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Version v0.0.1"


[menu.main]
parent = "Versions"
name = "v0.0.1"
pre = "🏅"
weight = 5
identifier = "versions-v0.0.1"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

This release primarily focuses on prototyping Monteur fulfilling its objective.
It is the proof-of-concept of the technology.




## Supported Variants
Here are the list of supported Monteur variants available for different
platforms and CPU architectures.

* `linux-amd64` - `.deb`




## Notable Changes
Here are some of the notable changes for this version.



### Backwards Compatible
Among the backward compatible changes are:

1. Added Setup API - [Issue 1](https://gitlab.com/zoralab/monteur/-/issues/1).
2. Added Build API - [Issue 2](https://gitlab.com/zoralab/monteur/-/issues/2).
3. Added Release API - [Issue 3](https://gitlab.com/zoralab/monteur/-/issues/3).
4. Added Package API - [Issue 4](https://gitlab.com/zoralab/monteur/-/issues/4).
5. Added Test API - [Issue 6](https://gitlab.com/zoralab/monteur/-/issues/6).
6. Added Clean API - [Issue 7](https://gitlab.com/zoralab/monteur/-/issues/7).
7. Added Custom Headers for HTTPS download - [Issue 8](https://gitlab.com/zoralab/monteur/-/issues/9).
8. Added Secrets data Parsing - [Issue 10](https://gitlab.com/zoralab/monteur/-/issues/10).
9. Added Publish API - [Issue 11](https://gitlab.com/zoralab/monteur/-/issues/11).
10. Added Compose API - [Issue 17](https://gitlab.com/zoralab/monteur/-/issues/17).



### Non-Backwards Compatible
No non-backwards compatible changes found in this version.




## What's Next
Here are the key development for the next iterations:

1. Support `darwin-amd64` and `darwin-arm64` for Monteur.
2. Native `.deb` package library support.
3. Native `.zip` package library support.
4. Native `.tar.gz` package library support.
5. Native `sha256` file checksum support.




## Epilogue
That's all for this release. If you have any question, please feel free to
raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
