+++
date = "2022-01-03T16:32:00+08:00"
title = "Dependencies Checking"
description = """
Monteur has a built-in dependency checking function that is used across all jobs
for reproducible build purposes. This section explains the checking function
in details.
"""
keywords = [
	"dependencies",
	"checking",
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
url = "/en/internals/dependencies-checking/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Dependencies Checking Function"

[thumbnails.1]
url = "/en/internals/dependencies-checking/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Dependencies Checking Function"

[thumbnails.2]
url = "/en/internals/dependencies-checking/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Dependencies Checking Function"


[menu.main]
parent = "CI Internals"
name = "Dependencies Checking"
pre = "🧮"
weight = 5
identifier = "internals-dependencies-checking"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Main Purpose
The purpose is to ensure all external programs (dependencies) are checked to
be available **before executing** the list of
[Commands Execution Units]({{< link "/internals/commands/" "this"
"url-only" />}}).

That way, other Monteur users of your project can go through the list in order
to make sure everything is up-front ready instead of executing the job blindly
for a frankenstein result.



### Not to Confuse With Product Dependencies
**DO NOT CONFUSE this internal function with your products' dependencies** (as
in, your end-product app for your customers needs `library-X` to run). Monteur
can only covers the app manufacturing dependencies. It is upto you to package
your customers' dependencies using supported packagers (e.g. AppImage, Snaps,
and etc).



### Optional Implementation
It is **ENTIRELY OPTIONAL** to implement this function so any CI-Jobs
(especially those native functions supported jobs) that has 0 dependency can
leave the data table blank.




## Data Pattern
The data pattern for dependencies starts with the tag `[Dependencies]`. It is
an array of items so it usually comes with extra square braces that looks like
`[[Dependencies]]`.

A full data structure pattern with 2 entires are shown as follows:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[Dependencies]]
Name = 'Hugo'
Condition = 'all-all'
Type = 'command'
Command = 'hugo'

[[Dependencies]]
Name = 'Git Version Control Software for Changelog Generations'
Condition = 'all-all'
Type = 'command'
Command = 'git'
```

Each dependency has the following fields:
* `Name` - name of the dependency. Mainly for logging and self-referencing use.
* `Condition` - the compute [Platform Identification]({{< link
"/internals/platform-identification/" "this" "url-only" />}}) that specify the
dependency is applied to.
* `Type` - the dependency type. Currently Monteur supports:
  * `command` - the external command type.
* `Command` - the command that will be used for execution.




## Epilogue
That's all for Monteur's internal function for dependency checking. If you have
any queries, please proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
