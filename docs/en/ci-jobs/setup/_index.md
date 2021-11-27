+++
date = "2021-11-25T12:54:54+08:00"
title = "Setup"
description = """
Montuer natively facilitates the `Setup` CI job with the sole purpose to setup
the repository for development and testing, building the releasable version,
package and sign the release version. This section covers entirely on setting up
`Setup` job.
"""
keywords = [
	"setup",
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
url = "/en/ci-jobs/setup/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Setup CI Job"

[thumbnails.1]
url = "/en/ci-jobs/setup/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Setup CI Job"

[thumbnails.2]
url = "/en/ci-jobs/setup/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Setup CI Job"


[menu.main]
parent = "CI Jobs"
name = "Setup"
pre = "🧩"
weight = 5
identifier = "ci-jobs-setup"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of the job is simple: **to setup the repository all the way for
other jobs including development without minimal to no further instructions**.




## Overall Configurations
To configure the job for execution, you need to supply and modify
`.configs/monteur/setup/config.toml` file. These are the various settings for
customizations.



### Downloads
To configure the download behavior, you can include or modify the existing
`[Downloads]` table. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Downloads]
Limit = 3
Timeout = 120000000000 # nanosecond
```


#### `Limit`
The `Limit` is to instruct Monteur to control the maximum number of download
task at a time.

The data type is `unsigned integer`.

Monteur executes download tasks asynchonously so it will cause a spike of
download bandwidth should there be many downloads. Hence, the `Limit` is set in
place to ensure Monteur does not cause a network distruption.

#### `Timeout`
The `Timeout` is to instruct a sourcing task to stop and to report itself as
error due to unknown long waiting time. The unit is `nanosecond` so 2 minutes
is `120 * 1000 * 1000 * 1000`.

The data type is positive `64-bit integer`.

Setting the value to `0` will revert to default timeout which is `2 minutes`.




## Program Configurations
Monteur accepts one program configuration file per dependent program type (e.g.
one `go.toml` for `Go programming language`). However, the internal operations
allow many programs to be setup simultenously and asynchonously (e.g.
`go.toml`, `hugo.toml`, `golangci-lint.toml`, ...). Each program
configuration file shares the same file structure.



### Storing Location
All program configuration files **SHALL** be stored inside
`.configs/monteur/setup/programs/` directory.



### File Nature
The configuration file **MUST** have the file extension. Otherwise, it will be
ignored. Currently the following formats are supported sorted by priority
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
Name = 'Go'
Description = """
Go programming language for compiling Go source codes.
"""
Type = 'https-download'
```

The `Name` field will be used for various internal configurations so Monteur
recommends **keeping it concise and without any weird symbol (space will be
replaced by short dash)**.

The `Description` is mainly for logging and the config file comprehension
purposes. You can write a short description for it.

Among the supported sourcing `Type` values are:

* `https-download` - download via HTTPS remote server.
* `local-system`   - check local system has the program ready for use.


#### `[Variables]`
This table houses all the data that are subjected to change overtime. It takes
a `key-value` combination to house your value for later use. Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
Version = '1.17.3'
BaseURL = 'https://golang.org/dl/'
```

Note that the values **SHALL be a non-formatting value (e.g. without
`{{- .SomeKey -}}`)**. To use those, see `[FMTVariables]` section.

The values can be any data types as long as it is sensible for direct
replacement in a variable formatting activities.


#### `[FMTVariables]`
This table houses all the data that are subjected to change overtime and
containing formatting clause (e.g. `{{- .Version -}}`). These variables shall be
processed after the `[Variables]` table and all clauses shall be replaced with
the given variables. The processed output `key-value` data shall be backfilled
or be overwritten back into `[Variables]` table.

Example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[FMTVariables]
SourceDir = '{{- .WorkingDir -}}/public'
```

Will be processed and stored as:

```go {linenos=table,hl_lines=[],linenostart=1}
Variables["SourceDir"] ➤ "/home/u0/Documents/projects/tmp/setup/public"
```

The **arrangement of the elements does not dictates the order of processing**
due to the `Key-Value` data nature. Hence, please be creative with the `Key`
naming and **only creates `FMTVariable` independent of each other**.

{{< note info "For Your Information" >}}
The formatting clauses are strictly using Go's
[text/template](https://pkg.go.dev/text/template) package where the `key` of the
variable denotes the value replacement. Keep in mind that without the leading
period (e.g. `{{ Version }}`), the clause is invalid.

To fill in, simple append a period in front a desired key. Example, for
`Version = '1.16.3'` variable, the clause can be any of the following:

1. `{{ .Version }}.tar.gz` - direct replacement without any whitespace trimming
2. `{{- .Version }}.tar.gz` - direct replacement and trim leading whitespace
3. `{{ .Version -}}.tar.gz` - direct replacement and trim tailing whitespace
4. `{{- .Version -}}.tar.gz` - direct replacement and trim leading and tailing
   whitespaces

All the above shall be formatted into: `1.16.3.tar.gz` as its final output
and depending on your whitespace trimming request, can join into previous or
after.
{{< /note >}}

Monteur does supply a number of default variables for formatting. See
[Variables Processing]({{< link "/internals/variables-processing" "this"
"url-only" />}}) section.


#### `[[Dependencies]]`
The `[Dependencies]` is an array of program data for checking all required
programs only for `Setup` job. Therefore, it has extra square braces when
defining its data.

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
command **SHALL ALWAYs** be the executable name: `git`.

`Condition` is the Montuer's `OS-ARCH` platform identification ID. See
[Platform Identification]({{< link "/internals/platform-identification" "this"
"url-only" />}}) for more info of its value.

Currently, Monteur supports the following `Type` values:

* `command` - identify by terminal command (e.g. `hugo`, `git`, ...)




#### `[Sources.OS-ARCH]`
The `[Sources.OS-ARCH]` table is a list of sourcing instructions for Monteur
to source the program across various operating system (`OS`) and CPU
architecture (`ARCH`).

Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Sources.all-all]
Format = 'tar.gz'
URL = '{{- .BaseURL -}}{{- .Archive -}}'
Method = 'GET'
[Sources.all-all.Checksum]
Type = 'sha256'
Format = 'hex'

[Sources.darwin-amd64]
Archive = 'go{{- .Version -}}.darwin-amd64.{{- .Format -}}'
[Sources.darwin-amd64.Checksum]
Value = '765c021e372a87ce0bc58d3670ab143008dae9305a79e9fa83440425529bb636'

[Sources.darwin-arm64]
Archive = 'go{{- .Version -}}.darwin-arm64.{{- .Format -}}'
[Sources.darwin-arm64.Checksum]
Value = 'ffe45ef267271b9681ca96ca9b0eb9b8598dd82f7bb95b27af3eef2461dc3d2c'

...

[Sources.windows-arm64]
Format = 'zip'
Archive = 'go{{- .Version -}}.windows-arm64.{{- .Format -}}'
[Sources.windows-arm64.Checksum]
Value = '4e7f9a19af8a96e81b644846f27d739344375f9c69bad2e673406ab8e8a01101'
```

The `[Sources.{OS}-{ARCH}]` naming is based on Montuer's `OS-ARCH` platform
identification values. See
[Platform Identification]({{< link "/internals/platform-identification" "this"
"url-only" />}}) for more info.

Monteur allows overwriting to `all-all` platform ID data for the sole purpose of
avoiding duplicating values especially for programs supporting wide arrays of
operating system and CPU architectures. Hence, **it is always recommended to
define the common values into the `all-all` platform ID before defining the
platform specific ones** as shown above.

Then, based on the program's downloadable portal, the specific values are mapped
back accordingly. Should there be any conflicting fields from `all-all`
platform ID, it shall be overwritten by the platform specific ones. Based on the
example above, to setup the program on `windows-arm64` platform, the final
`Format` value is `.zip` and not the `tar.gz`.

Monteur also supports checksum function for security and integrity purpose. To
declare a checksum, simply states the `Checksum` after the platform table.
Following the example above, it would be:
`[Sources.darwin-arm64.Checksum]`. Similarly, one can also define common
checksum settings in the `all-all` common platform ID for avoiding duplications.

In the end, a source for a platform with an CPU architecture should comply
to the following data structure:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Sources.darwin-amd64]
Format = 'tar.gz'
Archive = 'go{{- .Version -}}.darwin-amd64.{{- .Format -}}'
URL = '{{- .BaseURL -}}{{- .Archive -}}'
Method = 'GET'
[Sources.darwin-amd64.Checksum]
Type = 'sha256'
Format = 'hex'
Value = '765c021e372a87ce0bc58d3670ab143008dae9305a79e9fa83440425529bb636'
```

`Format` is to define the downloaded compressed archive format. Currently,
Monteur supports the following formats:

* `tar.gz` - Tar Gunzip archive commonly used on UNIX platforms.
* `zip` - Zip archive commonly used on Windows platform.
* `raw` - Not an archive. Do nothing.

`Archive` is the naming format of the program after being sourced. This is for
unarchive purposes and since there is no standardization across all naming
conventions, it needs one to define the pattern for it. Variables formatting
clauses are allowed for this field.

{{< note info "Why Archive Field?" >}}
Take a look at both:
1. https://github.com/golangci/golangci-lint/releases/tag/v1.43.0
2. https://github.com/gohugoio/hugo/releases/tag/v0.89.4

You can't blame them for being customer-oriented to their respective specific
customers:
 * [Hugo](https://gohugo.io/) is for non-technical publishers (not Go
   developer/programmer).
 * [Golangci-Lint](https://golangci-lint.run/) is specifically for Go
   developer/programmer.
{{< /note >}}

`URL` is the location of the source in URI with protocol format. For local
directory or files, simply use the local file protocol (`file://`). Variables
formatting cluses are allowed for this field.

`Method` is the action of sourcing. For `https-download` type, `Method` is the
case-sensitive HTTP Method.

`Checksum.Type` is the algorithm used for generating the checksum token. Monteur
currently supports the following algorithms:
* `sha256`
* `sha512`
* `md5`

`Checksum.Format` is the checksum algorithm format used to parse the given
`Checksum.Value`. Monteur supports the following formats:
* `base64`
* `base64-url`
* `hex`

`Checksum.Value` is the checksum value for that particular source in `string`
format.


#### `[[Setup]]`
`[Setup]` is basically an array of instructions for installing the unarchived
program into Monteur's localized binary directory. Hence, this is why it has
extra square braces.

Its values are complying to Monteur's [Commands Execution Units]({{< link
"/internals/commands/" "this" "url-only" />}}). Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[Setup]]
Type = 'move'
Condition = 'all-all'
Source = '{{- .WorkingDir -}}/go'
Target = '{{- .BinDir -}}/golang'
```

For programs requiring build from source, you can build your chain of commands
here.


#### `[Config]`
The config is the last step of `Setup` job where it generates the config
sourcing script specific to the operating system. These scripts are for the
customers to source from Monteur's local filesystem to develop the repository.

Here is an example of a specific config file in linux for Go Programming
Language looks like:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Config]
linux = """
#!/bin/bash
export LOCAL_BIN="${LOCAL_BIN:-{{- .BinDir -}}}"
export GOROOT="${LOCAL_BIN}/golang"
export GOPATH="${LOCAL_BIN}/gopath"
export GOBIN="${GOPATH}/bin"
export GOCACHE="${LOCAL_BIN}/gocache"
export GOENV="${LOCAL_BIN}/goenv"

stop_go() {
        PATH=:${PATH}:
        GOROOT="${GOROOT}/bin"
        PATH=${PATH//:$GOROOT:/:}
        PATH=${PATH//:$GOBIN:/:}
        PATH=${PATH//:$GOPATH:/:}
        PATH=${PATH%:}
        unset GOROOT GOPATH GOBIN GOCACHE GOENV
}

case $1 in
--stop)
        stop_go
        ;;
*)
        export PATH="${PATH}:${GOROOT}/bin:${GOPATH}:${GOBIN}"

        if [ ! -z "$(type -p go)" ] && [ ! -z "$(type -p gofmt)" ]; then
                1>&2 printf "[ DONE  ] localized Go started.\\n"
        else
                1>&2 printf "[ ERROR ] localized Go failed to initalized.\\n"
                stop_go
        fi
        ;;
esac
"""
```

Monteur will seek out the right config data based on the platfrom OS. Should you
need to be specific with a CPU architecture, please script it inside your
config file.




## Known Templates
Now that you understand how Monteur executes `Setup` CI Job, fortunately,
Monteur maintains a number of recipes for you to kick-start your deployment.
Here are some currently maintained templates for different deployments:
