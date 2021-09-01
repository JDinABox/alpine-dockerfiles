mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))
root_dir :=  $(realpath $(mkfile_dir)../)

# wireguard client/server dir
wg_dir := $(root_dir)/wireguard/

gen-wireguard:
	go run ./gen-dockerfile/main.go \
		-m=github.com/jdinabox/alpine-dockerfiles/wireguard \
		-r=$(wg_dir) \
		-w=data/wireguard \
		-a="wireguard-tools"
	go run ./gen-makefile/main.go \
		-n="alpine-wireguard" \
		-r=$(wg_dir)
