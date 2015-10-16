FROM scratch

COPY monitoring /monitoring

CMD ["/monitoring", "--port", "80"]

EXPOSE 80
