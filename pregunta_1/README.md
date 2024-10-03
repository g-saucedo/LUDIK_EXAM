<p align="center"><a href="https://laravel.com" target="_blank"><img src="https://raw.githubusercontent.com/laravel/art/master/logo-lockup/5%20SVG/2%20CMYK/1%20Full%20Color/laravel-logolockup-cmyk-red.svg" width="400" alt="Laravel Logo"></a></p>


## Prueba tecnica pregunta 1

Dentro de la prueba se ha trabajado bajo la versión [10.48.22]('') en donde se ha desarrollado lo siguiente.

- Creación de la base de datos dentro de un docker file.
- A travez del comando ´php artisan make:model -m´ se a creado la tabla de ´products´.
- Dentro del archivo ´App\Models´ se ha añadido una configuración el cual protege la tabla de la base de datos.
- Como siguiente paso se creado el controlador correspondiente para realizar las operaciones CRUD.
- Una vez ya creado las operaciones CRUD, se crea las rutas correspondientes por operación con el fin de acceder a las funciones, .

## Instrucciones de ejecución

Para ejecutar el proyecto por primera vez ese necesario tener las siguientes consideraciones:

1. Tener instalado php y composer, en las versiones 8.3.12 y 2.7.9 respectivamente.
2. Una vez instalado todo, es necesario crear la base de datos dentro de mysql con el siguiente nombre 'laravel_db'.
3. Ya creada la base de datos necesitamos configurar los enviroments, dentro del archivo '.ENV', seguir las pautas segun el archvio empezando con el nombre de la base de datos.

    DB_CONNECTION=mysql
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_DATABASE=laravel_db
    DB_USERNAME=symfony_user
    DB_PASSWORD=symfony_password

4. Ya configurados nuestro env. abriremos nuestra terminal y ejecutaremos los siguietes comandos.

- composer install. para instalar dependencias.
- php artisan migrate. para migrar las tablas en la base de datos.
- php artisan serve. para poder ejecutar el proyecto

## apis

- **[GET]** lista total de productos http://127.0.0.1:8000/api/products

- **[GET]** obtener un producto segun id http://127.0.0.1:8000/api/products/{id}.

- **[POST]** crear un producto http://127.0.0.1:8000/api/products.
   JSON DE ENVIO => {
    "name": "Producto 1",
    "description": "Descripción del producto 1",
    "price": 100.50,
    "stock": 10
}

- **[PUT]** actualizar un producto segun id http://127.0.0.1:8000/api/products/{id}.
   JSON DE ENVIO => {
    "name": "Producto 1",
    "description": "Descripción del producto 1",
    "price": 100.50,
    "stock": 10
}

- **[DELETE]** eliminar un producto segun id;
