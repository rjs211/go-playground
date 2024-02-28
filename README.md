# go-playground


1. create base docker image using `docker build -f BaseDockerfile   -t vatsa/gopl-sweph-base:latest .`
2. run docker compouse using the build flag so it builds everytime. `docker compose up --build`