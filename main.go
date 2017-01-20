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
  // "strings"
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
// var m json
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
type content struct{
    entity string `json:"entity"`
    Type string `json:"Type"`
}

type Data struct{
    resultType string `json:"resultType"`
    resultQuestion string `json:"resultTypeQuestion"`
    resultContent []content `json:"resultContent"`
    requirementType string `json:"requirementType"`
}

var d Data
// type actions struct {
//        type string
//        label string
//        text string
// }

// type template struct { 
//        type string
//        text string
//        actions [2]actions
// }

// type confirm struct {
//        type string 
//        altText string
//        template template
// }

// var FF template

func main() {
	var err error
 //  var y = map[string]interface{}{
 //        "type": "message",
 //        "label": "Yes",
 //        "text": "yes",
 //      }
 //  var n = map[string]interface{}{
 //        "type": "message",
 //        "label": "No",
 //        "text": "no",
 //    }
 //  // var t = template{"confirm","FF",{n,y}}
	// // json.Unmarshal([]byte(GG), &FF)
 //  var cacheContent = map[string]interface{}{
 //    "type": "confirm",
 //    "text": "Are you sure?",
 //    "actions":map[string]interface{}{
 //        "type": "message",
 //        "label": "Yes",
 //        "text": "yes",
 //      },
 //  }  
  // m, err := json.Marshal(cacheContent)
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
    // u := map[string]interface{}{}
    body, err := ioutil.ReadAll(resp.Body)
    // er := json.NewDecoder(strings.NewReader(body)).Decode(ff)
    json.Unmarshal(body, &d)
    echo = string(d.resultType)

    if(q=="give me bottun"){
      echo = "bottun"
    }
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
  leftBtn := linebot.NewMessageTemplateAction("left", "left clicked")
  rightBtn := linebot.NewMessageTemplateAction("right", "right clicked")

  template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

  templatemessgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
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
        if(echo == "bottun"){
           _, err = bot.ReplyMessage(event.ReplyToken, templatemessgage).Do()
        } else {
           _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(echo)).Do()
        }
				
          
        
			}
		}
	}
}
