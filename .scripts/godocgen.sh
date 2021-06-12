#!/bin/bash
##############################################################################
# USER INPUTS VARIABLES
##############################################################################
VERSION="1.0.0"
GODOCGEN_OUTPUT_PATH="${GODOCGEN_OUTPUT_PATH:-"${PWD}/docs/en-us/go"}"
GODOCGEN_SHIFT_LIST=(
)

##############################################################################
# APP VARIABLES
##############################################################################
repo_path="${repo_path:-"$PWD"}"
programs_path="${repo_path}/.bin"
tmp_path="${repo_path}/tmp/godocgen"

godocgen_template_path="${repo_path}/.godocgen/templates/md"
godocgen_output_filename="_index.md"

exit_code=0
go_config="${programs_path}/goconfig"
godocgen=""

old_GOROOT="${GOROOT:-""}"
old_GOPATH="${GOPATH:-""}"
old_GOBIN="${GOBIN:-""}"
old_GOCACHE="${GOCACHE:-""}"
old_GOENV="${GOENV:-""}"

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
_check_go_availability() {
	_print_status info "checking go availability..."

	if [ -f "$go_config" ]; then
		_print_status info "detected local go configs. Deploying..."
		unset GOROOT GOPATH GOBIN GOCACHE GOENV
		source "$go_config"
		_print_status ok
	fi

	if [ -z "$GOROOT" ]; then
		_print_status error "missing GOROOT."
		exit 1
	fi

	if [ -z "$GOPATH" ]; then
		_print_status error "missing GOPATH."
		exit 1
	fi
}


_check_godocgen_availability() {
	_print_status info "checking godocgen availability..."

	godocgen="$(type -p godocgen)"
	if [ ! -z "$godocgen" ]; then
		_print_status info "Deploying $godocgen..."
		_print_status ok
	else
		_print_status error "no available godocgen detected!"
		exit_code=1
		exit 1
	fi
}


_close_app() {
	_print_status info "closing gopher helper..."

	_print_status info "stopping local go configurations..."
	if [ -f "$go_config" ]; then
		source "$go_config" --stop
		_print_status ok
	fi

	_print_status info "restoring GOROOT GOPATH GOBIN GOCACHE GOENV..."
	if [ ! -z "$old_GOROOT" ]; then
		export GOROOT="$old_GOROOT"
	else
		unset GOROOT
	fi

	if [ ! -z "$old_GOPATH" ]; then
		export GOPATH="$old_GOPATH"
	else
		unset GOPATH
	fi

	if [ ! -z "$old_GOBIN" ]; then
		export GOBIN="$old_GOBIN"
	else
		unset GOBIN
	fi

	if [ ! -z "$old_GOCACHE" ]; then
		export GOCACHE="$old_GOCACHE"
	else
		unset GOCACHE
	fi

	if [ ! -z "$old_GOENV" ]; then
		export GOENV="$old_GOENV"
	else
		unset GOENV
	fi

	_print_status ok
	exit $exit_code
}


_generate_go_docs() {
	_print_status info "generating go documentations..."
	$godocgen --filename "$godocgen_output_filename" \
		--template "$godocgen_template_path" \
		--output "$GODOCGEN_OUTPUT_PATH" \
		--input "${repo_path}/..."
	if [ $? -eq 0 ]; then
		_print_status ok
	else
		_print_status error "error occured!"
		exit_code=1
	fi
}


_restore_artifacts() {
	if [ ${#GODOCGEN_SHIFT_LIST[@]} -eq 0 ]; then
		return 0
	fi

	_print_status info "restoring artifacts..."
	mkdir -p "$GODOCGEN_OUTPUT_PATH"

	for _item in "${GODOCGEN_SHIFT_LIST[@]}"; do
		_print_status info "restoring $_item ..."
		mv "${tmp_path}/${_item}" "${GODOCGEN_OUTPUT_PATH}/${_item}"
		_print_status ok
	done
	unset _item
}


_shift_artifacts() {
	if [ ${#GODOCGEN_SHIFT_LIST[@]} -eq 0 ]; then
		return 0
	fi

	_print_status info "preserving artifacts..."
	rm -rf "$tmp_path"
	mkdir -p "$tmp_path"

	for _item in "${GODOCGEN_SHIFT_LIST[@]}"; do
		_print_status info "preserving $_item ..."
		mv "${GODOCGEN_OUTPUT_PATH}/${_item}" "$tmp_path"
		_print_status ok
	done
	unset _item
}


##############################################################################
# PUBLIC FUNCTIONS
##############################################################################
run() {
	_print_status info "executing Godocgen..."
	_check_go_availability
	_check_godocgen_availability
	trap _close_app EXIT
	_shift_artifacts
	_generate_go_docs
	_restore_artifacts
	_print_status ok
}


print_version() {
	echo "$VERSION"
}


##############################################################################
# CLI PARAMETERS AND HELP
##############################################################################
print_help() {
	echo "\
GOPHER
A Go repository helper to keep the processes sane.
--------------------------------------------------------------------------------
To use: $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -h, --help			print this program's help messages.


2. -r, --run			run godocgen for this repo.


3. -v, --version		print app version.
"
}

run_action() {
case "$action" in
"h")
	print_help
	;;
"r")
	run
	;;
"v")
	print_version
	;;
*)
	_print_status error "invalid command"
	exit 1
	;;
esac
}


process_parameters() {
while [[ $# != 0 ]]; do
case "$1" in
-h|--help)
	action="h"
	;;
-r|--run)
	action="r"
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
