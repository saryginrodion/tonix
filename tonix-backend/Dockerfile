FROM golang:bookworm

WORKDIR /tonix-backend

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

ENTRYPOINT ["go", "run", "."]
