+++
date = "2021-11-26T07:02:07+08:00"
title = "Commands Execution Unit"
description = """
Monteur supplies its own Commands Execution Unit. To ensure its portability
across various platforms, most of the commands are natively implemented. This
section explains how Monteur operates its Commands Execution Unit.
"""
keywords = [
	"commands",
	"Monteur",
	"execution unit",
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
url = "/en/internals/commands/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur's Commands Execution Unit"

[thumbnails.1]
url = "/en/internals/commands/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur's Commands Execution Unit"

[thumbnails.2]
url = "/en/internals/commands/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur's Commands Execution Unit"


[menu.main]
parent = "Internals"
name = "Commands Execution Units"
pre = "🧑‍✈️"
weight = 5
identifier = "internals-commands-execution-units"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Data Structure
Monteur's Commands Execution Unit uses a set of variables arranged in an array
of instructions. Each unit has a standardized list of fields shown as follows:

```toml {linenos=table,hl_lines=[],linenostart=1}
[[CMD]]
Name = 'Get Publish Branch First Commit for Cleaning'
Condition = [ 'all-all' ]
Type = 'command'
Location = '{{- .WorkingDir -}}'
Source = 'git rev-list --max-parents=0 --abbrev-commit HEAD'
Target = ''
Save = 'FirstCommitID'
SaveRegex = '(\w*)'
SaveStderr = true
ToSTDOUT = '{{- FirstCommitID -}}'
ToSTDERR = 'Completed obtaining first commit ID'
```

The Commands Execution Unit definition is declared by the CI Job. Example, in
Setup, you get `[[Setup]]` instead of `[[CMD]]`. Please refers to the CI Job
documentation.

* `Name`
  1. Compulsory.
  2. Mainly for logging and referencing purposes.
* `Condition`
  1. Compulsory.
  2. the list of supported [Platform ID]({{< link
     "/internals/platform-identification/" "this" "url-only" />}}) where the
     command shall be executed when Monteur is operating on matching platform.
  3. Should `all-all` exists in the array list, it will overwrite all platform
     IDs, both specified and unspecified, as always `true`.
* `Type`
  1. Compulsory.
  2. The Monteur's Commands Execution ID.
* `Location`
  1. Optional.
  2. the directory location for executing the command.
  3. [Variables formatting]({{< link
     "/internals/variables-processing/#variables-formatting" "this"
     "url-only" />}}) is available for this field.
  4. When available, Monteur shall change into that directory before executing
     the command.
* `Source`
  1. Compulsory/Optional depending on Command's `Type`.
  2. usually taken as main input of the command. Please refers to the command's
     specification for its purpose.
  3. [Variables formatting]({{< link
     "/internals/variables-processing/#variables-formatting" "this"
     "url-only" />}}) is available for this field.
* `Target`
  1. Compulsory/Optional depending on Command's `Type`.
  2. usually taken as secondary input of the command or output. Please refers to
     the command's specification for its purpose.
  3. [Variables formatting]({{< link
     "/internals/variables-processing/#variables-formatting" "this"
     "url-only" />}}) is available for this field.
* `Save`
  1. Optional.
  2. The `Key` value for saving the output (e.g. `STDOUT`) of the executed
     commands into the [Variables]({{< link "/internals/variables-processing/"
     "this" "url-only" />}}) list for later use.
* `SaveRegex`
  1. Optional.
  2. A regular expression filter to process the data prior saving.
  3. Use capturing group in your regex using parenthesis `(...)`. Every captured
     groups shall be concatenated from left to right. Example:
     `(\w{1,}):(\s{1}).{1,}\((\d+.\d+%)\)` for `total:    statement  (50.0%)`
     yields `total 50.0%`.
  4. **DO NOT** include the regex slashes (as in `/ YOUR_REGEX /`).
  4. Only works if `Save` is enabled (Not empty) and itself is not empty.
* `SaveStderr`
  1. Optional.
  2. Instruct Monteur to save `STDERR` instead of the default `STDOUT` data.
  3. Only works if `Save` is enabled (not empty) and itself is set to `true`
     (boolean).
* `ToSTDOUT`, `ToSTDERR`
  1. Optional.
  2. Dump written string to `STDERR` and `STDOUT` respectively after the command
     is executed successfully.
  3. [Variables formatting]({{< link
     "/internals/variables-processing/#variables-formatting" "this"
     "url-only" />}}) is available for this field **including its own save
     variable**.
  4. No action will be taken respectively if `ToSTDOUT` and `ToSTDERR` are empty
     string (`""`) after formatting.




## Execution Sequences
This section explains Monteur's steps of executing each of its available
commands. There are 2 main activities:

1. The chain of commands are parsed and validated that its `Type` is supported
   by the runtime Monteur.
2. The executions of the chain of commands.


### Step 0: Validate Monteur's Supporting Commands
During the TOML data file parsing for chain of commands, Monteur will check
through each command's `Type` is within its supporting range. Otherwise, Monteur
shall raise an error without executing any of the commands.

Otherwise, Monteur will begin executing the chain of commands.



### Step 1: Process the Current Command
The Commands Execution Unit will then process `Location`, `Source`, and `Target`
with [Variables formatting]({{< link
"/internals/variables-processing/#variables-formatting" "this" "url-only" />}}).

At this point, should there be any processing errors (e.g. bad formatting
clauses), Monteur shall raise that error and the chain of commands is stopped.



### Step 2: Executing The Command
Once all validation and processing are completed, Monteur will then execute
the command and generate its output data (can be structural depending on command
`Type`).



### Step 3: Process Save
If `Save` is set, the generated output is then saved into the Variables List
based on the command `Type`. Should there be any existing variable with the same
`Key` value (with `Save`), Monteur shall overwrite the existing value with the
command output.



### Step 4: Repeat Step 1 for Next Command
The Step 1 process is then repeated for the next command in line. Otherwise,
the chain of commands is completed.




## Available Commands
Monteur shall continuously add new commands from time to time. These are the
currently available commands sorted by `Type`.
