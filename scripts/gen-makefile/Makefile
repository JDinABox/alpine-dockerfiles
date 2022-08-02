build:
	go build -ldflags="-s -w" -o ./cmd/app.so ./cmd

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))
root_dir :=  $(realpath $(mkfile_dir)../)

docker-build:
	# docker build --tag {{docker-name}}:latest . 
	cd $(root_dir); docker build -f {{codeModule}}.Dockerfile --tag {{docker-name}}:latest . 