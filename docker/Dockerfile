FROM ubuntu:24.04

RUN apt-get update && apt-get install -y curl python3 python3-pip python3-venv

COPY secyourflow /usr/local/bin/secyourflow

EXPOSE 8080

VOLUME /usr/local/bin/secyourflow_data

HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

COPY docker/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]