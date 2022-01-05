+++
date = "2022-01-05T09:03:15+08:00"
title = "Native Changelog Deb Prepare API"
description = """
To ensure all `.deb` changelog are aligned consistently with the repository,
Monteur provides a recipe to systematically update all `.deb` changelog across
all `.deb` package variants. This recipe allows Monteur to execute in a highly
customizable manner easily and seamlessly.
"""
keywords = [
	"Native deb Changelog",
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
url = "/en/ci-jobs/prepare/changelog-deb/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Prepare with deb Changelog"

[thumbnails.1]
url = "/en/ci-jobs/prepare/changelog-deb/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Prepare with deb Changelog"

[thumbnails.2]
url = "/en/ci-jobs/prepare/changelog-deb/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Prepare with deb Changelog"


[menu.main]
parent = "Native - Package DEB"
name = "Prepare (DEB Changelog)"
pre = "🧶"
weight = 65
identifier = "ci-jobs-monteur-prepare-changelog-deb"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to update all the `.deb`
packages' persistent changelog file so that they are aligned to the repository
logs in a consistent manner with minimal to no further instructions**.

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Prepare API recipe for natively `.deb` changelog update
work. Please read through your selected version's details on what has changed,
what is required, and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 1.0.0
Version 1.0.0 `changelog-deb` Prepare API is available for download here:
{{< link "/ci-jobs/prepare/changelog-deb/changelog-deb-v1p0p0.toml"
	"this" "" "" "button" "" "download" >}}
changelog-deb-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.2`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/prepare/jobs/`
   directory with the name `changelog-deb.toml`.
2. Once done, edit the configuration file for:
   1. `ChangelogFrom` - the future branch for git log differentiation.
   2. `ChangelogTo` - the target branch to be marged into for git log
      differenciation.
   3. `[Packages.XXX]` - list all deb packages.
        1. `Packages.XXX.OS` - change package operating system.
        2. `Packages.XXX.Arch` - change package CPU architecture.
        3. `Packages.XXX.Distribution` - change package distro.
   4. `Packages.XXX.Changelog` - persistent changelog filepath.

For detailed information about each fields, visit:
[Prepare Specification Data Structure]({{< link
"/ci-jobs/prepare/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.




## Epilogue
That's all for Monteur updating `.deb` changelog with its native functions. If
you found a bug or have any questions about the recipe, please feel free to
raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
