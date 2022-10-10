protoc:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		src/main.proto

build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o target/main src/*.go

zip: build
	zip -j target/function.zip target/main

create: zip
	aws lambda create-function \
		--function-name feature_flags_go \
		--zip-file fileb://target/function.zip \
		--runtime go1.x \
		--role arn:aws:iam::885495516625:role/cargo-lambda-role-ece8c25b-520c-452c-bf1d-eba8db7383db \
		--handler main

update: zip
	aws lambda update-function-code \
		--function-name feature_flags_go \
		--zip-file fileb://target/function.zip 

