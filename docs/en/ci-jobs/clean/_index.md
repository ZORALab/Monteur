+++
date = "2021-12-22T16:38:08+08:00"
title = "Clean"
description = """
Montuer natively facilitates the `Clean` CI Job with the sole purpose to
clean up the repository from all the jobs' artifacts. This job is mainly for
cleaning up repository with customizable deletion easily and seamlessly.
"""
keywords = [
	"clean",
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
url = "/en/ci-jobs/clean/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Clean CI Job"

[thumbnails.1]
url = "/en/ci-jobs/clean/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Clean CI Job"

[thumbnails.2]
url = "/en/ci-jobs/clean/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Clean CI Job"


[menu.main]
parent = "CI Jobs"
name = "Clean"
pre = "🧹"
weight = 20
identifier = "ci-jobs-clean"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of the job is simple: **to clean the repository up from
development artifacts in a consistent manner with minimal to no further
instructions**.




## Overall Configurations
To configure the job for execution, you need to supply and modify
`.configs/monteur/clean/config.toml` file. These are the various settings for
customizations.



### `[Variables]`
To configure job-wide variables for all cleaning tasks, you can include or
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
To configure job-wide formattable variables for all composing tasks, you can
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




## Cleaners' Configurations
Monteur accepts one cleaner configuration file per publishing channel (e.g.
one `build-log.toml` for `cleaning build job's log directory`). However, the
internal operations allow many cleaner works to be executed simultenously and
asynchonously (e.g. `build-tmp.toml`, `release-log.toml`, `release-tmp.toml`,
`...`). Each configuration file shares the same file structure.



### Storing Location
All cleaning configuration files **SHALL** be stored inside
`.configs/monteur/clean/jobs/` directory.



### File Nature
The configuration file **MUST** have the file extension. Otherwise, it will be
ignored. Currently the following formats are supported and sorted by priority
sequences:

1. `TOML` (`recipe.toml`) - https://github.com/toml-lang/toml

The filename does not affect any of the parsed configurations so feel free to
name it according to your own pattern.



### Data Structure
Each configuration file follows the following data structure:


#### `[Metadata]`
This table houses all the information about the recipe metadata such as its
name, description, and how to source the program. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Metadata]
Name = 'Clean Up Build Log Artifacts'
Description = """
Clean up all artifacts for a job.
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
"url-only" />}}) **specific to this cleaner**. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
Job = 'build'
MonteurPath = '.monteurFS'
```

All the variables are either `create` or `overwrite` to the existing variables
list.


#### `[FMTVariables]`
This table houses all [Formattable Variables Definition]({{< link
"/internals/variables-processing/#formattable-variables-definition" "this"
"url-only" />}}) (e.g. `{{- .Version -}}`) **specific to this cleaner**.
Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[FMTVariables]
Target = '{{- .RootDir -}}/{{- .MonteurPath -}}/log/{{- .Job -}}'
```

All formatted variables are either `create` or `overwrite` to the existing
variables list.


#### `[[Dependencies]]`
The `[Dependencies]` is an array of program data for checking all required
programs only for this cleaner. Therefore, it has extra square braces when
defining its data.

It is compliant to Monteur's internal [Dependencies Checking]({{< link
"/internals/dependencies-checking/" "this" "url-only" />}}) function.

Here is an example for defining a list of dependencies:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[Dependencies]]
Name = 'Remove in Linux'
Condition = 'linux-all'
Type = 'command'
Command = 'rm'
```


#### `[[CMD]]`
`[CMD]` is basically an array of instructions for executing the cleaner's
cleaning work. Hence, this is why it has extra square braces.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[CMD]]
Name = "Remove Target Directory"
Type = 'delete-recursive-quiet'
Condition = [ 'all-all' ]
Source = '{{- .Target -}}'

[[CMD]]
Name = "Create Target Directory"
Type = 'create-path'
Condition = [ 'all-all' ]
Source = '{{- .Target -}}'
```




## Known Templates
Now that you understand how Monteur executes `Clean` CI Job, fortunately,
Monteur maintains a number of recipes for you to kick-start your deployment.
Here are some currently maintained templates for different deployments:
