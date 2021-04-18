ARG BASE_TAG

FROM golang:${BASE_TAG}
COPY ecmake /ecmake
ENTRYPOINT ["/ecmake"]