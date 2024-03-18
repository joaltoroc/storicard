.PHONY: help database-up database-down local run

help:
	@echo "Available targets:"
	@echo "  make database-up   	- Start the database container"
	@echo "  make database-down 	- Stop and remove the database container"
	@echo "  make local         	- Run the application locally"
	@echo "  make run           	- Start the database, start the application locally"
	@echo "  make down           	- Shutdown the database"

# This target waits for the MySQL container to become available
wait-for-mysql:
	@echo "Waiting for MySQL container to start..."
	@until docker compose exec mysql-db mysql -umysql -ppwd -hlocalhost -e "SELECT 1"; do \
		sleep 6; \
	done
	@echo "MySQL is up and running!"

database-up: 
	docker compose up mysql-db -d

service-up:
	docker compose up storicard-app -d

docker-down:
	docker compose down 

run: database-up service-up

down : docker-down
