package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func generatePassword() string {
	NumAndChars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	ret := []byte{}
	for i := 0; i < 16; i++ {
		ra := rand.New(rand.NewSource(time.Now().UnixNano()))
		ret = append(ret, NumAndChars[ra.Intn(len(NumAndChars))])
	}
	StdChars := []byte("~!@#$%^&*()_+`-={}|[]:<>?,./")
	for i := 0; i < 3; i++ {
		ra := rand.New(rand.NewSource(time.Now().UnixNano()))
		ret[ra.Intn(10)] = StdChars[ra.Intn(len(StdChars))]
	}
	return string(ret)
}

// GetUserPasswordSalt get user password salt, when using by register user or change password
func getUserPasswordSalt(length int) (string, error) {
	buffer := make([]byte, length)
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	if n, err := ra.Read(buffer); err != nil || n != length {
		return "", err
	}
	salt := fmt.Sprintf("%x", buffer)
	return salt, nil
}

func TestGenerateNewUser(t *testing.T) {
	tmp := `INSERT INTO "booking_account" ("created_at","updated_at","deleted_at","username","password","salt") VALUES ('%s','%s',NULL,'%s','%s','%s')`
	salt, err := getUserPasswordSalt(12)
	if err != nil {
		t.Fatal(err)
	}
	user := "wayfarer"
	// password := generatePassword()
	password := "wayfarer2020"
	savePassword, err := HashUserPassword(salt, password)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	t.Log(user, password)
	t.Logf(tmp, now, now, user, savePassword, salt)
}

func TestGetRedisKey(t *testing.T) {
	Convey("TestGetRedisKey", t, func() {
		Convey("input 1(int), should return booking-service_1", func() {
			So(GetRedisKey(1), ShouldEqual, "booking-service_1")
		})

		Convey("input 1(string), should return booking-service_1", func() {
			So(GetRedisKey("1"), ShouldEqual, "booking-service_1")
		})

		Convey("input 1.1(float), should return booking-service_1.1", func() {
			So(GetRedisKey(1.1), ShouldEqual, "booking-service_1.1")
		})

		Convey("input 1(int), 2(string), should return booking-service_1_2", func() {
			So(GetRedisKey(1, "2"), ShouldEqual, "booking-service_1_2")
		})

		Convey("input 1(int), 2(string), 3(float), should return booking-service_1_2_3.2", func() {
			So(GetRedisKey(1, "2", 3.2), ShouldEqual, "booking-service_1_2_3.2")
		})
	})
}

func TestHashUserPassword(t *testing.T) {
	Convey("TestHashUserPassword", t, func() {
		Convey("input empty string, should return e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", func() {
			pass, err := HashUserPassword("", "")
			So(pass, ShouldEqual, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
			So(err, ShouldBeNil)
		})

		Convey("input salt is empty, password is '123456', should return 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", func() {
			pass, err := HashUserPassword("", "123456")
			So(pass, ShouldEqual, "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92")
			So(err, ShouldBeNil)
		})

		Convey("input salt is 'abcdef', password is empty, should return bef57ec7f53a6d40beb640a780a639c83bc29ac8a9816f1fc6c5c6dcd93c4721", func() {
			pass, err := HashUserPassword("abcdef", "")
			So(pass, ShouldEqual, "bef57ec7f53a6d40beb640a780a639c83bc29ac8a9816f1fc6c5c6dcd93c4721")
			So(err, ShouldBeNil)
		})

		Convey("input salt is 'abcdef', password is '123456', should return 8fa0acf6233b92d2d48a30a315cd213748d48f28eaa63d7590509392316b3016", func() {
			pass, err := HashUserPassword("abcdef", "123456")
			So(pass, ShouldEqual, "8fa0acf6233b92d2d48a30a315cd213748d48f28eaa63d7590509392316b3016")
			So(err, ShouldBeNil)
		})
	})
}

type Message struct {
	Name string `json:"name,omitempty"`
	Sex  string `json:"sex,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func (m *Message) GetName() string {
	return m.Name
}
func (m *Message) GetSex() string {
	return m.Sex
}

type NewMessage interface {
	GetName() (Name string)
	GetSex() (Sex string)
}

func TestZhuan(t *testing.T) {
	a := `{"name": "laowang", "sex": "man", "age": 10}`
	var test Message
	_ = json.Unmarshal([]byte(a), &test)
	var bb NewMessage
	bb = &test
	fmt.Println(bb.GetName())

	aa := bb.(*Message)
	fmt.Printf("aa: %+v", aa)
}
