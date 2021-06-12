#!/bin/bash
PACKAGE_NAME="Hugo"
PACKAGE_VERSION="0.75.0"

__generate_package_name() {
	__machine="$1"
	__arch="$2"

	__name="hugo_${PACKAGE_VERSION}_${__machine}-${__arch}"

	if [ "${__machine}" == "windows" ]; then
		__name="${__name}.zip"
	else
		__name="${__name}.tar.gz"
	fi

	echo "$__name"
	unset __name
}


__generate_url() {
	_base_url="https://github.com/gohugoio/hugo/releases/download"
	echo "${_base_url}/v${PACKAGE_VERSION}/${package_name}"
}


__supported_version() {
	case "$arch" in
	arm64)
		arch=ARM64
		case "$machine" in
		linux)
			machine=Linux
			;;
		*)
			_print_status error "unsupported operating system."
			exit 1
			;;
		esac
		;;
	armv5|armv6l|armv7)
		arch=ARM
		case "$machine" in
		freeBSD)
			machine=FreeBSD
			;;
		linux)
			machine=Linux
			;;
		netbsd)
			machine=NetBSD
			;;
		openbsd)
			machine=OpenBSD
			;;
		*)
			_print_status error "unsupported operating system."
			exit 1
			;;
		esac
		;;
	amd64)
		arch=64bit
		case "$machine" in
		dragonfly)
			machine=DragonFlyBSD
			;;
		darwin)
			machine=macOS
			;;
		freeBSD)
			machine=FreeBSD
			;;
		linux)
			machine=Linux
			;;
		netbsd)
			machine=NetBSD
			;;
		openbsd)
			machine=OpenBSD
			;;
		windows)
			machine=Windows
			;;
		*)
			_print_status error "unsupported operating system."
			exit 1
			;;
		esac
		;;
	i386)
		arch=32bit
		case "$machine" in
		freeBSD)
			machine=FreeBSD
			;;
		linux)
			machine=Linux
			;;
		netbsd)
			machine=NetBSD
			;;
		openbsd)
			machine=OpenBSD
			;;
		windows)
			machine=Windows
			;;
		*)
			_print_status error "unsupported operating system."
			exit 1
			;;
		esac
		;;
	*)
		_print_status error "unsupported cpu architecture."
		exit 1
		;;
	esac
}

__wrap_up() {
	rm "${output_path}/LICENSE" &> /dev/null
	rm "${output_path}/README.md" &> /dev/null
}
