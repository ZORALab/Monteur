#!/bin/bash
##############################################################################
# APP VARIABLES
##############################################################################
VERSION="3.0.0"
action=""
publish_branch="gh-pages"
publish_mode=""

machine=""
arch=""
repo_path="$(
	__path="$PWD"
	while true; do
		if [ ! -d "${__path}/.bin" ]; then
			__path="${__path%/*}"
			continue
		elif [ -z "$__path" ]; then
			echo ""
			return 1
		fi
		break
	done
	echo "$__path"
)"
program_path="${repo_path}/.bin"
script_path="${repo_path}/.scripts"
tmp_path="${repo_path}/tmp"

bissetii="${script_path}/bissetii.sh"
godocgen="${script_path}/godocgen.sh"
hugo="${program_path}/hugo"


program_dependencies=(
	"$hugo"
)

##############################################################################
# FUNCTIONS LIBRARY
##############################################################################
_print_status() {
	_status_mode="$1" && shift 1

	# process status message
	_status_message=""
	case "$_status_mode" in
	error)
		_status_message="[  ERROR  ] $@"
		;;
	warning)
		_status_message="[  WARNING  ] $@"
		;;
	info)
		_status_message="[  INFO  ] $@"
		;;
	plain)
		_status_message="$@"
		;;
	*)
		return 0
		;;
	esac

	1>&2 echo "${_status_message}"
	unset _status_mode _status_message
}


__start_bissetii() {
	$bissetii -r
}


##############################################################################
# PRIVATE FUNCTIONS
##############################################################################
_clean_repo() {
	"$script_path"/gopher.sh -c
}

_get_machines_properties() {
	# detect CPU
	case "$(uname -m)" in
	x86_64)
		arch="amd64"
		;;
	i386|i686|x86|i686-AT386)
		arch="i386"
		;;
	aarch64)
		arch="arm64"
		;;
	armv5*)
		arch="armv5"
		;;
	armv6*)
		arch="armv6l"
		;;
	armv7*)
		arch="armv7"
		;;
	BePC)
		arch="bepc"
		;;
	ppc)
		arch="ppc"
		;;
	ppc64)
		arch="ppc64le"
		;;
	sparc64)
		arch="sparc64"
		;;
	*)
		_print_status error "unknown CPU architecture."
		exit 1
		;;
	esac
	_print_status info "detected architecture: $arch"


	# detect OS
	case "$(uname -s | tr '[:upper:]' '[:lower:]')" in
	darwin)
		machine="darwin"
		;;
	dragonfly)
		machine="dragonfly"
		;;
	freebsd)
		machine="freebsd"
		;;
	linux)
		machine="linux"
		;;
	android)
		machine="android"
		;;
	nacl)
		machine="nacl"
		;;
	netbsd)
		machine="netbsd"
		;;
	openbsd)
		machine="openbsd"
		;;
	plan9)
		machine="plan9"
		;;
	solaris)
		machine="solaris"
		;;
	windows)
		machine="windows"
		;;
	*)
		_print_status error "unknown operating system."
		exit 1
		;;
	esac
	_print_status info "detected operating system: $machine"
}


_check_supported_machine() {
	for _program in "${program_dependencies[@]}"; do
		if [ ! -e "$_program" ]; then
			_print_status error "missing $_program. Aborted."
			exit 1
		fi
		_print_status info "found dependency: $_program"
	done
}


_farewell_message() {
	_print_status info "===CI COMPLETED==="
}


_process_godocgen_build() {
	_job="build go documentations"
	_print_status info "'$_job' job included. Executing..."

	$godocgen -r
	if [ $? -ne 0 ]; then
		_print_status error "stopping CI. Godocgen build failed."
		exit 1
	fi

	_print_status info "==DONE== '$_job' job completed successfully."
}


_process_bissetii_build() {
	_job="build bissetii artifacts"
	_print_status info "'$_job' job included. Executing..."

	$bissetii -B
	if [ $? -ne 0 ]; then
		_print_status error "stopping CI. Hugo build failed."
		exit 1
	fi

	_print_status info "==DONE== '$_job' job completed successfully."
}


_publish_artifacts() {
	_job="publish bissetii artifact"
	_print_status info "'$_job' job included. Executing..."

	$bissetii -P clean -t "$publish_branch"
	if [ $? -ne 0 ]; then
		_print_status error "stopping CI. Publish failed."
		exit 1
	fi

	_print_status info "==DONE== '$_job' job completed successfully."
}


_start_building() {
	"$script_path"/gopher.sh -b
	if [ $? -ne 0 ]; then
		exit 1
	fi
}


_start_develop() {
	_print_status info "starting development system..."

	_print_status info "execute one testing round..."
	"$script_path"/gopher.sh --test

	_print_status info "initialize gopher web report interfaces..."
	"$script_path"/gopher.sh --present

	_print_status info "initialize go development environment..."
	"$script_path"/gopher.sh --develop

	_print_status info "!!! SUCCESS !!! Enjoy your development."
}


_stop_develop() {
	_print_status info "stopping development system..."
	"$script_path"/gopher.sh --stop-develop
	_print_status success ok
}


_setup_repository() {
	rm -rf "$program_path" &> /dev/null
	"$script_path"/helix.sh -i
}


_start_presentation() {
	"$script_path"/gopher.sh -p
	if [ $? -ne 0 ]; then
		exit 1
	fi
}


_start_testing() {
	_oldPath="$PWD" && cd "${repo_path}/pkg"

	"$script_path"/gopher.sh -t
	if [ $? -ne 0 ]; then
		exit 1
	fi

	"$script_path"/gopher.sh -B
	if [ $? -ne 0 ]; then
		exit 1
	fi

	cd "$_oldPath" && unset _oldPath
}


##############################################################################
# PUBLIC FUNCTIONS
##############################################################################
bad_command() {
	_print_status error "invalid command."
	exit 1
}


clean() {
	_get_machines_properties
	_check_supported_machine
	_clean_repo
	_farewell_message
}


develop() {
	_get_machines_properties
	_check_supported_machine
	_start_develop
}


build() {
	_get_machines_properties
	_check_supported_machine
	_start_building
	_farewell_message
}


release() {
	_get_machines_properties
	_check_supported_machine
	_farewell_message
}


setup() {
	_setup_repository
	_farewell_message
}


stop_develop() {
	_stop_develop
	_farewell_message
}


testing() {
	_get_machines_properties
	_check_supported_machine
	_start_testing
	_farewell_message
}


present() {
	_get_machines_properties
	_check_supported_machine
	_start_presentation
}


print_version() {
	echo $VERSION
}


publish() {
	_get_machines_properties
	_check_supported_machine
	_process_godocgen_build
	_process_bissetii_build
	if [ "$publish_mode" != "only-build" ]; then
		_publish_artifacts
	fi
	_farewell_message
}


print_help() {
	echo "\
CI AUTOMATOR
A program continuous integration automation tool.

-------------------------------------------------------------------------------
To use: $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -c, --clean                  execute CI clean.

2. -d, --develop                start the development environment.

3. -h, --help                   print help for this script app.

4. -b, --build                  execute CI build.

5. -p, --publish [MODE]         execute CI publish. [MODE] values can be:
                                  1. \"\"
                                  2. \"only-build\"

6. --present                    launch all presentation interfaces manually.

7. -r, --release                execute CI release.

8. -s, --setup                  execute CI setup.

9. -sd, --stop-develop          stop the development environment.

10. -t, --test                   execute CI test.

11. -v, --version               print this app version.
"
}

##############################################################################
# MAIN CLI
##############################################################################
run_action() {
case "$action" in
"c")
	clean
	;;
"d")
	develop
	;;
"h")
	print_help
	;;
"b")
	build
	;;
"p")
	publish
	;;
"present")
	present
	;;
"r")
	release
	;;
"s")
	setup
	;;
"sd")
	stop_develop
	;;
"t")
	testing
	;;
"v")
	print_version
	;;
*)
	bad_command
	;;
esac
}


process_parameters() {
while [[ $# != 0 ]]; do
case "$1" in
-b|--build)
	action="b"
	;;
-c|--clean)
	action="c"
	;;
-d|--develop)
	action="d"
	;;
-h|--help)
	action="h"
	;;
-p|--publish)
	action="p"
	if [ "$2" != "" ] && [ "${2:1}" != "-" ]; then
		publish_mode="$2"
		shift 1
	fi
	;;
--present)
	action="present"
	;;
-r|--release)
	action="r"
	;;
-s|--setup)
	action="s"
	;;
-sd|--stop-develop)
	action="sd"
	;;
-t|--test)
	action="t"
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
}


if [[ $BASHELL_TEST_ENVIRONMENT != true ]]; then
	main $@
fi
