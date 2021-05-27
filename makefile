int-test: int-test-docker-compose-up int-test-docker-compose-exec int-test-docker-compose-down

int-test-docker-compose-down:
	docker-compose down

int-test-docker-compose-up:
	docker-compose up -d
	docker-compose exec -T redis sh < ./pre-test-script/redis.sh

int-test-docker-compose-exec:
	docker-compose exec go-app sh -c "ginkgo -r -tags integration"
