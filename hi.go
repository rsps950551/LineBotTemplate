// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	
	"log"
	"fmt"
	"io/ioutil"
	"database/sql"
	"net/http"
	//"net/url"

	// "github.com/line/line-bot-sdk-go/linebot"
	_ "github.com/go-sql-driver/mysql"
)
 

func main() {
 	httpGet()
}

func mysql(){
	var db, err = sql.Open("mysql","wmlab:wmlab@tcp(140.115.54.82:3306)/wmlab?charset=utf8")
	if err != nil {
		fmt.Println(err)
         // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

 	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from test")
	if err != nil {
		log.Println(err)
	}
 
	defer rows.Close()
	var col1 int
	for rows.Next() {
		err := rows.Scan(&col1)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(col1)
	}
}
func httpGet() {
	//encodeurl:= url.QueryEscape("http://140.115.54.82/luis.php?question="+q)
    resp, err := http.Get("http://140.115.54.84/gg.php")
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    log.Println(string(body));
}

