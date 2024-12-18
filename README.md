## Aplicación CLI para CRUD en Base de Datos MySQL

Esta aplicación CLI permite realizar operaciones CRUD en una base de datos MySQL. Está preparada para ser desplegada utilizando Docker, lo que facilita su configuración y uso en distintos entornos.

### Características

- Realiza operaciones CRUD (Crear, Leer, Actualizar y Eliminar) en una base de datos MySQL.
- Dockerizada para una fácil implementación.
- Conectada a un contenedor MySQL a través de una red `bridge` de Docker.

### Descarga e Instalación

Puedes descargar la imagen Docker de la aplicación desde Docker Hub:

[Repositorio Docker Hub](https://hub.docker.com/repository/docker/zxspooky/docker-repo/general)

### Configuración

1. **Base de Datos MySQL:**
   - Asegúrate de tener un contenedor MySQL en ejecución.
   - Usa los scripts proporcionados en el repositorio para crear la base de datos y las tablas necesarias.

2. **Variables de Entorno:**
   - Dentro de los archivos del proyecto, encontrarás un archivo `.env` que contiene las variables necesarias para la conexión a la base de datos. Asegúrate de configurarlo correctamente antes de ejecutar la aplicación.
3. **Creacion de images**
   -Dentro de los archivos hay un txt con el script para crear los contenedores y su creacion de la bridge para que ambos containers esten conectados entre si.

### Gracias!
