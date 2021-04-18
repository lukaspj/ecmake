FROM scratch
COPY ecmake /ecmake
ENTRYPOINT ["/ecmake"]