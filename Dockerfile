FROM quay.io/experimentalplatform/ubuntu:latest

COPY platform-monitoring /monitoring

CMD ["dumb-init", "/monitoring", "--port", "80"]

EXPOSE 80
