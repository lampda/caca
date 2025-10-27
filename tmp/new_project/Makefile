r:compile run
compile:
	@go build -o ./bin/tmp/new_project ./cmd/tmp/new_project
	
run:
	@clear
	./bin/tmp/new_project

compile_test:
	@go test -v ./cmd/tmp/new_project

test:
	go test -v ./cmd/tmp/new_project
	
test_f:
	go test -v -run $(FN) ./cmd/tmp/new_project

debug:debug
	env --chdir=./cmd/tmp/new_project gdlv debug &
	
debugt:debugt
	env --chdir=./cmd/tmp/new_project gdlv test &
	
debug_f:
	env --chdir=./cmd/tmp/new_project gdlv test -run $(FN) &
