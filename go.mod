module github.com/storezhang/drone-plugin-mcu

go 1.16

require (
	github.com/mcuadros/go-defaults v1.2.0
	github.com/storezhang/glog v1.0.8
	github.com/storezhang/gox v1.5.0
	github.com/storezhang/replace v1.0.7
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/gox => ../pangu
