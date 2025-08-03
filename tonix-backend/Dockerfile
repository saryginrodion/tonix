FROM golang:bookworm

WORKDIR /tonix-backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "."]
