package main

import (
	"flag"
	"fmt"
	"github.com/Aoana/go-star-sim/internal/pkg/stars"
	"github.com/shirou/gopsutil/cpu"
	"log"
	"time"
)

var velocityFunc func()

func main() {

	// Parse arguments
	nstars := flag.Int("stars", 12, "Number of stars in cluster")
	nsteps := flag.Int("steps", 100, "Number of time steps")
	velocity := flag.Float64("velocity", 0.5, "Initial velocity of stars [0 - 5.0]")
	model := flag.String("model", "Exact", "\"Exact\", \"ExactGR\", \"BarnesHut\" or \"BarnesHutGR\"")
	width := flag.Int("width", 1920, "Grid width size")
	height := flag.Int("height", 1080, "Grid height size")
	file := flag.String("file", "/tmp/output", "Path to output file")
	flag.Parse()

	// Store simulation header data
	stars.Data.Width = *width
	stars.Data.Height = *height
	stars.Data.Steps = *nsteps
	stars.Data.Model = *model
	stars.Data.Velocity = *velocity

	switch *model {
	case "Exact":
		velocityFunc = stars.VelocityExact
	case "ExactGR":
		velocityFunc = stars.VelocityExactGR
	case "BarnesHut":
		velocityFunc = stars.VelocityBarnesHut
	case "BarnesHutGR":
		velocityFunc = stars.VelocityBarnesHutGR
	default:
		log.Fatal("Unknown gravity model")
	}

	// Spawn all stars
	stars.StartValues(*nstars)

	// Create coordinate slice for storing positions
	pixels := make([]stars.Pixel, len(stars.StarList))
	for i := range pixels {
		pixels[i].Px = make([]uint16, stars.Data.Steps)
		pixels[i].Py = make([]uint16, stars.Data.Steps)
	}

	// Run simulation
	log.Println("Simulation starting")
	log.Printf("Stars=%d, Model=%s, Grid=%dx%d, Timesteps=%d", len(stars.StarList), stars.Data.Model, stars.Data.Width, stars.Data.Height, stars.Data.Steps)

	// Measure the time
	start := time.Now()

	for steps := 0; steps < stars.Data.Steps; steps++ {
		// Move all stars
		stars.Move()
		// Physical velocity calculation (based on method)
		velocityFunc()
		// Store position for post processing
		for s := range stars.StarList {
			pixels[s].Px[steps] = uint16(stars.StarList[s].X)
			pixels[s].Py[steps] = uint16(stars.StarList[s].Y)
		}
	}

	stars.Data.Time = time.Since(start)
	stars.Data.Stars = pixels

	var cpuModel string
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Println("Unable to get CPU info", err)
		cpuModel = "CPU not recognized"
	} else {
		cpuModel = cpuStat[0].ModelName
	}

	stars.Data.Summary = fmt.Sprintf("Stars: %d\nGrid: %dx%d\nModel: %s\nSteps: %d\nVelocity: %0.2f\n\nBuild Info:\n%s\nCalculation Time %0.2f minutes",
		len(stars.Data.Stars), stars.Data.Width, stars.Data.Height, stars.Data.Model, stars.Data.Steps, stars.Data.Velocity, cpuModel, stars.Data.Time.Minutes())

	log.Printf("Simulation complete, took %0.2f minutes, storing to file %s", stars.Data.Time.Minutes(), *file)
	err = stars.Write(*file)
	if err != nil {
		log.Fatal("Unable to create file", err)
	}
}
