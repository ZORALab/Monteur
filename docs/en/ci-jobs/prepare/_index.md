+++
date = "2022-01-03T09:07:54+08:00"
title = "Prepare"
description = """
Monteur natively facilitates the `Prepare` CI Job solely to update repository
for next build, package, and release that are otherwise not suitable to be in
any of them. This job is mainly for executing consistent packaging whenever and
whereever is needed easily and seamlessly.
"""
keywords = [
	"prepare",
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
url = "/en/ci-jobs/prepare/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Prepare CI Job"

[thumbnails.1]
url = "/en/ci-jobs/prepare/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Prepare CI Job"

[thumbnails.2]
url = "/en/ci-jobs/prepare/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Prepare CI Job"


[menu.main]
parent = "CI Jobs"
name = "Prepare"
pre = "🧶"
weight = 65
identifier = "ci-jobs-prepare"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of the job is simple: **to execute all packaging for various
deployments seamlessly and consistently on a single host machine with minimal to
no further instructions**.

This job was made available since Monteur `v0.0.2`.




## Overall Configurations
To configure the job for execution, you need to supply and modify
`.configs/monteur/prepare/config.toml` file. These are the various settings for
customizations.



### `[Variables]`
To configure job-wide variables for all packaging tasks, you can include or
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
To configure job-wide formattable variables for all packaging tasks, you can
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




## Prepare's Configurations
Monteur accepts one preparation recipes file per packaging variant (e.g. one
`deb-changlog.toml` for `linux-amd64`, `linux-arm64`, `linux-riscv`, and etc).
However, the internal operations allow many packaging recipes file to run
simultenously and asynchonously (e.g. `deb-changelog.toml`, `md-changelog.toml`,
`txt-changelog.toml`, ...). Each configuration file shares the same file
structure.



### Storing Location
All packaging recipe configuration files **SHALL** be stored inside
`.configs/monteur/prepare/jobs` directory.




### File Nature
The configuration file **MUST** have the file extension. Otherwise, it will be
ignored. Currently the following formats are supported and sorted by priority
sequences:

1. `TOML` (`recipe.toml`) - https://github.com/toml-lang/toml

The filename does not affect any of the parsed configurations so feel free to
name it according to your own pattern. Monteur recommends **using packaging
recipe name** to keep sorting work sane. Example: `deb.toml` for all
`.deb` packages.



### Data Structure
Each configuration file follows the following data structure:


#### `[Metadata]`
This table houses all the information about the packaging recipe metadata such
as its name, description, and its supported types. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
Name = 'DEB Changelog'
Description = """
prepare .deb changelog for all releases.
"""
Type = 'deb'
```

The `Name` field will be used for various internal configurations for Monteur's
packaging recipe identifications and logging purposes.

The `Description` is mainly for logging and the config file comprehension
purposes. You can write a short description for it.

The `Type` is the type of supported packaging modes for Monteur to execute
the recipes. The currently supported modes are:

* `deb` - generates/updates the `.deb` package changelog data.
* `md` - generates/updates the `.md` chnagelog file.


#### `[Variables]`
This table houses all [Plain Variables Definition]({{< link
"/internals/variables-processing/#plain-variables-definition" "this"
"url-only" />}}) **specific to this packaging recipes**. It shall appears onto
all listed packages. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
ChangelogFrom = 'origin/staging'
ChangelogTo = 'origin/main'
```

All the variables are either `create` or `overwrite` to the existing variables
list.


#### `[FMTVariables]`
This table houses all [Formattable Variables Definition]({{< link
"/internals/variables-processing/#formattable-variables-definition" "this"
"url-only" />}}) (e.g. `{{- .Version -}}`) **specific to this packaging
recipe**. It shall appears onto all listed packages. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[FMTVariables]
Lintian = '{{- .DataDir -}}/debian/lintian'
Profile = '{{- .DataDir -}}/debian/profiles'
```

All formatted variables are either `create` or `overwrite` to the existing
variables list.


#### `[[Dependencies]]`
The `[Dependencies]` is an array of programs due for existence checking
specifically **meant for the entire packaging recipe**. Therefore, it has extra
square braces when defining its data.

It is compliant to Monteur's internal [Dependencies Checking]({{< link
"/internals/dependencies-checking/" "this" "url-only" />}}) function.

Here is an example for defining a list of dependencies:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[Dependencies]]
Name = 'Git for Changelog Generation'
Condition = 'all-all'
Type = 'command'
Command = 'git'
```


#### `[Changelog]`
The `[Changelog]` is the table for generating the current version's changelog
entries. This `[Changelog]` shall be executed first before packaging executions.

Should its `Changelog.CMD` remains empty, this changelog update execution shall
be skipped entirely.

`[Changelog]` has its own set of important fields to process the generated
outcomes. Here is the current supported fields:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Changelog]
LineBreak = "\n" # This field must be in double quote for special characters.
Regex = ''
```

* `LineBreak` - the line-breaking character for breaking a long changelog string
  into list of changelog entries string.
* `Regex` - optional field designed for filtering each changelog entry line. The
  opening and closing forward slashes (`/`) are not required.

##### `[[Changelog.CMD]]`
The list of instructions for sourcing the changelog. Hence, this is why it has
extra square braces.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[9],linenostart=1}
[[Changelog.CMD]]
Name = "Get Changelog From Git Log Comparisons"
Type = 'command'
Condition = [ 'all-all' ]
Source = """git --no-pager log \
"{{- .ChangelogTo -}}..{{- .ChangelogFrom -}}" \
--pretty="format:%h %s"
"""
Save = "ChangelogEntries"
```

Regardless how long your instruction is, `[Changelog]` shall **ONLY** read the
final output, a single long `string` from the `ChangelogEntries` variables.
Hence, remember to indicate `Save =` field for the last instruction to output
the changelog entries data.

Note that you **DO NOT** need to manually split the entries. Monteur will
split them natively using the `Changelog.LineBreak` symbol.

##### Manually File Changelog Entries
If, for any edge case reason that you really need to fill in the changelog
entries manually, simply create a `Entries` field with an array of strings.
Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Changelog]
Entries = [
	'changed feature A',
	'changed feature B',
	...,
]
```

Only do this at the last resort since it involves a lot of manual labor-work.


#### `[Packages.XXX]`
`[Package]` are the list of packages for generating or updating the changelog
across different supporting compute system for a single packaging variant.
Example: `.deb` for `linux-amd64`, `linux-arm64`, `linux-386`, and etc.
All these packages shall be iterated with the `[CMD]` table with the given
variables above, including a now processed list of `ChangelogEntries` value
(now a `[]string`).

Depending on `Metadata.Type`, Monteur will format the changelog entries before
for the package. This reduces the need to build large `[CMD]` commands list and
promotes consistency.

The code example is as follows:

```toml {linenos=table,hl_lines=["1-11"],linenostart=1}
[Packages.001]
OS = [ 'linux' ]
Arch = [ 'amd64' ]
Changelog = '{{- .DataDir -}}/debian/changelog-amd64'
Distribution = [
        'stable',
]
BuildSource = false

[Packages.001.Files]
'{{- .PackageDir -}}/monteur' = '{{- .BuildDir -}}/linux-amd64'

[Packages.002]
OS = [ 'linux' ]
Arch = [ 'arm64' ]
Changelog = '{{- .DataDir -}}/debian/changelog-arm64'
Distribution = [
       'stable',
]
BuildSource = false

[Packages.002.Files]
'{{- .PackageDir -}}/monteur' = '{{- .BuildDir -}}/linux-arm64'
```

* `[Package.XXX]` - the package tag. `XXX` can be anything since it is only used
  to populate the list.
* `OS` - the list of supported operating system. The **first (1st)** shall be
  used for primary filling (e.g. only 1 OS from the list).
* `Arch` - the list of supported CPU architecture. The **first (1st)** shall be
  used for primary filling (e.g. only 1 Architecture from the list).
* `Changelog` - the filepath location for prepending new changelog entries data.
    * Formattable variables are available for dynamic formatting.
* `Distribution` - the supported distributions of the operating system (E.g.
  `stable`, `unstable`, `experimental`, `debian`, `ubuntu`, ...). When in
  doubts, sticks to `stable`, `unstable`, or `experimental`.
* `BuildSource` - the decision to build source packge instead of binary package
  for supported types of packaging recipe types (e.g. `.deb` package).
* `[Package.XXX.Files]` - the list of files to be copied over during
  package preparations stage (right before executing `[CMD]`). The `Key` is the
  destination while the `Value` is the source of the file. If we follows the
  example above, `{{- .BuildDir -}}/linux-amd64` shall be copied to
  `{{- .PackageDir -}}/monteur`.
    * Formattable variables are available for dynamic formatting for both `key`
      and `value`.


#### `[[CMD]]`
The list of instructions for sourcing the changelog. Hence, this is why it has
extra square braces.

The purpose is for anyone else to prepare anything else aside changelog.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
Name = 'Placeholder'
Type = 'placeholder'
Condition = [ 'all-all' ]
Source = ''
Target = ''
```




## Known Templates
Now that you understand how Monteur executes `Prepare` CI Job, fortunately,
Monteur maintains a number of recipes for you to kick-start your deployment.
Here are some currently maintained templates for different deployments: