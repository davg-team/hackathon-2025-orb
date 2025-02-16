FROM --platform=linux/amd64 golang:1.22

WORKDIR /app

EXPOSE 8081
# Download all dependencies.
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .
# RUN ls
RUN go build -o bin ./cmd/main.go

CMD [ "./bin" ]
