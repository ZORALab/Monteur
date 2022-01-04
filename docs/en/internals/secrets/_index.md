+++
date = "2022-01-04T17:06:51+08:00"
title = "Managing Secrets"
description = """
Monteur allows secrets data to be used inside its CI Jobs with high
customizations and known safety controls. This section explains how to use
secret or private data in details.
"""
keywords = [
	"manage",
	"secrets",
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
url = "/en/internals/secrets/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Secret Management"

[thumbnails.1]
url = "/en/internals/secrets/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Secret Management"

[thumbnails.2]
url = "/en/internals/secrets/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Secret Management"


[menu.main]
parent = "Z) Monteur Internals"
name = "Managing Secrets"
pre = "🙊"
weight = 5
identifier = "internals-secrets"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

This function is available since Monteur `v0.0.2`.




## Customizable Pathing with Overriding Mechanism
Monteur secrets management allows you to fully customize your secrets
configurations data files from multiple directories. They are listed inside
`Filesystem.SecretsDir` array field in `.configs/monteur/workspace.toml`.

Here is an example:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Filesystem]
...
SecretsDir = [
        '{{ .HomeDir }}/.secrets',
        '{{ .RootDir }}/.configs/monteur/secrets',
]
```

[Formatting with Variables ability]({{< link "/internals/variables-processing/"
"this" "url-only" />}}) is available for all the paths but **is strictly limited
to all the `[Filesystem]` fields except `SecretsDir`**.

The positioning of the directories is important as the latter's data fields
**shall overwrites the former**. If we use the example above:

1. If `{{ .HomeDir }}/.secrets` has a field called `App.Color = 'Green'`
2. And `{{ .RootDir }}/.configs/monteur/secrets` has the same field called
   `App.Color = 'Red'`
3. The final queriable output for `App.Color` is `Red`.
4. Should there be 2 files in the same directory each having `App.Color` with
   different values, the output is unpredictable. Hence, please doing such
   action.

File extension **is strictly required for format identification purposes**.
Monteur currently supports the following configuration file formats:

1. `TOML` (`.toml`) - main configuration file.

The overriding mechanism allows one to re-use default secrets data and overrides
required fields with a separate one from a different location.




## Logging Protections
Note that all Monteur's output to all `STDOUT`, `STDERR`, and log files are
always filtered with all the secrets' value to redact them for secrecy
protection. The values are directly redacted using `strings.ReplaceAll`
function in Go with the redaction label.

The source code for the filtration is located in `liblog` internal package where
the first thing to do in any print function is to filter out secret data before
any output, regardless the logging package supplied by 3rd party vendors:

1. `next` - https://gitlab.com/zoralab/monteur/-/blob/next/gopkg/monteur/internal/liblog/Logger.go
2. `staging` - https://gitlab.com/zoralab/monteur/-/blob/staging/gopkg/monteur/internal/liblog/Logger.go
3. `main` - https://gitlab.com/zoralab/monteur/-/blob/main/gopkg/monteur/internal/liblog/Logger.go

### Caveat
While we do all our best to secure your secret data, **Monteur CANNOT redact
processed secret data (E.g. regex extraction via `[CMD]` and etc)** and we
really **DO NOT want to complicate the filter** function to the point of slowing
down the entire Monteur performance.

There is no way for Monteur developer to guess all the permutations and
combinations of Monteur cases. Hence, YOU HAVE BEEN WARNED that:

1. **ONLY use Monteur's secrets feature for direct data insertion, NOT for
   data processing**!
2. **DO NOT share your raw log files**!
3. In any cases of sharing or reporting to upstream, **please peer review your
   log files with your colleague before sharing out to anyone**. If possible,
   only expose only the business-need-to-know log data fragment instead of the
   entire file.
4. Please don't abuse the secrecy feature. Make a feature request in our
   [Issues Section](https://gitlab.com/zoralab/monteur/-/issues) if you found
   something useful but not wanting to deal with sensitive data.


#### Dumb Redaction
Should your secret data becomes overly common and simple (e.g. `0`), you will
notice that Monteur will redact all `0` from your log and stdout output in a
dumb and blind manner.

We're still working on how to smartly and securely redact sensitive information.
As for now, we're happy to redact first rather than being really sorry and
regretful.




## Querying Secret Data
Monteur flattens the Secret data into a single and simple `key:value` data
table. To query the data, when a field has [Formattable Variables]({{< link
"/internals/variables-processing/" "this" "url-only" />}}) ability, you can
query it using `GetSecret` function, such that:

```toml {linenos=table,hl_lines=[],linenostart=1}
Authentication = 'token {{ GetSecret "Github.Token" -}}'
```



#### Flattened Data Structure
Monteur intentionally flattens the secret data structure to prioritize output
filtering performance. You can still query the data using the dot (`.`)
connection. Example, for the following data structure:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Sample]
Type = 'squirrel'

[[Sample.Favourites]]
Food = [
	"Apple",
	"Pineapple",
]

[[Sample.Hates]]
Food = [
	"Beef",
	"Shrimp",
]

[[Sample.Favourites]]
Toy = [
	"Tree",
	"Honey Feeder",
]
```

The queries string (left side) would be:

```txt {linenos=table,hl_lines=[],linenostart=1}
"Sample.Type"                 = 'squirrel'
"Sample.Favourites.0.Foods.0" = 'Apple'
"Sample.Favourites.0.Foods.1" = 'Pineapple'
"Sample.Favourites.1.Toy.0"   = 'Tree'
"Sample.Favourites.1.Toy.1"   = 'Honey Feeder'
"Sample.Favourites.1.Foods.0" = '<no data>'
"Sample.Hates.0.Foods.0"      = 'Beef'
"Sample.Hates.0.Foods.1"      = 'Shrimp'
```

However, the values retain its natural data type that can be any of the
following:

1. `string`
2. `int`, `uint`
3. `float64` / `float32` depending on your host CPU architecture.
4. `bool`

Obviously list types (`map` and `arrays`) got flatten out.

Should any invalid query appears, the string `<no value>` shall appear as
replacement.




## Security Reviews Assistances
For those that wants to review Monteur's source codes, you can review the
following internal packages:

1. `gopkg/monteur/internal/secrets` - 3rd-party vendor parser.
2. `gopkg/monteur/internal/libsecrets` - the interfacing internal package
   between 3rd-party packages and monteur.
3. `gopkg/monteur/internal/libworkspace` - where the secrets is initialized.
4. `gopkg/monteur/apiCommand.go` - where the secrets are being shifted.
5. `gopkg/monteur/internal/libcmd` - where the secrets are used to create job
   loggers.

By default and always, only `libsecrets.Secrets` is created as a `struct`
pointer and being passed around. The data is safely stored privately inside the
structure and can only be quried using its `Query` and `Filter` methods.

Monteur welcomes new reviews and suggestions to strengthen its security for the
benefits of all.




## Epilogue
That's all for Monteur's secrets data management. If you have any queries,
please proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
