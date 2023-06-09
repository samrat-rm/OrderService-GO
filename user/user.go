package user

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB
// var err error

// const DNS = "root:@tcp(localhost:3306)/quickmart?charset=utf8mb4&parseTime=True&loc=Local"

// type User struct {
// 	gorm.Model
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// 	Email     string `json:"email"`
// }

// func InitialMigration() {
// 	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Cannot connect to DB")
// 	}
// 	DB.AutoMigrate(&User{}) // creates table if no there
// }

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var users []User
// 	DB.Find(&users)
// 	json.NewEncoder(w).Encode(users)
// }

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var user User
// 	DB.First(&user, params["id"])
// 	json.NewEncoder(w).Encode(user)
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	DB.Create(&user)
// 	json.NewEncoder(w).Encode(user)
// }

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var user User
// 	DB.First(&user, params["id"])
// 	json.NewDecoder(r.Body).Decode(&user)
// 	DB.Save(&user)
// 	json.NewEncoder(w).Encode(user)
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var user User
// 	DB.Delete(&user, params["id"])
// 	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
// }
