+++
date = "2022-01-04T06:03:52+08:00"
title = "Init API"
description = """
To keep Monteur setup in your application seamlessly, Monteur has a dedicated
Init API just for the job. This section explains how to utilize this Init API
and setup your repository in a single go.
"""
keywords = [
	"init",
	"repo",
	"api",
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
url = "/en/internals/init-api/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Init API"

[thumbnails.1]
url = "/en/internals/init-api/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Init API"

[thumbnails.2]
url = "/en/internals/init-api/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Init API"


[menu.main]
parent = "Z) Monteur Internals"
name = "Init API"
pre = "🛫"
weight = 5
identifier = "internals-init-api"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Initialize Your Repository
To initialize your repository, **go to your repository's root directory**, and
simply issue the following command:

```bash {linenos=table,hl_lines=[],linenostart=1}
monteur init
```

Monteur shall properly scan for its existence and if it is absent, Monteur shall
do the rest of the config files creation automatically.

{{< note info "note" >}}
Init API is only available after Monteur version `v0.0.2`.
{{< /note >}}




## Fill in Your App Metadata
Once all files are created, it's time to update your application metadata once
will do. There are a few critical files to update:



### App `metadata.toml`
The main metadata TOML configuration file is located in
`.configs/monteur/app/config/en/metadata.toml`. This file houses all the core
metadata of your application that will be used across many packagers and
releasers system.

Among the information you need to update when appropriate are:

```toml {linenos=table,hl_lines=["2-13", "16-17", "22-26", "31-35", "40-44"],linenostart=1}
[Software]
Name = '<directory Name in title case>'
Command = '<directory name in filepath format>'
ID = '<directory name in filepath format>'
Version = 'v0.0.1' # version number
Category = 'devel' # software category
Suite = ''         # software suite (e.g. MSFT Office for MSFT Word)
Abstract = 'software XYZ tool in one single app.'
Description = """
Describe your software as much as you want. Please make sure you end it with
its unique selling pitch.
"""
Website = 'https://example.com/'

[Software.Contact]
Name = 'ACME Entity'
Email = [ 'hello@example.com' ]




[Software.Maintainers.ACME]
Name = 'ACME Entity'
Email = [ 'tech@example.com' ]
[Software.Maintainers.ACME.JointTime]
Year = '2021'




[Software.Contributors.ACME]
Name = 'ACME Entity'
Email = [ 'tech@example.com' ]
[Software.Contributors.ACME.JointTime]
Year = '2021'




[Software.Sponsors.ACME]
Name = 'ACME Entity'
Email = [ 'tech@example.com' ]
[Software.Sponsors.ACME.JointTime]
Year = '2021'
```

* `Software.Name` and `Software.ID` - Monteur automatically takes your directory
  name as the application name. Hence, if you desired a different naming, please
  edit accordingly. For `Software.ID`, it **MUST** be in alphanumeric and dash
  (`-`) only.
* `Software.Command` - the command name to call your application.
* `Software.Version` - controls the version of your app.
* `Software.Category` - the common category of your software (e.g. `devel`,
  `games`, ...)
* `Software.Suite` - the sub-component of a bigger software bundle like
  `Microsoft Office` if the app is `Microsoft Word`.
* `Abstract` - a short, no article ("a", "an", "the") description
   (max 70 characters).
* `Description` - long description of your app. It shall continues from the
  `Abstract`.
* `Website` - the official website of your app.
* `Software.Contact.Name` - the general contactable entity (person or org)'s
   name.
* `Software.Contact.Email` - the general contactable entity's email.
* `Software.[TYPE].[ID]` - These are the list of contributors contacts that will
  be used in credits and citation. You can add more entries by duplicating the
  section with a different and unique `[ID]` of the `[Type]`. The `[ID]` (in the
  above example: `ACME`) is mainly used for uniquely identify an entity across
  all lists. You can name it however you want except using period (`.`). For
  `[TYPE]`, these are the currently supported categories:
  * `Maintainers` - the current maintainers who has executive roles on app
    development and releases.
  * `Contributors` - the minor contributors who contribute code changes into the
    repository.
  * `Sponsors` - the financial contributors to the app development and releases.
* `Software.[Type].[ID].Name` - the entity's legal name.
* `Software.[Type].[ID].Email` - the entity's reachable email address.
* `Software.[Type].[ID].JointTime.Year` - the year that the entity joint.

{{< note info Info >}}
`Software.Contact` is used for general communications while
`Software.[Type].[ID]` are mainly used for hall of fames and legal binding
documentations such as automated license generations.
{{< /note >}}



### App `help.toml`
Although not required, you should look into the `help.toml` which contains all
the general help guides for various help documentations like man pages and etc.
The default is good enough to automate most of the known documentations. If you
are not using it, you can leave it as it is.

The file is located in `.configs/monteur/app/config/en/help.toml`. It has the
following template that will be automatically formatted based on your app
`metadata.toml`.

```toml {linenos=table,hl_lines=[],linenostart=1}
[Help]
Command = '$ {{ .App.Command }} help'
Description = '{{ .App.Abstract }}'
Resources = 'Visit {{ .App.Website }} for detailed documentations.'

[Help.Manpage]
Lv1 = """
.\\" {{ .App.Name }} - Lv1 Manpage
.\\" Contact {{ index .App.Contact.Email 0 }} for errors or typos.
.TH man 1 "{{- .App.Time.Day }} {{ .App.Time.Month }} {{ .App.Time.Year -}}" "{{- .App.Version -}}" "{{- .App.ID }} man page"

.SH NAME
{{ .App.Name }} - {{ .App.Abstract }}

.SH SYNOPSIS
{{ .App.Help.Command }}

.SH DESCRIPTION
{{ .App.Description -}}

.SH OPTIONS
{{ .App.Help.Description }}

.SH SEE ALSO
{{ .App.Help.Resources }}

#.SH AUTHORS
#{{ range $index, $owner := .App.Maintainers -}}
#       {{- $owner.Name }} ({{- index $owner.Email 0 -}})
#{{ end -}}
"""
```

* `Help.Command` - the command with argument for triggering the program's help.
* `Help.Description` - the abstract of the your app.
* `Help.Resources` - the statement for finding more help resources.
* `Help.Manpage.[KEY]` - the content for generating the different levels of
  manpage files. The `[KEY]` shall be used as the name or extension of the
  manpage depending on the packager.



### App `debian.toml`
These are Debian's `.deb` specific app metadata for packaging your app to the
Debian-based operating system. You may need to refer
[Debian Policy Specification](https://www.debian.org/doc/debian-policy/) for the
fields' descriptions.

The configuration file is located in
`.configs/monteur/app/config/en/debian.toml` and it is currently providing the
following template:

```toml {linenos=table,hl_lines=["35", "58-62", "64"],linenostart=1}
[DEB]
Compat = 11
Rules = """
#!/usr/bin/make -f

# Uncomment this to turn on verbose mode.
#export DH_VERBOSE=1

%:
	dh $@

override_dh_auto_build:
	echo "nothing to override"
"""

[DEB.Control]
Essential = false
PackageType = 'deb'
Priority = 'optional'
RulesRequiresRoot = 'binary-targets'
Standards = '4.6.0'
Section = 'devel'

[DEB.Relationships]
'Build-Depends' = [
	'debhelper (>= 11)',
]
'Depends' = [
]

# More Info:
#  https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-vcs-fields
[DEB.VCS]
Type = 'Vcs-Git'
URL = 'https://example.com/git'
Branch = 'main'
#Path = '.'

[DEB.Testsuite]
Paths = [
	# 'relative/path/to/debTestScript',
]

[DEB.Copyright]
Format = 'https://www.debian.org/doc/packaging-manuals/copyright-format/1.0/'
Disclaimer = ''
Comment = ''

[DEB.Changelog]
Urgency = 'low'

[DEB.Source]
Format = '3.0 (native)'
LocalOptions = """
"""
Options = """
"""
LintianOverrides = """
# supply the license file in case OS did not supply one

# we are go package so it is okay not to specify Depends:
"""

[DEB.Install]
```

Depending on which packager you're using (e.g. `debuild`), certain information
are compulsory. Monteur itself is using this packager most of the settings are
up-to-date.

Among the attention needed from you are:

* `DEB.VCS.URL` - change it into your git URL.
* `DEB.Source.LintianOverrides` - add any lintian overrides that you found false
  positive (can be added later once you established your repository.
* `DEB.Install` - The list of installed files in `key:value` format. The `key`
  shall designate the directory path for installing while the value is the
  assembled content (executable, files, etc). You can add more if you have more
  files in it.
  * Example: `'usr/bin' = 'monteur'` where the `monteur` program is installed
    into `usr/bin` directory of the target operating system.



### App Copyright `<license>.toml`
Lastly, you should update your app copyright's license metadata. This file holds
a license's full metadata for generating various data files across the
repository. A default license TOML file was generated in
`.configs/monteur/app/config/en/copyrights/copyrights.toml`. Please rename it
to the appropriate license filename like `apache-2.toml` for
[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) license.

Since app can be mutli-licensed (e.g. Creative Commons for non-coding materials
and Code Licensing), you can create as many `<license>.toml` file as you deemed
appropriate.

The default TOML license file has the following template:
```toml {linenos=table,hl_lines=[],linenostart=1}
[Copyright]
Name = 'License Name' # license full name
ID = 'license-name'   # identifiable license ID
Comment = ''
Materials = [ '*' ]
Holders = [
        # A string with year, then name, then email with diamond braces. Ex.:
        #
        #   '2021 John S. Smith <john.s.smith@email.com>',
        #
        # If the list is empty, all maintainers, contributors, and sponsors
        # shall be attached to it.
]
Notice = """
License notice usually prepend in front of source codes.
"""
Text = """
Full license text body.
"""
```

* `Copyright.Name` - Name of the license.
* `Copyright.ID` - The unique ID of the license. Make you use common identifier
  like the [SPDX Licenses List](https://spdx.org/licenses/).
* `Comment` - optional comments for the license.
* `Materials` - the materials where the license is applied. Use asterisk (`*`)
  to indicate all. For example, to indicate all source codes with extension
  `.go`, it is `*.go` across the repository.
* `Holders` - the optional list of copyright holders. When left empty, all
  contributors from the app's `metadata.toml` shall be sorted out and listed as
  the copyright holders. This field is only meant for overwrite said list for
  special cases' license.
* `Notice` - The short copyright notice that is usually appened in front of the
  source code.
* `Text` - The full text body of the license file.




## Build Your Repository's CI Jobs
Now that you have your app metadata setup, it's time to setup your Monteur
CI Jobs recipes. Different CI Jobs has different recipe files but they are
following the same filesystem which are housing under
`.configs/monteur/${ci-job}/jobs/` directory.

Each recipe has its own configurations so please go through them accordingly.
The initialization shall be completed once you had fully configured all the
recipes to your app building.

To start, please head over to:

{{< link "/ci-jobs" "this" "" "" "button" >}}
CI Jobs Catalog
{{< /link >}}




## Epilogue
That's all for Monteur's Init API. If you have any queries,
please proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
