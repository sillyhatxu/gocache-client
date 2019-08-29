package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type UserInfo struct {
	Id                  string    `json:"id" mapstructure:"id"`
	MobileNumber        string    `json:"mobile_number" mapstructure:"mobile_number"`
	Name                string    `json:"Name" mapstructure:"Name"`
	Paid                bool      `json:"Paid" mapstructure:"Paid"`
	FirstActionDeviceId string    `json:"first_action_device_id" mapstructure:"first_action_device_id"`
	TestNumber          int       `json:"test_number" mapstructure:"test_number"`
	TestNumber64        int64     `json:"test_number_64" mapstructure:"test_number_64"`
	TestDate            time.Time `json:"test_date" mapstructure:"test_date"`
	Member              *UserInfo `json:"member" mapstructure:"member"`
}

var Client *CacheConfig

func init() {
	Client = NewCacheConfig(5*time.Second, 10*time.Second)
}

func TestCache(t *testing.T) {
	value, found := Client.Get("test1")
	assert.EqualValues(t, found, false)
	assert.Nil(t, value)
	Client.Set("test1", "test1-haha", DefaultExpiration)
	Client.Set("test2", "test2-lala", NoExpiration)

	value, found = Client.Get("test1")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test1-haha")

	value, found = Client.Get("test2")
	assert.EqualValues(t, found, true)
	assert.EqualValues(t, value, "test2-lala")

	member := &UserInfo{Id: "ID_2222", MobileNumber: "m_555555", Name: "m_test", Paid: false, FirstActionDeviceId: "m_deviceid", TestNumber: 11, TestNumber64: 22, TestDate: time.Now()}
	userinfo := &UserInfo{Id: "ID_1001", MobileNumber: "555555", Name: "test", Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now(), Member: member}
	Client.Set("test-object", userinfo, NoExpiration)
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

	Client.Set("test3", "test3-heihei", DefaultExpiration)
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

	Client.Set("test-object", userinfo, NoExpiration)

	value, found := Client.Get("test-object")
	assert.EqualValues(t, found, true)
	result := value.(*UserInfo)
	assert.EqualValues(t, result, userinfo)
}