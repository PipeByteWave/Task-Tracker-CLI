package cmd

import (
	"RastreadorDeTareas/crud"
	"fmt"
	"github.com/spf13/cobra"
)

func UpdateTask() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Actualiza una tarea existente",
		Run: func(cmd *cobra.Command, args []string) {
			// Obtener valores de las banderas
			id, _ := cmd.Flags().GetUint("id")  // Usando GetUint para obtener un uint
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")

			// Validar que las banderas estén presentes
			if id == 0  {
				fmt.Println("El ID es obligatorio")
				return
			}
			if title == "" && description == "" {
				fmt.Println("Es nesesario que cambies por lo menos un item para efectuar la actualización")
			}

			// Llamar a la función CRUD para actualizar la tarea
			err := crud.UpdateTask(fileName, id, title, description)
			if err != nil {
				fmt.Println("Error al actualizar la tarea:", err)
			} else {
				fmt.Println("Tarea actualizada exitosamente.")
			}
		},
	}
}
