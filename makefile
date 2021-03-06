VERSION = "unset"
DATE=$(shell date -u +%Y-%m-%d-%H:%M:%S-%Z)
DOCKER_IMAGE ?=	jackharley7/api-gateway
IMAGE_TAG ?= latest

# .PHONY: codegen
# codegen:
# 	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
# 	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
# 	go install $(shell go list -f '{{ .Dir }}' -m github.com/golang/protobuf)/protoc-gen-go
# 	protoc \
# 		--proto_path=${GOPATH}/src \
# 		--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
# 		--grpc-gateway_out=logtostderr=true:. \
# 		--swagger_out=logtostderr=true:. \
# 		--go_out=plugins=grpc:. -I . proto/conversation.proto proto/link.proto proto/invitation.proto \
# 	gofmt -w proto

codegen:
	cp ../conversationservice/proto/*.proto ./
	cp ../notificationservice/proto/*.proto ./
	cp ../userservice/proto/*.proto ./
	cp ../cmmservice/proto/*.proto ./
	cp ../commentservice/proto/*.proto ./
	cp ../upvoteservice/proto/*.proto ./
	cp ../hashtag/proto/*.proto ./
	cp ../subscriptionservice/proto/*.proto ./
	# GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	# GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	# go install $(shell go list -f '{{ .Dir }}' -m github.com/golang/protobuf)/protoc-gen-go
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		conversation.proto link.proto invitation.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		notification.proto counts.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		profile.proto user.proto userTemp.proto education.proto workexperience.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		cmm.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		comment.proto commentUpvote.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		upvote.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		hashtag.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/pkg/mod \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--govalidators_out=. \
		subscription.proto