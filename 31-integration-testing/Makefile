.PHONY: up down restart test

up:
	docker-compose up -d --build

down:
	docker-compose down

restart: down up

test:
	set -e ;\
	docker-compose -f docker-compose.test.yml up --build -d ;\
	test_status_code=0 ;\
	docker-compose -f docker-compose.test.yml run integration_tests go test || test_status_code=$$? ;\
	docker-compose -f docker-compose.test.yml down ;\
	exit $$test_status_code ;\

test-cleanup:
  	docker-compose down \
		--rmi local \
		--volumes \
		--remove-orphans \
		--timeout 60; \
  	docker-compose rm -f
