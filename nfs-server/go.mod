module github.com/jdinabox/alpine-dockerfiles/nfs-server

go 1.17

// TODO: Wait for https://github.com/kubernetes/klog/pull/242
replace github.com/go-logr/logr => github.com/go-logr/logr v0.4.0

require (
	github.com/allocamelus/allocamelus v0.0.0-20210524065912-74122e28a3c0
	github.com/jdinabox/alpine-dockerfiles/wireguard v0.0.0-20210901175522-c3ffa137c8d7
	github.com/jdinabox/go-await v0.0.0-20210901041928-61062ac5156f
	github.com/jdinabox/tool-server v0.0.0-20210901182225-6268cea610b6
	github.com/json-iterator/go v1.1.11
	k8s.io/klog/v2 v2.10.0
)

require (
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/go-logr/logr v1.1.0 // indirect
	github.com/gofiber/fiber/v2 v2.18.0 // indirect
	github.com/klauspost/compress v1.13.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.29.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20210831042530-f4d43177bf5e // indirect
)
