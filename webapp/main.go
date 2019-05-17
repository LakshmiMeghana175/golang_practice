package main

import (
	"database/sql"
	"fmt"

	"html/template"
	"io/ioutil"
	"mux"
	_ "mysql"
	"net/http"
	"sessions"
	"strconv"
)

type employee struct{
	Empname string
	Empemail string
	Empcontact int64
	Empbloodgroup string
	Empskills string
	Emplocation string
	Emppassword string

}

var first string = "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><body>"
var last string = "</body></html>"

var encryptionkey = "vfdhgfdfgsd"
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
	session,_:= usersession.Get(r, "authouser1")
	//session.Values["username"]=""
	//fmt.Println(session.Values["username"])
	if session.Values["username"]!="" {

		fmt.Println("login again")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}


	var temppass string
	if len(Uname) != 0 && len(Upassword) != 0 {
		session, _= usersession.New(r, "authouser1")

		selectque, err := db.Query("SELECT emppassword from employee where empname=?", Uname)
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
	      // w.Write([]byte(first+"sucessfully logged in....."+last))
			http.Redirect(w, r, "/index", 307)
		} else {
			http.Redirect(w, r, "/register", http.StatusMovedPermanently)

		}

	} else {
		http.Redirect(w, r, "/register", http.StatusMovedPermanently)

	}
	defer db.Close()

}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	showhandler(w,r)

	//bytes, _ := ioutil.ReadFile("index.html")

	//fmt.Fprint(w, string(bytes))

}
func registerpagehandler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadFile("register.html")

	fmt.Fprint(w, string(bytes))

}
func registerhandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	temp1,_:=strconv.ParseInt(r.FormValue("contact"),10,64)


	e := employee{r.FormValue("name"),
		 r.FormValue("email"),
		 temp1,
		 r.FormValue("bloodgroup"),
		 r.FormValue("skills"),
		 r.FormValue("location"),
		 r.FormValue("password")	}
	temp2:=r.FormValue("confirmpassword")

	if len(e.Empname) != 0 && len(e.Emppassword) != 0 && len(e.Empemail) != 0 && len(temp2) != 0 && len(e.Empbloodgroup) != 0 && len(e.Empskills) != 0 && len(e.Emplocation) != 0{
		if e.Emppassword==temp2 {

		insform, err := db.Prepare("INSERT INTO employee VALUES(?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insform.Exec(e.Empname, e.Empemail, e.Empcontact,e.Empbloodgroup,e.Empskills,e.Emplocation,e.Emppassword)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		fmt.Fprintln(w, "Sucesfully registered ..please login")

	}else {
			fmt.Println( "Password and Confirm Password should be same")
			http.Redirect(w, r, "/register", http.StatusMovedPermanently)

		}
}else {
		fmt.Println("These fields should not be empty please register again")
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
	session,_:= usersession.Get(r, "authouser1")


	 Uname := session.Values["username"]
	 var tempname string

	temp1,_:=strconv.ParseInt(r.FormValue("contact"),10,64)
	selectque, err := db.Query("SELECT empname from employee where empname=?", Uname)
	if err != nil {
		panic(err.Error())
	}


	for selectque.Next() {

		err = selectque.Scan(&tempname)
		if err != nil {
			panic(err.Error())
		}

	}


	e := employee{tempname,
		r.FormValue("email"),
		temp1,
		r.FormValue("bloodgroup"),
		r.FormValue("skills"),
		r.FormValue("location"),
		r.FormValue("password")	}
	if len(e.Empname) != 0 && len(e.Emppassword) != 0 && len(e.Empemail) != 0 && len(e.Empbloodgroup) != 0 && len(e.Empskills) != 0 && len(e.Emplocation) != 0 {

		insform, err := db.Prepare("UPDATE employee set  empemail=?,empcontact=?, empbloodgroup=?,empskills=?, emplocation=?, emppassword=? where empname=?")
		if err != nil {
			panic(err.Error())
		}

			insform.Exec(e.Empemail, e.Empcontact,e.Empbloodgroup,e.Empskills,e.Emplocation,e.Emppassword,Uname)
			tmp1:=template.Must(template.ParseFiles("update1.html"))
			tmp1.Execute(w,e)
			//fmt.Fprintln(w, first+"<center>Updated user",Uname,"Email: ",Uemail,"<br/><br/><a href=\"/index\">home</a></center>"+last)

			//http.Redirect(w,r,"/index",http.StatusMovedPermanently)

	} else {
		http.Redirect(w, r, "/update", http.StatusMovedPermanently)

	}

	defer db.Close()
}

func deletepagehandler(w http.ResponseWriter, r *http.Request){

	bytes, _:= ioutil.ReadFile("delete.html")

	fmt.Fprint(w, string(bytes))
}
func deletehandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w,r,"delete.html")
	Upassword:=r.FormValue("password")


	db := dbConn()

	session, _:= usersession.Get(r, "authouser1")
	Uname := session.Values["username"]

	if session == nil {
		fmt.Println("session invalid")
	} else {
		Uname = session.Values["username"]



		insform, err := db.Prepare("DELETE FROM employee where empname=? and emppassword=?")
		if err != nil {
			panic(err.Error())
		}
		insform.Exec(Uname,Upassword)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}

	defer db.Close()
}


func showhandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()


	e:=employee{}
	session, _ := usersession.Get(r, "authouser1")
	if session == nil {
		fmt.Println("session invalid")
	} else {
		Uname := session.Values["username"]

		selectque, err := db.Query("SELECT * from employee where empname=?", Uname)
		if err != nil {
			panic(err.Error())
		}
		for selectque.Next() {

			err = selectque.Scan(&e.Empname,&e.Empemail,&e.Empcontact,&e.Empbloodgroup,&e.Empskills,&e.Emplocation,&e.Emppassword)
			if err != nil {
				panic(err.Error())
			}




			tmp1:=template.Must(template.ParseFiles("index.html"))
			tmp1.Execute(w,e)

		}

	}

	defer db.Close()
}
func logouthandler(w http.ResponseWriter, r *http.Request){
	session, _ := usersession.Get(r, "authouser1")
	session.Values["username"]=""
	_=session.Save(r,w)
	fmt.Println(session.Values["username"],"sdgfd")
	http.Redirect(w,r,"/",http.StatusMovedPermanently)


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
	router.HandleFunc("/delete",deletepagehandler).Methods("GET")

	router.HandleFunc("/logou", logouthandler)


	http.ListenAndServe(":9000", router)

}
