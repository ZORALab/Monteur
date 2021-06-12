#!/bin/bash
##############################################################################
# USER VARIABLES
##############################################################################
VERSION="1.2.0"
TEST_TIMEOUT="${TEST_TIMEOUT:-"14400s"}" # 4 hours
BENCHMARK_TARGET="${BENCHMARK_TARGET:-.}"
BENCHMARK_TIMEOUT="${BENCHMARK_TIMEOUT:-"1s"}" # 1 second

##############################################################################
# APP VARIABLES
##############################################################################
action=""
exit_code=0
file_timestamp="$(date +"%Y-%m-%d-%H-%M-%SZ%z")"

repo_path="$PWD"
programs_path="${repo_path}/.bin"
builds_path="${repo_path}/build/go"
configs_path="${repo_path}/.configs/gopher"
outputs_path="${repo_path}/test"
tests_path="${outputs_path}/testing"
benchmarks_path="${outputs_path}/benchmark/${file_timestamp}"
releases_path="${configs_path}/releases"

benchmark_profiles_path="${benchmarks_path}/profiles"
benchmark_html_path="${benchmarks_path%/*}/benchmark.html"
benchmark_block_svg_path="${benchmarks_path}/benchmark-block.svg"
benchmark_cpu_svg_path="${benchmarks_path}/benchmark-cpu.svg"
benchmark_mem_svg_path="${benchmarks_path}/benchmark-mem.svg"
benchmark_mutex_svg_path="${benchmarks_path}/benchmark-mutex.svg"
benchmark_log_profile_path="${benchmark_profiles_path}/bench-log.profile"
benchmark_block_profile_path="${benchmark_profiles_path}/bench-block.profile"
benchmark_cpu_profile_path="${benchmark_profiles_path}/bench-cpu.profile"
benchmark_mem_profile_path="${benchmark_profiles_path}/bench-mem.profile"
benchmark_mutex_profile_path="${benchmark_profiles_path}/bench-mutex.profile"

release_config_path="${releases_path}/configs"
release_target_path="${releases_path}/targets"

test_coverages_path="${tests_path}/coverages"
test_profiles_path="${tests_path}/profiles"
test_log_path="${tests_path}/test-log.html"
test_html_path="${tests_path}/test-coverage-map.html"
test_coverage_path="${tests_path}/test-coverage.html"
test_static_path="${tests_path}/test-static-analysis.html"
test_profile_path="${test_profiles_path}/test.profile"
test_coverage_profile_path="${test_profiles_path}/test-coverage.profile"

go_config="${programs_path}/goconfig"
go=""
gofmt=""
golangci_lint=""
benchstat=""

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


__print_html_head() {
	__given_location="$1"
	__target_path="$2"
	__title="$3"

	echo "\
<!DOCTYPE html>
<head>
	<title>${__title}</title>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name=generator content="BASH $BASH_VERSION" />
	<style>
		*,
		*:after,
		*:before {
			box-sizing: inherit;
		}
		html {
			height: 100%;
			box-sizing: border-box;
			font-size: 62.5%; // similar to 1.6rem = 16px
		}
		body {
			font-family: Roboto, Helvetica Neue, Helvetica, Arial, sans-serif;
			font-size: 1.6rem;
			font-weight: 365;
			letter-spacing: .01rem;
			line-height: 1.6;
			max-width: 100vw;
			width: 100%;
			padding: 0 5rem;
			margin: 0;
		}
		pre {
			overflow-x: auto;
		}
		table {
			display: block;
			overflow-x: auto;
			white-space: nowrap;
			border-collapse: collapse;
			width: 100%;
			max-width: -moz-fit-content;
			max-width: fix-content;
			margin: 0 auto 2.5rem;
			padding: 0;
		}
		table thead {
			background-color: black;
			color: white;
		}
		table tr {
			border-bottom: .1rem solid black;
		}
		table th,
		table td {
			vertical-align: middle;
			padding: 1.5rem;
		}
	</style>
</head>
<body>
<h1>${__title}</h1>
<p>
	Results generated using various <code>Go</code> programming language
	tools. Please refresh the page to get the latest results.
</p>
<h2>Test Configurations</h2>
<table>
	<thead>
		<tr>
			<th>Aspects</th>
			<th>Values</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Target</td>
			<td><pre>${__given_location}</pre></td>
		</tr>
		<tr>
			<td>Run Time</td>
			<td><pre>$(date)</pre></td>
		</tr>
		<tr>
			<td>Test Output Location</td>
			<td><pre>$tests_path</pre></td>
		</tr>
		<tr>
			<td>Compiler Location</td>
			<td><pre>$go</pre></td>
		</tr>
		<tr>
			<td>Formatter Location</td>
			<td><pre>$gofmt</pre></td>
		</tr>
		<tr>
			<td>Go Environment</td>
			<td><pre>$(go env)</pre></td>
		</tr>
	</tbody>
</table>
<br/><br/>
<hr/>" > "$__target_path"

	unset __target_path __given_location __title
}


__print_html_tail() {
	__target_path="$1"

	if [ -f "$__target_path" ]; then
		echo "</p></body>" >> "$__target_path"
	fi

	unset __target_path
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

	go="$(type -p go)"
	if [ -z "$go" ]; then
		_print_status error "missing go compiler. Did you install go?"
		exit 1
	fi

	gofmt="$(type -p gofmt)"
	if [ -z "$gofmt" ]; then
		_print_status error "missing gofmt. Did you install go?"
		exit 1
	fi
}


_check_benchstat_availability() {
	_print_status info "checking benchstat availability..."

	benchstat="$(type -p benchstat)"
	if [ ! -z "$benchstat" ]; then
		_print_status info "Deploying $benchstat..."
		_print_status ok
	else
		_print_status warning "no available benchstat detected!"
	fi
}


_check_golangci_lint_availability() {
	_print_status info "checking golangci-lint availability..."

	golangci_lint="$(type -p golangci-lint)"
	if [ -e "${programs_path}/golangci-lint" ]; then
		_print_status info "detected local golangci-lint. Deploying..."
		golangci_lint="${programs_path}/golangci-lint"
		_print_status ok
	elif [ ! -z "$golangci_lint" ]; then
		_print_status info "Deploying $golangci_lint..."
		_print_status ok
	else
		_print_status warning "no available golangci-lint detected!"
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


_run_clean() {
	_print_status info "clean up build artifacts..."
	find "$repo_path" \( \
		-name '*.test' -o \
		-name '*.out' -o \
		-name '*.d' -o \
		-name '*.o' \
		-name '*.elf' -o \
		-name '*.hex' -o \
		-name '*.pyc' -o \
		-name '*.bin' \
		\) \
		-type f \
		-delete
	_print_status ok

	_print_status info "clean up test artifacts (${outputs_path})..."
	rm -rf "$outputs_path"
	_print_status ok

	_print_status info "clean up build artifacts (${builds_path})..."
	rm -rf "$builds_path"
	_print_status ok
}


_present_test_results() {
	__list=(
		"$test_html_path"
		"$test_log_path"
		"$test_coverage_path"
		"$test_static_path"
		"$benchmark_html_path"
	)

	for __file in "${__list[@]}"; do
		_print_status info "opening (${__file##*$repo_path/})..."
		if [ ! -f "$__file" ]; then
			_print_status warning "missing $__file. Skipped."
			continue
		fi

		xdg-open "$__file" &> /dev/null
		if [ $? -ne 0 ]; then
			_print_status error "unable to open file."
		else
			_print_status ok
		fi
	done

	unset __list __file
}


_run_benchmark() {
	_print_status info "running benchmark for: $BENCHMARK_TARGET"

	# prepare benchmark output path
	mkdir -p "$benchmark_profiles_path"

	# generate HTML header
	__print_html_head "$BENCHMARK_TARGET" \
		"$benchmark_html_path" \
		"Go Benchmark Results"

	# execute benchmark
	echo "<h2>Current Log Output</h2><p><pre>" >> "$benchmark_html_path"
	go test -run=none \
		-bench="$BENCHMARK_TARGET" \
		-benchmem \
		-benchtime "$BENCHMARK_TIMEOUT" \
		-blockprofile "$benchmark_block_profile_path" \
		-cpuprofile "$benchmark_cpu_profile_path" \
		-memprofile "$benchmark_mem_profile_path" \
		-mutexprofile "$benchmark_mutex_profile_path" \
		2> /dev/null \
		| tee -a "$benchmark_log_profile_path" "$benchmark_html_path"
	echo "</pre></p>" >> "$benchmark_html_path"

	# process svg graphs
	echo "<h2>Benchmark SVG</h2>" >> "$benchmark_html_path"
	if [ ! -z "$(type -p gvgen)" ]; then
		__list=(
			"block"
			"cpu"
			"mem"
			"mutex"
		)

		for __name in "${__list[@]}"; do
			# construct code list
			__profile_path="benchmark_${__name}_profile_path"
			__svg_path="benchmark_${__name}_svg_path"
			__profile_path="${!__profile_path}"
			__svg_path="${!__svg_path}"

			if [ -f "$__profile_path" ]; then
				_print_status info "generating $__svg_path"

				echo "<h3>$__name</h3>" \
					>> "$benchmark_html_path"

				go tool pprof \
					--output "$__svg_path" \
					-svg "$__profile_path" \
					&> /dev/null

				echo "\
<img src='$__svg_path' loading='lazy' />
<a href='$__svg_path' target='_blank'>Click to Enlarge</a>
<br/><br/>" >> "$benchmark_html_path"

				_print_status success ok
			fi

			unset __profile_path __svg_path
		done
	else
		__message="missing graphviz. No svg is generated!"
		echo "<p>$__message</p>" >> "$benchmark_html_path"
		_print_status warning "$__message"
		unset __message
	fi

	# process delta statistics
	echo "<h2>Benchmark Delta</h2>" >> "$benchmark_html_path"
	if [ ! -z "$benchstat" ]; then
		__has_content=false
		for _archive in "${benchmarks_path%/*}"/*; do
			if [ "$_archive" == "$benchmarks_path" ] || \
				[ -f "$_archive" ]; then
				continue
			fi

			__new_profile="$benchmark_log_profile_path"
			__old_profile="${__new_profile##*${benchmarks_path}/}"
			__old_profile="${_archive}/${__old_profile}"
			if [ ! -f "$__old_profile" ]; then
				continue
			fi
			__has_content=true

			__old_timestamp="${_archive##*${benchmarks_path%/*}/}"
			echo "<h3>$__old_timestamp</h3><p><pre>" \
				>> "$benchmark_html_path"
			_print_status info "comparing with: $__old_timestamp"

			$benchstat "$__old_profile" "$__new_profile" \
				2> /dev/null \
				| tee -a "$benchmark_html_path"

			echo "</p></pre><br/>" >> "$benchmark_html_path"
			unset __new_profile __old_profile __old_timestamp
		done

		if [ "$__has_content" == "false" ]; then
			__message="Nothing to compare."
			echo "<p>$__message</p>" >> "$benchmark_html_path"
			_print_status info "$__message"
			unset __message
		fi
	else
		__message="missing benchstat. No benchmark delta!"
		echo "<p>$__message</p>" >> "$benchmark_html_path"
		_print_status warning "$__message"
		unset __message
	fi

	# clean up
	echo "</body></html>" >> "$benchmark_html_path"
	rm "${repo_path}/"*".test" &> /dev/null
	unset __list __svg_list __has_content
	_print_status ok
}


__run_build_subject() {
	__target="$1"
	__output_path="$2"
	__release_config_path="$3"
	__did_build=false
	__os="$GOOS"
	__arch="$GOARCH"

	for __config in "${__release_config_path}/"*; do
		__did_build=true
		GOOS=""
		GOARCH=""
		__name="${__config##*${__release_config_path}/}"
		source "$__config"

		if [ -z "$GOOS" ]; then
			_print_status error "missing \$GOOS."
			exit_code=1
			continue
		fi

		if [ -z "$GOARCH" ]; then
			_print_status error "missing \$GOARCH."
			exit_code=1
			continue
		fi

		_print_status info "building $__target $__name"
		$go build -ldflags "-s -w" \
			-o "${__output_path}/${__name}" ./...

		if [ $? -ne 0 ]; then
			_print_status error "an error occured!"
			exit_code=1
		else
			_print_status ok
		fi
	done

	# clean up
	if [ "$__did_build" == "false" ]; then
		_print_status info "nothing to build."
	fi

	if [ ! -z "$__os" ]; then
		GOOS="$__os"
	fi

	if [ ! -z "$__arch" ]; then
		GOARCH="$__arch"
	fi
	unset __did_build __os __arch __output_path __config __target
}


_run_build() {
	_print_status info "running build sequences..."

	if [ ! -f "$release_target_path" ]; then
		_print_status error "missing ${release_target_path##*${repo_path}/}"
		exit_code=1
		return 1
	fi
	source "$release_target_path"

	if [ ! -d "$release_config_path" ]; then
		_print_status error "missing release configs. Stopped."
		exit_code=1
		return 1
	fi

	for __target in "${BUILD_TARGETS[@]}"; do
		__output_path="${builds_path}/${__target##*/}"

		_print_status info ""
		_print_status info "building for $__target ..."

		if [ ! -d "$__target" ]; then
			_print_status error "Not a directory. Skipped."
			exit_code=1
			continue
		fi

		ret=($(find "$__target" -maxdepth 1 -name "*.go"))
		if [ ${#ret[@]} -eq 0 ]; then
			_print_status error "Not a go package. Skipped."
			exit_code=1
			continue
		fi

		_print_status info "cleaning up $__output_path ..."
		rm -f "$__output_path" &> /dev/null
		mkdir -p "$__output_path"
		_print_status ok

		cd "$__target"
		__run_build_subject "$__target" \
			"$__output_path" \
			"$release_config_path"
		cd "$repo_path"
		_print_status info ""
	done
}


_run_static_analysis() {
	__location="$1"

	_print_status info "running static analysis for: $__location"

	# prepare test output file
	mkdir -p "${test_static_path%/*}"

	# begin static analysis
	__print_html_head "$__location" \
		"$test_static_path" \
		"GolangCI Lint Static Analysis"
	echo "<h2>Current Log Output</h2><p><pre>" >> "$test_static_path"

	ret=$(2>&1 "$golangci_lint" run "$__location")
	if [ -z "$ret" ]; then
		ret="no error found. All Cleared!"
	fi
	echo "$ret" | tee -a "$test_static_path"

	echo "</pre></p>" >> "$test_static_path"
	__print_html_tail "$test_static_path"

	_print_status ok
}


_run_test() {
	_print_status info "running go test for: $__location"
	__location="$1"

	# prepare test output file
	mkdir -p "${test_profiles_path}"
	__print_html_head "$__location" "$test_log_path" "Go Test Results"
	echo "<h2>Current Log Output</h2><p><pre>" >> "$test_log_path"

	# validate location if it is a single package
	if [ "${arg##*/}" != "..." ]; then
		ret=($(find $arg -maxdepth 1 -name "*.go"))
		if [ ${#ret[@]} -eq 0 ]; then
			__message="$__location has no .go files"
			_print_status error "$__message"
			echo "$__message</pre></p>" >> "$test_log_path"
			__print_html_tail "$test_log_path"

			unset __message
			exit_code=1
			return 1
		fi
	fi

	# run test
	go test -timeout "$TEST_TIMEOUT" \
		-coverprofile "$test_profile_path" \
		-race \
		-v \
		"$__location" \
		| tee -a "$test_log_path"
	if [ ${PIPESTATUS[0]} -ne 0 ]; then
		exit_code=1
	fi


	echo "</pre></p>" >> "$test_log_path"
	__print_html_tail "$test_log_path"

	if [ $? -ne 0 ]; then
		unset __location
		exit_code=1
		return 1
	fi

	# process results
	if [ -f "$test_profile_path" ]; then
		2>&1 go tool cover \
			-html="$test_profile_path" \
			-o "$test_html_path" \
			> /dev/null
		if [ $? -ne 0 ]; then
			exit_code=1
		fi
	fi

	# wrapping up
	unset __location
	_print_status ok
}


_run_test_coverage() {
	__location="${1:-.}"

	_print_status info "running go test coverage statistics: $__location"
	__packages=($(go list "${__location}/..."))
	__mode="count"
	__verdict=0


	# prepare coverage workspaces
	mkdir -p "${test_profiles_path}"
	rm -rf "$test_coverages_path" &> /dev/null
	mkdir -p "$test_coverages_path"
	__print_html_head "$__location" \
		 "$test_coverage_path" \
		"Go Test Coverage Statistics"
	echo "<h2>Current Log Output</h2><p><pre>" >> "$test_coverage_path"


	# execute individual test coverages
	for __package in "${__packages[@]}"; do
		__subject="${test_coverages_path}/${__package//\//-}.cover"

		go test -timeout 20m \
			-covermode="$__mode" \
			-coverprofile="$__subject" \
			"$__package" \
			| tee -a "$test_coverage_path"
		if [ ${PIPESTATUS[0]} -ne 0 ]; then
			__verdict=1
		fi
	done
	unset __package __subject __packages


	# consolidate data
	echo "mode: $__mode" > "$test_coverage_profile_path"
	grep -h -v "^mode:" "$test_coverages_path"/*.cover \
		>> "$test_coverage_profile_path"
	go tool cover -func="$test_coverage_profile_path" \
		| tee -a "$test_coverage_path"


	# wrapping up
	echo "</pre></p>" >> "$test_coverage_path"
	__print_html_tail "$test_coverage_path"
	unset __location __mode

	if [ $__verdict -eq 0 ]; then
		_print_status ok
	else
		_print_status error "error occured!"
		exit_code=1
		exit 1
	fi
}


##############################################################################
# PUBLIC FUNCTIONS
##############################################################################
benchmark_repo() {
	_print_status info "starting gopher helper benchmark sequences..."
	_check_go_availability
	_check_benchstat_availability
	trap _close_app EXIT
	_run_benchmark
	_print_status success "Done. Goodbye!"
}


build_repo() {
	_print_status info "starting gopher helper build sequences..."
	_check_go_availability
	_run_build
	trap _close_app EXIT
	_print_status success "Done. Goodbye!"
}


clean_repo() {
	_print_status info "starting gopher helper clean sequences..."
	_run_clean
	_print_status success "Done. Goodbye!"
}


print_version() {
	echo $VERSION
}


present_data() {
	_print_status info "starting gopher helper present sequences..."
	_present_test_results
	_print_status success "Done. Goodbye!"
}


start_develop() {
	_print_status info "starting gopher development environment..."
	_check_go_availability
	_check_golangci_lint_availability
	_check_benchstat_availability
	_print_status success "Done. Goodbye!"
}


stop_develop() {
	_print_status info "stopping gopher development environment..."
	_close_app
	_print_status success "Done. Goodbye!"
}


test_repo() {
	_exit_code=0
	_print_status info "starting gopher helper test sequences..."
	_check_go_availability
	_check_golangci_lint_availability
	trap _close_app EXIT
	_run_static_analysis "${repo_path}/..."
	_run_test "${repo_path}/..."
	_run_test_coverage "$repo_path"
	_print_status success "Done. Goodbye!"
}


################################
# CLI Parameters and Help      #
################################
print_help() {
	echo "\
GOPHER
A Go repository helper to keep the processes sane.
--------------------------------------------------------------------------------
To use: $0 [ACTION] [ARGUMENTS]

ACTIONS
1. -B, --benchmark		start gopher benchmark sequences.


2. -b, --build			start gopher build sequences.


3. -c, --clean			start gopher clean up sequences.


4. -d, --develop		start gopher development environment.


5. -h, --help			print this program's help messages.


6. -p, --present		present gopher data.


7. -sd, --stop-develop		stop gopher development environment.


8. -t, --test			start gopher test sequences.


9. -v, --version		print app version.
"
}

run_action() {
case "$action" in
"B")
	benchmark_repo
	;;
"d")
	start_develop
	;;
"b")
	build_repo
	;;
"c")
	clean_repo
	;;
"h")
	print_help
	;;
"p")
	present_data
	;;
"sd")
	stop_develop
	;;
"t")
	test_repo
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
-B|--benchmark)
	action="B"
	if [ "$2" != "" ] && [ "${2:1}" != "-" ]; then
		BENCHMARK_TARGET="$2"
		shift 1
	fi
	;;
-d|--develop)
	action="d"
	;;
-b|--build)
	action="b"
	;;
-c|--clean)
	action="c"
	;;
-h|--help)
	action="h"
	;;
-p|--present)
	action="p"
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
