FROM alpine:3.18
LABEL authors="alikon"
WORKDIR /kubez
COPY kubez /kubez/
# 安装gilbc兼容层
RUN apk add gcompat
EXPOSE 8080
CMD ["./kubez"]
