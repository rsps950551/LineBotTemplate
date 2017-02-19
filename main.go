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
  "container/list"
  //"bytes"
	"net/http"
	"encoding/json"
	"net/url"
  "strings"
	// "database/sql"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
	// _ "github.com/go-sql-driver/mysql"
)


var bot *linebot.Client
var echo string 
var op string
var bottun bool
// var FF = []byte(`{
//       "Type": "confirm",
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
// 	}`)
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

// const GG =`{
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
// 	}`

//
type Data struct{
    resultType string `json:"resultType"`
    resultQuestion string `json:"resultQuestion"`
    resultContent []content `json:"resultContent"`
    requirementType string `json:"requirementType"`
}

type content struct{
    entity string `json:"entity"`
    Type string `json:"Type"`
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
  bottun = false
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
    // var resultType string
    // var resultQuestion string
    // var resultContent string 
    // var requirementType string
    echo = "OK"
    bottun = false
    resp, err := http.PostForm("http://140.115.54.82/luis.php",url.Values{"question": {q}})
    if err != nil {
        // handle error
       panic(err.Error())
    }
    defer resp.Body.Close()
    //var r Data
    var r =  map[string]interface{}{}
    // var resultContent []content
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
       panic(err.Error())
    }
    // er:=json.NewDecoder(resp.Body).Decode(&D)
    // if er != nil {
    //     // handle error
    //   echo = "er. Error()"
    // } else {
    //   echo = "FK"
    // }
    var tempString string
    tempString =string(body) 
    //var entity []string
    //var Type []string

    temp1 := strings.Split(tempString,"entity")
    temp2 := strings.Split(tempString,"\"Type")
    entity := list.New()
    Type := list.New()
    for i := 0; i < len(temp1); i++ {
      if i>=1{
        entity.PushBack( strings.Split(strings.Split(temp1[i],"\",")[0],":\"")[1] )
      }
    }
    for i := 0; i < len(temp2); i++ {
      if i>=1{
        Type.PushBack( strings.Split(strings.Split(temp2[i],"\"}")[0],":\"")[1] )
      }
    }
    //temp2 := strings.Split(tempString,)

    //temp2 := strings.Split(temp1[1],",")

    json.Unmarshal(body, &r)
    // for n, a := range r["resultContent"] {  
    //   echo = echo + n + a
    // }  

    if r["resultType"].(string) == "none" {
      echo = "我不了解你在說什麼～@@"
    } else if r["resultType"].(string) == "greeting" {
      echo = "你好！我是LUIS！我可以提供您數學的教材或是練習題喔！"
    } else if r["resultType"].(string) == "appreciation" {
      echo = "歡迎您再次使用LUIS!我很樂意再次提供您服務！"
    } else if r["resultType"].(string) == "connectionError" {
      echo = "對不起，我出了點問題，現在沒辦法回答你問題@@"
    } else if r["resultType"].(string) == "unknown" {
      echo = "不好意思，我不知道你問的定理是什麼QQ"
    } else if r["resultType"].(string) == "question" {
      if r["requirementType"] == "none" {
        bottun = true
        for e:= entity.Front();e!=nil;e = e.Next(){
         op += e.Value.(string)+" "
        }
      } else {
        OP = ""
      }
    }
    //test~~~~~~
    // echo = "resultType: "+ r["resultType"].(string) +"\n"
    // echo += "resultQuestion: "+r["resultQuestion"].(string) +"\n"
    // echo += "resultContent:" +"\n"
    // t:= Type.Front()
    // for e:= entity.Front();e!=nil;e = e.Next(){
    //   echo += e.Value.(string)+" "+t.Value.(string)+"\n"
    //   t = t.Next()
    // }
    // echo += "requirementType: " + r["requirementType"].(string)
    

    
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

  leftBtn := linebot.NewMessageTemplateAction("練習題", "我要練習題 "+op)
  rightBtn := linebot.NewMessageTemplateAction("教材", "我要教材 "+op)

  template := linebot.NewConfirmTemplate("請問是需要練習題還是教材?", leftBtn, rightBtn)

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
        if bottun {
           // _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(echo)).Do()
           _, err = bot.ReplyMessage(event.ReplyToken, templatemessgage).Do()

        } else {
           _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(echo)).Do()
           
        }
				
          
        
			}
		}
	}
}
