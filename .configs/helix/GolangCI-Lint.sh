#!/bin/bash
PACKAGE_NAME="Golang CI Lint"
PACKAGE_VERSION="1.31.0"


__generate_package_name() {
	__machine="$1"
	__arch="$2"

	__name="golangci-lint-${PACKAGE_VERSION}-${__machine}-${__arch}"

	if [ "$__machine" == "windows" ]; then
		__name="${__name}.zip"
	else
		__name="${__name}.tar.gz"
	fi

	echo "$__name"
	unset __name
}


__generate_url() {
	__baseurl="https://github.com/golangci/golangci-lint/releases/download"
	echo "${__baseurl}/v${PACKAGE_VERSION}/${package_name}"
}


__supported_version() {
	__unsupported_os="unsupported operating system."
	__unsupported_arch="unsupported cpu architecture."

	case "$arch" in
	armv7)
		case "$machine" in
		freebsd)
			;;
		linux)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	armv6l)
		arch=armv6
		case "$machine" in
		freebsd)
			;;
		linux)
			;;
		*)
			_print_status error "$__unsupported_os"
			exit 1
			;;
		esac
		;;
	amd64)
		case "$machine" in
		freebsd)
			;;
		darwin)
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
	*)
		_print_status error "$__unsupported_arch"
		exit 1
		;;
	esac
}


__wrap_up() {
	__target="${output_path}/golangci-lint"
	__reside="${output_path}/golangci-lint-pkg"

	mv "$__target"* "$__reside"
	echo "\
#!/bin/sh
${__reside}/golangci-lint \$@" > "$__target"
	chmod +x "$__target"

	unset __target __reside
}
