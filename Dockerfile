FROM alpine:latest
WORKDIR /build

COPY main ./
# COPY go.mod go.sum ./

# EXPOSE 3000
# EXPOSE 8000

# COPY ./.env ./build/
# RUN go mod tidy

# RUN go mod download

# RUN GOOS=linux GOARCH=amd64 go build -v -o ./main ./cmd  

# RUN ls -l
# RUN go build main.go
# COPY ../../.env /
CMD ["./main"]