package api

// Download protoc compile from here
// https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1
// then go get -u github.com/golang/protobuf/protoc-gen-go

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/common/common.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/common/common.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/common/common.proto

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/admin/admin.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/admin/admin.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/admin/admin.proto

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/cpc/cpc.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/cpc/cpc.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/cpc/cpc.proto

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/debug/debug.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/debug/debug.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/debug/debug.proto

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/miner/miner.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/miner/miner.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/miner/miner.proto

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/net/net.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/net/net.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/net/net.proto

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/personal/personal.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/personal/personal.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/personal/personal.proto

//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/txpool/txpool.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/txpool/txpool.proto
//go:generate protoc -I. -Iv1/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true,Mcommon/common.proto=bitbucket.org/cpchain/chain/api/v1/common:. ./v1/txpool/txpool.proto
