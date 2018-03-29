#!/bin/bash
#
# Script to send custom header to a specific host
# Created by David Riedel <david.riedel@adesso.de> and
# External dependencies:

# - wget
# - mktemp
#

function usage() {
	echo 
	echo "Usage: ${0##*/} -u URL -h header -s Header size in Byte -i Number of Requests"
	echo 
	echo "Example: ./big-header-tester.sh -u serviceintern.muc.allianz -h svrid -s 100000 -i 100"
	echo
}


while getopts ":u:h:s:i:" options; do
  case $options in
    u) URL=${OPTARG};;
    h) HEADER=${OPTARG};;
    s) SIZE=${OPTARG};;
    i) ITERATIONS=${OPTARG};;
    ## option without a required argument
    :)
       echo "ERROR - option -${OPTARG} requires an argument"
       usage
       exit 1
       ;;
    ## unknown option
    \?)
       echo "ERROR - unknown option -${OPTARG}"
       usage
       exit 1
       ;;
    ## this should never happen
     *)
       echo "ERROR - there's an error in the matrix!"
       usage
       exit 1
       ;;
  esac 
done

if [[ ${OPTIND} -eq 1 ]]; then
	echo "ERROR - no command line option specified"
	usage 
	exit 1
fi


echo -e "header = $HEADER:\c" > .wgetrc

headerval="A"
echo "Generating header value"
for i in `seq 1 $SIZE`;
   do
	headerval=$headerval"A"
	#echo -e "A\c">>.wgetrc
   done

echo '' > response-code.txt

header="$HEADER: $headerval"

echo $header > response-code.txt

echo "Loop URL $URL"
for i in `seq 1 $ITERATIONS`;
   do
         sleep 2
	 curl -ki -H "$header" $URL >> response-code.txt 2>&1
   done

echo "Done - check File response-code.txt for Results" 

rm .wgetrc 