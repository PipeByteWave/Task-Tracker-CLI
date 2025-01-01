package cmd

import (
	"RastreadorDeTareas/crud"
	"fmt"

	"github.com/spf13/cobra"
)

func RealizedTask() *cobra.Command {
	return &cobra.Command{
		Use: "realized",
		Short: "Marca una tarea como realizada",
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := cmd.Flags().GetUint("id")

			if id == 0 {
				fmt.Println("El ID es obligatorio")
				return
			}

			err := crud.RealizedTask(fileName, id)
			if err != nil {
				return
			}

			fmt.Println("Genial!!, tarea terminada")

		},
	}
}