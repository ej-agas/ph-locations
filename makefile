docker:
	docker build -t ph-locations:latest . -f build/Dockerfile.dev
	docker compose -f build/docker-compose.yaml up -d --build

app:
	docker build -t ph-locations:latest . -f build/Dockerfile.dev
	docker compose -f build/docker-compose.yaml up -d --build app

prod:
	docker build -t ph-locations:latest . -f build/Dockerfile.prod
	docker compose -f build/docker-compose.yaml up -d --build
