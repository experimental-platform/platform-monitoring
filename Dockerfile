FROM golang:onbuild

CMD ["/go/bin/app", "--port", "80"]

EXPOSE 80
