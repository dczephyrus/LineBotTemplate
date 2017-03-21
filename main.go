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
	"log"
	"net/http"
	"os"
	"strings"
	
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var gradeschoolreply := [5]Strings{
  "中2發言1",
  "中2發言2",
  "中2發言3",
  "中2發言4",
  "中2發言5",
}
func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
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

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyMessage := parseMessage(message.Text)
				if(replyMessage!=""){					
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage+" !")).Do(); err != nil {
					log.Print(err)
					}
				}
			}
		}
	}
}
func parseMessage(input Strings)Strings{
	resultString:=""
	if(Strings.Contains(input,"你好")) {
		resultString="你好,迷途的羔羊"
	}else if(Strings.Contains(input,"存在X")||Strings.Contains(input,"存在x")){
		resultString="中2發言預計區"
	}else{ 	//in group behave guite
		resultString=""
	}
	return resultString
}
