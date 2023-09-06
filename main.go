package main

import "github.com/fatih/color" // Указываем путь до нужного пакета внутри репозитория

func main() {
	red := color.New(color.FgRed).PrintFunc()
	red("Warning")
}
