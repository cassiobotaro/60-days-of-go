# Grand Finale

## Steps to reproduce

### Compile go into a shared c library

`go build -o awesome.so -buildmode=c-shared awesome.go`

### Compile the C code, specifying the shared object library

`gcc -o awesome awesome.c ./awesome.so`


### Run

`./awesome`
