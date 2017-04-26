package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

func updateClocks(writer *uilive.Writer, clocks []*Clock) {
	// Why two for loops?
	// uilive reflects the screen pretty fast. Once you combine `Update()` with
	// writing the results, it will display a broken result sometimes.
	for _, clock := range clocks {
		clock.Update()
	}

	writer.Start()
	defer writer.Stop()

	for _, clock := range clocks {
		fmt.Fprintln(writer, clock.DateString)
		fmt.Fprintln(writer, clock.TimeFiglet)
	}
}

func main() {
	timezones := []string{"America/Los_Angeles", "Etc/UTC", "Asia/Shanghai", "Asia/Seoul"}

	var clocks []*Clock
	for _, timezone := range timezones {
		clock, err := NewClock(timezone)
		if err != nil {
			panic(err)
		}
		clocks = append(clocks, clock)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	writer := uilive.New()
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		updateClocks(writer, clocks)
		for {
			select {
			case <-ticker.C:
				updateClocks(writer, clocks)
			}
		}
	}()

	go func() {
		defer wg.Done()

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}()

	wg.Wait()
}
