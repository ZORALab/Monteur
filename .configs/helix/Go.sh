#!/bin/bash
PACKAGE_NAME="Go"
PACKAGE_VERSION="1.16.5"
_GO_GET_LIST=(
	"golang.org/x/perf/cmd/benchstat"
	"gitlab.com/zoralab/godocgen/cmd/godocgen"
)

__generate_package_name() {
	__machine="$1"
	__arch="$2"

	__name="go${PACKAGE_VERSION}.${__machine}-${__arch}"

	if [ "$__machine" == "windows" ]; then
		__name="${__name}.zip"
	else
		__name="${__name}.tar.gz"
	fi

	echo "$__name"
	unset __name __machine __arch
}


__generate_url() {
	echo "https://golang.org/dl/${package_name}"
}


__supported_version() {
	__unsupported_os="unsupported operating system."
	__unsupported_arch="unsupported cpu architecture."

	case "$arch" in
	amd64)
		case "$machine" in
		darwin)
			;;
		freebsd)
			;;
		linux)
			;;
		windows)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	386)
		case "$machine" in
		freebsd)
			;;
		linux)
			;;
		windows)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	armv6l)
		case "$machine" in
		linux)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	arm64)
		case "$machine" in
		linux)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	ppc64le)
		case "$machine" in
		linux)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	s390x)
		case "$machine" in
		linux)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	*)
		_print_status error "$__unsupported_arch"
		exit 1
		;;
	esac
}


__wrap_up() {
	__config_path="${output_path}/configs/goconfig"
	__golang_path="${output_path}/golang/bin"
	__goroot="${output_path}/golang"
	__gopath="${output_path}/gopath"
	__gocache="${output_path}/gocache"
	__goenv="${output_path}/goenv"
	__gobin="${__gopath}/bin"

	mv "${output_path}/go" "${output_path}/golang"
	mkdir -p "${__gobin%/*}" "$__gocache" "$__goenv"

	# write pathing script
	echo "\
#!/bin/bash
export GOLANG_PATH=\"${__golang_path}\"
export GOROOT=\"${__goroot}\"
export GOPATH=\"${__gopath}\"
export GOBIN=\"${__gobin}\"
export GOCACHE=\"${__gocache}\"
export GOENV=\"${__goenv}\"
export PATH=\"\${PATH}:\${GOROOT}/bin:\${GOPATH}:\${GOBIN}\"

stop_go() {
	PATH=:\${PATH}:
	GOROOT=\"\${GOROOT}/bin\"
	PATH=\${PATH//:\$GOLANG_PATH:/:}
	PATH=\${PATH//:\$GOROOT:/:}
	PATH=\${PATH//:\$GOBIN:/:}
	PATH=\${PATH//:\$GOPATH:/:}
	PATH=\${PATH%:}
	unset GOLANG_PATH GOROOT GOPATH GOBIN GOCACHE GOENV
}

case "\$1" in
-h|--help|help)
	2>&1 echo \"HELP
1) to start: $ source \$0
2) to stop : $ \$0 --stop\"
	;;
--stop)
	stop_go
	1>&2 echo \"localized Go stopped.\"
	;;
--start)
	if [ ! -z \"\$(type -p go)\" ] && [ ! -z \"\$(type -p gofmt)\" ]; then
		1>&2 echo \"localized Go started.\"
	else
		1>&2 echo \"[ ERROR ] Go failed to initalized.\"
		stop_go
		exit 1
	fi
	;;
*)
	1>&2 echo \"[ ERROR ] unknown command.\"
	return 1
	;;
esac" > "$__config_path"
	chmod +x "$__config_path"

	# go get all the key software
	source "$__config_path" --start
	for _repo in "${_GO_GET_LIST[@]}"; do
		_print_status info "go getting $_repo"
		go get "$_repo"
	done
	"$__config_path" --stop


	# check root program dependencies
	_list=(
		"gcc    # gcc package"
		"gvgen  # graphviz package"
	)
	for _software in "${_list[@]}"; do
		if [ -z "$(type -p $_software)" ]; then
			_print_status warning "missing $_software in this OS."
		fi
	done


	unset __config_path __goroot __gopath __gocache __gobin __goenv
}
