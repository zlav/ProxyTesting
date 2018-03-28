#!/bin/bash

# Shell script automatically updates the gorouter with the specified route
# every 60 seconds



# while getopts ":u:" options; do
#   case $options in
#     u) URL=${OPTARG};;
#     :)
#       echo "ERROR -option -${OPTARG} requires argument"
#   esac
# done
#
# if [[ ${OPTIND} -eq 1 ]]; then
#   echo "ERROR - no option specifies"
#   exit 1
# fi

gnatsd &

# Path on Pivotal workstations
~/workspace/routing-release/bin/gorouter &
sleep 6

while :
do
  nats-pub 'router.register' '{"host":"127.0.0.1","port":4567,"uris":["my_first_url.vcap.me","my_second_url.vcap.me"],"tags":{"another_key":"another_value","some_key":"some_value"}}'
  sleep 60
done
