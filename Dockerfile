FROM scratch

COPY platform-monitoring /monitoring

CMD ["/monitoring", "--port", "80"]

EXPOSE 80
