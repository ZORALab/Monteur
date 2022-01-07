+++
date = "2022-01-07T09:40:58+08:00"
title = "Getting Started with GitLab Integrations"
description = """
Monteur is set to be integration friendly with GitLab CI so that no additional
learning debt is needed. This allows Monteur to isolate vendor locked-in factor,
allowing you to easily swap your repository with other CI providers. This
section guides you on how to integrate Monteur with GitLab.
"""
keywords = [
	"getting started",
	"gitlab",
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
url = "/en/getting-started/gitlab/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Getting Started with GitLab Integrations"

[thumbnails.1]
url = "/en/getting-started/gitlab/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Getting Started with GitLab Integrations"

[thumbnails.2]
url = "/en/getting-started/gitlab/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Getting Started with GitLab Integrations"


[menu.main]
parent = "Platform - GitLab"
name = "Getting Started"
pre = "🛫"
weight = 5
identifier = "getting-started-gitlab"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Prerequisite
Before proceeding with constructing your `.gitlab-ci.yml` integration file,
you need to setup Monteur and ensures it works properly and locally. In short,
You should treat GitLab CI as your virtual employee who will execute your
Monteur automation tasks same as you.

If you have not setup Monteur yet, please visit the following section to get
started:

{{< link "/getting-started" "this" "" "" "button" >}}
Getting Started with Monteur
{{< /link >}}

Do not worry, you can always re-visit this section later using our navigation
panel.




## Objectives
GitLab integrations is heavily relying on its GitLab CI's `.gitlab-ci.yml` file.
Hence, the goal of this section is to construct one based on your repository's
goal. You do not need all the GitLab-CI tasks so choose wisely.

The guide here uses GitLab CI Runner with
[Docker Executor](https://docs.gitlab.com/runner/executors/docker.html) for
maximum [capability coverage](https://docs.gitlab.com/runner/executors/#selecting-the-executor)
and ease to use and using [debian:latest](https://hub.docker.com/_/debian)
Docker image to avoid paywall constraints.

However, since Monteur is for isolating vendor-specific CI configurations,
you can modify the the recipes below as per your repository and GitLab CI runner
needs.




## `.gitlab-ci.yml` Components
Here are the code fragments for constructing yout GitLab CI's `.gitlab-ci.yml`
file.

By minimum, the `.gitlab-ci.yml` **MUST** at least contain the following main
components:

```yaml {linenos=table,hl_lines=[],linenostart=1}
image: debian:latest

stages:
    - test
    - build
    - package
    - deploy
    - docs

variables:
    TERM: "xterm"
    GITLAB_CI: "true"
```

You need the docker `image` specification, the standard CI stages, and common
variables for outputs.

For CI `stages`, Monteur recommends `test`, `build`, and `docs` are sufficient.

For `package` and `deploy` CI `stages`, usually they involves secrecy data such
as key signing and etc. Since Monteur can now handle these jobs locally at your
side, there is no need to expose these secrecy to external contractors, further
protecting your secrecy data's confidentiality.



### Test
Coming soon.



### Build
Coming soon.



### Package
Coming soon.



### Deploy
Coming soon.



### Compose
Coming soon.



### Publish
Since Monteur uses `gh-pages` branch as
[GitLab Pages](https://docs.gitlab.com/ee/user/project/pages/) its primary
implementation, it's very easy to implement the
[Monteur Publish API]({{< link "/ci-jobs/publish/gitlab/" "this"
"url-only" />}}) protocol. The configurations are shown as follow:

```yaml {linenos=table,hl_lines=[],linenostart=1}
pages:
    stage: docs
    tags:
        - debian
    environment:
        name: production
    only:
        refs:
            - gh-pages
    artifacts:
        paths:
            - public
        expire_in: 1 day
    script:
        - mkdir -p public
        - shopt -s extglob
        - mv !(public|.*) public
```

What the code does is moving all the contents in the branch into `public/`
directory for GitLab CI to artifact it. Currently, there is
[no way](https://gitlab.com/gitlab-org/gitlab-runner/-/issues/1057) to artifact
the root repository of the `gh-pages` so some file movements is needed.

**DO NOT** rename the job (`pages`) to anything else as it is a special job
recognized by GitLab CI.

Also, **DO NOT** remove the artifact expiry time (`expire_in`) as GitLab can
easily bloat your repository
[storage limit](https://docs.gitlab.com/ee/user/gitlab_com/index.html#account-and-limit-settings).




## Epilogue
That's all for integrating Monteur with GitLab. If you have any queries, please
proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
