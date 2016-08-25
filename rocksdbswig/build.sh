set -e

cd cc

swig -I/usr/local/include -go -c++ -cgo -intgosize 64 -outdir ../ rocksdbswig.i

g++ -c -fPIC -O2 -std=c++11 rocksdbswig_wrap.cxx

ar -r librocksdbswigwrap.a rocksdbswig_wrap.o

cd ..

python add_cgo_flags.py rocksdbswig.go
