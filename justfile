TestImageName := "a10n/aarch64codegen-testenv"
TestDockerfile := "Dockerfile.testenv"
Workdir := "/aarch64codegen"
TestCommand := "go test ./... -coverprofile=coverage.out"
CoverageFile := "coverage.out"
DockerPlatforms := "linux/amd64,linux/arm64"

# run the package tests inside the testing docker environment container
test:
    #!/usr/bin/env bash
    set -euxo pipefail

    Container=$(docker run -di {{TestImageName}})

    docker cp . $Container:{{Workdir}}
    docker exec $Container /bin/sh -c "{{TestCommand}}"
    docker cp $Container:{{Workdir}}/{{CoverageFile}} {{CoverageFile}}

    docker kill $Container
    docker rm $Container

# build and push the test environment docker image
build-testenv:
    docker buildx build --platform {{DockerPlatforms}} \
        -f {{TestDockerfile}} \
        -t {{TestImageName}} \
        --push \
        .
