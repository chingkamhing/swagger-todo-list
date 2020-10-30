#!/bin/bash
#
# Script file to run todo-list-server.
#

OPTS=""
NUM_ARGS=0

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to run todo-list-server."
	echo
	echo "Usage: $SCRIPT_NAME [description]"
	echo "Options:"
	echo " -p                           Static file server's root path"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
	"p")
		OPTS="$OPTS --static $2"
		shift
		;;
	"h")
		Usage
		exit
		;;
	esac
	shift
done

if [ "$#" -ne "$NUM_ARGS" ]; then
    echo "Invalid parameter!"
	Usage
	exit 1
fi

# run todo-list-server
# note: invoke "todo-list-server --help" for all the flag description
./todo-list-server --scheme http --port 8888 $OPTS
