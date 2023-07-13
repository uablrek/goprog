#! /bin/sh
##
## build.sh --
##
##
## Commands;
##

prg=$(basename $0)
dir=$(dirname $0); dir=$(readlink -f $dir)
tmp=/tmp/${prg}_$$

die() {
    echo "ERROR: $*" >&2
    rm -rf $tmp
    exit 1
}
help() {
    grep '^##' $0 | cut -c3-
    rm -rf $tmp
    exit 0
}
test -n "$1" || help
echo "$1" | grep -qi "^help\|-h" && help

log() {
	echo "$prg: $*" >&2
}
dbg() {
	test -n "$__verbose" && echo "$prg: $*" >&2
}

##   env
##     Print environment.
cmd_env() {
	test "$envread" = "yes" && return 0
	envread=yes

	if test -z "$__version"; then
		# Build a *correct* semantic version from date and time (no leading 0's)
		__version=$(date +%Y.%_m.%_d+%H.%M | tr -d ' ')
	fi
	
	if test "$cmd" = "env"; then
		opts="version"
		set | grep -E "^(__($opts)|ARCHIVE)="
		return 0
	fi
}
##   dynamic [--version=]
##     Build with dynamic linking
cmd_dynamic() {
	cmd_env
	mkdir -p $dir/_output
	go build -o $dir/_output -ldflags "-X main.version=$__version" \
		./cmd/... || die
	strip $dir/_output/*
}
##   static [--version=]
##     Build with static linking
cmd_static() {
	cmd_env
	mkdir -p $dir/_output
	CGO_ENABLED=0 GOOS=linux go build -o $dir/_output \
		-ldflags "-extldflags '-static' -X main.version=$__version" \
		./cmd/... || die
	strip $dir/_output/*
}


##
# Get the command
cmd=$1
shift
grep -q "^cmd_$cmd()" $0 $hook || die "Invalid command [$cmd]"

while echo "$1" | grep -q '^--'; do
	if echo $1 | grep -q =; then
		o=$(echo "$1" | cut -d= -f1 | sed -e 's,-,_,g')
		v=$(echo "$1" | cut -d= -f2-)
		eval "$o=\"$v\""
	else
		if test "$1" = "--"; then
			shift
			break
		fi
		o=$(echo "$1" | sed -e 's,-,_,g')
		eval "$o=yes"
	fi
	shift
done
unset o v
long_opts=`set | grep '^__' | cut -d= -f1`

# Execute command
trap "die Interrupted" INT TERM
cmd_$cmd "$@"
status=$?
rm -rf $tmp
exit $status
