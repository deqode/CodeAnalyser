#!/bin/sh
root=$PWD
path[0]="plugin/go/detectRuntime"
path[1]="plugin/go/env"
path[2]="plugin/go/db/postgres/V_1_X"
path[3]="plugin/go/framework/beego/v1_x"
path[4]="plugin/go/getDependencies"
path[5]="plugin/go/orm/gorm/V_1_X"
path[6]="plugin/go/libraries/kafka/V1X"
path[7]="plugin/go/preDetectCommand"
path[8]="plugin/go/staticAssets"
path[9]="plugin/go/buildDirectory"
path[10]="plugin/go/testCasesCommands"
path[11]="plugin/globalDetectors/docker"
path[12]="plugin/globalDetectors/makeFile"
path[13]="plugin/globalDetectors/procFile"
path[14]="plugin/go/commands"

for p in "${path[@]}"; do
    rm ${p}/plugin -f
#    git rm -r ${p}/plugin
done

 for p in "${path[@]}"; do
     cd ${p}
     go build   -o plugin
     cd ${root}
 done