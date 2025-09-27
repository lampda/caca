compile:
	go build -o ./bin/uwu -gcflags='all=-N -l' ./cmd/uwu   2> ./errors.err
	
c:
	go build -o ./bin/uwu -gcflags='all=-N -l' ./cmd/uwu
	
run:
	./bin/uwu

compile_test:
	go test -v ./cmd/uwu 2> ./errors.err

test:
	go test -v ./cmd/uwu
	
test_f:
	go test -v -run $(FN) ./cmd/uwu

debug:debug
	env --chdir=./cmd/uwu gdlv debug
	
debugt:debugt
	env --chdir=./cmd/uwu gdlv test
	
debug_f:
	env --chdir=./cmd/uwu gdlv test -run $(FN)
