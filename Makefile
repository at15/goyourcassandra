.PHONY: run-c2 run-c3 down
run-c2:
	docker-compose up c2
run-c3:
	docker-compose up c3
down:
	docker-compose down
