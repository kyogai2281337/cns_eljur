FROM alpine
RUN mkdir /app
WORKDIR /app

COPY ./auth_alpine .
COPY ./auth.toml .
CMD ["./auth_alpine"]
EXPOSE 6987