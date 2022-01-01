+++
date = "2021-11-27T15:15:18+08:00"
title = "Publish"
description = """
Montuer natively facilitates the `Publish` CI Job with the sole purpose to
publish its documentations from the repository. This job is mainly for products
that has built-in documentation generator to update itself easily and
seamlessly.
"""
keywords = [
	"publish",
	"CI job",
	"monteur",
	"configurations",
	"documentation",
]
draft = false
type = ""
# redirectURL=""
layout = "list"


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
url = "/en/ci-jobs/publish/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Publish CI Job"

[thumbnails.1]
url = "/en/ci-jobs/publish/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Publish CI Job"

[thumbnails.2]
url = "/en/ci-jobs/publish/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Publish CI Job"


[menu.main]
parent = "CI Jobs"
name = "Publish"
pre = "📖"
weight = 110
identifier = "ci-jobs-publish"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of the job is simple: **to publish the repository's documentations
in a consistent manner, with minimal to no further instructions**.




## Overall Configurations
To configure the job for execution, you need to supply and modify
`.configs/monteur/publish/config.toml` file. These are the various settings for
customizations.



### `[Variables]`
To configure job-wide variables for all publishing tasks, you can include or
modify the existing `[Variables]` table. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
BuildOptions = "--minify"
```

This table only accepts [Plain Variables Definition]({{< link
"/internals/variables-processing/#plain-variables-definition" "this"
"url-only" />}}).

The values can be any data types as long as it is sensible for direct
replacement in a variable formatting activities.



### `[FMTVariables]`
To configure job-wide formattable variables for all publishing tasks, you can
include or modify the existing `[FMTVariables]` table. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[FMTVariables]
SourceDir = '{{- .WorkingDir -}}/public'
```

This table accepts [Formattable Variables Definition]({{< link
"/internals/variables-processing/#formattable-variables-definition" "this"
"url-only" />}}) (e.g. `{{- .Version -}}`).

These variables shall be processed after the `[Variables]` table and all the
formatting clauses shall be replaced with the given variables. The processed
`Key-Value` output data shall then be backfilled either create or overwrite back
into `[Variables]` table depending on its `Key-Value` existence.




## Publishers' Configurations
Monteur accepts one publisher configuration file per publishing channel (e.g.
one `gitlab-pages.toml` for `GitLab Pages`). However, the internal operations
allow many programs to publish simultenously and asynchonously (e.g.
`nginx.toml`, `gitlab-pages.toml`, `github-pages.toml`, ...). Each publishers
configuration file shares the same file structure.



### Storing Location
All program configuration files **SHALL** be stored inside
`.configs/monteur/publish/jobs/` directory.



### File Nature
The configuration file **MUST** have the file extension. Otherwise, it will be
ignored. Currently the following formats are supported and sorted by priority
sequences:

1. `TOML` (`program.toml`) - https://github.com/toml-lang/toml

The filename does not affect any of the parsed configurations so feel free to
name it according to your own pattern.



### Data Structure
Each configuration file follows the following data structure:


#### `[Metadata]`
This table houses all the information about the program metadata such as its
name, description, and how to source the program. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Metadata]
Name = 'Gitlab Pages'
Description = """
Publishing using Hugo static site generator directly onto GitLab Pages.
"""
```

The `Name` field will be used for various internal configurations so Monteur
recommends **keeping it concise and without any weird symbol (space will be
replaced by short dash)**.

The `Description` is mainly for logging and the config file comprehension
purposes. You can write a short description for it.


#### `[Variables]`
This table houses all [Plain Variables Definition]({{< link
"/internals/variables-processing/#plain-variables-definition" "this"
"url-only" />}}) **specific to this publisher**. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
MainLang = 'en'
PublishBranch = 'gh-pages'
FirstCommitID = 'Will be overwritten by publish command sequences'
```

All the variables are either `create` or `overwrite` to the existing variables
list.


#### `[FMTVariables]`
This table houses all [Formattable Variables Definition]({{< link
"/internals/variables-processing/#formattable-variables-definition" "this"
"url-only" />}}) (e.g. `{{- .Version -}}`) **specific to this publisher**.
Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[FMTVariables]
SourceDir = '{{- .WorkingDir -}}/public'
DestinationDir = '{{- .WorkingDir -}}/{{- .PublishBranch -}}'
```

All formatted variables are either `create` or `overwrite` to the existing
variables list.


#### `[[Dependencies]]`
The `[Dependencies]` is an array of program data for checking all required
programs only for this `Publish` task. Therefore, it has extra square braces
when defining its data.

**DO NOT CONFUSE `[Dependencies]` with your product's dependencies**. This
design was meant for any external programs that has additional setup
configurations (e.g. building from source codes requires a bunch of libraries
and make programs available as commands and etc). Hence, it's entirely optional
to have this table.

Here is an example for defining a list of dependencies:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[Dependencies]]
Name = 'Hugo'
Condition = 'all-all'
Type = 'command'
Command = 'hugo'

[[Dependencies]]
Name = 'Git'
Condition = 'all-all'
Type = 'command'
Command = 'git'
```

For each dependency, the `Name` is mainly for logging and referencing purposes.
The `Command` however, is the program's command name or library name depending
on `Type`. On UNIX system, shell's `alias`es are not supported and will never be
visible to Monteur.

Example: for `Git Version Control Software`, the `Name` can be anything but the
`Command` **SHALL ALWAYS** be the executable name: `git`.

`Condition` is the Montuer's `OS-ARCH` platform identification ID. See
[Platform Identification]({{< link "/internals/platform-identification" "this"
"url-only" />}}) for more info about its value.

Currently, Monteur supports the following `Type` values:

* `command` - identify by terminal command (e.g. `hugo`, `git`, ...)


#### `[[CMD]]`
`[CMD]` is basically an array of instructions for publishing the repository's
documentations to the publisher. Hence, this is why it has extra square braces.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[CMD]]
Name = 'Check Artifact Directory Exists and Ready'
Condition = 'all-all'
Type = 'is-exists'
Source = '{{- .SourceDir -}}/index.html'

[[CMD]]
Name = 'Remove Git Workspace for Publishing Branch'
Condition = 'all-all'
Type = 'command-quiet'
Source = 'git worktree remove "{{- .DestinationDir -}}"'

[[CMD]]
Name = 'Delete Publishing Directory Regardlessly'
Condition = 'all-all'
Type = 'delete-recursive-quiet'
Source = '{{- .DestinationDir -}}'

...
```




## Known Templates
Now that you understand how Monteur executes `Publish` CI Job, fortunately,
Monteur maintains a number of recipes for you to kick-start your deployment.
Here are some currently maintained templates for different deployments:
