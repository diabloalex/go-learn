# Go Learn

Proyecto de aprendizaje en Go con un servidor HTTP simple.

## Ejecutar el proyecto

```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

## Endpoints

- `GET /` - Retorna "hola mundo"
- `POST /post` - Retorna "Respuesta del handler POST"

## Ejemplos de uso

### GET
```bash
curl http://localhost:8080/
```

### POST
```bash
curl -X POST http://localhost:8080/post
```

