package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)
var s = gocron.NewScheduler(time.UTC)
var count1 = 0
var count2 = 0

var wg sync.WaitGroup
func task1() {
    
    count1 += 1
    fmt.Println("I am at task1", count1)
    if count1 == 10 {
        wg.Done()
        s.RemoveByTag("tag1")
    }
    
}
func task2() {
    
    count2 += 1
    fmt.Println("I am at task2", count2)
    if count2 == 20 {
        wg.Done()
        s.RemoveByTag("tag2")
    }
}
func task3() {
    fmt.Println("I am at task3")
    wg.Done()
    s.RemoveByTag("tag3")
}

func main() {
    s.Every(2).Seconds().Tag("tag1").Do(task1)
    s.Every(1).Seconds().Tag("tag2").Do(task2)
    s.Every(1).Seconds().Tag("tag3").Do(task3)
    
    wg.Add(3)
    go s.StartBlocking()
    wg.Wait()
    fmt.Println("task completed!!!!!")
}