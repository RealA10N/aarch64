export DOCKER_TEST_IMAGE := "aarch64codegen-test"

test: build
    docker run --rm -it $DOCKER_TEST_IMAGE

build:
    docker build -f ./Dockerfile.test -t $DOCKER_TEST_IMAGE .
