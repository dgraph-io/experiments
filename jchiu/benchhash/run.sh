set -e

NUMCPU=5

mkdir -p results
for q in 10 100 1000; do
  go test -cpu $NUMCPU -benchn 100000 -benchq $q -bench=. > results/benchhash.q$q.txt
done

