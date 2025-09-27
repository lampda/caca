compile:
	go build -o ./bin/nya -gcflags='all=-N -l' ./cmd/nya   2> ./errors.err
	
c:
	go build -o ./bin/nya -gcflags='all=-N -l' ./cmd/nya
	
run:
	./bin/nya

compile_test:
	go test -v ./cmd/nya 2> ./errors.err

test:
	go test -v ./cmd/nya
	
test_f:
	go test -v -run $(FN) ./cmd/nya

debug:debug
	env --chdir=./cmd/nya gdlv debug
	
debugt:debugt
	env --chdir=./cmd/nya gdlv test
	
debug_f:
	env --chdir=./cmd/nya gdlv test -run $(FN)
