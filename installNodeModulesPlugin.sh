#!/bin/sh
root=$PWD
path[0]="plugin/js/commands"
path[1]="plugin/js/db/cassandra/v1_x"
path[2]="plugin/js/db/mariadb/v1_x"
path[3]="plugin/js/db/microsoftSQLServer/v1_x"
path[4]="plugin/js/db/mongodb/v1_x"
path[5]="plugin/js/db/mysql/v1_x"
path[6]="plugin/js/db/oracle/v1_x"
path[7]="plugin/js/db/postgres/v1_x"
path[8]="plugin/js/db/redis/v1_x"
path[9]="plugin/js/db/sqlite/v1_x"
path[10]="plugin/js/detectRunTime"
path[11]="plugin/js/env"
path[12]="plugin/js/framework/angular/v1_x"
path[13]="plugin/js/framework/express/v1_x"
path[14]="plugin/js/framework/gatsby/v1_x"
path[15]="plugin/js/framework/nest/v1_x"
path[16]="plugin/js/framework/next.js/v1_x"
path[17]="plugin/js/framework/vue/v1_x"
path[18]="plugin/js/getDependencies"
path[19]="plugin/js/testCasesCommands"
path[20]="plugin/js/libraries/kafka/v1_x"
path[21]="plugin/js/libraries/knex/v1_x"
path[22]="plugin/js/libraries/react/v1_x"
path[23]="plugin/js/preDetectCommands"
path[24]="plugin/js/orm/bookshelf/v1_x"
path[25]="plugin/js/libraries/mongoose/v1_x"
path[26]="plugin/js/orm/sequelize/v1_x"
path[27]="plugin/js/orm/typeorm/v1_x"
# path[28]=""
# path[29]=""

for p in "${path[@]}"; do
    cd ${p}
    npm install
    cd ${root}
done
