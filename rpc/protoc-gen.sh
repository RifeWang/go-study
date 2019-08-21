# --proto_path 源路径，--go_out 输出路径，一定要指明 plugins=grpc
protoc --proto_path=grpc --go_out=plugins=grpc:grpc test.proto