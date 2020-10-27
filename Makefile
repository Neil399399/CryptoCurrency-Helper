bitcoin-testnet:
	docker-compose -f docker-compose.yaml up -d bitcoin-testnet;

bitcoin-testnet-down:
	docker stop bitcoin-testnet && docker rm bitcoin-testnet

docker-omni:
	docker build -t neil/omnicore:latest -f ./docker/dockerfile.omni .

start-vault:
	docker-compose -f docker-compose.yaml up -d vault; \
	sleep 5;
	./vault/script/init.sh ; \
	./vault/script/deploy.sh

stop-vault:
	docker stop vault && docker rm vault

start-omni:
	docker-compose -f docker-compose.yaml up -d omni-testnet;

stop-omni:
	docker stop omni-testnet && docker rm omni-testnet