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

{{< note warning "Heads Up" >}}
While Monteur trying to test itself on Raspberry Pi 3B GitLab Runner, we
encountered so many unknown errors to the point that **we do not recommend using
Raspberry Pi as GitLab Runner at all**, least for integrating Monteur.
{{< /note >}}

{{< note warning "Heads Up" >}}
Monteur currently only works with Shell Executor. See
https://docs.gitlab.com/runner/executors/ for more info.
{{< /note >}}



## `.gitlab-ci.yml` Components
Here are the code fragments for constructing yout GitLab CI's `.gitlab-ci.yml`
file.

By minimum, the `.gitlab-ci.yml` **MUST** at least contain the following main
components:

```yaml {linenos=table,hl_lines=[],linenostart=1}
image:
  name: debian:latest
  entrypoint: [ "" ]

variables:
  MONTEUR_CONFIG: ".monteurFS/config/main"

stages:
  - setup
  - test
  - build
  - docs

cache:
  - key: "MonteurFS"
    paths:
      - ".monteurFS/"

before_script:
  - apt-get update -y
  - apt-get upgrade -y
  - |
    apt-get --no-install-recommends install \
      curl \
      gnupg2 \
      ca-certificates \
      -y
  - |
    curl https://www.zoralab.com/pubkey.gpg \
      | gpg --yes --dearmor --output /usr/share/keyrings/zoralab-keyring.gpg
  - |
    echo 'deb [signed-by=/usr/share/keyrings/zoralab-keyring.gpg] https://monteur.zoralab.com/releases/deb next experimental' \
        >  /etc/apt/sources.list.d/zoralab-monteur.list
  - apt-get update -y
  - apt-get install monteur -y
```

This is the standard way of setting up Monteur. To reduce the Monteur setup
time, simply use the cache across all the jobs would be suffice (1 time setup
and use across all.

For CI `stages`, Monteur recommends `setup`, `test`, `build`, and `docs` are
sufficient.

For Monteur's `Package` and `Release` API, they usually involve secret data such
as key signing and etc. Since Monteur can now handle these jobs locally at your
side, there is no need to expose these secrecy to external contractors, further
protecting your secrecy data's confidentiality.

Obviously, the caching method is used as `setup` stage shall assemble your
Monteur local filesystem (`MONTEUR_CONFIG`). Then, the filesystem is cached
and used across all the remaining jobs, saving time and bandwidth for not
repeatedly calling Monteur Setup API for each job.



### Test
For Monteur Test API, assuming the test coverage text output is unchanged,
the recommended settings would be as follows:

```yaml {linenos=table,hl_lines=[],linenostart=1}
test:
  stage: test
  tags:
    - linux
  environment:
    name: production
  except:
    refs:
      - gh-pages
  interruptible: true
  coverage: '/TOTAL\s*TEST\s*COVERAGE:\s*(\d+.?\d*%?)/'
  script:
    - source "$MONTEUR_CONFIG"
    - monteur test
```



### Build
For Monteur Build API, the recommended settings would be as follows:

```yaml {linenos=table,hl_lines=[],linenostart=1}
build:
  stage: build
  tags:
    - linux
  environment:
    name: production
  except:
    refs:
      - gh-pages
  environment:
    name: production
  script:
    - source "$MONTEUR_CONFIG"
    - monteur prepare
    - monteur build
```



### Compose
For Compose GitLab CI Job, the recommended settings would be as follows:

```yaml {linenos=table,hl_lines=[],linenostart=1}
compose:
  stage: docs
  tags:
    - linux
  environment:
    name: production
  except:
    refs:
      - gh-pages
  script:
    - source "$MONTEUR_CONFIG"
    - monteur compose
```

Note that this Job only wants to test the Monteur Compose API is working fine.
It does not publish the output.





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
    - linux
  environment:
    name: production
  only:
    refs:
      - gh-pages
  cache: []
  artifacts:
    paths:
      - public
    expire_in: 1 day
  before_script:
    - mkdir -p public
    - shopt -s extglob
    - mv !(public|.*) public
  script:
    - printf "[ DONE ] Nothing to implement.\n"
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
