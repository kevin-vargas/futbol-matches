DIR=${CURDIR}/scripts

dev:
	go run backend/main.go
start:
	docker compose up -d
stop:
	docker compose down
test:
	cd backend; \
	${DIR}/test.sh;
coverage:
	cd backend; \
	${DIR}/coverage.sh; 
deploy-prod:
	@./publish/build.sh registry.cloud.okteto.net/kevin-vargas prod
deploy-dev:
	@./publish/build.sh docker.fast.ar dev
sec:
	kubectl apply -f ./secrets/secrets.yml