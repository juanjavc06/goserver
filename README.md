# Goserver

Este proyecto es un ejemplo sencillo de servidor HTTP escrito en Go que implementa una pequeña API para administrar roles de usuario. No se utilizan librerías externas, solo la librería estándar de Go.

## Estructura de carpetas

- **internal/domain**: Entidades de dominio y errores.
- **internal/infrastructure**: Implementaciones concretas de repositorios (en memoria).
- **internal/application**: Servicios y DTOs que encapsulan la lógica de negocio.
- **internal/interfaces**: Handlers HTTP que exponen los endpoints REST.

## Librerías utilizadas y su función

- **`net/http`**: crea el servidor HTTP y permite registrar funciones que atienden cada ruta.
- **`encoding/json`**: serializa y deserializa estructuras para comunicarnos en JSON con los clientes.
- **`sync`**: proporciona el tipo `Mutex` que asegura que el repositorio en memoria sea seguro en entornos concurrentes.
- **`testing`**: facilita la escritura de pruebas unitarias para verificar el comportamiento de los módulos.

## Ejecutar el proyecto

```bash
# Iniciar el servidor
$ go run main.go
```

## Ejecutar pruebas

```bash
$ go test ./...
```

La salida debería mostrar que todas las pruebas finalizan correctamente.
