build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go

zip: build
	zip function.zip main

create: zip
	aws lambda create-function --function-name feature_flags_go --zip-file fileb://function.zip --runtime go1.x --role arn:aws:iam::885495516625:role/cargo-lambda-role-ece8c25b-520c-452c-bf1d-eba8db7383db --handler main

update: zip
	aws lambda update-function-code --function-name feature_flags_go --zip-file fileb://function.zip 

