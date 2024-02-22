FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=mms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=mms.yaml

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY --from=builder /app/mms-rpc ./
COPY --from=builder /app/etc/${CONFIG_FILE} ./etc/

EXPOSE 9104

ENTRYPOINT ./mms-rpc -f etc/${CONFIG_FILE}