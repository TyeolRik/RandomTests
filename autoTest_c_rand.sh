#!/bin/bash
# For Dieharder test

# c_rand
testName=c_rand
echo "*********************"
echo "*    ${testName} test    *"
echo "*********************"
time_start="$(date -u +%s)"
./${testName}.out 12345 | pv | dieharder -a -g 200 >> ./test_results/${testName}.txt
time_end="$(date -u +%s)"
elapsed=$(($time_end-$time_start))
echo "$elapsed seconds costs for SEED = 12345 Dieharder -a -g 200" >> ./test_results/${testName}.txt
echo "Test ends in ${elapsed} seconds!" | mail -s "Test${i} End" kino6147@gmail.com -A ./test_results/${testName}.txt