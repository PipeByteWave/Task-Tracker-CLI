package crud

import (
	"RastreadorDeTareas/utils"
	"errors"
	"fmt"
	"time"
)


func CreateTask(fileName string, task utils.Task) error {
	// Leer las tareas existentes
	tasks, err := utils.ReadJson(fileName)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %v", err)
	}


	var maxId uint = 0
	for _, t := range tasks {
		if t.Id > maxId {
			maxId = t.Id
		}
	}

	task.Id = maxId + 1
	task.CreatedAt = time.Now()

	// Agregar la nueva tarea a la lista
	tasks = append(tasks, task)

	// Escribir las tareas actualizadas al archivo
	err = utils.WriteJson(fileName, tasks)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	return nil
}

func GetAllTasks(fileName string) ([]utils.Task, error) {
	// Leer el archivo JSON
	tasks, err := utils.ReadJson(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Si no hay tareas, se retorna un slice vacío sin error
	if len(tasks) == 0 {
		return nil, nil
	}

	// Retornamos las tareas leídas
	return tasks, nil
}


func UpdateTask(fileName string, id uint, title, description string) error {
	// Leer las tareas existentes
	tasks, err := utils.ReadJson(fileName)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %v", err)
	}

	// Buscar la tarea por ID
	var task *utils.Task
	for i := range tasks {
		if tasks[i].Id == id {
			task = &tasks[i]
			break
		}
	}

	if task == nil {
		return errors.New("no se encontró una tarea con el ID proporcionado")
	}

	if title != "" {
		task.Title = title
	}
	if description != "" {
		task.Description = description
	}

	now := time.Now()
	task.UpdateAt = &now

	// Escribir las tareas actualizadas al archivo
	err = utils.WriteJson(fileName, tasks)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	return nil
}

func DeleteTask(fileName string, id uint) error {
	// Leer las tareas existentes
	tasks, err := utils.ReadJson(fileName)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %v", err)
	}

	// Buscar y eliminar la tarea por ID
	var taskDeleted bool
	for i, t := range tasks {
		if t.Id == id {
			// Eliminar la tarea
			tasks = append(tasks[:i], tasks[i+1:]...)
			taskDeleted = true
			break
		}
	}

	// Si no se encontró la tarea
	if !taskDeleted {
		return fmt.Errorf("tarea con ID %d no encontrada", id)
	}

	// Escribir las tareas actualizadas al archivo
	err = utils.WriteJson(fileName, tasks)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	return nil
}

func FindTaskById(tasks []utils.Task, id uint) (*utils.Task, error) {
	for i := range tasks { // Cambié la variable de _ a i para poder modificar el slice
		if tasks[i].Id == id {
			return &tasks[i], nil // Devolvemos la dirección del elemento dentro del slice
		}
	}

	return nil, fmt.Errorf("task with id %d not found", id)
}

func FindAndDeleteTask(tasks *[]utils.Task, id uint) error {
	var taskIndex int
	var found bool
	for i, task := range *tasks {
		if task.Id == id {
			taskIndex = i
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	*tasks = append((*tasks)[:taskIndex], (*tasks)[taskIndex+1:]...)

	return nil
}

func RealizedTask(fileName string, id uint) error {
	// Leer las tareas existentes
	tasks, err := utils.ReadJson(fileName)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %v", err)
	}

	// Buscar la tarea por ID
	var taskUpdated bool
	for i, t := range tasks {
		if t.Id == id {
			// Marcar la tarea como realizada
			tasks[i].Status = "Realizada"
			taskUpdated = true
			break
		}
	}

	// Si no se encontró la tarea
	if !taskUpdated {
		return fmt.Errorf("tarea con ID %d no encontrada", id)
	}

	// Escribir las tareas actualizadas al archivo
	err = utils.WriteJson(fileName, tasks)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	return nil
}
