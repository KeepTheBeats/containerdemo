#!/bin/env bash

# example: bash run.sh "./containerdemo" "-inputfrom" "config" "-config" "para4Sh.conf"
# This script will firstly render placeholders, and then execute the command.


CURRENT_DIR=$(cd $(dirname $0); pwd)

# render the parameters in the config file
cp -f ${CURRENT_DIR}/para4ShOri.conf ${CURRENT_DIR}/para4Sh.conf
sed -i "s/\[\[PARAMETER1\]\]/${PARAMETER1}/g" ${CURRENT_DIR}/para4Sh.conf
sed -i "s/\[\[PARAMETER2\]\]/${PARAMETER2}/g" ${CURRENT_DIR}/para4Sh.conf

# run the application
cd ${CURRENT_DIR}
exec "$@"
