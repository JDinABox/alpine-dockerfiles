mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))
root_dir :=  $(realpath $(mkfile_dir)../)

gen_all: gen-wireguard gen-nfs-server

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

# nfs server dir
nfs_server_dir := $(root_dir)/nfs-server/

gen-nfs-server:
	go run ./gen-dockerfile/main.go \
		-m=github.com/jdinabox/alpine-dockerfiles/nfs-server \
		-r=$(nfs_server_dir) \
		-w=data/nfs \
		-a="wireguard-tools nfs-utils"
	go run ./gen-makefile/main.go \
		-n="alpine-nfs-server" \
		-r=$(nfs_server_dir)