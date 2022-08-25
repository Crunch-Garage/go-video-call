#DEV

build-dev:
	docker build -t Crunch-Garage/go-video-call -f containers/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn .

clean-dev:
	docker-compose -f containers/composes/dc.dev.yml.down

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up