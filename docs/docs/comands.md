*** Comandos Disponibles ***

# create
- ``Descripción``: Permite crear una tarea con un título y una descripción
- ``Uso``: ./task create --title "titulo de la tarea" --description "descripción de la tarea"

# list
- ``Descripción``: Nos permite visualizar las tareas creadas
- ``Uso``: ./task list

# update
- ``Descripción``: Permite actualizar el título y/o la descripción de una tarea mediante su ID
- ``Uso``: ./task update --id 1 --title " nuevo titulo de la tarea" --description "nueva descripción de la tarea"
- ``Aclaración``: debes ingresar por lo menos un item para efectuar la actualización

# delete
- ``Descripción``: Permite eliminar una tarea mediante su ID
- ``Uso``: ./task delete --id 1

# realized
- ``Descripción``: Permite marcar como realizada una tarea
- ``Uso``: ./task realized --id 1


*** Flags ***

``Las flags te ayudan a usar la aplicación con mayor fluidez``

* --id > -i
* --title > -t
* --description > -d

