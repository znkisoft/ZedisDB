// all examples are copied from go-redis/v8/example_test.go
package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         ":7000",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func TestServerConnection(t *testing.T) {
	ch := make(chan string, 10)

	// open connections
	for i := 0; i < 10; i++ {
		go func() {
			msg, err := rdb.Ping(ctx).Result()
			if err != nil {
				t.Errorf("can'not ping server: %s", err)
			}
			ch <- msg
		}()
	}

	// loop over channel and check if the value equal "PONG"
	for i := 0; i < 10; i++ {
		if msg := <-ch; msg != "PONG" {
			t.Errorf("expected message PONG, got %s", msg)
		}
	}
}

func TestGetAndSet(t *testing.T) {
	// test SET
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Error(err)
	}

	// test GET
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	if val != "value" {
		t.Errorf("expect value is 'value', but got '%s'", val)
	}

	// test GET with missing key
	val2, err := rdb.Get(ctx, "missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		t.Error(err)
	} else {
		t.Errorf("expect nil, but got '%s'", val2)
	}
}

func TestSetWithExpirationKey(t *testing.T) {
	// Last argument is expiration. Zero means the key has no expiration time.
	err := rdb.Set(ctx, "key3", "value3", 0).Err()
	if err != nil {
		t.Error(err)
	}

	// key4 will expire in an hour.
	err = rdb.Set(ctx, "key_expire", "value_expire_in_one_hour", time.Hour).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestIncr(t *testing.T) {
	// test incr
	result, err := rdb.Incr(ctx, "counter").Result()
	if err != nil {
		t.Error(err)
	}

	if result != 1 {
		t.Errorf("expect result is 1, but got %d", result)
	}
}

// func TestConn(t *testing.T) {
// 	conn := rdb.Conn(context.Background())
//
// 	err := conn.ClientSetName(ctx, "foobar").Err()
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	// Open other connections.
// 	for i := 0; i < 10; i++ {
// 		go rdb.Ping(ctx)
// 	}
//
// 	s, err := conn.ClientGetName(ctx).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(s)
// 	// Output: foobar
// }

// func ExampleClient_BLPop() {
// 	if err := rdb.RPush(ctx, "queue", "message").Err(); err != nil {
// 		panic(err)
// 	}

// 	// use `rdb.BLPop(0, "queue")` for infinite waiting time
// 	result, err := rdb.BLPop(ctx, 1*time.Second, "queue").Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result[0], result[1])
// 	// Output: queue message
// }

// func ExampleClient_Scan() {
// 	rdb.FlushDB(ctx)
// 	for i := 0; i < 33; i++ {
// 		err := rdb.Set(ctx, fmt.Sprintf("key%d", i), "value", 0).Err()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	var cursor uint64
// 	var n int
// 	for {
// 		var keys []string
// 		var err error
// 		keys, cursor, err = rdb.Scan(ctx, cursor, "key*", 10).Result()
// 		if err != nil {
// 			panic(err)
// 		}
// 		n += len(keys)
// 		if cursor == 0 {
// 			break
// 		}
// 	}

// 	fmt.Printf("found %d keys\n", n)
// 	// Output: found 33 keys
// }

// func ExampleClient_ScanType() {
// 	rdb.FlushDB(ctx)
// 	for i := 0; i < 33; i++ {
// 		err := rdb.Set(ctx, fmt.Sprintf("key%d", i), "value", 0).Err()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	var cursor uint64
// 	var n int
// 	for {
// 		var keys []string
// 		var err error
// 		keys, cursor, err = rdb.ScanType(ctx, cursor, "key*", 10, "string").Result()
// 		if err != nil {
// 			panic(err)
// 		}
// 		n += len(keys)
// 		if cursor == 0 {
// 			break
// 		}
// 	}

// 	fmt.Printf("found %d keys\n", n)
// 	// Output: found 33 keys
// }

// // ExampleStringStringMapCmd_Scan shows how to scan the results of a map fetch
// // into a struct.
// func ExampleStringStringMapCmd_Scan() {
// 	rdb.FlushDB(ctx)
// 	err := rdb.HMSet(ctx, "map",
// 		"name", "hello",
// 		"count", 123,
// 		"correct", true).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get the map. The same approach works for HmGet().
// 	res := rdb.HGetAll(ctx, "map")
// 	if res.Err() != nil {
// 		panic(err)
// 	}

// 	type data struct {
// 		Name    string `redis:"name"`
// 		Count   int    `redis:"count"`
// 		Correct bool   `redis:"correct"`
// 	}

// 	// Scan the results into the struct.
// 	var d data
// 	if err := res.Scan(&d); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(d)
// 	// Output: {hello 123 true}
// }

// // ExampleSliceCmd_Scan shows how to scan the results of a multi key fetch
// // into a struct.
// func ExampleSliceCmd_Scan() {
// 	rdb.FlushDB(ctx)
// 	err := rdb.MSet(ctx,
// 		"name", "hello",
// 		"count", 123,
// 		"correct", true).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	res := rdb.MGet(ctx, "name", "count", "empty", "correct")
// 	if res.Err() != nil {
// 		panic(err)
// 	}

// 	type data struct {
// 		Name    string `redis:"name"`
// 		Count   int    `redis:"count"`
// 		Correct bool   `redis:"correct"`
// 	}

// 	// Scan the results into the struct.
// 	var d data
// 	if err := res.Scan(&d); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(d)
// 	// Output: {hello 123 true}
// }

// func ExampleClient_Pipelined() {
// 	var incr *redis.IntCmd
// 	_, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
// 		incr = pipe.Incr(ctx, "pipelined_counter")
// 		pipe.Expire(ctx, "pipelined_counter", time.Hour)
// 		return nil
// 	})
// 	fmt.Println(incr.Val(), err)
// 	// Output: 1 <nil>
// }

// func ExampleClient_Pipeline() {
// 	pipe := rdb.Pipeline()

// 	incr := pipe.Incr(ctx, "pipeline_counter")
// 	pipe.Expire(ctx, "pipeline_counter", time.Hour)

// 	// Execute
// 	//
// 	//     INCR pipeline_counter
// 	//     EXPIRE pipeline_counts 3600
// 	//
// 	// using one rdb-server roundtrip.
// 	_, err := pipe.Exec(ctx)
// 	fmt.Println(incr.Val(), err)
// 	// Output: 1 <nil>
// }

// func ExampleClient_TxPipelined() {
// 	var incr *redis.IntCmd
// 	_, err := rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
// 		incr = pipe.Incr(ctx, "tx_pipelined_counter")
// 		pipe.Expire(ctx, "tx_pipelined_counter", time.Hour)
// 		return nil
// 	})
// 	fmt.Println(incr.Val(), err)
// 	// Output: 1 <nil>
// }

// func ExampleClient_TxPipeline() {
// 	pipe := rdb.TxPipeline()

// 	incr := pipe.Incr(ctx, "tx_pipeline_counter")
// 	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)

// 	// Execute
// 	//
// 	//     MULTI
// 	//     INCR pipeline_counter
// 	//     EXPIRE pipeline_counts 3600
// 	//     EXEC
// 	//
// 	// using one rdb-server roundtrip.
// 	_, err := pipe.Exec(ctx)
// 	fmt.Println(incr.Val(), err)
// 	// Output: 1 <nil>
// }

// func ExampleClient_Watch() {
// 	const maxRetries = 1000
//
// 	// Increment transactionally increments key using GET and SET commands.
// 	increment := func(key string) error {
// 		// Transactional function.
// 		txf := func(tx *redis.Tx) error {
// 			// Get current value or zero.
// 			n, err := tx.Get(ctx, key).Int()
// 			if err != nil && err != redis.Nil {
// 				return err
// 			}
//
// 			// Actual opperation (local in optimistic lock).
// 			n++
//
// 			// Operation is committed only if the watched keys remain unchanged.
// 			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
// 				pipe.Set(ctx, key, n, 0)
// 				return nil
// 			})
// 			return err
// 		}
//
// 		for i := 0; i < maxRetries; i++ {
// 			err := rdb.Watch(ctx, txf, key)
// 			if err == nil {
// 				// Success.
// 				return nil
// 			}
// 			if err == redis.TxFailedErr {
// 				// Optimistic lock lost. Retry.
// 				continue
// 			}
// 			// Return any other error.
// 			return err
// 		}
//
// 		return errors.New("increment reached maximum number of retries")
// 	}
//
// 	var wg sync.WaitGroup
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
//
// 			if err := increment("counter3"); err != nil {
// 				fmt.Println("increment error:", err)
// 			}
// 		}()
// 	}
// 	wg.Wait()
//
// 	n, err := rdb.Get(ctx, "counter3").Int()
// 	fmt.Println("ended with", n, err)
// 	// Output: ended with 100 <nil>
// }
//
// func ExamplePubSub() {
// 	pubsub := rdb.Subscribe(ctx, "mychannel1")
//
// 	// Wait for confirmation that subscription is created before publishing anything.
// 	_, err := pubsub.Receive(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// Go channel which receives messages.
// 	ch := pubsub.Channel()
//
// 	// Publish a message.
// 	err = rdb.Publish(ctx, "mychannel1", "hello").Err()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	time.AfterFunc(time.Second, func() {
// 		// When pubsub is closed channel is closed too.
// 		_ = pubsub.Close()
// 	})
//
// 	// Consume messages.
// 	for msg := range ch {
// 		fmt.Println(msg.Channel, msg.Payload)
// 	}
//
// 	// Output: mychannel1 hello
// }
//
// func ExamplePubSub_Receive() {
// 	pubsub := rdb.Subscribe(ctx, "mychannel2")
// 	defer pubsub.Close()
//
// 	for i := 0; i < 2; i++ {
// 		// ReceiveTimeout is a low level API. Use ReceiveMessage instead.
// 		msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
// 		if err != nil {
// 			break
// 		}
//
// 		switch msg := msgi.(type) {
// 		case *redis.Subscription:
// 			fmt.Println("subscribed to", msg.Channel)
//
// 			_, err := rdb.Publish(ctx, "mychannel2", "hello").Result()
// 			if err != nil {
// 				panic(err)
// 			}
// 		case *redis.Message:
// 			fmt.Println("received", msg.Payload, "from", msg.Channel)
// 		default:
// 			panic("unreached")
// 		}
// 	}
//
// 	// sent message to 1 rdb
// 	// received hello from mychannel2
// }

// func ExampleClient_SlowLogGet() {
// 	const key = "slowlog-log-slower-than"
//
// 	old := rdb.ConfigGet(ctx, key).Val()
// 	rdb.ConfigSet(ctx, key, "0")
// 	defer rdb.ConfigSet(ctx, key, old[1].(string))
//
// 	if err := rdb.Do(ctx, "slowlog", "reset").Err(); err != nil {
// 		panic(err)
// 	}
//
// 	rdb.Set(ctx, "test", "true", 0)
//
// 	result, err := rdb.SlowLogGet(ctx, -1).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(len(result))
// 	// Output: 2
// }
