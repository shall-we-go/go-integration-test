int-test: int-test-docker-compose-up int-test-run int-test-docker-compose-down

int-test-docker-compose-down:
	docker-compose down

int-test-docker-compose-up:
	docker-compose up --detach
	docker-compose exec -T redis sh < ./pre-test-script/redis.sh

int-test-run:
	docker-compose exec go-app sh -c "ginkgo -r -tags integration"
