package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdateAt    *time.Time `json:"updated_at,omitempty"`
}


func ReadJson(fileName string) ([]Task, error) {
	// Leer el contenido del archivo directamente
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %v", err)
	}

	// Deserializar los datos JSON en una variable de tipo []Task
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("error al deserializar JSON: %v", err)
	}

	// Devolver la lista de tareas
	return tasks, nil
}


func WriteJson(fileName string, tasks []Task) error {
	// Serializar los datos a JSON con formato legible
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar los datos a JSON: %v", err)
	}

	// Escribir los datos serializados al archivo
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %v", err)
	}

	return nil
}