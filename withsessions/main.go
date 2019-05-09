package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"mux"
	_ "mysql"
	"net/http"
	"sessions"
)

var first string = "<html><body>"
var last string = "</body></html>"
var encryptionkey = "secret very "
var usersession = sessions.NewCookieStore([]byte(encryptionkey))

func init() {
	usersession.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 3,
		HttpOnly: true,
	}
}
func dbConn() (db *sql.DB) {
	dbName := "meghana"
	dbUser := "root"
	dbDriver := "mysql"
	dbPass := "meghana"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func loginpagehandler(w http.ResponseWriter, r *http.Request) {

	bytes, _ := ioutil.ReadFile("login.html")

	fmt.Fprint(w, string(bytes))

}

func loginhandler(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	Uname := r.FormValue("name")
	Upassword := r.FormValue("password")

	var temppass string
	if len(Uname) != 0 && len(Upassword) != 0 {
		session, _ := usersession.New(r, "authuser")
		//http.Redirect(w,r,uredirect,345)
		selectque, err := db.Query("SELECT passwor from emp1 where nam=?", Uname)
		if err != nil {
			panic(err.Error())
		}

		for selectque.Next() {

			err = selectque.Scan(&temppass)
			if err != nil {
				panic(err.Error())
			}

		}
		if temppass == Upassword {
			session.Values["username"] = Uname

			err := session.Save(r, w)
			if err != nil {
				fmt.Println("session writing error")
			}

			http.Redirect(w, r, "/index", http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, "/register", http.StatusMovedPermanently)

		}

	} else {
		http.Redirect(w, r, "/register", http.StatusMovedPermanently)

	}
	defer db.Close()

}

func indexhandler(w http.ResponseWriter, r *http.Request) {

	bytes, _ := ioutil.ReadFile("index.html")

	fmt.Fprint(w, string(bytes))

}
func registerpagehandler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadFile("register.html")

	fmt.Fprint(w, string(bytes))

}
func registerhandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	Uname := r.FormValue("name")
	Uemail := r.FormValue("email")
	Upassword := r.FormValue("password")
	Uconfirmpassword := r.FormValue("confirmpassword")

	if len(Uname) != 0 && len(Upassword) != 0 && len(Uemail) != 0 && len(Uconfirmpassword) != 0 {
		insform, err := db.Prepare("INSERT INTO emp1(nam,email, passwor) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insform.Exec(Uname, Uemail, Upassword)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		fmt.Fprintln(w, "Sucesfully registered ..please login")

	} else {
		fmt.Fprintln(w, "These fields should not be empty please register again")
		http.Redirect(w, r, "/register", http.StatusMovedPermanently)

	}
	defer db.Close()
}
func updatepagehandler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadFile("update.html")

	fmt.Fprint(w, string(bytes))
}

func updatehandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	Upassword := r.FormValue("password")
	Uemail := r.FormValue("email")
	//var uredirect string="/"

	if len(Upassword) != 0 && len(Uemail) != 0 {
		//http.Redirect(w,r,uredirect,345)
		insform, err := db.Prepare("UPDATE emp1 set passwor=?, email=? where nam=?")
		if err != nil {
			panic(err.Error())
		}
		session, err := usersession.Get(r, "authuser")

		if session == nil {
			fmt.Println("update session invalid", err)
		} else {
			Uname := session.Values["username"]
			insform.Exec(Upassword, Uemail, Uname)
			fmt.Fprintln(w, first+"<h5>Updated user",Uname,"</h5><h6>Email: ",Uemail,"<br/></h6><a href=\"/index\">HOME</a>"+last)
			//http.Redirect(w,r,"/index",http.StatusMovedPermanently)
		}
	} else {
		http.Redirect(w, r, "/update", http.StatusMovedPermanently)

	}

	defer db.Close()
}

/*func deletepagehandler(w http.ResponseWriter, r *http.Request){
	bytes, _ := ioutil.ReadFile("delete.html")

	fmt.Fprint(w, string(bytes))
}*/
func deletehandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	session, _ := usersession.Get(r, "auth user")
	if session == nil {
		fmt.Println("session invalid")
	} else {
		Uname := session.Values["username"]

		//http.Redirect(w,r,uredirect,345)
		insform, err := db.Prepare("DELETE FROM emp1 where nam=?")
		if err != nil {
			panic(err.Error())
		}
		insform.Exec(Uname)
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)

	}

	defer db.Close()
}

/*func showpagehandler(w http.ResponseWriter, r *http.Request){
	bytes, _ := ioutil.ReadFile("show.html")

	fmt.Fprint(w, string(bytes))
}*/
func showhandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//Uname:=r.FormValue("name")
	session, _ := usersession.Get(r, "auth user")
	if session == nil {
		fmt.Println("session invalid")
	} else {
		Uname := session.Values["username"]

		selectque, err := db.Query("SELECT * from emp1 where nam=?", Uname)
		if err != nil {
			panic(err.Error())
		}
		for selectque.Next() {
			var tempname, tempemail, temppass string
			err = selectque.Scan(&tempname, &tempemail, &temppass)
			if err != nil {
				panic(err.Error())
			}
			fmt.Fprintln(w, tempname)
			fmt.Fprintln(w, tempemail)
		}

	}

	defer db.Close()
}

func main() {

	var router = mux.NewRouter()

	router.HandleFunc("/login", loginhandler).Methods("POST")
	router.HandleFunc("/index", indexhandler)
	router.HandleFunc("/", loginpagehandler).Methods("GET")
	router.HandleFunc("/register", registerhandler).Methods("POST")
	router.HandleFunc("/register", registerpagehandler).Methods("GET")
	router.HandleFunc("/update", updatehandler).Methods("POST")
	router.HandleFunc("/update", updatepagehandler).Methods("GET")
	router.HandleFunc("/delete", deletehandler).Methods("POST")
	//router.HandleFunc("/delete",deletepagehandler).Methods("GET")
	router.HandleFunc("/show", showhandler).Methods("POST")
	//router.HandleFunc("/show",showpagehandler).Methods("GET")

	http.ListenAndServe(":9000", router)

}
