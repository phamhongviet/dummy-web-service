FROM busybox
COPY dummy-web-service /
ENTRYPOINT ["/dummy-web-service"]
CMD [""]
