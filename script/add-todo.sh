#!/bin/bash
#
# Script file to use curl command to add a to-do task.
#

URL="http://localhost"
PORT="8888"
ENDPOINT="api/todo"
OPTS="-X POST"
NUM_ARGS=1

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to use curl command to add a to-do task."
	echo
	echo "Usage: $SCRIPT_NAME [description]"
	echo "Options:"
	echo " -k                           Allow https insecure connection"
	echo " -u  [url]                    IMS Customer Portal URL"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
	"k")
		OPTS="$OPTS -k"
		;;
	"u")
		URL=$2
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

# trim URL trailing "/"
if [ "$PORT" = "" ]; then
	URL="$(echo -e "${URL}" | sed -e 's/\/*$//')"
else
	URL="$(echo -e "${URL}:${PORT}" | sed -e 's/\/*$//')"
fi

DESCRIPTION=$1

# list all users' info
curl $OPTS -vd "{\"description\":\"$DESCRIPTION\"}" -H "Content-Type: application/json" -H "Accept: application/json" ${URL}/${ENDPOINT}
