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
  //"unicode/utf8"
  //"container/list"
  //"bytes"
	"net/http"
	//"encoding/json"
	//"net/url"
  //"strings"
	// "database/sql"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
	// _ "github.com/go-sql-driver/mysql"
)


var bot *linebot.Client
var echo string 
var op string
var bottun bool

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


func main() {
	var err error
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
func httpGet(q string , id string) {
	
    echo = "OK"
    bottun = false
    op = ""

    strings.Replace(q," ", ",", -1)
    
    resp, err := http.Get("http://140.115.54.93:54321/chatbot?q="+q+"&id="+id)
    if err != nil {
        // handle error
       panic(err.Error())
    }
    defer resp.Body.Close()
    
    
     
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
       panic(err.Error())
    }
    
    echo = string(body) 

    if echo == "正在為您轉接"{
        _, err = bot.PushMessage("Uf6263c4b814700c680228b8b64a27dd6", linebot.NewTextMessage(q)).Do()
        if err != nil {
        // handle error
        echo = err.Error()
        }   
    }
   
    // _, err = bot.PushMessage("Uf6263c4b814700c680228b8b64a27dd6", linebot.NewTextMessage(echo)).Do()
   
    
    //------------for Luis
    // var r =  map[string]interface{}{}
    // var tempString string
    // tempString =string(body) 
   

    // temp1 := strings.Split(tempString,"entity")
    // temp2 := strings.Split(tempString,"\"Type")
    // entity := list.New()
    // Type := list.New()
    // for i := 0; i < len(temp1); i++ {
    //   if i>=1{
    //     entity.PushBack( strings.Split(strings.Split(temp1[i],"\",")[0],":\"")[1] )
    //   }
    // }
    // for i := 0; i < len(temp2); i++ {
    //   if i>=1{
    //     Type.PushBack( strings.Split(strings.Split(temp2[i],"\"}")[0],":\"")[1] )
    //   }
    // }
    // json.Unmarshal(body, &r)
    

    // if r["resultType"].(string) == "none" {
    //   echo = "我不了解你在說什麼～@@"
    // } else if r["resultType"].(string) == "greeting" {
    //   echo = "你好！我是LUIS！我可以提供您數學的教材或是練習題喔！"
    // } else if r["resultType"].(string) == "appreciation" {
    //   echo = "歡迎您再次使用LUIS!我很樂意再次提供您服務！"
    // } else if r["resultType"].(string) == "connectionError" {
    //   echo = "對不起，我出了點問題，現在沒辦法回答你問題@@"
    // } else if r["resultType"].(string) == "unknown" {
    //   echo = "不好意思，我不知道你問的定理是什麼QQ"
    // } else if r["resultType"].(string) == "question" {
    //   if r["requirementType"].(string) == "none" {
    //     bottun = true
    //     for e:= entity.Front();e!=nil;e = e.Next(){
         
    //      op += e.Value.(string) 
    //     }
    //   } 
    // }
    
    //-----------------for luis
    
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
        
        if op!= ""{
          echo ="OK"
          bottun = false


        } else if event.Source.UserID == ""{
            _, err = bot.PushMessage("Uf6263c4b814700c680228b8b64a27dd6", linebot.NewTextMessage(message.Text)).Do()
            if err != nil {
            // handle error
            //    ..
            echo = err.Error()
            }   
        } else {

          httpGet(message.Text,event.Source.UserID)
        }
				// mysql()
				//message.ID+":"+message.Text
        if bottun {
           //_, err = bot.ReplyMessa  ge(event.ReplyToken, linebot.NewTextMessage(echo)).Do()
            var ff string
            var gg string
            ff = "我要"+op+"的練習題"
            gg = "我要"+op+"的教材"
            leftBtn := linebot.NewMessageTemplateAction("練習題", ff)
            rightBtn := linebot.NewMessageTemplateAction("教材", gg)

            template := linebot.NewConfirmTemplate("請問是需要練習題還是教材?", leftBtn, rightBtn)

            templatemessgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
            _, err = bot.ReplyMessage(event.ReplyToken, templatemessgage).Do()
           //op=""

        } else {
           _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(echo)).Do()
           // _, err = bot.PushMessage(event.ReplyToken, linebot.NewTextMessage(echo)).Do()
           op=""
        }
				
          
        
			}
		}
	}
}
