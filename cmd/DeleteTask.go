package cmd

import (
	"RastreadorDeTareas/crud"
	"fmt"
	"github.com/spf13/cobra"
)


func DeleteTask() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Elimina una tarea por su ID",
		Run: func(cmd *cobra.Command, args []string) {
			force, _ := cmd.Flags().GetBool("force")
			// Obtener valor de la bandera
			id, _ := cmd.Flags().GetUint("id") // Usando GetUint para obtener un uint

			if id == 0 {
				fmt.Println("El ID de la tarea es obligatorio.")
				return
			}

			if !force {
				var confirm string
				fmt.Print("¿Estás seguro de que deseas eliminar la tarea? (s/n): ")
				fmt.Scanln(&confirm)
				if confirm != "s" && confirm != "S" {
					fmt.Println("Eliminación cancelada.")
					return
				}
			}

			// Llamar a la función del CRUD para eliminar la tarea
			err := crud.DeleteTask(fileName, uint(id))
			if err != nil {
				// Si hubo un error, lo mostramos
				fmt.Println("Error al eliminar la tarea:", err)
				return
			}

			// Si no hubo errores, la tarea fue eliminada exitosamente
			fmt.Println("Tarea eliminada exitosamente.")
		},
	}
	cmd.Flags().Bool("force", false, "Eliminar la tarea sin confirmación")

	return cmd
}
