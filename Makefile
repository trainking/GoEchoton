all: init userrpc authserver initctl

PACKET_PATH = ./deployments/packet
RPC_PATH = ./internal/rpc
SERVER_PATH = ./internal/server
CMD_PATH = ./cmd
BIN_PATH = ./deployments/packet/bin

init:
	cp ./scripts/deploy/start_all.sh $(PACKET_PATH)/
	cp ./scripts/deploy/stop_all.sh $(PACKET_PATH)/
	go mod tidy

userrpc:
	go build -o $(PACKET_PATH)/userrpc/userrpc $(RPC_PATH)/user/userrpc.go

authserver:
	go build -o $(PACKET_PATH)/authserver/authserver $(SERVER_PATH)/auth/authserver.go
	
initctl:
	go build -o $(BIN_PATH)/initctl $(CMD_PATH)/initctl/initctl.go

clean:
	rm -rf $(PACKET_PATH) 