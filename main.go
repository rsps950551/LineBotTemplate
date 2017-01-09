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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"net/url"
	// "database/sql"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	// _ "github.com/go-sql-driver/mysql"
)


var bot *linebot.Client
var echo string
var FF = []byte(`{
      "Type": "confirm",
      "text": "Are you sure?",
      "actions": [
          {
            "type": "message",
            "label": "Yes",
            "text": "yes"
          },
          {
            "type": "message",
            "label": "No",
            "text": "no"
          }
      ]
	}`)
var m json
// const confirm = `
// {
//   "type": "template",
//   "altText": "this is a confirm template",
//   "template": {
//       "type": "confirm",
//       "text": "Are you sure?",
//       "actions": [
//           {
//             "type": "message",
//             "label": "Yes",
//             "text": "yes"
//           },
//           {
//             "type": "message",
//             "label": "No",
//             "text": "no"
//           }
//       ]
//   }
// }`

const GG =`{
      "type": "confirm",
      "text": "Are you sure?",
      "actions": [
          {
            "type": "message",
            "label": "Yes",
            "text": "yes"
          },
          {
            "type": "message",
            "label": "No",
            "text": "no"
          }
      ]
	}`

type actions struct {
       Type string
       label string
       text string
}

type template struct { 
       Type string
       text string
       actions [2]actions
}

// type confirm struct {
//        type string
//        altText string
//        template template
// }

// var FF template

func main() {
	var err error
  var n = actions{"message","NO","no"}
  var y = actions{"message","yes","yes"}
  var t = template{"confirm","FF",[n,y]}
	// json.Unmarshal([]byte(GG), &FF)
  m, err := json.Marshal(t)
	//json.Unmarshal(FF, &m)
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)	

}

// func mysql(){
// 	var db, err = sql.Open("mysql","wmlab:wmlab@tcp(140.115.54.82:3306)/wmlab?charset=utf8")
// 	if err != nil {
// 		fmt.Println(err)
//          // Just for example purpose. You should use proper error handling instead of panic
//     }
//     defer db.Close()

//  	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rows, err := db.Query("select * from test")
// 	if err != nil {
// 		log.Println(err)
// 	}
 
// 	defer rows.Close()
// 	var col1 int
// 	for rows.Next() {
// 		err := rows.Scan(&col1)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		dbinfo=col1
// 	}
// }
func httpGet(q string) {
	//encodeurl:= url.QueryEscape("http://140.115.54.82/luis.php?question="+q)
    resp, err := http.PostForm("http://140.115.54.82/luis.php",url.Values{"question": {q}})
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    echo = string(body)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	//GG
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				httpGet(message.Text)
				// mysql()
				//message.ID+":"+message.Text
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("FF",m)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
