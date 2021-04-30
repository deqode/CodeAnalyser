#!/bin/sh
root=$PWD
path[0]="plugin/js/commands"
path[1]="plugin/js/db/postgres/v1_x"
path[2]="plugin/js/db/redis/v1_x"
path[3]="plugin/js/detectRunTime"
path[4]="plugin/js/env"
path[5]="plugin/js/framework/express/v1_x"
path[6]="plugin/js/framework/nest/v1_x"
path[7]="plugin/js/getDependencies"
path[8]="plugin/js/preDetectCommands"
path[9]="plugin/js/libraries/kafka/v1_x"
path[10]="plugin/js/libraries/mongoose/v1_x"
path[11]="plugin/js/db/mongodb/v1_x"

for p in "${path[@]}"; do
    cd ${p}
    npm install
    cd ${root}
done
