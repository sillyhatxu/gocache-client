package cacheclient

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

func TestCacheSetGet(t *testing.T) {
	client, err := NewCacheClient()
	if err != nil {
		panic(err)
	}
	err = client.Set("test1", []byte("testhaha"))
	if err != nil {
		panic(err)
	}
	testsrc, err := client.Get("test1")
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, testsrc, "testhaha")
}

func TestCacheSetGetSrc(t *testing.T) {
	client, err := NewCacheClient()
	if err != nil {
		panic(err)
	}
	err = client.SetSrc("test1", "testhaha")
	if err != nil {
		panic(err)
	}
	testsrc, err := client.GetSrc("test1")
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, testsrc, "testhaha")
}

func TestCacheSetGetObj(t *testing.T) {
	client, err := NewCacheClient()
	if err != nil {
		panic(err)
	}
	member := UserInfo{Id: "ID_2222", MobileNumber: "m_555555", Name: "m_test", Paid: false, FirstActionDeviceId: "m_deviceid", TestNumber: 11, TestNumber64: 22}
	userinfo := UserInfo{Id: "ID_1001", MobileNumber: "555555", Name: "test", Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, Member: &member}
	err = client.SetObj("userinfo", userinfo)
	if err != nil {
		panic(err)
	}
	var user UserInfo
	err = client.GetObj("userinfo", &user)
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, user.Id, "ID_1001")
	assert.EqualValues(t, user.Name, "test")
	assert.EqualValues(t, user.MobileNumber, "555555")
	assert.EqualValues(t, user.Member.Id, "ID_2222")
	assert.EqualValues(t, user.Member.Name, "m_test")
	assert.EqualValues(t, user.Member.Member, "m_555555")
	assert.EqualValues(t, user, userinfo)
}

func TestCacheSetGetArray(t *testing.T) {
	client, err := NewCacheClient()
	if err != nil {
		panic(err)
	}
	var array []UserInfo
	array = append(array, UserInfo{Id: "ID_1001", MobileNumber: "MOBILE_1001", Name: "NAME_1001"})
	array = append(array, UserInfo{Id: "ID_1002", MobileNumber: "MOBILE_1002", Name: "NAME_1002"})
	array = append(array, UserInfo{Id: "ID_1003", MobileNumber: "MOBILE_1003", Name: "NAME_1003"})
	array = append(array, UserInfo{Id: "ID_1004", MobileNumber: "MOBILE_1004", Name: "NAME_1004"})
	array = append(array, UserInfo{Id: "ID_1005", MobileNumber: "MOBILE_1005", Name: "NAME_1005"})
	err = client.SetObj("array", array)
	if err != nil {
		panic(err)
	}

	var resultArray []UserInfo
	err = client.GetObj("array", &resultArray)
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, len(resultArray), len(array))
	assert.EqualValues(t, resultArray, array)
}

//func TestStuctCache(t *testing.T) {
//	member := &UserInfo{Id: "ID_2222", MobileNumber: "m_555555", Name: "m_test", Paid: false, FirstActionDeviceId: "m_deviceid", TestNumber: 11, TestNumber64: 22, TestDate: time.Now()}
//	userinfo := &UserInfo{Id: "ID_1001", MobileNumber: "555555", Name: "test", Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now(), Member: member}
//
//	Client.SetWithExpiration("test-object", userinfo, NoExpiration)
//
//	value, found := Client.Get("test-object")
//	assert.EqualValues(t, found, true)
//	result := value.(*UserInfo)
//	assert.EqualValues(t, result, userinfo)
//}
//
//func TestCacheClient_IncrementInt64(t *testing.T) {
//	go func() {
//		for i := 0; i < 200; i++ {
//			result, err := Client.IncrementInt64WithExpiration("test", 2*time.Second)
//			assert.Nil(t, err)
//			log.Printf(fmt.Sprintf("result : %d", result))
//			time.Sleep(200 * time.Millisecond)
//		}
//	}()
//	go func() {
//		for i := 0; i < 200; i++ {
//			result, err := Client.IncrementInt64WithExpiration("test", 2*time.Second)
//			assert.Nil(t, err)
//			log.Printf(fmt.Sprintf("result : %d", result))
//			time.Sleep(200 * time.Millisecond)
//		}
//	}()
//	for i := 0; i < 300; i++ {
//		result, err := Client.IncrementInt64WithExpiration("test", 2*time.Second)
//		assert.Nil(t, err)
//		log.Printf(fmt.Sprintf("result : %d", result))
//		time.Sleep(200 * time.Millisecond)
//	}
//}
