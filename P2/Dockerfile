# Etapa de compilación
FROM golang:1.23 AS builder
WORKDIR /app
# Copia los archivos de módulos y descarga dependencias
COPY go.mod  ./
RUN go mod download
# Copia el resto del código fuente
COPY . .
# Compila la aplicación en modo estático
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Etapa de ejecución
FROM scratch
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
