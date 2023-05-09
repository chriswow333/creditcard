
#Stage 1 - Install dependencies and build
FROM --platform=linux/amd64 golang:1.17.6-alpine as builder

#RUN mkdir -p /app

WORKDIR /go/src
ENV DOCKER_DEFAULT_PLATFORM=linux/amd64

COPY . ./
# COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o appv1


# Stage 2 - Create the run-time image
FROM --platform=linux/amd64 scratch

ENV DOCKER_DEFAULT_PLATFORM=linux/amd64


ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=z20339
ENV POSTGRES_HOST=10.17.128.3
ENV POSTGRES_PORT=5432


ENV GIN_MODE=release

WORKDIR /server

COPY --from=builder /go/src/appv1 ./

EXPOSE 8080

CMD ["./appv1"]

#EXPOSE 8080

#USER chris

#RUN chmod +x appv1

#ENTRYPOINT ["./appv1"]
#CMD [ "go run main.go" ]

# docker build -t apiv1 .

# docker run -it --rm -p 8080:8080 apiv1

# docker buildx build --platform linux/amd64 .

# docker run --rm -it --add-host=host.docker.internal:host-gateway 

# docker run --name postgres-test -v ~/test_project/postgres-data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=z20339 -p 5432:5432 -d postgres

# sudo docker run -e POSTGRES_PASSWORD=z20339 -v ~/pgdata:/var/lib/postgresql/data -p 5432:5432 -d postgres
#  docker tag appv1 asia-east1-docker.pkg.dev/woven-woodland-373807/pincard-api/appv1