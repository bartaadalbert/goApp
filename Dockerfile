from golang as builder

WORKDIR /devops

COPY cmd .

RUN CGO_ENABLED=0 go build -o test_app

FROM scratch

ADD ./html /html

COPY --from=builder /devops/test_app .

ENTRYPOINT [ "/test_app" ]

EXPOSE 8080