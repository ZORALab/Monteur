+++
date = "2021-12-22T18:26:41+08:00"
title = "GitLab (Pages) Publish API"
description = """
Monteur is a strong supporter for GitLab Pages and is currently using GitLab to
manage its software publications. This recipe allows repository integrated with
GitLab Pages to publish its repository's website artifact in a highly
customizable manner easily and seamlessly.
"""
keywords = [
	"GitLab Pages",
	"publish CI Job",
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
url = "/en/ci-jobs/publish/gitlab/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Publish with GitLab Pages"

[thumbnails.1]
url = "/en/ci-jobs/publish/gitlab/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Publish with GitLab Pages"

[thumbnails.2]
url = "/en/ci-jobs/publish/gitlab/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Publish with GitLab Pages"


[menu.main]
parent = "Platform - GitLab"
name = "Publish (GitLab Pages)"
pre = "📖"
weight = 110
identifier = "ci-jobs-publish-gitlab"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to make publishing to
GitLab Pages carried out in a consistent manner with minimal to no further
instructions**.

GitLab Pages is available at: https://docs.gitlab.com/ee/user/project/pages/

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## One-time Setup
Due to the recipe nature of using `gh-pages` branch for publications, you need
to setup that branch with a 1-time setup before use. Otherwise, the recipe shall
fail spectacularly.

To begin, simply execute the following git commands to create an empty
`gh-pages` branch:

{{< note warning "Be Careful" >}}
Make sure you perform the following with a git repository on a clean status
branch (no staging). Otherwise, you will lose files as one of the command
actually deletes everything.
{{< /note >}}

```bash {linenos=table,hl_lines=[],linenostart=1}
$ git checkout --orphan gh-pages
$ git reset --hard
$ git clean -fd
$ git commit --allow-empty -m "Init"
$ git push origin gh-pages:gh-pages
$ git checkout <back to your work branch>
```




## Recipe Versions
Here are the available Publish API recipe for GitLab Pages integrations. Please
read through your selected version's details on what has changed, what is
required, and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 1.0.0
Version 1.0.0 GitLab Pages API is available for download here:
{{< link "/ci-jobs/publish/gitlab/gitlab-pages-v1p0p0.toml" "this" "" ""
"button" "" "download" >}}
gitlab-pages-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                      |
|:---------------------|----------------------------:|
| Monteur Version      | `v0.0.1`                    |
| Supported Platforms  | follows Git's availability  |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/publish/jobs/` directory as `gitlab-pages.toml`.
2. Once done, verify that:
   1. `FMTVariables.SourceDir` is the directory holding the website artifact
       (e.g. `public/` directory).
3. Add any additional `[[Dependencies]]` if you're customizing the default
  `[[CMD]]` commands list accordingly.
4. Customize `[[CMD]]` commands list as per your need.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.
3. *Backward Compatible* - Uses `gh-pages` for compatibility across
   [Github Pages](https://pages.github.com/) and
   [GitLab pages](https://docs.gitlab.com/ee/user/project/pages/).




## Epilogue
That's all for Monteur operating on GitLab platform with GitLab Pages. If you
found a bug or have any questions about the recipe, please feel free to raise
your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
