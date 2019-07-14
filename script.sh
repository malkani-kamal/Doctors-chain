#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "up 		- It will bring the existing network up"
echo "down 		- Bring the running network down"
echo "restart 	- Restarting the network"
echo "gen		- Generating the crypto"
echo "validate	- Checking all components of the network"
echo "------------------------------------------------------------"


MODE=$1
echo "$1  $MODE" 
# Determine whether starting, stopping, restarting, generating or upgrading
if [ $MODE = "up" ]; then
  echo "Bringing the existing network up"
elif [ "$MODE" == "down" ]; then
  echo "Bringing the running network down"
elif [ "$MODE" == "restart" ]; then
  echo "Restarting the network"
elif [ "$MODE" == "generate" ]; then
  echo "Generating the crypto"
elif [ "$MODE" == "validate" ]; then
  echo "Checking all components of the network"
else
#  printHelp
  exit 1
fi

exit 0