+++
date = "2021-12-22T13:16:23+08:00"
title = "Release"
description = """
Montuer natively facilitates the `Release` CI Job with the sole purpose to
upstream various packages of the application or libraries for various platform
system. This job is mainly for executing consistent upstreaming tasks whenever
and whereever is needed easily and seamlessly.
"""
keywords = [
	"release",
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
url = "/en/ci-jobs/release/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Release CI Job"

[thumbnails.1]
url = "/en/ci-jobs/release/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Release CI Job"

[thumbnails.2]
url = "/en/ci-jobs/release/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Release CI Job"


[menu.main]
parent = "CI Jobs"
name = "Release"
pre = "🌾"
weight = 7
identifier = "ci-jobs-release"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of the job is simple: **to execute all the release processes
easily and seamlessly whenever requested in a consistent manner with minimal to
no further instructions**.




## Overall Configurations
To configure the job for execution, you need to supply and modify
`.configs/monteur/package/config.toml` file. These are the various settings for
customizations.



### `[Variables]`
To configure job-wide variables for all release tasks, you can include or
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
To configure job-wide formattable variables for all release tasks, you can
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




## Releases' Configurations
Monteur accepts one release configuration file per releasing variant (e.g. one
`reprepro.toml` for `linux-amd64.deb`, `linux-arm64.deb`, `linux-riscv.deb` and
etc). However, the internal operations allow many programs to execute release
tasks simultenously and asynchonously (e.g. `reprepro.toml`, `flatpak.toml`,
`appimagehub.toml`, ...). Each configuration file shares the same file
structure.



### Storing Location
All releasing recipe configuration files **SHALL** be stored inside
`.configs/monteur/release/releases/` directory.




### File Nature
The configuration file **MUST** have the file extension. Otherwise, it will be
ignored. Currently the following formats are supported and sorted by priority
sequences:

1. `TOML` (`recipe.toml`) - https://github.com/toml-lang/toml

The filename does not affect any of the parsed configurations so feel free to
name it according to your own pattern. Monteur recommends **using releasing
recipe name** to keep sorting work sane. Example: `reprepro.toml` for all
`.deb` packages.




### Data Structure
Each configuration file follows the following data structure:


#### `[Metadata]`
This table houses all the information about the program metadata such as its
name, description, and how to source the program. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Metadata]
Name = 'Reprepro'
Description = """
Monteur's .deb packagers released to hosting repository via Reprepro.
"""
Type = 'manual'
```

The `Name` field will be used for various internal configurations for Monteur's
packaging recipe identifications and logging purposes.

The `Description` is mainly for logging and the config file comprehension
purposes. You can write a short description for it.

The `Type` is the type of supported releasing modes for Monteur to execute the
recipes. The current supported modes are:

* `manual` - release the packages completely manually using commands. It is for
  those who wants complete manual controls.


#### `[Variables]`
This table houses all [Plain Variables Definition]({{< link
"/internals/variables-processing/#plain-variables-definition" "this"
"url-only" />}}) **specific to this releasing recipes**. It shall appears onto
all listed packages. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
GPGID = 'hello@zoralab.com'
GPGExistence = ''
Distribution = '' # to be filled later
```

All the variables are either `create` or `overwrite` to the existing variables
list.


#### `[FMTVariables]`
This table houses all [Formattable Variables Definition]({{< link
"/internals/variables-processing/#formattable-variables-definition" "this"
"url-only" />}}) (e.g. `{{- .Version -}}`) **specific to this releasing
recipe**. It shall appears onto all listed packages. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
DataPath = '{{- .DataDir -}}/debian/reprepro'
```

All formatted variables are either `create` or `overwrite` to the existing
variables list.


#### `[[Dependencies]]`
The `[Dependencies]` is an array of programs due for existence checking
specifically **meant for the entire packaging recipe**. Therefore, it has extra
square braces when defining its data.

**DO NOT CONFUSE `[Dependencies]` with your product's dependencies**. This
design was meant for any external programs that has additional setup
configurations (e.g. building from source codes requires a bunch of libraries
and make programs available as commands and etc). Hence, it's entirely optional
to have this table.

Here is an example for defining a list of dependencies:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[Dependencies]]
Name = 'Reprepro'
Condition = 'all-all'
Type = 'command'
Command = 'reprepro'

[[Dependencies]]
Name = 'GPG Tool For Signing'
Condition = 'all-all'
Type = 'command'
Command = 'gpg'
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


#### `[Releases]`
`[Releases]` are the list of packaged file data to be released by executing the
`[CMD]` table onto each of the item. All these packages shall be iterated with
the `[CMD]` table supplied with the given variables above.

Depending on `Metadata.Type`, Monteur will execute a preparation executions
before executing the `[CMD]` for a package. This reduces the need to build
large `[CMD]` commands list and promotes consistency.

The code example is as follows:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Releases]
Target = '{{- .ReleaseDir -}}/deb'

[Releases.Packages.linux-amd64]
Source = '"{{- .PackageDir -}}/monteur-linux-amd64/monteur_"*"_amd64.deb"'

[Releases.Packages.linux-arm64]
Source = '"{{- .PackageDir -}}/monteur-linux-arm64/monteur_"*"_arm64.deb"'
```

* `[Releases]` - the table tag for the all the packages ready for release.
* `Target` - the directory path for housing the released data. This is used by
  local repository releases such as `reprepro`. This value shall be applied to
  all `Releases.Packages` without their `Target` value set.
* `[Releases.Packages]` - the list of individual packages.
* `[Releases.Packages.XXX]` - the package tag `XXX` can be anything as it only
  meant for identification and logging purposes in case there is an error.
* `Releases.Packages.XXX.Source` - the filepath / directory path for the
  packaged file.
* `Releases.Packages.XXX.Target` - the release destination of filepath /
  directory path for the packaged file. If this is not set, the `Target` in
  `[Releases]` table shall fill in automatically.


#### `[[CMD]]`
`[CMD]` is basically an array of releasing commands or instructions for
releasing the packaged software file to the upstream. Hence, this is why it has
extra square braces.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[CMD]]
Name = "Get GPG Secret Key for Verifications"
Type = 'command'
Condition = [ 'all-all' ]
Source = 'gpg --list-secret-keys "{{- .GPGID }}"'
Save = 'GPGExistence'

[[CMD]]
Name = "Verify GPG Secret Key Must Exists For Signing"
Type = 'is-not-empty'
Condition = [ 'all-all' ]
Source = '{{- .GPGExistence -}}'

[[CMD]]
Name = "Create Necessary Conf Data Directory"
Type = 'create-path'
Condition = [ 'all-all' ]
Source = '{{- .DataPath -}}/conf'

[[CMD]]
Name = 'Get Current Branch'
Type = 'command'
Condition = [ 'all-all' ]
Source = 'git branch --show-current'
Save = 'Distribution'

[[CMD]]
Name = "Verify Distribution is Not Empty"
Type = 'is-not-empty'
Condition = [ 'all-all' ]
Source = '{{- .Distribution -}}'

[[CMD]]
Name = "Release Using Reprepro"
Type = 'command'
Condition = [ 'all-all' ]
Source = """reprepro --basedir {{ .DataPath }} \
--outdir {{ .Target }} \
includedeb {{ .Distribution }} \
{{ .Source }}
"""
```




## Known Templates
Now that you understand how Monteur executes `Release` CI Job, fortunately,
Monteur maintains a number of recipes for you to kick-start your deployment.
Here are some currently maintained templates for different deployments:
