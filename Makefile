# github.com/Mariiel/classmoods/cmd/caca
# for testing individual modules
# go test -run TestMyFunction ./...
go: compile run
	@echo 'Compiled and runned'

compile:
	go build -o ./bin/caca -gcflags='all=-N -l' ./cmd/caca   2> ./errors.err
	
c:
	go build -o ./bin/caca -gcflags='all=-N -l' ./cmd/caca
	
run:
	@env --chdir=./test ../bin/caca steins_gate lampda

compile_test:
	go test -v ./cmd/caca 2> ./errors.err

test:
	go test -v ./cmd/caca
	
test_f:
	go test -v -run $(FN) ./cmd/caca

debug:debug
	env --chdir=./cmd/caca gdlv debug steins_gate lampda
	
debugt:debugt
	env --chdir=./cmd/caca gdlv test
	
debug_f:
	env --chdir=./cmd/caca gdlv test -run $(FN)
	
	
# implement to gdlv debugging a specific test function
# debug_f:debug
# 	dlv test --build-flags='github.com/Mariiel/classmoods/cmd/google2' -- -test.run ^TestSiiaScrapper$
