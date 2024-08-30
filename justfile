test:
    docker run --rm -it $(docker build -q -f ./Dockerfile.test .)