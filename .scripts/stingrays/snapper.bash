#!/bin/bash
################################
# User Variables               #
################################
VERSION="0.3.0"

################################
# App Variables                #
################################
action=""
version=""
snap_path=""
workspace_path=""
control_path=""
snapcraft_yaml="snapcraft.yaml"
install_type=""

install_dependency() {
	sudo apt-get install snapd -y
	if [[ "$(which snapcraft)" == "" ]]; then
		sudo snap install snapcraft --classic
	else
		sudo snap refresh snapcraft
	fi
}

install() {
	case "$install_type" in
	packages)
		install_dependency
		;;
	*)
		echo "[ ERROR ] not specified to install."
		exit 1
		;;
	esac
}

update_version() {
	if [[ "$version" == "" ]]; then
		echo "[ ERROR ] unknown version number."
		exit 1
	fi

	if [[ ! -d "$snap_path" ]]; then
		echo "[ ERROR ] missing snap folder."
		exit 1
	fi
	control_path="${snap_path}/$snapcraft_yaml"

	if [[ ! -f "$control_path" ]]; then
		echo "[ ERROR ] missing $snapcraft_yaml file."
		exit 1
	fi

	ret="$(cat "$control_path" | grep "version")"
	if [[ "$ret" == "" ]]; then
		echo "[ ERROR ] missing version field."
		exit 1
	fi

	sed -i '/^version:/c\'"version: '$version'" $control_path
}

compile() {
	# prepare
	if [[ ! -d "$workspace_path" ]]; then
		echo "[ ERROR ] - invalid workspace path."
		return 1
	fi

	if [[ "$output_path" != "" ]]; then
		mkdir -p "$output_path"
	fi

	# compile
	current_dir="$PWD"
	cd "$workspace_path"
	snapcraft clean
	snapcraft
	cd "$current_dir"

	# export
	snap_path="$workspace_path/$(ls "$workspace_path" | grep ".snap")"
	if [[ ! -f "$snap_path" ]]; then
		echo "[ ERROR ] - faled to build $workspace_path"
		return 1
	fi

	if [[ "$output_path" != "" ]]; then
		mv "$snap_path" "$output_path"
	fi
}

upstream() {
	if [[ ! -f "$snap_path" ]]; then
		echo "[ ERROR ] - invalid file."
		return 1
	fi

	snapcraft push "$snap_path"
}

print_version() {
	echo $VERSION
}

print_help() {
	echo "\
SNAPPER
The Snapcraft semi-automatic script for building and updating snap packages
-------------------------------------------------------------------------------
To use: $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -h, --help			print help. Longest help is up
				to this length for terminal
				friendly printout.

2. -v, --version		print app version.

3. -uv, --update-version	update the version for a specified
				control file.
				COMPULSORY VALUES:
				1. -uv [new version number]
				       e.g. -uv 1.20.1

				COMPULSORY ARGUMENTS:
				1. -sp, --snap [path to snap folder]
				       e.g. -cp ./snap

4. -c, --compile		compile the snap.
				COMPULSORY ARGUMENTS:
				1. -w, --workspace-path [path/to/workspace]
				     the path to workspace that builds the snap
				     binary.

				OPTIONAL ARGUMENTS:
				1. -o, --output [path/to/output]
				    output snap directory. It will create a
				    snaps folder to store all the snap binaries.

5. -u, --upstream		upload the .snap file to snapcraft.io store.
				COMPULSORY VALUES:
				1. -u [path/to/file.snap]
				    the path to the .snap binary file for
				    upstream.

6. -i, --install		install snap related matters.
				COMPULSORY VALUES:
				1. -i packages
				  install all snap and snapcraft packages.
				  E.g.:
				  $ program -i packages
"
}

run_action() {
case "$action" in
"i")
	install
	;;
"uv")
	update_version
	;;
"c")
	compile
	;;
"u")
	upstream
	;;
"h")
	print_help
	;;
"v")
	print_version
	;;
*)
	echo "[ ERROR ] - invalid command."
	return 1
	;;
esac
}

process_parameters() {
while [[ $# != 0 ]]; do
case "$1" in
-uv|--update-version)
	action="uv"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		version="$2"
		shift 1
	fi
	;;
-i|--install)
	action="i"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		install_type="$2"
		shift 1
	fi
	;;
-c|--compile)
	action="c"
	;;
-sp|--snap)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		snap_path="$2"
		shift 1
	fi
	;;
-w|--workspace-path)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		workspace_path="$2"
		shift 1
	fi
	;;
-o|--output-path)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		output_path="$2"
		shift 1
	fi
	;;
-u|--upstream)
	action="u"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		snap_path="$2"
		shift 1
	fi
	;;
-h|--help)
	action="h"
	;;
-v|--version)
	action="v"
	;;
*)
	;;
esac
shift 1
done
}

main() {
	process_parameters $@
	run_action
	if [[ $? != 0 ]]; then
		exit 1
	fi
}

if [[ $BASHELL_TEST_ENVIRONMENT != true ]]; then
	main $@
fi
