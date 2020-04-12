FROM golang:1.14 AS build

COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 go build

FROM golang:1.14 as final
COPY --from=build /src/ecmake /ecmake

ENTRYPOINT ["/ecmake"]