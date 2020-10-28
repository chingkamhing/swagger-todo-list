#!/bin/bash
#
# Script file to use curl command to get a list of todo tasks.
#

URL="http://localhost"
PORT="8888"
ENDPOINT="api/todo"
OPTS="-X GET"
NUM_ARGS=0

# default parameters
SINCE=$(date +%s)
LIMIT=10

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to use curl command to get a list of todo tasks."
	echo
	echo "Usage: $SCRIPT_NAME [description]"
	echo "Options:"
	echo " -s                           Unix time of since when to get the todo list in reverse-chronological order (default: now)"
	echo " -l                           Number of limit (default: $LIMIT)"
	echo " -k                           Allow https insecure connection"
	echo " -u  [url]                    IMS Customer Portal URL"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
	"s")
		SINCE=$2
		shift
		;;
	"l")
		LIMIT=$2
		shift
		;;
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

# convert parameter array to http parameter string
PARAMS=( \
	"since=$SINCE" \
	"limit=$LIMIT" \
)
PARAMS_STRING=`IFS="&";echo "${PARAMS[*]}";IFS=$`

# list all users' info
# curl $OPTS -v -H "Content-Type: application/json" -H "Accept: application/json" -H "Authorization: Basic YWRtaW46TXlQYXNzd29yZA==" ${URL}/${ENDPOINT}?${PARAMS_STRING}
curl $OPTS -v -H "Content-Type: application/json" -H "Accept: application/json" -H 'MY-API-KEY: MySecureAPIKey' ${URL}/${ENDPOINT}?${PARAMS_STRING}
