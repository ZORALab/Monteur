#!/bin/bash
################################
# User Variables               #
################################
VERSION="0.6.0"

################################
# App Variables                #
################################
action=""
deb_path=""
workspace_path=""
output_path=""
gpg_id=""
gpg_compile=true
source_only=false
app_name=""
app_output_path=""
install_type=""
lintian_path=""
deb_profile="debian"
upstream_tag=""
dput_path=""

print_version() {
	echo $VERSION
}

install_dependency() {
	sudo apt-get install build-essential fakeroot devscripts -y
	sudo apt-get install gnupg ubuntu-dev-tools apt-file -y
}

install() {
	case "$install_type" in
	packages)
		install_dependency
		;;
	*)
		echo "[ ERROR ] nothing specified to install."
		exit 1
		;;
	esac
}

verify_gpg() {
	if [[ "$gpg_id" == "" ]]; then
		gpg_compile=false
		return 0
	fi

	ret="$(gpg --list-secret-keys | grep "$gpg_id")"
	if [[ "$ret" == "" ]]; then
		echo "[ WARNING ] - GPG key not found. Compile without GPG."
		gpg_compile=false
		return 0
	fi
}

prepare_workspace_path() {
	if [[ ! -d "$workspace_path" ]]; then
		echo "[ ERROR ] - workspace is not a directory."
		exit 1
	fi

	if [[ ! -d "${workspace_path}/debian" ]]; then
		echo "[ ERROR ] - workspace is not a valid debian directory."
		exit 1
	fi

	app_output_path=$(dirname "$workspace_path")
	app_name=$(basename "$app_output_path")
}

prepare_output_path() {
	if [[ "$output_path" == "" ]]; then
		return 0
	fi

	mkdir -p "$output_path"
}

export_output() {
	ls -p "$app_output_path" \
		| grep -v "/" \
		| xargs -I{} cp "${app_output_path}/"{} "${output_path}/."
}

compile() {
	prepare_workspace_path
	verify_gpg
	prepare_output_path

	lastpath="$PWD"
	cd "$workspace_path"

	# formulate main arguments
	arguments=""
	if [[ "$gpg_compile" == "false" ]]; then
		arguments="-us -uc"
	elif [[ "$source_only" == "true" ]]; then
		arguments="-k$gpg_id -S"
	else
		arguments="-k$gpg_id"
	fi

	# formulate lintian arguments
	lintian=""
	if [[ "$lintian_path" != "" && -d "$lintian_path" ]]; then
		lintian="--include-dir $lintian_path"
	fi

	if [[ "$deb_profile" != "" ]]; then
		lintian="$lintian --profile $deb_profile"
	fi
	arguments="$arguments --lintian-opts $lintian"

	# execute
	debuild $arguments
	export_output
	cd "$lastpath"
	unset lastpath
}

upstream() {
	if [[ ! -f "$deb_path" ]]; then
		echo "[ ERROR ] missing .changes file: $deb_path"
		exit 1
	fi

	if [[ ! -f "$dput_path" ]]; then
		echo "[ ERROR ] missing dput.cf file: $dput_path"
		exit 1
	fi

	if [[ "$upstream_tag" == "" ]]; then
		echo "[ ERROR ] missing upstream tag: $upstream_tag"
		exit 1
	fi

	dput --config "$dput_path" "$upstream_tag" "$deb_path"
}

print_help() {
	echo "\
DEBBER
The Debian package semi-automatic script for building and updating .deb
packages.
-------------------------------------------------------------------------------
To use: $ $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -h, --help			print help. Longest help is up
				to this length for terminal
				friendly printout.


2. -v, --version		print app version.


3. -c, --compile		compile the deb for the given parameters.
				COMPULSORY ARGUMENTS:
				1. -w, --workspace-path [path to workspace]
				   provide the workspace path for deb to
				   compile. It is the app directory containing
				   the debian folder.
				   E.g.:
				     $ program -c -w path/to/tmp/app

				2. -o, --output [path to output]
				   provide the output path for exporting the
				   output deb or .changes files. It will
				   create a debs folder to store all the files.
				   E.g.:
				     $ program -c -o path/to/release

				OPTIONAL ARGUMENTS:
				1. -k, --gpg [gpg id]
				   GPG key id to use for PPA repo and signing.
				   E.g.:
				     $ program -c -k AB234AB

				2. -cs, --source-only
				   Compile 'source-only' deb package. E.g.:
				     $ program -c -cs

				3. -lp, --lintian-path path/to/lintian/data
				   Custom lintian data path. E.g.:
				     $ program -lp \$HOME/.lintian


				EXAMPLES:
				1. $ program \\
					-c \\
					-w ./tmp \\
					-o ./release \\
					-lp \$HOME/.lintian \\
					-k AB234AB \\
					-cs

				2. $ program \\
					-w ./tmp \\
					-o ./release \\
					-lp \$HOME/.lintian \\
					-k AB234AB

				3. $ program \\
					-w ./tmp \\
					-o ./release


4. -u, --upstream		upload the .deb file to upstream.
				COMPULSORY VALUE:
				1. -u [path to file]
				  path to the .changes for upstream.
				  E.g.:
				  $ program -u ./path/to/pack_sources.changes

				COMPULSORY ARGUMENTS:
				1. -ut, --upstream-tag [tag name]
				   the tag name for the dput.cf configurations.
				   It uses the provided dput.cf to search for
				   upstream details and perform dput upstream.
				   E.g.:
				   $ program -u sample_sources.changes \\
					     -ut ubuntu-launchpad

				2. -dp, --dput-path [path to file]
				   the path to the dput.cf. Manta requires you
				   to specify the location.
				   E.g.:
				   $ program -u sample_sources.changes \\
					     -ut ubuntu-launchpad \\
					     -dp ./dput.cf


5. -i, --install		install deb related matters.
				COMPULSORY VALUE:
				1. -i packages
				  install all debuild packages.
				  E.g.:
				  $ program -i packages
"
}

run_action() {
case "$action" in
"i")
	install
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
-deb|--debian)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		deb_path="$2"
		shift 1
	fi
	;;
-i|--install)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		install_type="$2"
		shift 1
	fi
	action="i"
	;;
-c|--compile)
	action="c"
	;;
-w|--workspace-path)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		workspace_path="$2"
		shift 1
	fi
	;;
-o|--output)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		output_path="$2"
		shift 1
	fi
	;;
-k|--gpg)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		gpg_id="$2"
		shift 1
	fi
	;;
-cs|--source-only)
	source_only=true
	;;
-u|--upstream)
	action="u"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		deb_path="$2"
		shift 1
	fi
	;;
-ut|--upstream-tag)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		upstream_tag="$2"
		shift 1
	fi
	;;
-dp|--dput-path)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		dput_path="$2"
		shift 1
	fi
	;;
-lp|--lintian-path)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		lintian_path="$2"
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
