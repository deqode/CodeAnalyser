#!/bin/sh
root=$PWD
path[0]="plugin/go/detectRuntime"
path[1]="plugin/go/env"
path[2]="plugin/go/db/postgres/V_1_X"
path[3]="plugin/go/framework/beego/v1_x"
path[4]="plugin/go/getDependencies"
path[5]="plugin/go/orm/gorm/V_1_X"

for p in "${path[@]}"; do
    rm ${p}/plugin -f
#    git rm -r ${p}/plugin
done

 for p in "${path[@]}"; do
     cd ${p}
     go build   -o plugin
     cd ${root}
 done