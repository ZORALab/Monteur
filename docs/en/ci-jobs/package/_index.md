+++
date = "2021-12-21T17:33:13+08:00"
title = "Package"
description = """
Montuer natively facilitates the `Package` CI Job with the sole purpose to pack
various variants of the application or libraries for various platform system
deployments. This job is mainly for executing consistent packaging whenever and
whereever is needed easily and seamlessly.
"""
keywords = [
	"package",
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
url = "/en/ci-jobs/package/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Package CI Job"

[thumbnails.1]
url = "/en/ci-jobs/package/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Package CI Job"

[thumbnails.2]
url = "/en/ci-jobs/package/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Package CI Job"


[menu.main]
parent = "CI Jobs"
name = "Package"
pre = "📦"
weight = 80
identifier = "ci-jobs-package"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of the job is simple: **to execute all packaging for various
deployments seamlessly and consistently on a single host machine with minimal to
no further instructions**.




## Overall Configurations
To configure the job for execution, you need to supply and modify
`.configs/monteur/package/config.toml` file. These are the various settings for
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




## Packages' Configurations
Monteur accepts one packaging recipes file per packaging variant (e.g. one
`deb.toml` for `linux-amd64`, `linux-arm64`, `linux-riscv`, and etc).
However, the internal operations allow many packaging recipes file to run
simultenously and asynchonously (e.g. `deb.toml`, `appimage.toml`, `targz.toml`,
...). Each configuration file shares the same file structure.



### Storing Location
All packaging recipe configuration files **SHALL** be stored inside
`.configs/monteur/package/jobs` directory.




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
Name = 'DEB'
Description = """
Package monteur into .deb package file using manual commands.
"""
Type = 'deb-manual'
```

The `Name` field will be used for various internal configurations for Monteur's
packaging recipe identifications and logging purposes.

The `Description` is mainly for logging and the config file comprehension
purposes. You can write a short description for it.

The `Type` is the type of supported packaging modes for Monteur to execute
the recipes. The currently supported modes are:

* `deb-manual` - compile to `.deb` package but using manual commands instead.
  Useful for external compilers like `debuild`, `sbuild`, or `fpm`.
* `manual` - compile package completely manually using commands. It is for those
  who wants complete manual controls.


#### `[Variables]`
This table houses all [Plain Variables Definition]({{< link
"/internals/variables-processing/#plain-variables-definition" "this"
"url-only" />}}) **specific to this packaging recipes**. It shall appears onto
all listed packages. Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
GPGID = 'hello@zoralab.com'
GPGExistence = ''  # to be filled later for verifications
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

**DO NOT CONFUSE `[Dependencies]` with your product's dependencies**. This
design was meant for checking any external programs needed for additional setup
configurations (e.g. building from source codes requires a bunch of libraries
and make programs available as commands and etc). Hence, it's entirely optional
to have this table.

Here is an example for defining a list of dependencies:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[Dependencies]]
Name = 'DEB Builder'
Condition = 'all-all'
Type = 'command'
Command = 'debuild'

[[Dependencies]]
Name = 'GPG Tool'
Condition = 'all-all'
Type = 'command'
Command = 'gpg'

[[Dependencies]]
Name = 'Git for Changelog Generation'
Condition = 'all-all'
Type = 'command'
Command = 'git'
```

For each dependency, the `Name` is mainly for logging and referencing purposes.
The `Command` however, is the program's command name or library name depending
on `Type`. On UNIX system, shell's `alias`es are not supported and will never be
visible to Monteur.

Example: for `Git for Changelog Generation`, the `Name` can be anything but the
`Command` **SHALL ALWAYS** be the executable name: `git`.

`Condition` is the Montuer's `OS-ARCH` platform identification ID. See
[Platform Identification]({{< link "/internals/platform-identification" "this"
"url-only" />}}) for more info about its value.

Currently, Monteur supports the following `Type` values:

* `command` - identify by terminal command (e.g. `hugo`, `git`, ...)


#### `[Packages.XXX]`
`[Package]` are the list of packages to be packaged by the `[CMD]` table. All
these packages shall be iterated with the `[CMD]` table with the given variables
above, including a now processed list of `ChangelogEntries` value (now a
`[]string`).

Depending on `Metadata.Type`, Monteur will execute a preparation executions
before executing the `[CMD]` for a package. This reduces the need to build large
`[CMD]` commands list and promotes consistency.

For `Metadata.Type` set to `manual`, to ensure you have complete control, this
recipe type shall not execute the preparation sequences and all the packages are
directly and completely controlled by `[CMD]`.

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
`[CMD]` is basically the array of packaging commands or instructions for
executing the packaging algorithms across each listed packages. Hence, this is
why it has extra square braces.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[CMD]]
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
Name = "Compile Deb Package"
Type = 'command'
Condition = [ 'all-all' ]
Location = '{{- .PackageDir -}}'
Source = 'debuild -b -k{{- .GPGID }} -a{{- index .PkgArch 0 }}'

...
```




## Known Templates
Now that you understand how Monteur executes `Package` CI Job, fortunately,
Monteur maintains a number of recipes for you to kick-start your deployment.
Here are some currently maintained templates for different deployments:
