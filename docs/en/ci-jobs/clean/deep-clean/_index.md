+++
date = "2021-12-23T17:05:21+08:00"
title = "Native Deep Clean API"
description = """
Monteur provides a deep clean recipe for cleaning up a particular directory
without depending on external programs. This recipe allows Monteur to deep clean
the repository in a highly customizable manner easily and seamlessly.
"""
keywords = [
	"Monteur Deep Clean",
	"Clean CI Job",
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
url = "/en/ci-jobs/clean/deep-clean/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Clean with Native Deep Clean"

[thumbnails.1]
url = "/en/ci-jobs/clean/deep-clean/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Clean with Native Deep Clean"

[thumbnails.2]
url = "/en/ci-jobs/clean/deep-clean/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Clean with Native Deep Clean"


[menu.main]
parent = "Native - Monteur"
name = "Clean (Deep Clean)"
pre = "🧹"
weight = 20
identifier = "ci-jobs-monteur-deep-clean"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to deep clean a given
directory path in the repository painlessly in a consistent manner with minimal
to no further instructions**.

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Clean API recipe for natively deep-cleaning. Please read
through your selected version's details on what has changed, what is required,
and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 1.0.0
Version 1.0.0 `deep clean` Clean API is available for download here:
{{< link "/ci-jobs/test/go/go-1p0p0.toml" "this" "" "" "button"
	"" "download" >}}
go-1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.1`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/clean/cleaners/`
   directory with the name pattern `purpose.toml` like `build-log.toml`.
2. Once done, edit the configuration file for:
   1. `FMTVariables.Path` - the directory path that you want to deep clean. You
      can make use of `Variables` to creatively construct the `Path` value in
      order to scale across multiple directories.

For detailed information about each fields, visit:
[Clean Specification Data Structure]({{< link
"/ci-jobs/clean/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.




## Epilogue
That's all for Monteur operating directory deep cleaning with its native
functions. If you found a bug or have any questions about the recipe, please
feel free to raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
