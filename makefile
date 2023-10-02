.PHONY test:
test:
	docker build -t gossiper .
	docker compose up