#!/bin/bash
# For Dieharder test

declare -a testList
testList=( "keccak256" )

for eachTest in ${testList[@]}; do
    echo "*********************"
    echo "*    ${eachTest} test    *"
    echo "*********************"
    time_start="$(date -u +%s)"
    ./bin/${eachTest} 12345 | pv | dieharder -a -g 200 >> ./test_results/${eachTest}.txt
    time_end="$(date -u +%s)"
    elapsed=$(($time_end-$time_start))
    echo "$elapsed seconds costs for SEED = 12345 Dieharder -a -g 200" >> ./test_results/${eachTest}.txt
    echo "Test ends in ${elapsed} seconds!" | mail -s "Test ${eachTest} End" kino6147@gmail.com -A ./test_results/${eachTest}.txt
done