build:
	cp -f ./nginx/Dockerfile .
	docker build -t test-nginx .
	rm -f Dockerfile
	cp -f ./php-fpm/Dockerfile .
	docker build -t test-php .
	rm -f Dockerfile
run:
	docker-compose up -d
stop:
	docker-compose down