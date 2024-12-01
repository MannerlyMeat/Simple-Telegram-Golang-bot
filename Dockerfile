FROM golang:1.23 AS build

RUN mkdir /app
WORKDIR /app
COPY . .
RUN ls -l

FROM alpine:latest
#В TOKEN вписывается токен телеграм бота
ENV TOKEN= 
COPY --from=build /app .

CMD ["./main"]