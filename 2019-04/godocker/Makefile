clean:
	docker system prune -f -a
	docker image rm godocker-onbuild
	docker image rm godocker-latest
	docker image rm godocker-scratch
	docker image rm godocker-static
	docker image rm godocker-staticssl
	docker image prune -f -a
	rm main

all: onbuild latest scratch static staticssl

onbuild:
	docker build -t godocker-onbuild -f Dockerfile.onbuild .

latest:
	docker build -t godocker-latest -f Dockerfile.latest .

scratch:
	go build -o main .
	docker build -t godocker-scratch -f Dockerfile.scratch .

static:
	CGO_ENABLED=0 GOOS=linux go build -a -o main .
	docker build -t godocker-static -f Dockerfile.scratch .

staticssl:
	CGO_ENABLED=0 GOOS=linux go build -a -o main .
	docker build -t godocker-staticssl -f Dockerfile.scratchssl .
