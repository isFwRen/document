## 安装go-zero需要的插件
- goctl安装文档：https://go-zero.dev/docs/tasks/installation/goctl
- goctl检查: goctl env check --install --verbose --force
- goctl
    - GO111MODULE=on go install github.com/zeromicro/go-zero/tools/goctl@latest
- protoc
    - 按这里进行安装：https://grpc.io/docs/protoc-installation/
    - PB_REL="https://github.com/protocolbuffers/protobuf/releases"
    - curl -LO $PB_REL/download/v25.1/protoc-25.1-linux-x86_64.zip
    - unzip protoc-25.1-linux-x86_64.zip -d $HOME/.local
    - goctl安装的protoc，不会附带include文件夹，而且goctl安装的protoc会放在go sdk的bin里
    - 需要手动下载protoc，解压后把include文件夹放到对应地方
    - /root/.gvm/pkgsets/go1.21.4/global/bin/protoc
    - include目录和bin同级
    - 将protoc的include文件夹，复制到：mv include /root/.gvm/pkgsets/go1.21.4/global
- protoc-gen-grpc-gateway
- protoc-gen-openapiv2
- protoc-gen-validate


```shell
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```

安装validate插件
```
go install github.com/envoyproxy/protoc-gen-validate@latest
```

卸载
```
go get github.com/golang/protobuf/protoc-gen-go@none
```