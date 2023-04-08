package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Todo: Add logging to program. IDK MAN
// Load Config File

type Configuration struct {
	Paths []string
	Days  int
}

type DiskStatus struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

//Note to self do logging.

func LoadConfig(jsonconf string) Configuration {
	var GoConfig Configuration
	raw, err := ioutil.ReadFile(jsonconf)
	if err != nil {
		log.Println("Error: Could not read config.")
	}
	json.Unmarshal(raw, &GoConfig)
	return GoConfig
}

// DiskUsage returns the disk usage information to DiskStatus struct.

func DiskUsage() (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		log.Fatal(err)
	}
	disk.Total = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.Total - disk.Free
	return
}

// Percent returns the current percentage of disk usage.

func Percent() (Usage float64) {
	disk := DiskUsage()
	percentUsed := float64(disk.Used) / float64(disk.Total) * 100
	return percentUsed
}

func Monitor() {
	DiskUsage()
	currentUsage := Percent()
	time.Sleep(3 * time.Second)
	if currentUsage > 80 {
		Config := LoadConfig("/workspaces/codespaces-blank/Conf/disk.json")
		var dt = time.Duration(Config.Days) * 24 * time.Hour
		for _, str := range Config.Paths {
			fmt.Printf("Checking %s for old files.\n", str)
			fileInfo, err := ioutil.ReadDir(str)
			if err != nil {
				log.Fatal(err)
			}
			now := time.Now()
			for _, file := range fileInfo {
				if diff := now.Sub(file.ModTime()); diff > dt {
					fmt.Printf("Deleting %s which is over 3 weeks old\n", file.Name())
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	} else {
		fmt.Printf("\rDisk usage is below 80%%")
	}
}

func main() {
	LoadConfig("/workspaces/codespaces-blank/Conf/disk.json")
	DiskUsage()
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGURG:
				fmt.Println("Ignoring SIGURG")
			default:
				os.Exit(1)
			}
		}
	}()
	// Mintoring Disk Space
	for {
		Monitor()
	}
}
