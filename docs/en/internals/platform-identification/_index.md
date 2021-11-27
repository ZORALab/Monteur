+++
date = "2021-11-26T07:01:06+08:00"
title = "Platform Identification"
description = """
Monteur, being created using Go Programming Language, has the ability to
identify a runtime operating system and its CPU architecture. The combind values
are known as platform condition. This section explains these conditions in
details.
"""
keywords = [
	"platform",
	"identification",
	"monteur",
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
url = "/en/internals/platform-identification/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Internal Platform Identification"

[thumbnails.1]
url = "/en/internals/platform-identification/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Internal Platform Identification"

[thumbnails.2]
url = "/en/internals/platform-identification/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Internal Platform Identification"


[menu.main]
parent = "Internals"
name = "Platform Identification"
pre = "🧿"
weight = 5
identifier = "internals-platform-id"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Supporting Cross-Platform Operations
To ensure Monteur enables its customers to be able to work on different devices
and platforms, Monteur has to specify all its runtime conditions at some points
for defining a configuration files.

Example, for [Commands Execution Unit]({{< link "/internals/commands/" "this"
"url-only" />}}), a customer can generate a list of commands that is specific
to a platform and CPU architecture (e.g. `windows` with `arm64` CPU).




## Data Pattern
Although Monteur is created using Go Programming Language and
`go tool dist list` provides a great deal of supported cross-platform
conditions (e.g. `linux/amd64`) but the values' format is not usable for
filename and others.

As such, Monteur modifies the output to use simple dash (`-`) so that it can
be deployed in both filenames and taken as values consistently. Example:

`linux/amd64` ➤ evolves into ➤ `linux-amd64`

Hence, in Monteur, the platform ID complies to the following pattern:

```yaml {linenos=table,hl_lines=[],linenostart=1}
Pattern:    [OS]-[ARCH]
Example:   linux-amd64
```



### Generic `all-all`
There are certain situation where a configuration is platform independent (e.g.
`go` command since it is using external binary compiler).

Hence, Monteur supplies a generic `all-all` value to represent it instead of
listing all the conditions (which is a horror experience).




## Deployment
Often, whenever a configuration requires a platform identification, it is
usually named under `Condition`. Then depending on the deployment, that
configuration will only be considered when Monteur is operating on a matching
platform.

Here is an example from [Commands Execution Unit]({{< link
"/internals/commands/" "this" "url-only" />}}) where each command unit has a
`Condition` to tell Monteur when to take them in:

```toml {linenos=table,hl_lines=[3],linenostart=1}
[[CMD]]
Name = 'Get Publish Branch First Commit for Cleaning'
Condition = 'all-all'
Type = 'command'
Location = '{{- .WorkingDir -}}'
Source = 'git rev-list --max-parents=0 --abbrev-commit HEAD'
Target = ''
Save = 'FirstCommitID'
```




## Supported Platforms
Now that we understand how Monteur performs platform identifications, here are
the list of Monteur supported platforms:

{{< cards "en.cards.monteur.platforms"
	"--grid-column-max-width:max-content;--grid-justify-content:center;" >}}




## Epilogue
That's all for Monteur's Platform Identification. If you have any queries,
please proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
