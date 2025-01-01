package cmd

import (
	"RastreadorDeTareas/crud"
	"fmt"

	"github.com/spf13/cobra"
)

func ListTasks() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lista todas las tareas",
		Run: func(cmd *cobra.Command, args []string) {
			// Obtener las tareas del CRUD
			tasks, err := crud.GetAllTasks(fileName)
			if err != nil {
				fmt.Println("Error al obtener las tareas:", err)
				return
			}

			// Verificar si no hay tareas
			if tasks == nil {
				fmt.Println("No hay tareas disponibles.")
				return
			}

			// Mostrar las tareas
			for _, task := range tasks {
				fmt.Printf("ID: %d | Título: %s | Descripción: %s | Estado: %s\n", task.Id, task.Title, task.Description, task.Status)
			}
		},
	}
}
