package main

import (
	"RastreadorDeTareas/cmd"
	"log"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",                                      // Nombre de la aplicación
	Short: "Gestor de tareas",                                 // Descripción corta
	Long: "Una aplicación CLI simple para gestionar tareas. Puedes crear, listar, actualizar y eliminar tareas con simples comandos.",
}

func init() {
	// Definir las banderas globales con alias
	rootCmd.PersistentFlags().UintP("id", "i", 0, "Id de la tarea") // "i" es el alias corto para "id"
	rootCmd.PersistentFlags().StringP("title", "t", "", "Título de la tarea (obligatorio)") // "t" es el alias corto para "title"
	rootCmd.PersistentFlags().StringP("description", "d", "", "Descripción de la tarea (obligatorio)") // "d" es el alias corto para "description"
}


func main() {
	// Registra los comandos específicos en el comando raíz
	rootCmd.AddCommand(cmd.CreateTask())
	rootCmd.AddCommand(cmd.ListTasks())
	rootCmd.AddCommand(cmd.UpdateTask())
	rootCmd.AddCommand(cmd.DeleteTask())
	rootCmd.AddCommand(cmd.RealizedTask())

	// Ejecuta el comando raíz
	if err := rootCmd.Execute(); err != nil {
		log.Println("Error al ejecutar el comando:", err)
		os.Exit(1) // Si algo falla, salimos con un código de error
	}
}
