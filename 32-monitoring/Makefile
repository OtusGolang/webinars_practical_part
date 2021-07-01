.PHONY: up

up:
	docker-compose up

cleanup:
	docker-compose down \
        --rmi local \
		--volumes \
		--remove-orphans \
		--timeout 60; \
  	docker-compose rm -f
