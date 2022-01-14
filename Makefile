all: init userrpc authserver

PACKET_PATH = ./deployments/packet
RPC_PATH = ./internal/rpc
SERVER_PATH = ./internal/server

init:
	go mod tidy

userrpc:
	go build -o $(PACKET_PATH)/userrpc/userrpc $(RPC_PATH)/user/userrpc.go

authserver:
	go build -o $(PACKET_PATH)/authserver/authserver $(SERVER_PATH)/auth/authserver.go

clean:
	rm -rf $(PACKET_PATH)