+++
date = "2022-01-05T07:17:30+08:00"
title = "Native Version File Scripting Prepare API"
description = """
To ensure all version files in the repository is consistent with Monteur's app
metadata, Monteur provides a recipe to script the version file so that user
only change version number at one location will do. This recipe allows Monteur
to script version file in a highly customizable manner easily and seamlessly.
"""
keywords = [
	"Native Version File Scripting",
	"Prepare CI Job",
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
url = "/en/ci-jobs/prepare/version-file-scripting/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Prepare with Version File Scripting"

[thumbnails.1]
url = "/en/ci-jobs/prepare/version-file-scripting/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Prepare with Version File Scripting"

[thumbnails.2]
url = "/en/ci-jobs/prepare/version-file-scripting/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Prepare with Version File Scripting"


[menu.main]
parent = "Native - Monteur"
name = "Prepare (Version File Scripting)"
pre = "🧶"
weight = 65
identifier = "ci-jobs-monteur-prepare-version-scripting"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to script version file
so that they are aligned to Monteur's app metadata's Version field, allowing
user to only update Version via Monteur's app metadata painlessly, in a
consistent manner with minimal to no further instructions**.

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Prepare API recipe for natively version file scripting.
Please read through your selected version's details on what has changed, what is
required, and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 1.0.0
Version 1.0.0 `version file scripting` Prepare API is available for download
here:
{{< link "/ci-jobs/prepare/version-file-scripting/version-file-v1p0p0.toml"
	"this" "" "" "button" "" "download" >}}
version-file-1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.2`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/prepare/jobs/`
   directory with the name `version-XXX.toml` where `XXX` is something you
   can differentiate if you have many files. Otherwise, `version-repo.toml` or
   `version-file.toml` is suffice.
2. Once done, edit the configuration file for:
   1. `CodeFilepath` - the filepath of the version file.

For detailed information about each fields, visit:
[Prepare Specification Data Structure]({{< link
"/ci-jobs/prepare/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.




## Epilogue
That's all for Monteur operating directory version file scripting with its
native functions. If you found a bug or have any questions about the recipe,
please feel free to raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
