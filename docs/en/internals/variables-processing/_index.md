+++
date = "2021-11-26T07:00:07+08:00"
title = "Variables Processing"
description = """
Monteur has its own unique way to process all the required variables from main
config file down to unit-defined config files. This section is to explain how
it works internally.
"""
keywords = [
	"variables",
	"processing",
	"monteur internals",
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
url = "/en/internals/variables-processing/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Internals - Variables Processing"

[thumbnails.1]
url = "/en/internals/variables-processing/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Internals - Variables Processing"

[thumbnails.2]
url = "/en/internals/variables-processing/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Internals - Variables Processing"


[menu.main]
parent = "CI Internals"
name = "Variables Processing"
pre = "🪢"
weight = 5
identifier = "internals-variable-processing"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Variables Formatting
The main purpose is to facilitate formatting using variables capability for a
managable data processing across various configuration fields. Here is an
example from [Commands Execution Unit]({{< link "/internals/commands/" "this"
"url-only" />}}) where `Location`, `Source`, and `Target` has variables
formatting abilities:

```toml {linenos=table,hl_lines=["5-7"],linenostart=1}
[[CMD]]
Name = 'Get Publish Branch First Commit for Cleaning'
Condition = 'all-all'
Type = 'command'
Location = '{{- .WorkingDir -}}'
Source = 'git rev-list --max-parents=0 --abbrev-commit HEAD'
Target = ''
Save = 'FirstCommitID'
```

Not all data fields have such capability enabled. Hence, please refers to the
documentations' details to know when to use one.

To fill in a `Value` using the formatting clause, simply prepend a period in
front its `Key`. Example, for `Version = '1.16.3'` variable, the clause can be
any of the following:

1. `{{ .Version }}.tar.gz` - direct replacement without any whitespace trimming
2. `{{- .Version }}.tar.gz` - direct replacement and trim leading whitespace
3. `{{ .Version -}}.tar.gz` - direct replacement and trim tailing whitespace
4. `{{- .Version -}}.tar.gz` - direct replacement and trim leading and tailing
   whitespaces

All the above shall always be formatted into: `1.16.3.tar.gz` as its final
output and depending on your whitespace trimming request, the final value can be
miraculously joining into previous or after text. Hence, use trimming when
sensible.




## Data Structure
Monteur supplies 2 types of variables definitions:

1. Plain Variables Definitions
2. Formattable Variables Definitions

For both types, its values is stored in a simple `Key-Value` data pattern where:

1. the `Key` is **case-sensitive and strictly a `string`**
2. the `Value` can be **any data types**

When the `Value` is a `string`, Monteur recommends **using single quote
(`'...'`) over double quote (`"..."`)** as some values may use double quote part
of its value. This is to make things easier without many characters
cancellations (e.g. `\\n` for newline (`\n`)). For long `string` value however,
use triple double quote (`"""`) for opening and closing indicators.

Monteur's variables carry out the `create or overwrite policy`. This means that:

1. For any variables with **a new `Key`**, Montuer **shall `create` it as a new
   element**.
2. For any variables with **the same `Key`**, Monteur **shall `overwrite` the
   existing element with the latest value**.


### Plain Variables Definition
As titled, the Plain Variables Definition allows one to define plain values
with no formatting into the variables list. The data structure is something
as follows:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Variables]
Version = '1.17.3'
BaseURL = 'https://golang.org/dl/'
```

The table indication is `[Variables]`. All its values are `Key-Value` in nature
and **it shall not have any sub-definitions (e.g. `[Variables.MyDef]`)**.

Monteur shall parse all plain variables definitions **AS IT IS**. These
variables will be used for formatting values later on including Formattable
Variables Definition (see below).



### Formattable Variables Definition
The Formattable Variables Definition allows one to define a formattable value
before being added to the variables list. The data structure is something as
follows:

```toml {linenos=table,hl_lines=[],linenostart=1}
[FMTVariables]
SourceDir = '{{- .WorkingDir -}}/public'
```

The table indication is `[FMTVariables]`. All its values are `Key-Value` in
nature and **it shall not have any sub-definitions (e.g.
`[FMTVariables.MyDef]`)**.

Depending on CI-Jobs, Monteur can supply pre-defined variables into the
variables list for formatting processing.

The formatting clauses are complying to the Go’s
[text/template](https://pkg.go.dev/text/template) package where the variables'
placement is denoted by its `Key` with a period (`.`) prefix. If we use the
example above, the processing sequences will be as follows:

```toml {linenos=table,hl_lines=[2],linenostart=1}
Variables["WorkingDir"] = "/home/u0/project"          # pre-defined variables
Variables["SourceDir"] = "{{- .WorkingDir -}}/public" # parsed fmtVariables
Variables["SourceDir"] = "/home/u0/project/public"    # formatted fmtVariables
		⥥
	...finalized into...
		⥥
Variables["WorkingDir"] = "/home/u0/project"
Variables["SourceDir"] = "/home/u0/project/public"
```

Due to the `Key-Value` nature of the variables list, the order/position of the
Formattable Variables Definition is not parsed in-sequence with a guarenteed
orderly manner. Hence, **ONLY create FMTVariables independent of one another**.

Formattable Variables Definitions takes processed and plain variables
`Key-Value` list as its processing input. Hence, it is always parsed **AFTER
parsing the plain variables definitions**.

The timing of the variable formatting depends on the CI Job. Hence, please
refers to the CI Job's documentation for its execution sequences.




## Parsing Sequences
Montuer facilitates various layers of variables definition for allowing you to
practice "Don't Repeat Yourself" mantra while keeping maintenances sane and
easy. This section covers all the parsing sequences in detailed.

This section is written in a sequenced flow for easier reading and understanding
the train of thoughts.



### Level 1: Main Variables Definition
Monteur starts off by parsing the variables from the main configuration files
located in `.configs/monteur/workspace.toml`. This will serves as the foundation
of the variables.

The variables are visible globally across all [CI Jobs]({{< link "/ci-jobs/"
"this" "url-only" />}}).

Montuer shall do the following in sequences:
1. Parse the plain variables definitions



### Level 2: Job Variables Definition
Monteur then parse the variables from the job configuration files located in
`.configs/monteur/<THE-CI-JOB>/config.toml`. This will append any new found
variables into the list while overwrite any conflicing variables with the latest
values.

The variables are visible only to that specific CI Job's every execution units.

Montuer shall do the following in sequences:
1. Parse the plain variables definitions
2. Merge the processed formatting variables definitions into variables list



### Level 3: Execution Unit Variables Definition
Monteur then parse the variables from the execution unit configuration file
located inside the execution unit configuration file (e.g. for `Setup` Job, it
is `.configs/monteur/setup/programs/myPrograms.toml`). This will append any new
found variables into the list while overwrite any conflicing variables with the
latest values.

The variables are visible only to that one particular execution unit.

Montuer shall do the following in sequences:
1. Parse the plain variables definitions
2. Parse the formatting variables definitions
3. Process the formatting variables definitions using the existing variables
4. Merge the processed formatting variables definitions into variables list
5. Finalized the variables list for deployment usage




## Epilogue
That's all for Monteur's Variables Processing. If you have any queries, please
proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
