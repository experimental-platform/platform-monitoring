FROM golang:onbuild

CMD ["/go/bin/monitoring", "--port", "80"]

EXPOSE 80
