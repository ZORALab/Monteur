+++
date = "2022-01-04T10:49:21+08:00"
title = "Package Meta Processing"
description = """
Monteur has a unilateral manner to process package metadata for certain CI Jobs.
Different jobs use different fields from the Package Meta. Hence, this section
explains the meta in details.
"""
keywords = [
	"package",
	"meta",
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
url = "/en/internals/package-meta/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Internal Package Meta Processing"

[thumbnails.1]
url = "/en/internals/package-meta/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Internal Package Meta Processing"

[thumbnails.2]
url = "/en/internals/package-meta/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Internal Package Meta Processing"


[menu.main]
parent = "Z) Monteur Internals"
name = "Package Meta Processing"
pre = "📋"
weight = 5
identifier = "internals-package-meta"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}




## Data Structure
Monteur processes all packages' metadata using a single unified data structure
across various CI Jobs.

An example (fully constructed) using Monteur own configurations is shown as
follows:

```toml {linenos=table,hl_lines=[],linenostart=1}
[Packages.XXX]
OS = [ 'linux' ]
Arch = [ 'amd64' ]
Name = '{{- .PkgName -}}-{{- .PkgVersion -}}-{{- .PkgOS -}}-{{- .PkgArch -}}'
Distribution = [
        'stable',
]
Changelog = '{{- .DataDir -}}/debian/changelog-{{- .PkgArch -}}'
BuildSource = false
Source = '{{- .PackageDir -}}/targz/{{- .PkgName -}}-{{- .PkgVersion -}}-{{- .PkgOS -}}-{{- .PkgArch -}}.tar.gz'
Target = '{{- .ReleaseDir -}}/archives'

[Packages.001.Files]
'{{- .PackageDir -}}/monteur' = '{{- .BuildDir -}}/{{- .PkgOS -}}-{{- .PkgArch -}}'
'{{- .PackageDir -}}/License.pdf' = '{{- .LicensePath -}}'
```

* `[Packages.XXX]` - the label of the package. You can label it with any name
  (replacing `XXX`) as long as it does not use period (`.`). Monteur recommends
  using the [Platform Identification]({{< link
  "/internals/platform-identification/" "this" "url-only" />}}) to keep things
  sane.
* `OS` - the list of supported operating system. The **FIRST (1st)** shall be
  used for single value filling.
* `Arch` - the list of supported CPU architecture. The **FIRST (1st)** shall be
  used for single value filling.
* `Name` - the package name without file extension.
    * [Formattable Variables]({{< link "/internals/variables-processing/#formattable-variables-definition" "this" "url-only" />}}) are available for dynamic formatting.
    * Depending on the packager program, this field can be overwritten (e.g.
      `debuild`).
* `Changelog` - the filepath location for prepending new changelog entries data.
    * [Formattable Variables]({{< link "/internals/variables-processing/#formattable-variables-definition" "this" "url-only" />}}) are available for dynamic formatting.
* `Distribution` - the supported distributions of the operating system (E.g.
  `stable`, `unstable`, `experimental`, `debian`, `ubuntu`, ...).
    * When in doubts, sticks to `stable`, `unstable`, or `experimental`.
* `BuildSource` - the decision to package source codes instead of the compiled
  binary programs for supported packagers.
* `[Package.XXX.Files]` - the list of files to be copied over during
  package preparations stage (right before executing `[CMD]`).
    * The `Key` is the destination while the `Value` is the source of the file.
    * [Formattable Variables]({{< link "/internals/variables-processing/#formattable-variables-definition" "this" "url-only" />}}) are available for dynamic formatting.
    * If we follows the example above, `{{- .BuildDir -}}/linux-amd64` shall be
      copied to `{{- .PackageDir -}}/monteur`.
* `Source` - the packaged filepath.
    * Introduced and used in [Release]({{< link "/ci-jobs/release/" "this"
      "url-only" />}}) to locate the package file.
* `Target` - the release destination directory.
    * Introduced and used in [Release]({{< link "/ci-jobs/release/" "this"
      "url-only" />}}) to house the released contents.

Depending on the CI Jobs, some fields **ARE COMPULSORY** while others are
optional. Please refers to the CI Job's documentations for their specific usage.




## Known CI-Jobs Deployments
Currently, Monteur deploys this package data to the following CI Jobs:

1. [Package]({{< link "/ci-jobs/package/" "this" "url-only" />}}) - for
   packaging your app.
2. [Release]({{< link "/ci-jobs/release/" "this" "url-only" />}}) - for
   releasing same type packages under a single release channel.
3. [Prepare]({{< link "/ci-jobs/prepare/" "this" "url-only" />}}) - for
   releasing same type packages under a single release channel.
   1. Available since Monteur `v0.0.2`.




## Design
The design was to ensure that each job are oriented towards the operator type
rather than the output package files based on hands-on experience.

*Case Study*:

A build variant can be packaged into various different packages for different
release channels:
 * `linux-amd64` ➤ `.appimage`, `.deb`, and `.tar.gz`
    * That's a total of 3 jobs with same package metadata (`total = 1 * 3 = 3`)
 * Image you now got a list of build variants like `linux-arm64`, `linux-arm`,
   `darwin-amd64`, `darwin-arm64`, `windows-amd64`, and `windows-arm64`,
   that would have a total of `18` jobs (`6 * 3`) sharing the same package
   metadata.
    * That's a lot of duplications!

This means that a build variant needs to be both vertically and horizontally
scaled in Monteur. After thorough testing between Monteur `v0.0.1` and Monteur
`v0.0.2`, the orientation against operator type (e.g. `appimage`,
`deb`, `targz`)  is more managable and maintainable than many output oriented
configuration files.




## Epilogue
That's all for Monteur's Platform Identification. If you have any queries,
please proceed to contact us via our
[Issues Section](https://gitlab.com/zoralab/monteur/-/issues) channel.
