package broker

import (
	"fmt"
	"sync"
	"testing"
)
func TestBroker(t *testing.T){
	b := newBroker()

	var wg sync.WaitGroup

	for i:=0;i<10;i++ {
		topic := fmt.Sprintf("test%d",i)
		payload := []byte(topic)
		ch, err := b.Subscribe(topic)
		if err!=nil{
			t.Fatal(err)
		}

		wg.Add(1)

		go func(){
			e := <-ch
			if string(e) != string(payload){
				t.Fatalf("%s expected %s got %s", topic, string(payload), string(e))
			}
			err := b.Unsubscribe(topic,ch)
			if err!=nil{
				t.Fatal(err)
			}

			wg.Done()
		}()

		er1 := b.Publish(topic,payload)
		if er1 !=nil{
			t.Fatal(er1)
		}
	}

	wg.Wait()
}