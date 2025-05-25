# ---- Etapa de Compilación ----
FROM golang:1.24-alpine AS builder
LABEL stage=gobuilder

# Variables de entorno para la compilación
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Establecer el directorio de trabajo
WORKDIR /build

# Copiar los archivos de módulos y descargar dependencias primero
# Esto aprovecha el cache de Docker si no cambian las dependencias
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Copiar todo el código fuente del proyecto
COPY . .

# Compilar la aplicación
# Reemplaza 'gerrit-wui' si el ejecutable debe tener otro nombre
RUN go build -ldflags="-w -s" -o /app/gerrit-wui ./cmd/gerrit-wui

# ---- Etapa Final ----
FROM alpine:latest
LABEL stage=production

# Crear un usuario y grupo no root para la aplicación
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Establecer el directorio de trabajo para la aplicación
WORKDIR /app

# Copiar el binario compilado desde la etapa de 'builder'
COPY --from=builder /app/gerrit-wui /app/gerrit-wui

# Asignar el binario al usuario no root
RUN chown appuser:appgroup /app/gerrit-wui

# Cambiar al usuario no root
USER appuser

# Exponer el puerto en el que tu aplicación escucha (si es un servidor web)
# Ejemplo: EXPOSE 8080

# Comando para ejecutar la aplicación
ENTRYPOINT ["/app/gerrit-wui"]

# Si tu aplicación necesita argumentos, puedes usar CMD
# CMD ["--config", "/etc/app/config.yaml"]