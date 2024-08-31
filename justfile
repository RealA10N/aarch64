TestImageName := "a10n/aarch64codegen-testenv"
TestDockerfile := "Dockerfile.testenv"
Workdir := "/aarch64codegen"
TestCommand := "go test ./... -coverprofile=coverage.out"
CoverageFile := "coverage.out"

test:
    #!/usr/bin/env bash
    set -euxo pipefail

    Container=$(docker run -di {{TestImageName}})

    docker cp . $Container:{{Workdir}}
    docker exec -it $Container /bin/sh -c "{{TestCommand}}"
    docker cp $Container:{{Workdir}}/{{CoverageFile}} {{CoverageFile}}

    docker kill $Container
    docker rm $Container

testenv-build:
    docker build -f {{TestDockerfile}} -t {{TestImageName}} .

testenv-push: testenv-build
    docker push {{TestImageName}}
