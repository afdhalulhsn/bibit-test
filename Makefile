.PHONY: clean protoc setup build grpc

clean:
	@echo "--- Cleanup all build and generated files ---"
	@rm -vf app/infrastructure/grpc/proto/*.pb.go
	@rm -vf app/infrastructure/grpc/proto/movie/*.pb.go
	@rm -vf app/infrastructure/grpc/proto/movie/*.pb.gw.go
	@rm -vf app/infrastructure/swagger/*.json

	@rm -vf ./main

protoc: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p app/infrastructure/grpc/proto/movie
	@mkdir -p app/infrastructure/swagger


	@echo "--- Compiling all proto files ---"
	@cd ./app/shared/proto/movie && protoc -I. --go_out=plugins=grpc:../../../../app/infrastructure/grpc/proto/movie --govalidators_out=../../../../app/infrastructure/grpc/proto/movie *.proto
	@cd ./app/shared/proto/movie && protoc -I . --grpc-gateway_out ../../../../app/infrastructure/grpc/proto/movie \
                                        --grpc-gateway_opt logtostderr=true \
                                        --grpc-gateway_opt paths=source_relative \
                                        *.proto
	@cd ./app/shared/proto/movie && protoc  --swagger_out=logtostderr=true:../../../../app/infrastructure/swagger movie.proto

	@gofmt -s -w .

run-app:
	@echo "--- running Aplication server in dev mode ---"
	@ go run cmd/server/main/main.go grpc