set -e

cd cc

swig -I/usr/local/include -go -c++ -cgo -intgosize 64 -outdir ../ tmp.i

g++ -c -fPIC -O2 -std=c++11 tmp_wrap.cxx

ar -r libtmp.a tmp_wrap.o

cd ..

python add_cgo_flags.py tmp.go

go build .
