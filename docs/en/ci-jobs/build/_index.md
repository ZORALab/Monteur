+++
date = "2021-12-02T08:22:42+08:00"
title = "Build"
description = """
Montuer natively facilitates the `Build` CI Job with the sole purpose to compile
various version of the application or libraries for various platform system.
This job is mainly for executing consistent build whenever and whereever is
needed easily and seamlessly.
"""
keywords = [
	"build",
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
url = "/en/ci-jobs/build/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Build CI Job"

[thumbnails.1]
url = "/en/ci-jobs/build/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Build CI Job"

[thumbnails.2]
url = "/en/ci-jobs/build/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Build CI Job"


[menu.main]
parent = "CI Jobs"
name = "Build"
pre = "🧰"
weight = 6
identifier = "ci-jobs-build"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of the job is simple: **to execute all the test suites easily
and seamlessly whenever requested in a consistent manner with minimal to no
further instructions**.




## Overall Configurations
To configure the job for execution, you need to supply and modify
`.configs/monteur/build/config.toml` file. These are the various settings for
customizations.



### `[Variables]`
To configure job-wide variables for all publishing tasks, you can include or
modify the existing `[Variables]` table. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
TestArguments = "--verbose"
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
MainDir = '{{- .RootDir -}}/gopkg'
```

This table accepts [Formattable Variables Definition]({{< link
"/internals/variables-processing/#formattable-variables-definition" "this"
"url-only" />}}) (e.g. `{{- .Version -}}`).

These variables shall be processed after the `[Variables]` table and all the
formatting clauses shall be replaced with the given variables. The processed
`Key-Value` output data shall then be backfilled either create or overwrite back
into `[Variables]` table depending on its `Key-Value` existence.




## Builds' Configurations
Monteur accepts one build configuration file per build variant (e.g. one
`linux-amd64.toml` for `app on Linux Operating System with x86_64 CPU`).
However, the internal operations allow many programs to build simultenously
and asynchonously (e.g. `darwin-amd64.toml`, `linux-arm64.toml`,
`linux-arm.toml`, ...). Each tester configuration file shares the same file
structure.



### Storing Location
All test configuration files **SHALL** be stored inside
`.configs/monteur/app/variants/` directory.

These config files are **also used across various build related CI Jobs** like
[Package API]({{< link "/ci-jobs/package/" "this" "url-only" />}}) and
[Release API]({{< link "/ci-jobs/release/" "this" "url-only" />}}).



### File Nature
The configuration file **MUST** have the file extension. Otherwise, it will be
ignored. Currently the following formats are supported and sorted by priority
sequences:

1. `TOML` (`program.toml`) - https://github.com/toml-lang/toml

The filename does not affect any of the parsed configurations so feel free to
name it according to your own pattern. Monteur recommends
**using [Platform ID]({{< link "/internals/platform-identification/" "this"
"url-only" />}})** to keep sorting work sane. Example: `linux-amd64.toml` for
`Linux Operating System` with `AMD64` CPU Architecture.



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
recommends **using [Platform ID]({{< link "/internals/platform-identification/"
"this" "url-only" />}})**.

The `Description` is mainly for logging and the config file comprehension
purposes. You can write a short description for it.

These variables are used across various build related CI Jobs like
[Package API]({{< link "/ci-jobs/package/" "this" "url-only" />}}) and
[Release API]({{< link "/ci-jobs/release/" "this" "url-only" />}}).


#### `[Variables]`
This table houses all [Plain Variables Definition]({{< link
"/internals/variables-processing/#plain-variables-definition" "this"
"url-only" />}}) **specific to this variant build**. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
MainLang = 'en'
```

All the variables are either `create` or `overwrite` to the existing variables
list.

These variables are used across various build related CI Jobs like
[Package API]({{< link "/ci-jobs/package/" "this" "url-only" />}}) and
[Release API]({{< link "/ci-jobs/release/" "this" "url-only" />}}).


#### `[FMTVariables]`
This table houses all [Formattable Variables Definition]({{< link
"/internals/variables-processing/#formattable-variables-definition" "this"
"url-only" />}}) (e.g. `{{- .Version -}}`) **specific to this variant build**.
Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[FMTVariables]
DataDir = '{{- .RootDir -}}/data'
```

All formatted variables are either `create` or `overwrite` to the existing
variables list.

These variables are used across various build related CI Jobs like
[Package API]({{< link "/ci-jobs/package/" "this" "url-only" />}}) and
[Release API]({{< link "/ci-jobs/release/" "this" "url-only" />}}).


#### `[[BuildDeps]]`
The `[BuildDeps]` is an array of program data for checking all required
programs only for this `Build` task. Therefore, it has extra square braces
when defining its data.

**DO NOT CONFUSE `[BuildDeps]` with your product's dependencies**. This
design was meant for any external programs that has additional setup
configurations (e.g. building from source codes requires a bunch of libraries
and make programs available as commands and etc). Hence, it's entirely optional
to have this table.

Here is an example for defining a list of dependencies:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[BuildDeps]]
Name = 'Go'
Condition = 'all-all'
Type = 'command'
Command = 'go'
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


#### `[[BuildCMD]]`
`[BuildCMD]` is basically an array of instructions for building the software of
a specific variant. Hence, this is why it has extra square braces.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[BuildCMD]]
Name = "Execute Go Build"
Type = 'command'
Condition = 'all-all'
Location = '{{- .RootDir -}}/gopkg'
Source = 'go build -ldflags "-s -w" -o "{{- .BuildPath -}}" "{{- .AppPath -}}"'

...
```




## Known Templates
Now that you understand how Monteur executes `Build` CI Job, fortunately,
Monteur maintains a number of recipes for you to kick-start your deployment.
Here are some currently maintained templates for different deployments:
