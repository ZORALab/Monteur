#!/bin/bash
################################
# User Variables               #
################################
VERSION="0.3.0"

################################
# App Variables                #
################################
action=""
priority="low"
version=""
reference_branch=""
target_branch=""
diff=""
timestamp=""
spins=(
)

# markdown changelog
md_changelog_path=""

# deb package changelog
deb_path=""
deb_changelog_path="${PWD}/changelog.debian"
deb_maintainer="DEBGUY"
deb_name="DEBNAME"
deb_dist_tag="DISTTAG"
deb_dist_series="DISTSERIES"


print_version() {
	echo $VERSION
}

validate_diff_params() {
	if [[ "$version" == "" ]]; then
		echo "[ ERROR ] - incorrect version number: $version"
		exit 1
	fi

	if [[ "$reference_branch" == "" ]]; then
		echo "[ ERROR ] - missing reference branch."
		exit 1
	fi

	if [[ "$target_branch" == "" ]]; then
		echo "[ ERROR ] - missing target branch."
		exit 1
	fi

	case "$priority" in
	low|medium|high)
		;;
	*)
		echo "[ ERROR ] - invalid priority."
		exit 1
	esac
}

fetch_latest() {
	current="$(git name-rev --name-only HEAD)"

	if [[ "$current" == "" ]]; then
		echo "[ ERROR ] - malformed git repository."
		exit 1
	fi

	if [[ "$current" == "$target_branch" ]]; then
		git pull
	else
		git fetch origin "$target_branch:$target_branch"
	fi

	if [[ "$current" == "$reference_branch" ]]; then
		git pull
	else
		git fetch origin "$reference_branch:$reference_branch"
	fi
}

get_diff() {
	validate_diff_params
	fetch_latest
	diff="$(git log \
		--oneline "origin/$target_branch..origin/$reference_branch")"
	timestamp=$(date -R)
	if [[ "$diff" == "" ]]; then
		echo "[ INFO ] Branches show 0 differences."
		exit 0
	fi
}

validate_log_params() {
	if [[ -f "$changelog_md_path" ]]; then
		echo "[ ERROR ] - CHANGELOG.md exists."
		exit 1
	fi

	if [[ "$version" == "" ]]; then
		echo "[ ERROR ] - missing version number."
		exit 1
	fi
}

get_log() {
	validate_log_params
	git pull
	diff="$(git log --oneline)"
	timestamp=$(date -R)
	if [[ "$diff" == "" ]]; then
		echo "[ ERROR ] Empty git log."
		exit 1
	fi
}

get_deb_data() {
	if [[ ! -d "$deb_path" ]]; then
		echo "[ ERROR ] missing debian folder. Abort deb building."
		return 1
	fi

	deb_control_path="$deb_path/control"
	deb_changelog_path="$deb_path/changelog"

	if [[ ! -f "$deb_control_path" ]]; then
		echo "[ WARNING ] missing debian/control. Using placeholder."
		return 1
	fi

	while IFS='' read -r line || [[ -n "$line" ]]; do
		field="${line%%:*}"
		value="${line#*: }"
		case "$field" in
		"Package")
			deb_name="$value"
			;;
		"Maintainer")
			deb_maintainer="$value"
			;;
		*)
			;;
		esac
	done < "$deb_control_path"
}

update_changelog_deb() {
	# header
	newlog="\
${deb_name} (${version}${deb_dist_tag}) ${deb_dist_series}; urgency=${priority}

"

	# update
	while read -r line; do
		newlog="$newlog  *${line:8}\n"
	done <<< "$diff"

	# tail
	newlog="${newlog}\n -- $deb_maintainer  $timestamp"

	# write to changelog.debian
	if [[ ! -f "$deb_changelog_path" ]]; then
		echo -e "$newlog" > "$deb_changelog_path"
		return 0
	fi
	echo -e "\
$newlog

$(cat "$deb_changelog_path")" > "$deb_changelog_path"
}

update_changelog_md() {
	# header
	newlog="\
# Version $version
## $timestamp - (Urgency: $priority)
--------------------------------------------------------------------------------"
	# update
	i=1
	while read -r line; do
		newlog="$newlog
$i. $line"
		((i+=1))
	done <<< "$diff"

	# tail

	# write to CHANGELOG.md
	if [[ ! -f "$md_changelog_path" ]]; then
		echo -e "$newlog" > "$md_changelog_path"
		return 0
	fi
	echo -e "\
$newlog

$(cat "$md_changelog_path")" > "$md_changelog_path"
}

update() {
	if [[ "$version" == "" || "${version:0:1}" = "-" ]]; then
		echo "[ ERROR ] no version specified."
		exit 1
	fi

	case "$action" in
	u)
		get_diff
		;;
	c)
		get_log
		;;
	*)
		exit 1
		;;
	esac

	if [[ "$md_changelog_path" != "" ]]; then
		update_changelog_md
	fi

	if [[ "$deb_path" != "" ]]; then
		get_deb_data
		update_changelog_deb
	fi
}

spin_deb_changelog() {
	for spin in "${spins[@]}"; do
		key="${spin%%=*}"
		value="${spin##*=}"
		sed -Ei "s|${key}|${value}|g" "$deb_changelog_path"
	done
}

spin() {
	if [[ ! -d "$deb_path" ]]; then
		echo "[ ERROR ] debian folder path not specified."
		exit 1
	fi

	deb_changelog_path="$deb_path/changelog"
	if [[ ! -f "$deb_changelog_path" ]]; then
		echo "[ ERROR ] missing changelog in debian folder."
		exit 1
	fi

	spin_deb_changelog
}

print_help() {
	echo "\
CHANGELOGGER
Update all official changelogs based on git log between the release branch and
your updated branch.
-------------------------------------------------------------------------------
To use: $ $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -h, --help			print help. Longest help is up
				to this length for terminal
				friendly printout.

2. -v, --version		print app version.

3. -u, --update			update all available changelogs.
				COMPULSORY VALUE:
				1. -u [new version number]
				       e.g. -u 1.0.1

				COMPULSORY ARGUMENTS:
				1. -ref, --reference [reference branch name]
				      git branch name containing the latest.
				      e.g. -u 1.0.1 -ref next

				2. -tgt, --target [target branch name]
				      git branch name for merging the changes.
				      e.g. -u 1.0.1 -tgt master

				OPTIONAL ARGUMENTS:
				1. -p, --priority [low/medium/high]
				      change update priority. Default is low.
				      Some packager like .deb uses it.
				      e.g. -u 1.0.1 -p high

				2. -deb, --debian [path to debian folder]
				      create deb package changelog inside the
				      debian folder. If debian/control path is
				      not available, placeholder is used.
				      e.g. -u 1.0.1 -deb ./debian

				3. -md, --markdown [path to output file]
				      create the markdown version of changelog.
				      if output path is given, the file will
				      be created at the designated location.
				      e.g. -u 1.0.1 -md ./CHANGELOG.md

				EXAMPLES:
				1. $ program -u 1.0.1 \\
					     -ref next \\
					     -tgt master \\
					     -md

				2. $ program -u 1.0.1 \\
					     -ref next \\
					     -tgt master \\
					     -md \"./CHANGELOG.md\"

				3. $ program -u 1.0.1 \\
					     -ref next \\
					     -tgt master \\
					     -p high \\
					     -md \"./CHANGELOG.md\" \\
					     -deb \"./app/debian\" \\

4. -c, --create			create a basic changelog during initialization.
				It shares the same arguments and values pattern
				with --update. The only difference is that it
				uses the local git branch to create the
				changelogs instead.
				COMPULSORY VALUE:
				1. -c [version number]
				  e.g. -c 1.0.1

				COMPULSORY ARGUMENTS
				- none

				OPTIONAL ARUGMENTS
				++ same as --update ++

				EXAMPLES:
				1. $ program -u 1.0.1 \\
					     -p high \\
					     -md \"./CHANGELOG.md\"

				2. $ program -u 1.0.1 \\
					     -p high \\
					     -md \"./CHANGELOG.md\" \\
					     -deb \"./app/debian\" \\

5. -s, --spin			spin the changelog placeholders with a set of
				given value.
				COMPULSORY VALUE:
				1. -s [path to debian folder]
				   path to the debian folder containing the
				   changelog
				   e.g. -s ./debian

				COMPULSORY ARGUMENTS:
				1. -k, --key-value [key=value]
				   the key-value pair for replacement. This
				   argument is repeatable for various
				   placeholders.
				   e.g. -s ./debian/changelog \\
					-k \"DISTSERIES=stable\" \\
					-k \"DISTTAG=1.2.1-0ubuntu1~artfulppa1\"

				EXAMPLES:
				1. $ program -s "./debian" \\
					-k \"DISTSERIES=artful\" \\
					-k \"DISTTAG=1.2.1-0ubuntu1~artfulppa1\"
"
}

run_action() {
case "$action" in
"h")
	print_help
	;;
"v")
	print_version
	;;
"u"|"c")
	update
	;;
"s")
	spin
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
-h|--help)
	action="h"
	;;
-v|--version)
	action="v"
	;;
-u|--update)
	action="u"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		version="$2"
		shift 1
	fi
	;;
-c|--create)
	action="c"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		version="$2"
		shift 1
	fi
	;;
-ref|--reference)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		reference_branch="$2"
		shift 1
	fi
	;;
-tgt|--target)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		target_branch="$2"
		shift 1
	fi
	;;
-p|--priority)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		priority="$2"
		shift 1
	else
		priority=""
	fi
	;;
-deb|--debian)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		deb_path="$2"
		shift 1
	fi
	;;
-md|--markdown)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		md_changelog_path="$2"
		shift 1
	fi
	;;
-s|--spin)
	action="s"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		deb_path="$2"
		shift 1
	fi
	;;
-k|--key-value)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		spins+=("$2")
		shift 1
	fi
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
