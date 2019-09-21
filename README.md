# gocache-client


This code dependent on referenced modules - [go-cache](https://github.com/patrickmn/go-cache)

## sample

[example](https://github.com/sillyhatxu/gocache-client/blob/master/client_test.go)

## new cache client

```
var Client = NewCacheConfig(Expiration(5*time.Second), CleanupInterval(10*time.Second))
```


## Demo

```
func TestCache(t *testing.T) {
	value, found := Client.Get("test1")
	assert.EqualValues(t, found, false)
	assert.Nil(t, value)
	Client.SetWithExpiration("test1", "test1-haha", DefaultExpiration)
	Client.SetWithExpiration("test2", "test2-lala", NoExpiration)

	value, found = Client.Get("test1")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test1-haha")

	value, found = Client.Get("test2")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test2-lala")

	member := &UserInfo{Id: "ID_2222", MobileNumber: "m_555555", Name: "m_test", Paid: false, FirstActionDeviceId: "m_deviceid", TestNumber: 11, TestNumber64: 22, TestDate: time.Now()}
	userinfo := &UserInfo{Id: "ID_1001", MobileNumber: "555555", Name: "test", Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now(), Member: member}
	Client.SetWithExpiration("test-object", userinfo, NoExpiration)
	time.Sleep(6 * time.Second)

	value, found = Client.Get("test1")
	assert.EqualValues(t, found, false)
	assert.Nil(t, value)

	value, found = Client.Get("test2")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test2-lala")

	value, found = Client.Get("test-object")
	assert.EqualValues(t, found, true)
	result := value.(*UserInfo)
	assert.EqualValues(t, result, userinfo)

	Client.SetWithExpiration("test3", "test3-heihei", DefaultExpiration)
	value, found = Client.Get("test3")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test3-heihei")
	time.Sleep(1 * time.Second)
	value, found = Client.Get("test3")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test3-heihei")
	time.Sleep(1 * time.Second)
	value, found = Client.Get("test3")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test3-heihei")
	time.Sleep(1 * time.Second)
	value, found = Client.Get("test3")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test3-heihei")
	time.Sleep(1 * time.Second)
	value, found = Client.Get("test3")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test3-heihei")
	time.Sleep(1 * time.Second)
	value, found = Client.Get("test3")
	assert.EqualValues(t, found, false)
	assert.Nil(t, value)
}

func TestStuctCache(t *testing.T) {
	member := &UserInfo{Id: "ID_2222", MobileNumber: "m_555555", Name: "m_test", Paid: false, FirstActionDeviceId: "m_deviceid", TestNumber: 11, TestNumber64: 22, TestDate: time.Now()}
	userinfo := &UserInfo{Id: "ID_1001", MobileNumber: "555555", Name: "test", Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now(), Member: member}

	Client.SetWithExpiration("test-object", userinfo, NoExpiration)

	value, found := Client.Get("test-object")
	assert.EqualValues(t, found, true)
	result := value.(*UserInfo)
	assert.EqualValues(t, result, userinfo)
}

func TestCacheClient_IncrementInt64(t *testing.T) {
	go func() {
		for i := 0; i < 200; i++ {
			result, err := Client.IncrementInt64WithExpiration("test", 2*time.Second)
			assert.Nil(t, err)
			log.Printf(fmt.Sprintf("result : %d", result))
			time.Sleep(200 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 200; i++ {
			result, err := Client.IncrementInt64WithExpiration("test", 2*time.Second)
			assert.Nil(t, err)
			log.Printf(fmt.Sprintf("result : %d", result))
			time.Sleep(200 * time.Millisecond)
		}
	}()
	for i := 0; i < 300; i++ {
		result, err := Client.IncrementInt64WithExpiration("test", 2*time.Second)
		assert.Nil(t, err)
		log.Printf(fmt.Sprintf("result : %d", result))
		time.Sleep(200 * time.Millisecond)
	}
}
```