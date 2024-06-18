FROM golang:1.22-bullseye

RUN mkdir /app
WORKDIR /app
RUN cd /app
COPY . .
RUN go build .

CMD ["./filmigobot"]