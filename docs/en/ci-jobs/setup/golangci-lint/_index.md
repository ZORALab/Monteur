+++
date = "2021-12-23T20:37:28+08:00"
title = "Golang-CI Linter Setup API"
description = """
Monteur provides a Golang-CI Linter setup recipe for setting up a localized
Go Programming Language static analysis linter for repository development. This
recipe allows Monteur to setup GolangCI-Lint locally inside the repository with
a highly customizable method, easily and seamlessly.
"""
keywords = [
	"GolangCI-Lint",
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
url = "/en/ci-jobs/setup/golangci-lint/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Setup with GolangCI-Lint"

[thumbnails.1]
url = "/en/ci-jobs/setup/golangci-lint/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Setup with GolangCI-Lint"

[thumbnails.2]
url = "/en/ci-jobs/setup/golangci-lint/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Setup with GolangCI-Lint"


[menu.main]
parent = "Software - GolangCI-Lint"
name = "Setup"
pre = "🧩"
weight = 10
identifier = "ci-jobs-setup-golangci-lint"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to setup `golangci-lint`
locally in a highly customizable manner (e.g. different version) consistently
with minimal to no further instructions**.

`golangci-lint` is available at: https://golangci-lint.run/

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Setup API recipe for `golangci-lint`. Please read through
your selected version's details on what has changed, what is required, and how
to customize and used them.

The arrangement are the latest at the top or first.



### Version 3.0.0
Version 3.0.0 `golangci-lint` Setup API is available for download here:
{{< link "/ci-jobs/setup/golangci-lint/golangci-lint-v3p0p0.toml" "this" "" ""
	"button" "" "download" >}}
golangci-lint-v3p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.3`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/setup/jobs/`
   directory with the name `golangci-lint.toml`.
2. That's all. Unless `golangci-lint` releases a new version, you will need to
   update:
   1. `Variables.Version` - the new version number.
   2. `Variables.BaseURL` - the base URL changes if Go development team decided
      to change again.
   2. `Sources.XXX.Checksum.Value` - update **each** checksum values to match
      their respective package checksum values (look for a text file in the
      release download portal).


#### Changes
1. **Non-Backward Compatible** - Refactored for friendly to `arm` development
   environment like using `.SourceSystem` instead of `.ComputeSystem` in `[CMD]`
   list.



### Version 2.0.0
Version 2.0.0 `golangci-lint` Setup API is available for download here:
{{< link "/ci-jobs/setup/golangci-lint/golangci-lint-v2p0p0.toml" "this" "" ""
	"button" "" "download" >}}
golangci-lint-v2p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.2`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/setup/jobs/`
   directory with the name `golangci-lint.toml`.
2. That's all. Unless `golangci-lint` releases a new version, you will need to
   update:
   1. `Variables.Version` - the new version number.
   2. `Variables.BaseURL` - the base URL changes if Go development team decided
      to change again.
   2. `Sources.XXX.Checksum.Value` - update **each** checksum values to match
      their respective package checksum values (look for a text file in the
      release download portal).


#### Changes
1. **Non-Backward Compatible** - Replace `[[Setup]]` into `[[CMD]]`.
2. **Non-Backward Compatible** - Changed to use new Monteur setup alogrithms.
3. *Backward Compatible* - supported continuous download between cancellations.
4. *Backward Compatible* - delete target only when needed (right before copy).
5. *Backward Compatible* - Added HTTPS headers table
   (`[Sources.all-all.headers]`) just in case of a needs.



### Version 1.0.0
Version 1.0.0 `golangci-lint` Setup API is available for download here:
{{< link "/ci-jobs/setup/golangci-lint/golangci-lint-v1p0p0.toml" "this" "" ""
	"button" "" "download" >}}
golangci-lint-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.1`                         |
| Supported Platforms  | native to Monteur                |


#### Installation Instructions
1. You should download and place the recipe into your `<config>/setup/programs/`
   directory with the name `golangci-lint.toml`.
2. That's all. Unless `golangci-lint` releases a new version, you will need to
   update:
   1. `Variables.Version` - the new version number.
   2. `Variables.BaseURL` - the base URL changes if Go development team decided
      to change again.
   2. `Sources.XXX.Checksum.Value` - update **each** checksum values to match
      their respective package checksum values (look for a text file in the
      release download portal).

For detailed information about each fields, visit:
[Setup Specification Data Structure]({{< link
"/ci-jobs/setup/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.
3. *Backward Compatible* - Updated to `golangci-lint` version `1.43.0`.




## Epilogue
That's all for Monteur setting up `golangci-lint` into the repository. If you
found a bug or have any questions about the recipe, please feel free to raise
your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
