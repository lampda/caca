r:compile run

compile:
	@go build -o ./bin/caca ./cmd/caca
	
run:
	@clear
	./bin/caca

compile_test:
	@go test -v ./cmd/caca

test:
	go test -v ./cmd/caca
	
test_f:
	go test -v -run $(FN) ./cmd/caca

debug:debug
	env --chdir=./cmd/caca gdlv debug &
	
debugt:debugt
	env --chdir=./cmd/caca gdlv test &
	
debug_f:
	env --chdir=./cmd/caca gdlv test -run $(FN) &
