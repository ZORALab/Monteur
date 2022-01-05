+++
date = "2021-12-23T11:10:57+08:00"
title = "Go Build API"
description = """
Monteur strongly supports `go` for building itself and other go applications.
This recipe allows repository equipped with go to seamlessly build go programs
in a highly customizable manner easily and seamlessly.
"""
keywords = [
	"Go",
	"Build CI Job",
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
url = "/en/ci-jobs/build/go/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Build with Go"

[thumbnails.1]
url = "/en/ci-jobs/build/go/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Build with Go"

[thumbnails.2]
url = "/en/ci-jobs/build/go/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Build with Go"


[menu.main]
parent = "Software - Go"
name = "Build"
pre = "🧰"
weight = 70
identifier = "ci-jobs-build-go"


[schema]
selectType = "WebPage"
+++

# {{% param "title" %}}
{{% param "description" %}}

The objective of distributing the recipe is simple: **to build a Go application
for a compute system variant painlessly in a consistent manner with minimal to
no further instructions**.

Go is available at:
https://go.dev/

All recipes are arranged based on its own
[semantic versioning](https://semver.org/) and is not directly related to
Monteur's actual release version. Hence, feel free to explore each versions
to suit your CI needs.




## Recipe Versions
Here are the available Build API recipe for `go` integrations. Please read
through your selected version's details on what has changed, what is required,
and how to customize and used them.

The arrangement are the latest at the top or first.



### Version 2.0.0
Version 2.0.0 `go` Build API is available for download here:
{{< link "/ci-jobs/build/go/go-v2p0p0.toml" "this" "" "" "button"
	"" "download" >}}
go-v2p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.2`                         |
| Supported Platforms  | follows `go`'s availability      |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/build/jobs` directory with the name pattern `OS-ARCH.toml` like
   `linux-amd64.toml` for `amd64` CPU architecture `linux` operating system.
2. Once done, edit the configuration file for:
   1. `Metadata.Name` - rename the name accordingly. Recommended `<OS>-<ARCH>`
      like `linux-amd64`.
   2. `Metadata.Description` - update accordingly by replacing the `<APP>`,
      `<OS>`, `<CPU>` to the appropriate values.
   3. `Variables.PlatformOS` - update the value accordingly to the target OS
      (e.g. `linux`).
   4. `Variables.PlatformCPU` - update the value accordingly to the target CPU
      architecture (e.g. `amd64`).
   5. `Variables.PlatformExt` - update the file extension if needed (e.g.
      `.exe`). Otherwise, leave it blank.
   6. `Variables.BuildFlags` - your build arguments matching the current
       variant. Please remove any non-applicable argument when needed. The
       default are designed for consumer-ready build configurations.
   7. `Variables.BuildConditions` - the build conditions of the `go build`. The
       default includes everything sensible. Please remove any non-applicable
       conditions when needed (like `CC=arm-linux-gnueabi-gcc`, `GOARM=7` for
       non-`arm` build).
       1. Keep `CGO_ENABLED=0` as minimum value even when you do not use `cgo`
          at all.
   8. `FMTVariables.SrcPath` - filepath to your `main.go`.
   9. `Dependencies` - remove non-applicable dependencies (e.g.
      `arm-linux-gnueabi-gcc` for non-`arm` build.

For detailed information about each fields, visit:
[Build Specification Data Structure]({{< link
"/ci-jobs/build/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Added `arm` build with their specific dependencies
   and build conditions.
2. *Backward Compatible* - Added `-trimpath` and its other variants into
   `Variables.BuildFlags`.
3. *Backward Compatible* - Added `Variables.BuildConditions` to enable
   CGO and arm configurable settings.
4. *Backward Compatible* - Added `CGO_ENABLED=0` to explictly instruct Go not to
   use `cgo`.



### Version 1.0.0
Version 1.0.0 `go` Build API is available for download here:
{{< link "/ci-jobs/build/go/go-v1p0p0.toml" "this" "" "" "button"
	"" "download" >}}
go-v1p0p0.toml
{{< /link >}}

| Min Requirements     | Values                           |
|:---------------------|---------------------------------:|
| Monteur Version      | `v0.0.1`                         |
| Supported Platforms  | follows `go`'s availability      |


#### Installation Instructions
1. You should download and place the recipe into your
   `<config>/app/variants/` directory with the name pattern `OS-ARCH.toml` like
   `linux-amd64.toml`.
2. Once done, edit the configuration file for:
   1. `Metadata.Name` - rename the name accordingly. Recommended `<OS>-<ARCH>`
      like `linux-amd64`.
   2. `Metadata.Description` - update accordingly by replacing the `<APP>`,
      `<OS>`, `<CPU>` to the appropriate values.
   3. `Variables.PlatformOS` - update the value accordingly to the target OS
      (e.g. `linux`).
   4. `Variables.PlatformCPU` - update the value accordingly to the target CPU
      architecture (e.g. `amd64`).
   5. `Variables.PlatformExt` - update the file extension if needed (e.g.
      `.exe`). Otherwise, leave it blank.
   6. `Variables.BuildFlags` - your build arguments matching the current
       variant. Please remove any non-applicable argument when needed (e.g.
       `-buildmode=pie` for `windows`).
   7. `FMTVariables.SrcPath` - filepath to your `main.go`.

For detailed information about each fields, visit:
[Build Specification Data Structure]({{< link
"/ci-jobs/build/#data-structure" "this" "url-only" />}}) for more info.


#### Changes
1. *Backward Compatible* - Created the base TOML configuration recipe.
2. *Backward Compatible* - Integrated with GitLab CI.
3. *Backward Compatible* - Uses `go` with `-buildmode=pie -ldflags "-s -w"` as
   default arguments.




## Epilogue
That's all for Monteur operating `go` software for building its binary. If you
found a bug or have any questions about the recipe, please feel free to raise
your question at our
[Issue Section](https://gitlab.com/zoralab/monteur/-/issues).
