+++
date = "2021-12-23T09:12:23+08:00"
title = "Debuild Package API"
description = """
Monteur supports `debuild` for packaging `.deb` packages across various app
variants that are compliant to the debian apt ecosystem. This recipe allows
repository equipped with debuild to seamlessly package `.deb` packages in a
highly customizable manner easily and seamlessly.
"""
keywords = [
	"Debuild",
	"Package CI Job",
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
url = "/en/ci-jobs/package/debuild/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Package with Debuild"

[thumbnails.1]
url = "/en/ci-jobs/package/debuild/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Package with Debuild"

[thumbnails.2]
url = "/en/ci-jobs/package/debuild/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Package with Debuild"


[menu.main]
parent = "Software - Debuild"
name = "Package"
pre = "📦"
weight = 80
identifier = "ci-jobs-package-debuild"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to make building `.deb`
packages painlessly in a consistent manner with minimal to no further
instructions**.

Debuild instructions are available at:
https://wiki.debian.org/BuildingTutorial

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## One Time Setup
There are a number of operating system setup based on your host Debian operating
system for cross-packaging.



### Cross-Build Tools
Starting from Debian Strecth onwards, you need to install your cross-build
essential packages to perform cross-compile and cross-packaging. These
cross-build packages are architecture specifics:

```bash {linenos=table,hl_lines=[],linenostart=1}
apt install crossbuild-essential-armhf
apt install crossbuild-essential-aarch64
...
```



### Overrides `dh_auto_build` and `dh_shlibdeps` When Applicable
In most cases where Monteur built the binary and its dependency upfront for
packaging, you might need to override `dh_auto_build` and `dh_shlibdeps`
sequences depending on what you're building. The overrides is available under
`<config>/monteur/app/config/<LANG>/debian.toml` field `DEB.Rules`.

Example: for a Go binary application which is a single static binary file
packaged into binary `.deb` package, you need to set the `DEB.Rules` to:

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
        echo "nothing to build"
override_dh_shlibdeps:
        echo "nothing to depend on"
"""
```




## Recipe Versions
Here are the available Package API recipe for `debuild` integrations. Please
read through your selected version's details on what has changed, what is
required, and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 2.0.0
Version 2.0.0 `debuild` Package API is available for download here:
{{< link "/ci-jobs/package/debuild/debuild-v2p0p0.toml" "this" "" "" "button"
	"" "download" >}}
debuild-v2p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.2`                         |
| Supported Platforms  | follows `debuild`'s availability |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/package/jobs/` directory with the name `debuild.toml`.
2. Once done, edit the configuration file for:
   1. `Variables.GPGID` - your GPG release singing key where private key is
      available.
   2. `Packages.XXX` - your .deb packages list item. Duplicate if there are more
      build variants.
   3. `Packages.XXX.OS` - list of supported operating system for this `.deb`
      package.
   4. `Packages.XXX.Arch` - list of supported CPU architecture for this `.deb`
      package.
   5. `Packages.XXX.Changelog` - the filepath for sourcing changelog data.
   6. `Packages.XXX.Distribution` - an array of supported distributions.
   7. `Packages.XXX.BuildSource` - instruct Monteur to package a source code
       pack or a binary pack (not both). If both are needed, duplicate
       the recipe since they are essentially 2 different packages.
   8. `[Packages.XXX.Files]` - the list of files to be assembled by Monteur
       for packaging.

For detailed information about each fields, visit:
[Package Specification Data Structure]({{< link
"/ci-jobs/package/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. **Non-Backward Compatible** - Removed `[Changelog]` as Monteur `v0.0.2` now
   ports it to [DEB Changelog Prepare API]({{< link
   "/ci-jobs/prepare/changelog-deb/" "this" "url-only" />}}).
2. *Backward Compatible* - Applied FMTVariables to many Packages fields instead
   of hardcoding to specific ones.



### Version 1.0.0
Version 1.0.0 `debuild` Package API is available for download here:
{{< link "/ci-jobs/package/debuild/debuild-v1p0p0.toml" "this" "" "" "button"
	"" "download" >}}
debuild-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.1`                         |
| Supported Platforms  | follows `debuild`'s availability |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/package/packagers/` directory with the name `debuild.toml`.
2. Once done, edit the configuration file for:
   1. `Variables.ChangelogFrom` - track your git changelog for your source
       branch.
   2. `Variables.ChangelogTo` - track your git changelog with the destination
       branch.
   3. `Variables.GPGID` - your GPG release singing key where private key is
      available.
   4. `Changelog.LineBreak` - your changelog line break characters.
   5. `Changelog.Regex` - optionally, if you need to filter each lines.
   6. `Packages.XXX` - your .deb packages list item. Duplicate if there are more
      build variants.
   7. `Packages.XXX.OS` - list of supported operating system for this `.deb`
      package.
   8. `Packages.XXX.Arch` - list of supported CPU architecture for this `.deb`
      package.
   9. `Packages.XXX.Changelog` - the filepath for updating and saving changelog.
   10. `Packages.XXX.Distribution` - an array of supported distributions.
   11. `Packages.XXX.BuildSource` - instruct Monteur to package a source code
       pack or a binary pack (not both). If both are needed, duplicate
       the recipe since they are essentially 2 different packages.
   12. `[Packages.XXX.Files]` - the list of files to be assembled by Monteur
       for packaging.

For detailed information about each fields, visit:
[Package Specification Data Structure]({{< link
"/ci-jobs/package/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.
3. *Backward Compatible* - Uses `git` to generate changelog entries between
   source and destination branches.




## Epilogue
That's all for Monteur operating `debuild` software. If you found a bug or have
any questions about the recipe, please feel free to raise your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
