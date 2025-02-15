FROM --platform=linux/amd64 golang:1.22

WORKDIR /app

# Download all dependencies.
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

EXPOSE 8080
COPY . .
# RUN ls
RUN go build -o bin ./cmd/main.go

CMD [ "./bin" ]