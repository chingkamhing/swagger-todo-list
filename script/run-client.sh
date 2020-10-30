#!/bin/bash
#
# Script file to run todo-list client.
#

OPTS=""
NUM_ARGS=0

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to run todo-list client."
	echo
	echo "Usage: $SCRIPT_NAME [description]"
	echo "Options:"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
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

# run todo-list client
# note: invoke "todo-list client --help" for all the flag description
./todo-list client --scheme http --port 8888
