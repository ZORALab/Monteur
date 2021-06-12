#!/bin/bash
##############################################################################
# USER VARIABLES
##############################################################################
VERSION="1.12.5"
BISSETII_GIT_URL="${BISSETII_GIT_URL:-"https://gitlab.com/zoralab/bissetii.git"}"
HOSTNAME="${HOSTNAME:-"localhost"}"
PORT="${PORT:-"8080"}"


##############################################################################
# APP VARIABLES
##############################################################################
action=""
object_type=""
arguments=""
target_branch=""
git_worktree_mounted=false
repo_path="$PWD"

dot_sites_path="${repo_path}/.sites"
configs_path="${dot_sites_path}/config"
themes_path="${dot_sites_path}/themes"
assets_path="${dot_sites_path}/assets"
bissetii_path="${themes_path}/bissetii"
docs_path="${repo_path}/docs"
data_path="${docs_path}/.data"
publish_path=""

hugo="$(type -p hugo)"
local_hugo="${repo_path}/.bin/hugo"
themes_config="${themes_path}/.themes"

report_ok="\033[1;32m---OK---\033[0m"

critical_directories_and_files=(
	".gitignore"
	"config"
	"archetypes"
	"i18n"
)

language_list=(
	"en-us"
	"zh-cn"
)


##############################################################################
# FUNCTIONS LIBRARY
##############################################################################
__print_status() {
	__status_mode="$1" && shift 1
	__status_message=""
	case "$__status_mode" in
	error)
		__status_message="\033[1;31m[  ERROR ] ${@}\033[0m"
		;;
	warning)
		__status_message="\033[1;33m[ WARNING ] ${@}\033[0m"
		;;
	info)
		__status_message="\033[1;34m[ INFO ] ${@}\033[0m"
		;;
	plain)
		__status_message="$@"
		;;
	*)
		return 0
		;;
	esac

	1>&2 echo -e "${__status_message}"
	unset __status_mode __status_message
}


##############################################################################
# PRIVATE FUNCTIONS
##############################################################################
_check_arguments() {
	__print_status info "checking given Hugo command's arguments..."
	if [ "$arguments" = "" ]; then
		__print_status error "the given arguments for Hugo is empty."
		exit 1
	fi
	__print_status info "$report_ok"
}


_check_available_git_branch() {
	_target="$1"
	__print_status info "checking '$_target' branch availability..."
	if [ -z "$_target" ]; then
		__print_status error "branch not specified. See --help."
		exit 1
	fi

	_remote_branch="$(git branch -r | grep "$_target")"
	_local_branch="$(git branch -l | grep "$_target")"
	if [ -z "$_remote_branch" ] && [ -z "$_local_branch" ]; then
		__print_status error "missing $_target branch."
		exit 1
	elif [ -z "$_remote_branch" ]; then
		__print_status error "$_target branch not available on remote."
		exit 1
	fi
	unset _remote_branch _local_branch _target
	__print_status info "$report_ok"
}


_check_dot_sites_existence() {
	__print_status info "checking local .sites existences..."
	if [ -d "${dot_sites_path}" ]; then
		__print_status error "there's already a local .sites generator."
		exit 1
	fi
	__print_status info "$report_ok"
}


_check_dot_sites_readiness() {
	__print_status info "checking local .sites existences..."
	if [ ! -d "${dot_sites_path}" ]; then
		__print_status error "missing .sites generator. Stopping Here."
		exit 1
	fi
	__print_status info "$report_ok"
}


_check_git_dependency() {
	__print_status info "checking Git..."
	if [ "$(type git)" = "" ]; then
		__print_status error "missing git program."
		exit 1
	fi
	__print_status info "$report_ok"
}


_check_hugo_dependency() {
	__print_status info "checking Hugo..."

	if [ -e "$local_hugo" ]; then
		__print_status info \
		"local ./${local_hugo##*${repo_path}/} detected. Deploying..."
		hugo="$local_hugo"
	fi

	if [ -z "$hugo" ]; then
		__print_status error "missing hugo program."
		exit 1
	fi

	if [ ! -z "$("$hugo" version | grep extended)" ]; then
		__print_status info "Hugo extended functions are available."
	fi

	__print_status info "$report_ok"
}


_check_publish_artifact_readiness() {
	if [ ! -d "$publish_path" ]; then
		__print_status error "missing ${publish_path}. Did you build?"
		exit 1
	fi
}


_close_git_worktree() {
	if [ $git_worktree_mounted = true ]; then
		__print_status info "unmounting worktree..."
		git worktree remove "$target_branch" &> /dev/null
		git_worktree_mounted=false
		rm -rf "$target_branch" &> /dev/null
		__print_status info "$report_ok"
	fi
}


__create_dot_themes_file() {
	echo "\
# list your theme url here separated by space. It follows this pattern:
#                    <URL> <branch> <optional: tag>
#
# if the line is not comforming to the pattern, it will be ignored.
#
#
# example:
# https://gitlab.com/ZORALab/bissetii master v1.0.0
################################################################################
https://gitlab.com/ZORALab/bissetii master \
" > "$themes_config"
}


_create_gitlabci_yaml_for_publish() {
	__print_status info "create .gitlab-ci.yml for publish branch..."
	echo "\
image: debian:latest

pages:
    stage: build
    tags:
        - linux
    environment:
        name: production
    only:
        refs:
            - $target_branch
    artifacts:
        paths:
            - public
    script:
        - mkdir -p public
        - shopt -s extglob
        - mv !(public|.*) public
" > "${target_branch}/.gitlab-ci.yml"
	__print_status info "$report_ok"
}


_get_publish_path() {
	config_filepath="${configs_path}/_default/config.toml"
	__print_status info "getting publish path from $config_filepath"
	ret=""
	old_IFS="$IFS"
	while IFS='' read -r line || [ -n "$line" ]; do
		if [[ "$line" == *publishDir* ]]; then
			ret="$line"
			break
		fi
	done < "$config_filepath"
	IFS="$old_IFS" && unset old_IFS

	ret="${ret#*\"}"
	publish_path="${ret%%\"*}"
	if [ "$publish_path" == "" ]; then
		__print_status error "missing publish path."
		exit 1
	fi

	publish_path="${publish_path#*./}"
	publish_path="${repo_path}/${publish_path}"
	__print_status info "publish path is: $publish_path"
	__print_status info "$report_ok"
}


_get_themes() {
	if [ ! -f "$themes_config" ]; then
		__print_status info "setting up theme config and directory..."
		mkdir -p "$themes_path"
		cd "$themes_path"
		__create_dot_themes_file
		__print_status info "$report_ok"
	fi

	__print_status info "getting all listed themes..."
	if [ ! -f "$themes_config" ]; then
		__print_status error "missing .themes file."
		exit 1
	fi
	cd "$themes_path"

	_verdict=0
	old_IFS="$IFS"
	while IFS='' read -r line || [ -n "$line" ]; do
		# skip comment line
		if [ "${line:0:1}" = "#" ]; then
			continue
		fi

		# process line data
		_url="${line%% *}"
		_remainder="${line#* }"
		_branch="${_remainder%% *}"
		_remainder="${_remainder#* }"
		_tag=""
		if [ "$_remainder" != "$_branch" ]; then
			_tag="${_remainder%% *}"
		fi

		# skip bad data
		if [ "$_url" = "" ] || [ "$_branch" = "" ]; then
			continue
		fi


		__print_status info "detected $_url. Setting up now..."
		_theme_name="${_url##*/}"
		_theme_path="${themes_path}/${_theme_name}"
		if [ ! -d "$_theme_path" ]; then
			cd "$_theme_directory"
			git clone "$_url"
			if [ $? -ne 0 ]; then
				__print_status error "unable to git clone."
				_verdict=1
				continue
			fi
		fi
		__print_status info "$report_ok"


		__print_status info "switching to $_branch branch..."
		cd "${_theme_path}"
		git checkout "$_branch"
		if [ $? -ne 0 ]; then
			__print_status error "unable to checkout branch."
			_verdict=1
			continue
		fi
		__print_status info "$report_ok"


		__print_status info "pull the latest updates from upstream..."
		git pull
		if [ $? -ne 0 ]; then
			__print_status error "unable to git pull."
			_verdict=1
			continue
		fi
		__print_status info "$report_ok"


		if [ "$_tag" != "" ]; then
			__print_status info "checkout $_tag tag..."
			git checkout tags/"$_tag"
			if [ $? -ne 0 ]; then
				__print_status error "unable to switch tag."
				_verdict=1
				continue
			fi
			__print_status info "$report_ok"
		fi
	done < "$themes_config"
	IFS="$old_IFS" && unset old_IFS

	if [ $_verdict -eq 1 ]; then
		__print_status error "One or more theme modules had failed."
		exit 1
	fi
}


_open_git_worktree() {
	__print_status info "setting up git worktree..."
	_close_git_worktree &> /dev/null
	rm -rf "$target_branch" &> /dev/null
	git worktree prune
	git worktree add "$target_branch" "$target_branch"
	git_worktree_mounted=true
	__print_status info "$report_ok"
}


_shift_404() {
	__print_status info "shifting 404.html to root directory..."
	_get_publish_path

	# get primary language path
	config_path="./config/_default/languages.toml"
	ret=""
	old_IFS="$IFS"
	while IFS='' read -r line || [ -n "$line" ]; do
		if [[ "$line" == [* ]]; then
			ret="$line"
			break
		fi
	done < "$config_path"
	IFS="$old_IFS" && unset old_IFS

	ret="${ret#*[}"
	lang="${ret%%]*}"
	if [ "$lang" == "" ]; then
		__print_status warning "FAILED - missing primary lang-code."
		return 0
	fi


	# copy 404.html to root location
	ret=($(find "$publish_path" -type f -name '404.html' | grep "$lang"))
	if [ "${#ret[@]}" != "0" ]; then
		__print_status info "copying ${ret[0]} to ${publish_path}/404.html"
		cp "${ret[0]}" "${publish_path}/404.html"
	fi

	__print_status info "$report_ok"
	unset ret lang
}


_unpack_bissetii_theme_module() {
	for i in "${critical_directories_and_files[@]}"; do
		__print_status info "unpacking $i"
		__pathing="${bissetii_path}/$i"
		if [ -d "$__pathing" ]; then
			cp -r "$__pathing" "$dot_sites_path"
		elif [ -f "$__pathing" ]; then
			cp "$__pathing" "$dot_sites_path"
		else
			mkdir -p "${dot_sites_path}/${i}"
			touch "${dot_sites_path}/${i}/.gitkeep"
		fi
		__print_status info "$report_ok"
		unset __pathing
	done


	__print_status info "unpacking assets/css"
	mkdir -p "${assets_path}/css"
	cp "${bissetii_path}/assets/css/"* "${assets_path}/css" 2> /dev/null
	__print_status info "$report_ok"

	__print_status info "unpacking docs"
	for lang in "${language_list[@]}"; do
		__pathing="${docs_path}/$lang"
		mkdir -p "$__pathing"
		touch "${__pathing}/.gitkeep"
		unset __pathing
	done
	__print_status info "$report_ok"


	__print_status info "unpacking docs/.data"
	if [ -d "${bissetii_path}/data" ]; then
		cp -r "${bissetii_path}/data" "$data_path"
	else
		mkdir -p "$data_path"
		touch "${data_path}/.gitkeep"
	fi
	__print_status info "$report_ok"
}


##############################################################################
# PUBLIC FUNCTIONS
##############################################################################
build() {
	__print_status info "building hugo artifacts..."
	_check_dot_sites_readiness
	_check_hugo_dependency
	_get_themes

	# starting the build
	cd "$dot_sites_path"
	$hugo --minify
	if [ $? -ne 0 ]; then
		__print_status error "build failed."
		exit 1
	fi

	_shift_404
	__print_status info "build completed."
}


close_app() {
	_close_git_worktree
	cd "$repo_path"
}
trap close_app EXIT


exec_command() {
	_check_dot_sites_readiness
	_check_hugo_dependency
	_check_arguments

	cd "$dot_sites_path"
	$hugo $arguments
	cd "$repo_path"
}


get_date() {
	date +"%Y-%m-%dT%T%:z"
}


get_themes() {
	_check_dot_sites_readiness
	_check_git_dependency
	_get_themes
}


install() {
	_check_git_dependency
	_check_dot_sites_existence
	_get_themes
	_unpack_bissetii_theme_module
	__print_status info "DONE - bissetii setup completed."
}


print_version() {
	echo $VERSION
}


publish() {
	_check_dot_sites_readiness
	_check_git_dependency
	_get_publish_path
	_check_available_git_branch "$target_branch"
	_check_publish_artifact_readiness
	_open_git_worktree


	if [ "$arguments" == "clean" ]; then
		__print_status info "cleaning up publishing branch..."
		cd "$target_branch"
		git reset --hard "$(git log '--format=format:%H' | tail -1)"
		git clean -fd
		cd ..
		__print_status info "$report_ok"
	fi


	__print_status info "placing latest artifacts..."
	cp -r "${publish_path}"/* "${target_branch}"
	__print_status info "$report_ok"


	if [ -e "${repo_path}/.gitlab-ci.yml" ]; then
		__print_status info "detected local .gitlab-ci.yml."
		_create_gitlabci_yaml_for_publish
	fi


	__print_status info "publish artifact to publishing branch..."
	cd "$target_branch"
	git add .
	git commit -m "Build output as of $(git log '--format=format:%H' -1)"
	git push -f origin "$target_branch"
	cd ..
	__print_status info "$report_ok"


	_close_git_worktree
}


run() {
	_check_dot_sites_readiness
	_check_hugo_dependency

	__print_status info "launching hugo server..."
	cd "$dot_sites_path"

	$hugo server --buildDrafts \
		--disableFastRender \
		--bind "$HOSTNAME" \
		--baseURL "$HOSTNAME" \
		--port "$PORT"
}


uninstall() {
	if [ -d "$dot_sites_path" ]; then
		rm -rf "$dot_sites_path" &> /dev/null
	fi

	if [ -d "$docs_path" ]; then
		__print_status info "ACTION NEEDED:
1. docs/ directory has user-specific documents data. Please delete it manually.
"
	fi
}


################################
# CLI Parameters and Help      #
################################
print_help() {
	echo "\
BISSETII
A maintainable way to manage Bissetii Hugo Theme and Go Template Module.
--------------------------------------------------------------------------------
To use: $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -B, --build			build the Hugo static website.

2. -cmd, --command		execute hugo with the specified commands.

3. -D, --date			get Hugo compatible date.

4. -g, --get			get all latest themes.

5. -h, --help			print this program's help messages.

6. -i, --install		setup Bissetii local 'repo-docs' Hugo static
				website generator.

7. -P, --publish		publish the current built artifacts to the
				publishing branch.
				OPTIONAL VALUES
				1. ''  --as in nothing is provided
					publish will retain past history and add
					a new entry on top of it.
				2. 'clean'
					publish will remove all history and add
					only a single entry.

				COMPULSORY ARGUMENTS
				1. -t, --target [branch-name]
					the branch name responsible for
					publications (e.g. 'gl-pages' or
					'gh-pages').

7. -r, --run			run Hugo server for local 'repo-docs'.
				OPTIONAL ARGUMENTS
				1. -b, --bind [hostname/domain name]
					hostname or domain name.
					Default is \$HOSTNAME or 'localhost'
					Example:
					./manager.sh -r -b 'http://example.com'

				2. -p, --port [number]
					port number.
					Default is '8080'
					Example:
					./manager.sh -r -p '12345'

8. -u, --uninstall		uninstall the local bissetii Hugo repo.

9. -v, --version		print app version.
"
}

run_action() {
case "$action" in
"b")
	build
	;;
"cmd")
	exec_command
	;;
"d")
	get_date
	;;
"g")
	get_themes
	;;
"h")
	print_help
	;;
"i")
	install
	;;
"p")
	publish
	;;
"r")
	run
	;;
"u")
	uninstall
	;;
"v")
	print_version
	;;
*)
	__print_status error "invalid command"
	return 1
	;;
esac
}

process_parameters() {
while [[ $# != 0 ]]; do
case "$1" in
-b|--bind)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		HOSTNAME="$2"
		shift 1
	fi
	;;
-B|--build)
	action="b"
	;;
-cmd|--command)
	action="cmd"
	shift 1
	arguments="${@}"
	shift "${#@}"
	;;
-D|--date)
	action="d"
	;;
-g|--get)
	action="g"
	;;
-h|--help)
	action="h"
	;;
-i|--install)
	action="i"
	if [[ "$2" != "" && "${2:0:1}" != '-' ]]; then
		object_type="$2"
		shift 1
	fi
	;;
-p|--port)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		PORT="$2"
		shift 1
	fi
	;;
-P|--publish)
	action="p"
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		arguments="$2"
		shift 1
	fi
	;;
-r|--run)
	action="r"
	;;
-t|--target)
	if [[ "$2" != "" && "${2:0:1}" != "-" ]]; then
		target_branch="$2"
		shift 1
	fi
	;;
-u|--uninstall)
	action="u"
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
