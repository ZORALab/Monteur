#!/bin/bash
################################
# User Variables               #
################################
VERSION="0.0.1"

################################
# App Variables                #
################################
exit_code=0
action=""
base_path="$PWD"
sites_path=".sites"
configs_path=".configs"
release_path="release"
output_path="${base_path}/${sites_path}/static/releases/deb"
conf_path="${base_path}/${configs_path}/debian/conf"
db_path="${base_path}/${configs_path}/debian/db"
log_path="${base_path}/${configs_path}/debian/log"
packages_path="${base_path}/${release_path}/debians"
reprepro_cmd="reprepro --basedir $base_path \
		 --outdir $output_path \
		 --confdir $conf_path \
		 --dbdir $db_path \
		 --logdir $log_path"
################################
# Functions                    #
################################
print_version() {
	echo $VERSION
}

check_status() {
	$reprepro_cmd check "$1"
}

list() {
	$reprepro_cmd list "$1"
}

remove() {
	distribution_name="$1"
	package_name="$2"
	$reprepro_cmd remove "$distribution_name" "$package_name"
}

upstream() {
	distribution_name="$1"
	deb_path="$2"

	$reprepro_cmd includedeb "$distribution_name" "$deb_path"
	if [ $? -ne 0 ]; then
		exit_code=1
	fi

	unset distribution_name deb_path
}

run() {
	$reprepro_cmd --delete clearvanished

	for directory in "$packages_path"/*; do
		codename="${directory##*/}"

		deb_path="$(find "$directory" -type f -name "*.deb")"

		echo "[ INFO ] Found: ${codename}/${deb_path##*/}"
		echo "[ INFO ] Upstreaming..."
		upstream "$codename" "$deb_path"
		unset codename deb_path
		echo "[ INFO ] Done."
	done

	exit $exit_code
}


################################
# CLI Parameters and Help      #
################################
print_help() {
	echo "\
PROGRAM NAME
One liner description
-------------------------------------------------------------------------------
To use: $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -h, --help			print help. Longest help is up
				to this length for terminal
				friendly printout.

2. -r, --run			run the program. In this case,
				says the message.
				COMPULSORY VALUES:
				1. -r \"[message you want]\"

				COMPULSORY ARGUMENTS:
				1. -

				OPTIONAL ARGUMENTS:
				1. -

3. -v, --version		print app version.
"
}

run_action() {
case "$action" in
"r")
	run
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
-r|--run)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		message="${@:2}"
		shift 1
	fi
	action="r"
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
