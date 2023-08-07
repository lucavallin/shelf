FROM golang:latest

ARG PORT=80
COPY ./yak-webshop .
ENV PORT=${PORT}
EXPOSE ${PORT}

ENTRYPOINT ["./yak-webshop", "serve"]