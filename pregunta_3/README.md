# URL Processor

Este proyecto precesa una lista de URLs y encuentra aquellas que su dominio contenga la palabra "shop"

## Requisitos

- [Node.js version 18]("").

- Typescript (inslar con `npm install typescript -g`)

## Instalación

1. Clones el repositorio

2. Ejecutar `npm isntall` para instalar las dependencias.
3. Colocar el archivo con las urls en la carpeta `data/` con el nombre `urls.txt`

## Compilar y ejecutar

1. Ejecuta `npm run build` para compilar el proyecto.
2. Luego, ejecuta `npm start` para iniciar con el procesamiento.

## Instrucciones adicionales

- Asegurar el archivo `urls.txt` mantega el formato correcto para las URLs.

## Tecnicas y herramientas utilizadas

Se utilizo la paqueteria de npm en su version 18, por lo que se reviso una libreria o herramienta para poder leer archivos, por lo principal txt.
Se hace uso de `fs`, que es un modulo base de node que permite leer el contenido de los archivos, apoyado por `readline` que nos facilita la lectura de datos, que permite leer lineas individuales
Se utilizo `Set` para poder almacenar las URLs, que concuerdan con los patrones previamente creados usando `Regex`, lo cual `Set` tambien nos permite eliminar automaticamente los duplicados por lo que no hay que hace más interacciones por ese punto.
Para lograr optimizar el codigo se hace la lectura linea por linea del archivo utilizando un bucle `for` con el fin no agotar la memoria y al hacer uso de expresiones regulares solo esperamos una confimacion boolean en caso la url ya formateada cumple o no los requisitos.

## Referencias

- [fs](https://www.knowledgehut.com/blog/web-development/npm-install-fs-in-node-js)
- [readline](https://nodejs.org/api/readline.html#readlinepromisescreateinterfaceoptions)
- [Regex](https://regexr.com/)
