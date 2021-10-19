#!/bin/bash

printf "| %30s | %10s | %10s | %10s | %10s |\n" "TEST NAME" "SEED" "PASSED" "WEAK" "FAILED"
for eachTests in ./test_results/*
do
    passed=$(grep -o -i PASSED ${eachTests} | wc -l)
    weak=$(grep -o -i WEAK ${eachTests} | wc -l)
    failed=$(grep -o -i FAILED ${eachTests} | wc -l)
    fields=$(echo ${eachTests} | cut -d "/" -f 3)
    fields=$(echo ${fields} | cut -d "." -f 1)
    testName=$(echo $fields | awk -F"__" '/__/{print $1}')
    seeds=$(echo $fields | awk -F"__" '/__/{print $2}')
    seeds=$(echo ${seeds} | cut -d "_" -f 2)
    printf "| %30s | %10d | %10s | %10s | %10s |\n" ${testName} ${seeds} ${passed} ${weak} ${failed}
done
