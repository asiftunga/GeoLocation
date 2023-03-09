package main

import (
	router "MapProject/api/routes"
	_ "MapProject/docs"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/shirou/gopsutil/process"
	"net"
	"os"
	"runtime"
	"sync"
	"time"
)

// @title GeoLocation Project
// @version 1.0
// @description This api written by Asif Tunga Mubarek.
// @termsOfService http://swagger.io/terms/
// @contact.name Asif Tunga Mubarek
// @contact.email asiftunga@hotmail.com
// @host localhost:3030
// @BasePath /
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,DELETE,PUT",
		ExposeHeaders: "Authorization",
	}))
	fmt.Println("")
	fmt.Println("Developed By : ")
	fmt.Println("")
	fmt.Println("╔══╦═╦╦═╗╔══╗─────────╔═╦═╗─╔╗───────╔╗")
	fmt.Println("║╔╗║═╬╣═╣╚╗╔╬╦═╦╦═╦═╗─║║║║╠╦╣╚╦═╗╔╦╦═╣╠╗")
	fmt.Println("║╠╣╠═║║╔╝─║║║║║║║╬║╬╚╗║║║║║║║╬║╬╚╣╔╣╩╣═╣")
	fmt.Println("╚╝╚╩═╩╩╝──╚╩═╩╩═╬╗╠══╝╚╩═╩╩═╩═╩══╩╝╚═╩╩╝")
	fmt.Println("────────────────╚═╝")
	fmt.Println("for more info contact -> asiftunga@hotmail.com")
	fmt.Println("")
	fmt.Println("Lutfen Api Dokumentasyonu icin http://127.0.0.1:3030/swagger/index.html adresine gidiniz")
	router.SetupRoutes(app)

	var wg sync.WaitGroup
	wg.Add(2)

	// Start the server in a separate Goroutine
	go func() {
		defer wg.Done()
		err := app.Listen(":3030")
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			conn, err := net.DialTimeout("tcp", "localhost:3030", 1*time.Second)
			if err == nil {
				conn.Close()
				break
			}
			time.Sleep(100 * time.Millisecond)
		}

		process, err := process.NewProcess(int32(os.Getpid()))
		if err != nil {
			panic(err)
		}
		var memStats runtime.MemStats
		var prevUsage = uint64(0)

		for {
			cpuPercent, err := process.CPUPercent()
			if err != nil {
				panic(err)
			}
			cpUsage := cpuPercent * float64(time.Second) / 1024 / 1024 / 1024
			runtime.ReadMemStats(&memStats)
			meUsage := memStats.Alloc - prevUsage
			prevUsage = memStats.Alloc

			fmt.Printf("\rCode's CPU Usage: %.7f KB | Code's RAM Usage: %.7f KB ", cpUsage, float64(meUsage)/1024/1024/1024)
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
