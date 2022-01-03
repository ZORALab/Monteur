+++
date = "2021-12-23T21:34:34+08:00"
title = "Releases"
description = """
Monteur has multiple release versions across its birth time. Each advanced
version contains upgrades necessary to improve its capability and performances.
This section presents its download portal across all released versions.
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
url = "/en/versions/default-1200x628.png"
width = "1200"
height = "628"
alternateText = "Monteur Releases"

[thumbnails.1]
url = "/en/versions/default-1200x1200.png"
width = "1200"
height = "1200"
alternateText = "Monteur Releases"

[thumbnails.2]
url = "/en/versions/default-480x480.png"
type = "image/png"
width = "480"
height = "480"
alternateText = "Monteur Releases"


[menu.main]
parent = ""
name = "Downloads"
pre = "📥"
weight = 5
identifier = "versions"


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
echo 'deb [signed-by=/usr/share/keyrings/zoralab-keyring.gpg] https://monteur.zoralab.com/releases/deb stable main' \
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



### Downloadable Archives Files
Here are some currently available Monteur archived versions for download (either
in `tar.gz` or `zip`) depending on your operating system:
