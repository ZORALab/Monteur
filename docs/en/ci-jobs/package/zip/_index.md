+++
date = "2022-01-05T17:32:34+08:00"
title = "Zip Package API"
description = """
Monteur natively supports `Zip` for packaging `.zip` packages across various
app variants that are compliant to existing programs. This recipe allows the
repository to seamlessly package `.zip` application in a highly customizable
manner easily and seamlessly.
"""
keywords = [
	"Zip",
	"Package CI Job",
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
url = "/en/ci-jobs/package/zip/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Package with Zip"

[thumbnails.1]
url = "/en/ci-jobs/package/zip/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Package with Zip"

[thumbnails.2]
url = "/en/ci-jobs/package/zip/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Package with Zip"


[menu.main]
parent = "Native - Monteur"
name = "Package (Zip)"
pre = "📦"
weight = 80
identifier = "ci-jobs-package-zip"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to make building `.zip`
packages painlessly in a consistent manner with minimal to no further
instructions**.

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Package API recipe for `.zip` integrations. Please read
through your selected version's details on what has changed, what is required,
and how to customize and used them.

The arrangement are the latest at the top or first.




### Version 1.0.0
Version 1.0.0 `Zip` Package API is available for download here:
{{< link "/ci-jobs/package/zip/zip-v1p0p0.toml" "this" "" "" "button"
	"" "download" >}}
zip-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.2`                         |
| Supported Platforms  | follows `Monteur`'s availability |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/package/jobs/` directory with the name `zip.toml`.
2. Once done, edit the configuration file for:
   1. `Packages.XXX` - your `.zip` packages list. Duplicate if there are more
      build variants to be packaged into `.zip` pack.
   2. `Packages.XXX.OS` - list of supported operating system for this package.
   3. `Packages.XXX.Arch` - list of supported CPU architecture for this package.
   4. `Packages.XXX.Name` - filename without extension for this package.
   5. `Packages.XXX.Distribution` - list of supported distributions.
   11. `Packages.XXX.BuildSource` - instruct Monteur to package a source code
       pack or a binary pack (not both). If both are needed, duplicate
       the recipe since they are essentially 2 different packages.
   12. `[Packages.XXX.Files]` - the list of files to be assembled by Monteur
       for packaging.

For detailed information about each fields, visit:
[Package Specification Data Structure]({{< link
"/ci-jobs/package/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.




## Epilogue
That's all for Monteur operating `Zip` packaging. If you found a bug or have
any questions about the recipe, please feel free to raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).