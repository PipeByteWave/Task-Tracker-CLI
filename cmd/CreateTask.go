package cmd

import (
	"RastreadorDeTareas/crud"
	"RastreadorDeTareas/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var fileName = "task.json"

func CreateTask() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Crea una nueva tarea",
		Run: func(cmd *cobra.Command, args []string) {
			// Obtener valores de las banderas
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")

			// Validar que las banderas no estén vacías
			if title == "" || description == "" {
				fmt.Println("El título y la descripción son obligatorios.")
				return
			}

			// Crear la nueva tarea
			task := utils.Task{
				Title:       title,
				Description: description,
				Status:      "Pendiente",
			}

			// Intentar agregar la tarea
			if err := crud.CreateTask(fileName, task); err != nil {
				fmt.Println("Error al crear tarea:", err)
			} else {
				fmt.Println("Tarea creada exitosamente.")
			}
		},
	}
}