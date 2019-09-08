package main

import (
	"flag"
	"log"
	"time"

	_ "github.com/influxdata/influxdb1-client"
	influx "github.com/influxdata/influxdb1-client/v2"
	"github.com/jmhobbs/incubator/device"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {
	var (
		relayPin  *string        = flag.String("relay-pin", "GPIO22", "Pin on the relay is connected to.")
		bme280Bus *string        = flag.String("bme280-bus", "I2C1", "I2C bus the BME280 is on.")
		onTemp    *float64       = flag.Float64("on", 37.25, "Temperature to turn the relay on.")
		offTemp   *float64       = flag.Float64("off", 38.75, "Temperature to turn the relay off.")
		interval  *time.Duration = flag.Duration("interval", 5*time.Second, "Time between readings.")
	)
	flag.Parse()

	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	p := gpioreg.ByName(*relayPin)
	if p == nil {
		log.Fatalf("error finding GPIO pin %q", *relayPin)
	}
	relay := device.NewRelay(p)

	b, err := i2creg.Open(*bme280Bus)
	if err != nil {
		log.Fatal("error opening i2c bus:", err)
	}
	defer b.Close()
	dev := device.NewBME280(b)

	influxClient, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     "http://incubator.velvetcache.org:8086",
		Username: "incubator",
		Password: "super-duper-secret",
		Timeout:  time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	states := make(chan State)

	go func() {
		for state := range states {
			bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
				Database:  "incubator",
				Precision: "s",
			})
			if err != nil {
				log.Println("influx error:", err)
				continue
			}

			point, err := influx.NewPoint("environment", map[string]string{}, state.Map(), state.Time)
			if err != nil {
				log.Println("influx error:", err)
				continue
			}
			bp.AddPoint(point)
			if err = influxClient.Write(bp); err != nil {
				log.Println("influx error:", err)
			}
		}
	}()

	for {
		temperature, _, humdity, err := dev.Read()
		if err != nil {
			log.Fatal("error reading from bme280:", err)
		}

		if temperature <= *onTemp && relay.IsOpen() {
			relay.Close()
		} else if temperature >= *offTemp && relay.IsClosed() {
			relay.Open()
		}

		states <- State{
			time.Now(),
			temperature,
			humdity,
			relay.IsClosed(),
		}

		time.Sleep(*interval)
	}
}

type State struct {
	Time        time.Time
	Temperature float64
	Humidity    float64
	Lamp        bool
}

func (s State) Map() map[string]interface{} {
	return map[string]interface{}{
		"temperature": s.Temperature,
		"humidity":    s.Humidity,
		"lamp":        s.Lamp,
	}
}
