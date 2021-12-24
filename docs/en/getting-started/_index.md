+++
date = "2021-12-24T00:11:52+08:00"
title = "Getting Started"
description = """
Monteur can be quite intimidating to use if left alone in the dark. Hence, we
put up a getting started guide for you to begin with. This section guides you
on how to setup Monteur and using it for your repository.
"""
keywords = [
	"releases",
	"monteur",
	"configurations",
	"documentation",
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
url = "/en/getting-started/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Getting Started with Monteur"

[thumbnails.1]
url = "/en/getting-started/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Getting Started with Monteur"

[thumbnails.2]
url = "/en/getting-started/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Getting Started with Monteur"


[menu.main]
parent = ""
name = "Getting Started"
pre = "🛫"
weight = 5
identifier = "getting-started"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Installation
To install Monteur into your operating system, please find your operating system
guides and execute its instructions:



### Debian / Ubuntu
A dedicated apt repository was setup for Debian-based operating system. While
in `root` permission (as in `sudo` all the instructions here):

```bash {linenos=table,hl_lines=[],linenostart=1}
# get ZORALab public GPG Keys (skip if you done this before with other products)
curl https://www.zoralab.com/pubkey.gpg \
	| gpg --yes --dearmor --output /usr/share/keyrings/zoralab-keyring.gpg

# write source list file (choose either 'stable', 'unstable', 'experimental')
# Here, 'stable' was chosen.
echo 'deb [arch=amd64 signed-by=/usr/share/keyrings/zoralab-keyring.gpg] https://monteur.zoralab.com/releases/deb stable main' \
	| sudo tee /etc/apt/sources.list.d/zoralab-monteur.list

# perform apt update
apt update -y

# install monteur
apt install montuer -y

# exit root account (without 'sudo') and test out monteur
monteur help
```

The `apt` repository is always packed with the latest versions. You can check
them out in our
[Releases Section]({{< link "/versions/" "this" "url-only" />}}). Should any
later version is available, it will be rolled-out via the `apt update` and
`apt upgrade`.




## Setup Monteur Configuration Directories
{{< note note "For Your Info" >}}
A new API `monteur install` is currently under development to simplify this
section. See [Issue 28](https://gitlab.com/zoralab/monteur/-/issues/28).
{{< /note >}}


### Setup workspace configurations
Once monteur is available, in your repository, you will setup monteur
configurations files. While inside your root repository, here are the
instructions (presented in Linux BASH):

```bash {linenos=table,hl_lines=[],linenostart=1}
# create .configs/monteur directory
mkdir -p .configs/monteur

# create the workspace.toml
# (These are recomemnded defaults. Customize accordingly.)
echo "[Filesystem]
BaseDir = 'src/' # for jobs test, build, etc. Set to '.' if unused.
WorkingDir = '.monteurFS/tmp/'
BuildDir = '.monteurFS/build/'
ScriptDir = '.scripts/'
BinDir = '.monteurFS/bin/'
BinCfgDir = '.monteurFS/config'
LogDir = '.monteurFS/log'
DataDir = '.configs/monteur/app/data'
ReleaseDir = 'docs/.static/releases'
SecretsDir = [
        '{{ .HomeDir }}/.secrets',
        '{{ .RootDir }}/.configs/monteur/secrets',
]

[Language]
Name = 'English'
Code = 'en'

[Variables]
Placeholder = 'en'

[FMTVariables]
FMTPlaceholer = 'Language: \"{{- .Placeholder -}}\"'
" > .configs/monteur/workspace.toml
```

### Create API directories
Now that Monteur can recongize its primary configuration directories, it's time
to setup those API configuration directories. While inside your root repository,
do the following (presented in Linux BASH):

```bash {linenos=table,hl_lines=[],linenostart=1}
# make all API directories
mkdir -p .configs/monteur/{build,clean,compose,package,publish,release,setup}

# create the config.toml file for each of them
echo '[Variables]

[FMTVariables]' > .configs/monteur/{build,clean,compose,package,publish,release,setup}/config.toml

# create their recipes directories
mkdir -p .configs/monteur/{clean,compose,package,publish,release}ers
mkdir -p .configs/monteur/setup/programs

# create the app metadata configuration directories
# NOTE: the language 'en' **MUST** match the workspace.toml's Language.Code
# value.
mkdir -p .configs/monteur/app/config/en/copyrights
```



### Create App Metadata Directories
This is to hold and publish app metadata as a whole and across each publications
channel. That way, maintainer only needs to maintain a single copy and monteur
will assemble the metadata based on the publisher.


#### Create App `metadata.toml`
Create the metadata.toml content into
`.configs/monteur/app/config/<Language.Code>/metadata.toml`. This is a main
template from Monteur so please customize accordingly.

```toml {linenos=table,hl_lines=[],linenostart=1}
[Software]
Name = 'Monteur'
Command = 'monteur'
ID = 'monteur'
Version = 'v0.0.1'
Category = 'devel'
Suite = ''
Abstract = 'software manufacturing automation tool in one single app.'
Description = """
Monteur focuses on keeping the software development repository's continuous
improvements execution seamlessly scalable with great consistency.

Monteur - Getting the job done locally and remotely at scale!
"""
Website = 'https://monteur.zoralab.com/'



[Software.Contact]
Name = 'ZORALab'
Email = [ 'hello@zoralab.com' ]




[Software.Maintainers.zoralab]
Name = 'ZORALab Enterprise'
Email = [ 'tech@zoralab.com' ]
[Software.Maintainers.zoralab.JointTime]
Year = '2021'




[Software.Contributors.zoralab]
Name = 'ZORALab Enterprise'
Email = [ 'tech@zoralab.com' ]
[Software.Contributors.zoralab.JointTime]
Year = '2021'




[Software.Sponsors.zoralab]
Name = 'ZORALab Enterprise'
Email = [ 'tech@zoralab.com' ]
[Software.Sponsors.zoralab.JointTime]
Year = '2021'
' > .configs/monteur/app/configs/en/metadata.toml
```



#### Create App `help.toml`
Create the metadata.toml content into
`.configs/monteur/app/config/<Language.Code>/help.toml`. This is a main
template from Monteur so please customize accordingly.

```toml {linenos=table,hl_lines=[],linenostart=1}
[Help]
Command = '$ {{ .App.Command }} help'
Description = 'Use the program help command to get all the assistances'
Resources = 'Visit https://monteur.zoralab.com for detailed documentations.'

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

#### Create App `debian.toml`
Create the Debian specific app metadata content into
`.configs/monteur/app/config/<Language.Code>/debian.toml`. This is a main
template from Monteur so please customize accordingly.

```toml {linenos=table,hl_lines=[],linenostart=1}
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
URL = 'https://gitlab.com/zoralab/monteur'
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
monteur: copyright-not-using-common-license-for-apache2
monteur: copyright-file-contains-full-apache-2-license

# we are go package so it is okay not to specify Depends:
monteur: undeclared-elf-prerequisites
"""

[DEB.Install]
'usr/bin' = 'monteur'
```

#### Create App Copyright `<license>.toml`
Lastly, create all of your App licenses config data files into
`.configs/monteur/app/config/<Language.Code>/copyrights/` directory. It is
recommended to name the file as the license ID file (e.g. `apache-2.toml` for
`Apache 2.0 License`). You can have many config files if the project has various
licenses.

This is a template from Monteur for a `license.toml` so please customize
accordingly.
```toml {linenos=table,hl_lines=[],linenostart=1}
[Copyright]
Name = 'Apache 2.0'     # name of the license
ID = 'Apache-2'         # generally identifiable license ID
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
SHORT LICENSE NOTICE HERE
"""
Text = """
FULL LICENSE BODY HERE
"""
```



## Adding Your Recipes
Now that everything is done, you can proceed to add your development recipes
either by developing your own or use any of our existing templates. Monteur
maintains a number of recipes based on the CI Jobs and each has its own one-time
installation instructions. When you completed the recipe and have monteur tested
out working fine, you're generally done with Monteur. Please feel free to
explore the maintained templates here:

{{< align center middle >}}
{{< link "/ci-jobs/" "this" "" "" "button" >}}
CI Jobs Templates
{{< /link >}}
{{< /align >}}
