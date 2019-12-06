bitcoin-testnet:
	docker-compose -f docker-compose.yaml up -d bitcoin-testnet;

bitcoin-testnet-down:
	docker stop bitcoin-testnet && docker rm bitcoin-testnet