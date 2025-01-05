# Stage 1: Builder stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 


WORKDIR /app/src/api  

# LISTA OS ARQUIVOS PARA VERIFICAR (DEPURAÇÃO)
RUN ls -l

# Compila o binário (agora no diretório correto)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main . \
    && if [ ! -f main ]; then echo "Erro na compilação do binário. Verifique os logs detalhados." && exit 1; fi

# Stage 2: Final image
FROM alpine:latest

RUN addgroup -S appgroup && adduser -S -G appgroup appuser

WORKDIR /app

COPY --from=builder --chown=appuser:appgroup /app/src/api/main /app/main 

USER appuser

EXPOSE 8080

CMD ["/app/main"]