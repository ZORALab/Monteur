+++
date = "2021-12-23T11:10:57+08:00"
title = "Go Test API"
description = """
Monteur strongly supports `go` for testing itself and other go applications.
This recipe allows repository equipped with go to seamlessly test go programs
in a highly customizable manner easily and seamlessly.
"""
keywords = [
	"Go",
	"Test CI Job",
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
url = "/en/ci-jobs/test/go/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Test with Go"

[thumbnails.1]
url = "/en/ci-jobs/test/go/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Test with Go"

[thumbnails.2]
url = "/en/ci-jobs/test/go/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Test with Go"


[menu.main]
parent = "Software - Go"
name = "Test"
pre = "⚗️"
weight = 60
identifier = "ci-jobs-test-go"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to execute full-scale
go program testing painlessly in a consistent manner with minimal to no further
instructions**.

Go is available at:
https://go.dev/

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Test API recipe for `go` integrations. Please read
through your selected version's details on what has changed, what is required,
and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 1.0.0
Version 1.0.0 `go` Test API is available for download here:
{{< link "/ci-jobs/test/go/go-v1p0p0.toml" "this" "" "" "button"
	"" "download" >}}
go-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.1`                         |
| Supported Platforms  | follows `go`'s availability      |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/test/jobs/`
   directory with the name `go.toml`.
2. Once done, edit the configuration file for:
   1. `Variables.Timeout` - alter yout timeout timing. Default is `14400s`.

For detailed information about each fields, visit:
[Test Specification Data Structure]({{< link
"/ci-jobs/test/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.
3. *Backward Compatible* - Generate both terminal base test coverage percentile
   and test coverage mapping.




## Epilogue
That's all for Monteur operating `go` software for testing `go` source codes. If
you found a bug or have any questions about the recipe, please feel free to
raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
