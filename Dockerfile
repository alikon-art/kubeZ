FROM alpine:3.18
LABEL authors="alikon"
WORKDIR /kubez
COPY kubez /kubez/
EXPOSE 8080
CMD ["./kubez"]
