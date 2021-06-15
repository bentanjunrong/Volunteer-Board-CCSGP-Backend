package main

func main() {
	r := InitRouter()
	r.Run("localhost:8080") // TODO: put this in an env file
}
