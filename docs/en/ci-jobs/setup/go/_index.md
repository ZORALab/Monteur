+++
date = "2021-12-23T17:05:21+08:00"
title = "Go Setup API"
description = """
Monteur provides a Go setup recipe for setting up a localized Go programming
language developement repository. This recipe allows Monteur to setup Go
in a localized filesystem inside the repository with a highly customizable
method, easily and seamlessly.
"""
keywords = [
	"Go",
	"Setup CI Job",
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
url = "/en/ci-jobs/setup/go/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Setup with Go"

[thumbnails.1]
url = "/en/ci-jobs/setup/go/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Setup with Go"

[thumbnails.2]
url = "/en/ci-jobs/setup/go/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Setup with Go"


[menu.main]
parent = "Software - Go"
name = "Setup"
pre = "🧩"
weight = 10
identifier = "ci-jobs-setup-go"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to setup Go locally in a
highly customizable manner (e.g. different version) into the repository
painlessly in a consistent manner with minimal to no further instructions**.

Go is available at: https://go.dev/

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Setup API recipe for Go. Please read through your
selected version's details on what has changed, what is required, and how to
customize and used them.

The arrangement are the latest at the top or first.



### Version 1.1.0
Version 1.1.0 `go` Setup API is available for download here:
{{< link "/ci-jobs/setup/go/go-1p1p0.toml" "this" "" "" "button"
	"" "download" >}}
go-1p1p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.1`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/setup/programs/`
   directory with the name `go.toml`.
2. That's all. Unless Go releases a new version, you will need to update:
   1. `Variables.Version` - the new version number.
   2. `Variables.BaseURL` - the base URL changes if Go development team decided
      to change again.
   2. `Sources.XXX.Checksum.Value` - update **each** checksum values to match
      their respective package checksum values.

For detailed information about each fields, visit:
[Setup Specification Data Structure]({{< link
"/ci-jobs/setup/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Updated to Go version `1.17.5`.
2. *Backward Compatible* - Fixed freebsd `amd64` typo where the URL uses `386`.



### Version 1.0.0
Version 1.0.0 `go` Setup API is available for download here:
{{< link "/ci-jobs/setup/go/go-1p0p0.toml" "this" "" "" "button"
	"" "download" >}}
go-1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.1`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/setup/programs/`
   directory with the name `go.toml`.
2. That's all. Unless Go releases a new version, you will need to update:
   1. `Variables.Version` - the new version number.
   2. `Variables.BaseURL` - the base URL changes if Go development team decided
      to change again.
   2. `Sources.XXX.Checksum.Value` - update **each** checksum values to match
      their respective package checksum values.

For detailed information about each fields, visit:
[Setup Specification Data Structure]({{< link
"/ci-jobs/setup/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.
3. *Backward Compatible* - Updated to Go version `1.17.3`.



## Epilogue
That's all for Monteur setting up Go Programming Language into the repository.
If you found a bug or have any questions about the recipe, please
feel free to raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
