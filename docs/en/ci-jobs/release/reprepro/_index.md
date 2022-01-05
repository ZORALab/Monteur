+++
date = "2021-12-23T08:10:56+08:00"
title = "Reprepro Release API"
description = """
Monteur supports Reprepro for building a local hosting apt repository that is
fully compatible with Debian apt ecosystem. This recipe allows repository
equipped with Reprepro to seamlessly release .deb packages in a highly
customizable manner easily and seamlessly.
"""
keywords = [
	"Reprepro",
	"Release CI Job",
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
url = "/en/ci-jobs/release/reprepro/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Release with Reprepro"

[thumbnails.1]
url = "/en/ci-jobs/release/reprepro/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Release with Reprepro"

[thumbnails.2]
url = "/en/ci-jobs/release/reprepro/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Release with Reprepro"


[menu.main]
parent = "Software - Reprepro"
name = "Release"
pre = "🌾"
weight = 90
identifier = "ci-jobs-release-reprepro"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to make upstreaming `.deb`
packages in a consistent manner with minimal to no further instructions**.

Reprepro instructions are available at:
https://wiki.debian.org/DebianRepository/SetupWithReprepro

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## One-time Setup
Reprepro requires a 1-time setup to meet its requirements. All you need to do is
creating the `conf/distributions` file in your `.DataDir` (data directory) meant
for reprepro. The content of the file are something as such:

```yaml {linenos=table,hl_lines=["1-8"],linenostart=1}
Origin: monteur.zoralab.com/releases/deb
Label: ZORALab's Monteur Deb Package (Main Branch)
Codename: main
Suite: stable
Architectures: amd64
Components: stable
Description: software manufacturing automation tool in one single app.
SignWith: hello@zoralab.com

Origin: monteur.zoralab.com/releases/deb
Label: ZORALab's Monteur Deb Package (Staging)
Codename: staging
Suite: unstable
Architectures: amd64
Components: unstable
Description: software manufacturing automation tool in one single app.
SignWith: hello@zoralab.com

Origin: monteur.zoralab.com/releases/deb
Label: ZORALab's Monteur Deb Package (Next)
Codename: next
Suite: experimental
Architectures: amd64
Components: experimental
Description: software manufacturing automation tool in one single app.
SignWith: hello@zoralab.com
```

* `Origin` - the URI of the source.list. This is where users add the download
  location into `/etc/apt/source.list.d/` directory.
* `Label` - name of the package series.
* `Codename` - controlled field. Denotes it to your branch name.
* `Suite` - stick to either `stable`, `unstable`, or `experimental`. Otherwise,
   if you release it to official OS, stick to their official release codenames.
* `Architectures` - supported architectures separated by space (` `). (E.x:
  `amd64 arm64 i386`). You may update this field over-time as you scale.
* `Components` - Distribution components. When in doubt, keep it to the same
  as `Suite`.
* `Description` - the abstract of the package. Keep it short and max 80
  characters.
* `SignWith` - your GPG signing key ID.

This will make user add your `sources.list` as:

```
deb https://{{- .Origin }} {{ .Codename }} {{ .Components }}

# example above:
deb https://monteur.zoralab.com/releases/deb main stable
```

And apt search will provide a pattern as:

```
{{ .AppName }}/{{- .Suite }} {{ .PkgVersion }} {{ .PkgArch }}
  {{ .Description }}

# Example
monteur/stable 0.0.1 amd64
  software manufacturing automation tool in one single app.
```



## Recipe Versions
Here are the available Release API recipe for Hugo integrations. Please read
through your selected version's details on what has changed, what is required,
and how to customize and used them.

The arrangement are the latest at the top or first.


### Version 2.0.0
Version 2.0.0 Reprepro API is available for download here:
{{< link "/ci-jobs/release/reprepro/reprepro-v2p0p0.toml" "this" "" ""
"button" "" "download" >}}
reprepro-v2p0p0.toml
{{< /link >}}

| Min Requirements     | Values                          |
|:---------------------|--------------------------------:|
| Monteur Version      | `v0.0.2`                        |
| Supported Platforms  | follows Reprepro's availability |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/release/jobs/` directory with the name `reprepro.toml`.
2. Once done, edit the configuration file for:
   1. `Variables.GPGID` - your GPG release singing key where private key is
      available.
   2. `FMTVariables.DataPath` - your reprepro data path. Make sure it has
      visibility to the `conf/distributions` file setup earlier.
   3. `Releases.Target` - your reprepro output directory.
   4. `Releases.Packages.XXX` - list of `.deb` packages. Duplicate the table
   with different label (`XXX`) if there are more `.deb` package variants to be
   released in the same repository.
   4. `Releases.Packages.XXX.OS` - the compatible target operating system.
   5. `Releases.Packages.XXX.Arch` - the compatible target cpu architecture.
   4. `Releases.Packages.XXX.Source` - the location of your `.deb` package file.
3. Add any additional `[[Dependencies]]` if you're customizing the default
  `[[CMD]]` commands list accordingly.
4. Customize `[[CMD]]` commands list as per your need.


#### Changes
1. **Non-Backward Compatible** - Formatted `Releases.Packages.XXX.Source` to use
   new variables for matching `debuild` overwritten package filename.
2. **Non-Backward Compatible** - Added `Releases.Packages.XXX.OS` and
   `Releases.Packages.XXX.Arch` as required by Monteur `v0.0.2` package meta.



### Version 1.0.0
Version 1.0.0 Reprepro API is available for download here:
{{< link "/ci-jobs/release/reprepro/reprepro-v1p0p0.toml" "this" "" ""
"button" "" "download" >}}
reprepro-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                          |
|:---------------------|--------------------------------:|
| Monteur Version      | `v0.0.1`                        |
| Supported Platforms  | follows Reprepro's availability |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/release/releasers/` directory with the name `reprepro.toml`.
2. Once done, edit the configuration file for:
   1. `Variables.GPGID` - your GPG release singing key where private key is
      available.
   2. `FMTVariables.DataPath` - your reprepro data path. Make sure it has
      visibility to the `conf/distributions` file setup earlier.
   3. `Releases.Target` - your reprepro output directory.
   4. `Releases.Packages.XXX.Source` - the location of your `.deb` package file.
3. Add any additional `[[Dependencies]]` if you're customizing the default
  `[[CMD]]` commands list accordingly.
4. Customize `[[CMD]]` commands list as per your need.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.




## Epilogue
That's all for Monteur operating Reprepro software. If you found a bug or have
any questions about the recipe, please feel free to raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
