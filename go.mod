module wangsw.go

require github.com/armstrongli/go-bmi v0.0.1

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.3.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

//本地replace的方式
replace github.com/armstrongli/go-bmi => .\staging\src\github.com\armstrongli\go-bmi //把下载的代码进行修改，引用的时候还是使用之前的路径

// 引用一个新的包时，先写在import里，前面加_,然后写在require里，然后go mod tidy就会把这个包下载下来，放在modules引用的包目录，可以使用
// 如果是要本地拓展该包，则创建相应目录，完成代码后，replace该包，就可以引用了

go 1.17
