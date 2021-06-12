#!/bin/bash
##############################################################################
# USER PARAMS
##############################################################################
HELIX_DIRECTORY="${HELIX_DIRECTORY:-""}"
OUTPUT_DIRECTORY="${OUTPUT_DIRECTORY:-""}"
ROOT="${ROOT:-false}"

##############################################################################
# APP VARIABLES
##############################################################################
VERSION="2.0.0"
action=""

machine=""
arch=""
repo_path="$PWD"
output_path="${OUTPUT_DIRECTORY:-"${repo_path}/.bin"}"
shell_init_path="${output_path}/repo-init.sh"
package_config_path="${output_path}/configs"
default_output_path="$output_path"
workspace="/tmp"
package=""
package_name=""
package_title=""
package_version=""

##############################################################################
# FUNCTIONS LIBRARY
##############################################################################
_print_status() {
	__status_mode="$1" && shift 1
	__msg=""
	__stop_color="\033[0m"
	case "$__status_mode" in
	error)
		__msg="[ ERROR   ] ${@}"
		__start_color="\e[91m"
		;;
	warning)
		__msg="[ WARNING ] ${@}"
		__start_color="\e[93m"
		;;
	info)
		__msg="[ INFO    ] ${@}"
		__start_color="\e[96m"
		;;
	success)
		__msg="[ SUCCESS ] ${@}"
		__start_color="\e[92m"
		;;
	ok)
		__msg="[ INFO    ] == OK =="
		__start_color="\e[96m"
		;;
	plain)
		__msg="$@"
		;;
	*)
		return 0
		;;
	esac

	if [ $(tput colors) -ge 8 ]; then
		__msg="${__start_color}${__msg}${__stop_color}"
	fi

	1>&2 echo -e "${__msg}"
	unset __status_mode __msg __start_color __stop_color
}


##############################################################################
# PRIVATE FUNCTIONS
##############################################################################
_check_helix_directory() {
	if [ -z "$HELIX_DIRECTORY" ]; then
		if [ -d "${repo_path}/.configs/helix" ]; then
			HELIX_DIRECTORY="${repo_path}/.configs/helix"
		elif [ -d "${repo_path}/.helix" ]; then
			HELIX_DIRECTORY="${repo_path}/.helix"
		else
			_print_status error "missing HELIX_DIRECTORY."
			exit 1
		fi
	fi
}


_check_package_title() {
	package_title="${1##*/}"
	package_title="${package_title%.*}"
	package_title="${package_title^}"
	if [ -z "$package_title" ]; then
		package_title="Program"
	fi
}


_check_root_privilege() {
	if [ "$ROOT" == "true" ] && [ "$(id -u)" != "0" ]; then
		_print_status error "need root access for setup."
		_print_status info ""
		return 1
	fi
}


_check_target_version() {
	package_version=""
	if [ ! -z "$PACKAGE_VERSION" ]; then
		package_version="$PACKAGE_VERSION"
	fi

	if [ -z "$package_version" ]; then
		_print_status error "package_version not specified."
		exit 1
	fi
	_print_status info \
		"$package_title $package_version requested. Installing..."
}


_clean_up_output_directory() {
	rm -rf "$output_path" &> /dev/null
	mkdir -p "$output_path"
}


_create_init_script() {
	_print_status info "creating init shell script for repository..."


	# make sure the directory is there for script creation
	mkdir -p "$output_path"
	if [ ! -d "$output_path" ]; then
		_print_status error "$output_path is not available."
		exit 1
	fi


	# create init script
	echo "\
#!/bin/bash
export REPO_PATH="$output_path"
export PATH=\"\${PATH}:\${REPO_PATH}\"

case \"\$1\" in
-h|--help|help)
	2>&1 echo \"HELP
1) to start: $ source \$0
2) to stop : $ \$0 --stop\"
	;;
--stop)
	# stop all configs
	for i in \"\${REPO_PATH}/configs/\"*; do
		1>&2 echo \"rescinding: \${i}\"
		source \"\$i\" --stop
	done

	# unset repo pathing
	PATH=\":\${PATH}:\"
	PATH=\"\${PATH//:\$REPO_PATH:/:}\"
	PATH=\"\${PATH%:}\"
	unset REPO_PATH

	1>&2 echo \"localized cmd rescinded.\"
	;;
*)
	for i in \"\${REPO_PATH}/configs/\"*; do
		1>&2 echo \"sourcing: \${i}\"
		source \"\$i\" --start
	done

	1>&2 echo \"localized cmd initialized.\"
	;;
esac" > "$shell_init_path"
	chmod +x "$shell_init_path"


	_print_status success ok
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


_download_package() {
	package_name="$(__generate_package_name "$machine" "$arch")"
	package="${repo_path}/${package_name}"

	_print_status info "downloading ${package_name}..."
	rm "$package" &> /dev/null

	# download by available downloader
	_url="$(__generate_url)"
	if [ ! -z "$(type -p curl)" ]; then
		curl --silent \
			--fail \
			--location \
			--output "$package" \
			--remote-name \
			"$_url"
	elif [ ! -z "$(type -p wget)" ]; then
		wget --https-only \
			--output-document="$package" \
			"$_url"
	else
		_print_status error "no supported downloader available."
		exit 1
	fi
	unset _url

	# check output
	if [ $? -ne 0 ] || [ ! -f "$package" ]; then
		_print_status error "download installer package failed."
		exit 1
	fi

	_print_status success ok
}


_check_supported_version() {
	__function=__supported_version
	if [ -z "$(type "$__function")" ]; then
		_print_status error "@devloper: you must supply '$__function'"
		exit 1
	fi
	$__function
}


_check_input_functions() {
	__function=__generate_package_name
	if [ -z "$(type $__function)" ]; then
		_print_status error "@devloper: you must supply '$__function'"
		exit 1
	fi

	__function=__generate_url
	if [ -z "$(type $__function)" ]; then
		_print_status error "@devloper: you must supply '$__function'"
		exit 1
	fi
}


_farewell_message() {
	_print_status success "\
All programs are installed successfully.

You may find the program available at:
$output_path

----
Please Enjoy! See ya.
"
}


_setup_package() {
	_print_status info "unpacking package..."

	# unpack
	output_path="${OUTPUT_DIRECTORY:-"$default_output_path"}"
	mkdir -p "$output_path" "$package_config_path"

	# extract by extensions
	__name="${package##*/}"
	if [[ "$__name" == *".tar.gz"* ]]; then
		tar -C "$output_path" -xzf "$package"
	elif [[ "$__name" == *".zip"* ]]; then
		unzip "$package" -d "$output_path"
	else
		_print_status error "unknown compression tool: $__name"
		exit 1
	fi

	# clean up
	rm "$package"
	_print_status success ok
}


_wrap_up_setup() {
	if [ ! -z "$(type __wrap_up 2> /dev/null)" ]; then
		_print_status info "wrapping up..."
		__wrap_up
		_print_status success ok
	fi

	_print_status info "DONE"
	_print_status info ""
}


##############################################################################
# PUBLIC FUNCTIONS
##############################################################################
install() {
	_check_helix_directory
	_clean_up_output_directory

	for config in "$HELIX_DIRECTORY"/*.sh; do
		# reset variables
		ROOT=false
		OUTPUT_DIRECTORY=""

		# get configurations
		source "$config"
		_check_package_title "$config"
		_check_target_version
		_get_machines_properties
		_check_supported_version

		_check_root_privilege
		if [ $? -ne 0 ]; then
			continue
		fi

		# process
		_download_package
		_setup_package
		_wrap_up_setup
	done

	_create_init_script
	_farewell_message
}

print_version() {
	echo $VERSION
}

bad_command() {
	_print_status error "invalid command."
	exit 1
}

print_help() {
	echo "\
HELIX INSTALLER AUTOMATOR
This is an automatic script to setup and install all repository's dependencies
from known Git version control system like Github or GitLab. The installation
path is setup into $output_path directory.

-------------------------------------------------------------------------------
To use: $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -h, --help                          print help for this script app.

2. -i, --install [CONFIG_DIRECTORY]    install all listed software in
                                       [CONFIG_DIRECTORY]. By default, it reads
                                       the following directories in priorities:
                                           1. ./.helix/
                                           2. ./.configs/helix/

                                       EXAMPLES:
                                           1. $ ./program-setup.bash -i
                                           2. $ ./program-setup.bash -i ./helix

3. -v, --version                       print this app version.
"
}


##############################################################################
# MAIN CLI
##############################################################################
run_action() {
case "$action" in
"h")
	print_help
	;;
"i")
	install
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
-h|--help)
	action="h"
	;;
-i|--install)
	action="i"
	if [[ "$2" != "" && "${2:1}" != "-" ]]; then
		package_version="$2"
		shift 1
	fi
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
