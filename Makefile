.DEFAULT_GOAL := buildandrun

BIN_FILE=life

buildandrun:
	@go build -o "${BIN_FILE}" .
	./${BIN_FILE}

install:
	@go install .
