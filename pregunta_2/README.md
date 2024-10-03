# CRUD de productos en Go

## Requisitos

 - Go

## Instalación

1. clona el proyecto desde el repositorio.
2. Navega al directorio del proyecto.
3. Ejecuta `go get` para instalar las dependencias.
4. Ejecuta `go run main.go` para iniciar la aplicación

## Uso

Accede a la API en `http://localhost:8080/products`.

### Endpoints

- **POST /products**: Crear un producto.
- **GET /products**: Obtener todos los productos.
- **GET /products/{id}**: Obtener un producto en especifico.
- **PUT /products/{id}**: Actualizar un producto en especifico.
- **DELETE /products/{id}**: Elimina un producto en especifico.

### Validaciones

- El precio y la cantidad en stock deben ser números positivos.

## Como se utilizo go

En la busqueda de un framework o lenguaje que no conocia o habia desarrollador antes 3 opciones se me vinieron.

- Django
- Ruby
- Go

En el cual me decidi por Go, pues me llamaba la atención mucho más que python o ruby, 

Revise un poco la documentación de Go y me percate de primera instancia que la sintaxis era totalmente diferente, empezando por la comparaciones con elementos `nulos` el cual se reemplazaba con `nil`.
Trabajando muy de la mano con la documentación oficial, y ayuda de foros logre hacer una conexion con la base de datos que habia creado anteriormente. Lo siguiente era poder crear el modelo de datos de la tabla, para poder acceder a ella y hacer los crud respectivos.
Tuve algunas dificultades con los campos de `create_at` y `update_at`, por lo que termine descartandolos pues hasta donde se solo hace uso de ello laravel con su configuracion predeterminada por lo que termine de descartar la declaracion de esos campos.

una vez creado el module segui con la creacion de una api basica de POST, ya que necesitaba asegurane que los datos se envien de manera satisfactoria, lograndos seguir con la documentacion, fue relativamente sencillo.

## Recursos utilizados

- [Documentacion oficial de go]()
- Tutorial sobre la inserción de datos mediante la libreria de gorm

## Que desafico me enfrente

El principal desafio fue la sintaxis pues es algo que da un cambio de 360, con lo que estaba acostumbrado y se me hizo un poco lento hacer el desarrollo, pero practicando y siguiendo documentación, uno se acostumbra.
