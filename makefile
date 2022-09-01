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