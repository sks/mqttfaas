#!/bin/sh -e

OUTPUT_TO_TOPIC=${FIRED_BY}/output

function to_fahrenheit(){
    local tc=$(cat "-")
    tf=$(awk -v tc=$tc 'BEGIN { print (tc * 9 / 5) + 32 }')
    echo "{\"topic\":\"$OUTPUT_TO_TOPIC\",\"data\": \"${tf}\"}"
}

function to_celsius(){
    local tf=$(cat "-")
    tc=$(awk -v tf=$tf 'BEGIN { print (tf - 32 ) * 5 / 9 }')
    echo "{\"topic\":\"${OUTPUT_TO_TOPIC}\",\"data\": \"${tf}\"}"
}

function main(){
    to_fahrenheit $@
}

main $@