mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))
root_dir :=  $(realpath $(mkfile_dir)../)

gen_all: gen-wireguard gen-nfs-server

# wireguard client/server dir
wg_dir := $(root_dir)/wireguard/

core_mod := github.com/jdinabox/alpine-dockerfiles

gen-wireguard:
	go run ./gen-dockerfile/main.go \
		-c=$(core_mod) \
		-m=wireguard \
		-r=$(root_dir) \
		-w=data/wireguard \
		-a="wireguard-tools"
	go run ./gen-makefile/main.go \
		-n="alpine-wireguard" \
		-m=wireguard \
		-r=$(wg_dir)

# nfs server dir
nfs_server_dir := $(root_dir)/nfs-server/

gen-nfs-server:
	go run ./gen-dockerfile/main.go \
		-c=$(core_mod) \
		-m=nfs-server \
		-r=$(root_dir) \
		-w=data/nfs \
		-a="wireguard-tools nfs-utils"
	go run ./gen-makefile/main.go \
		-n="alpine-nfs-server" \
		-m=nfs-server \
		-r=$(nfs_server_dir)