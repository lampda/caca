r:compile run
compile:
	@go build -o ./bin/new_project ./cmd/new_project
	
run:
	@clear
	./bin/new_project

compile_test:
	@go test -v ./cmd/new_project

test:
	go test -v ./cmd/new_project
	
test_f:
	go test -v -run $(FN) ./cmd/new_project

debug:debug
	env --chdir=./cmd/new_project gdlv debug &
	
debugt:debugt
	env --chdir=./cmd/new_project gdlv test &
	
debug_f:
	env --chdir=./cmd/new_project gdlv test -run $(FN) &
