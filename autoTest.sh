#!/bin/bash
# For Dieharder test

declare -a testList
testList=( "awc_c_no_seed" "awc_no_seed" "kiss_2011_64bit" "mt19937_64bits" "swb1_no_seed" "swb2_no_seed" )

for eachTest in ${testList[@]}; do
    echo "*********************"
    echo "*    ${eachTest} test    *"
    echo "*********************"
    time_start="$(date -u +%s)"
    ./bin/${eachTest}.out 12345 | pv | dieharder -a -g 200 >> ./test_results/${eachTest}.txt
    time_end="$(date -u +%s)"
    elapsed=$(($time_end-$time_start))
    echo "$elapsed seconds costs for SEED = 12345 Dieharder -a -g 200" >> ./test_results/${eachTest}.txt
    echo "Test ends in ${elapsed} seconds!"
done