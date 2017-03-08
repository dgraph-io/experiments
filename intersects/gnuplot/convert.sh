cat ./tues-afternoon.bench | awk '{if (length($0) > 0) { print $2"."$4,$8,$10}}' > /tmp/bench.data
cat /tmp/bench.data | awk '{print > $1".dat"}'
