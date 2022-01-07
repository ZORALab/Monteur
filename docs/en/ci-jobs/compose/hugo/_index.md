+++
date = "2021-12-22T18:26:41+08:00"
title = "Hugo Compose API"
description = """
Monteur supplies Hugo compose recipe as its default web documentation composing.
This allows repository equipped with Hugo website builder to seamlessly compose
its website artifact in a highly customizable manner easily and seamlessly.
"""
keywords = [
	"hugo",
	"compose CI Job",
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
url = "/en/ci-jobs/compose/hugo/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Compose with Hugo"

[thumbnails.1]
url = "/en/ci-jobs/compose/hugo/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Compose with Hugo"

[thumbnails.2]
url = "/en/ci-jobs/compose/hugo/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Compose with Hugo"


[menu.main]
parent = "Software - Hugo"
name = "Compose"
pre = "📝"
weight = 100
identifier = "ci-jobs-compose-hugo"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to compose Hugo website
artifact in a consistent manner with minimal to no further instructions**.

Hugo software is available at https://gohugo.io/ and you can reach their
communities out at https://discourse.gohugo.io/.

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Compose API recipe for Hugo integrations. Please reach
through your selected version's details on what has changed, what is required,
and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 1.1.0
Version 1.1.0 Hugo Compose API is available for download here:
{{< link "/ci-jobs/compose/hugo/hugo-v1p1p0.toml" "this" "" "" "button" ""
	"download" >}}
hugo-v1p1p0.toml
{{< /link >}}

| Min Requirements     | Values                      |
|:---------------------|----------------------------:|
| Monteur Version      | `v0.0.1`                    |
| Supported Platforms  | follows Hugo's availability |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/compose/jobs/` directory.
2. Once done, verify that:
   1. `Variables.MainLang` is your selected main language. Default is `en`.
   2. `FMTVariables.SourceDir` is the directory to execute `hugo` command.
   3. `FMTVariables.DestinationDir` is the directory path to export the website
      artifacts.
3. Add any additional `[[Dependencies]]` if you're customizing the default
  `[[CMD]]` commands list accordingly.
4. Customize `[[CMD]]` commands list as per your need.
5. If you're not building poly-lingual website, you have to remove:
   1. the `Hugo Workaround with 404` command.
   2. the `MainLang` variable.
6. If you are not on GitLab Pages:
   1. Please remove the `Copy GitLab CI if available` command as it is meant for
      double safety ensuring GitLab CI actually publishes the artifacts for
      `gh-pages` branch.

#### Changes
1. *Backward Compatible* - Added quietly delete command before generation.
2. *Backward Compatible* - Make `404.html` workaround to be quiet.



### Version 1.0.0
Version 1.0.0 Hugo Compose API is available for download here:
{{< link "/ci-jobs/compose/hugo/hugo-v1p0p0.toml" "this" "" "" "button" ""
	"download" >}}
hugo-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                      |
|:---------------------|----------------------------:|
| Monteur Version      | `v0.0.1`                    |
| Supported Platforms  | follows Hugo's availability |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/compose/jobs/` directory.
2. Once done, verify that:
   1. `Variables.MainLang` is your selected main language. Default is `en`.
   2. `FMTVariables.SourceDir` is the directory to execute `hugo` command.
   3. `FMTVariables.DestinationDir` is the directory path to export the website
      artifacts.
3. Add any additional `[[Dependencies]]` if you're customizing the default
  `[[CMD]]` commands list accordingly.
4. Customize `[[CMD]]` commands list as per your need.
5. If you're not building poly-lingual website, you have to remove:
   1. the `Hugo Workaround with 404` command.
   2. the `MainLang` variable.
6. If you are not on GitLab Pages:
   1. Please remove the `Copy GitLab CI if available` command as it is meant for
      double safety ensuring GitLab CI actually publishes the artifacts for
      `gh-pages` branch.

#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.
3. *Backward Compatible* - Uses `gh-pages` for compatibility across
   [Github Pages](https://pages.github.com/) and
   [GitLab pages](https://docs.gitlab.com/ee/user/project/pages/).
4. *Backward Compatible* - Added Hugo's `404.html` workaround by default for
   poly-lingual website.



## Epilogue
That's all for Monteur operating Hugo software. If you found a bug or have any
questions about the recipe, please feel free to raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
