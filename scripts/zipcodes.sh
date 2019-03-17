#!/bin/sh

OS=`uname`
DL="wget"

if [ "$OS" == "FreeBSD" ]; then
        DL="fetch"
fi

CWD=`pwd`
TMP=`mktemp -d`
cd $TMP

$DL http://www2.census.gov/geo/docs/maps-data/data/gazetteer/2018_Gazetteer/2018_Gaz_zcta_national.zip
echo -n '{ "data": [' > extract.txt ; unzip -p 2018_Gaz_zcta_national.zip | tail -n +2 | awk '{printf "%s ", $1; if (NF > 1) print $(NF-1),$NF ; else print $NF; }' | awk '{printf "{\"zipcode\": \"%s\",\"latitude\":\"%s\",\"longitude\":\"%s\"},",$1,$2,$3}' >> extract.txt ; cat extract.txt | sed s'/.$//' > $CWD/zipcodes.json ; echo ']}' >> $CWD/zipcodes.json  
rm -rf $TMP

