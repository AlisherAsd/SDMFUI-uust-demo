FROM golang:tip-20260926-trixie

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8081

CMD ["go", "run", "main.go"]