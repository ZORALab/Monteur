+++
date = "2021-12-24T00:51:19+08:00"
title = "Version v0.0.2"
description = """
Monteur released its version `v0.0.2` with a bunch of new changes, notable
releases and supports, and upgrades. This section denotes all the important
notices related to this release.
"""
keywords = [
	"version",
	"v0.0.2",
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
url = "/en/versions/v0-0-2/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Version v0.0.2"

[thumbnails.1]
url = "/en/versions/v0-0-2/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Version v0.0.2"

[thumbnails.2]
url = "/en/versions/v0-0-2/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Version v0.0.2"


[menu.main]
parent = "Versions"
name = "v0.0.2"
pre = "🏅"
weight = 5
identifier = "versions-v0.0.2"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

This release primarily interate major feature improvements over the first
prototype to standardize its internal operations and external interactions.
There are non-backward compatible changes so please go through this document
clearly for what has changed.

{{< note warning "Heads Up" >}}
There is a bug where the Setup API failed to create the main MonteurFS
directories (See [Issue 44](https://gitlab.com/zoralab/monteur/-/issues/44)).

Please upgrade to Monteur `v0.0.3` or manually create those main directories
for Monteur to work.
{{< /note >}}



## Downloadable Contents
Here are some of the downloadable archived pack available for manual deployment.

{{< release "v0-0-2" />}}




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

1. Added native `TarGz` function -
   [Issue 8](https://gitlab.com/zoralab/monteur/-/issues/8).
2. Added support for `linux/arm` targeting Raspberry Pi 3 and above -
   [Issue 21](https://gitlab.com/zoralab/monteur/-/issues/21).
3. Added support for `darwin/amd64` and `darwin/arm64` targeting MacOS -
   [Issue 23](https://gitlab.com/zoralab/monteur/-/issues/23).
4. Added native archive packaging functions and recipes -
   [Issue 26](https://gitlab.com/zoralab/monteur/-/issues/26).
5. Added Init API for easy repo initialization -
   [Issue 28](https://gitlab.com/zoralab/monteur/-/issues/28).
6. Added Prepare API for repo data preparations -
   [Issue 31](https://gitlab.com/zoralab/monteur/-/issues/31).
7. Added native Arhicver release function for `TarGz` and `Zip` packages -
   [Issue 33](https://gitlab.com/zoralab/monteur/-/issues/33).
8. Fixed Concurrency bug for Package API -
   [Issue 34](https://gitlab.com/zoralab/monteur/-/issues/34).
9. Added secret data support and logger filtration -
   [Issue 35](https://gitlab.com/zoralab/monteur/-/issues/35).
10. Fixed `delete-quiet` command bug to actually delete things -
   [Issue 37](https://gitlab.com/zoralab/monteur/-/issues/37).
11. Added native `sha256`, `sha512`, and `sha512->sha256` checksum functions.



### Non-Backwards Compatible
Among the non-backward compatible changes are:

1. Renaming all the jobs directories (`compose/composers`, `publish/publishers`,
   ...) into a unified name called `jobs` (`compose/jobs`, `publish/jobs`, ...)
   See [Issue 32](https://gitlab.com/zoralab/monteur/-/issues/32) for more info.
   1. What you need to do: **rename all of the said directories to `jobs`**.
2. Shifted Build API's jobs directory from (`app/variants` to `build/jobs`).
   See [Issue 32](https://gitlab.com/zoralab/monteur/-/issues/32) for more info.
   1. What you need to do: **shift and rename the said directory**.




## What's Next
Here are the key development for the next iterations:

1. Native `.deb` package library support.
2. Native `Homebrew` Support.
3. GitLab CI recipe.




## Epilogue
That's all for this release. If you have any question, please feel free to
raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
