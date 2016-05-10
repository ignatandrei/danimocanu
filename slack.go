package aorel_mta

import (
    "github.com/nlopes/slack"
    "log"
)

const ANDREI_VERSACE = "http://a1.ro/uploads/image/versace2.jpg"
const DANI_MOCANU = "http://static4.libertatea.ro/wp-content/uploads/2016/03/dani_mocanu.jpg"

var SLACK *slack.Client

func init() {
    if CONFIG_SLACK_ENABLED == false {
        log.Println("aorel-mta slack integration disabled")
        return
    }
    
    log.Println("aorel-mta slack integration enabled")
    SLACK = slack.New(CONFIG_SLACK_TOKEN)
}

func AndreiVersaceSays(message string) {
    if CONFIG_SLACK_ENABLED == true {
        SLACK.PostMessage("aorel", message, slack.PostMessageParameters{
            Text:message,
            Username:"Andrei Versace",
            IconURL:ANDREI_VERSACE,
        })
    }
}
func DaniMocanuSays(message string) {
    if CONFIG_SLACK_ENABLED == true {
        SLACK.PostMessage("aorel", message, slack.PostMessageParameters{
            Text:message,
            Username:"Dani Mocanu",
            IconURL:DANI_MOCANU,
    
        })
    }
}
