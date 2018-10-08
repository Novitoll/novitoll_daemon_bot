package bot

/*
	Telegram WebHook request body

	Example: URL / Text
	{
	  "message": {
	    "from": {
	      "username": "novitoll",
	      "first_name": "novitoll",
	      "is_bot": false,
	      "id": 345019684,
	      "language_code": "en-US"
	    },
	    "text": "https://weproject.kz/articles/detail/o-tom-kak-zarabotat-4000-dollarov-za-12-dney-i-ne-sidet-v-ofise/",
	    "entities": [
	      {
	        "length": 101,
	        "type": "url",
	        "offset": 0
	      }
	    ],
	    "chat": {
	      "username": "novitoll",
	      "first_name": "novitoll",
	      "type": "private",
	      "id": 345019684
	    },
	    "date": 1537020424,
	    "message_id": 28
	  },
	  "update_id": 776799951
	}
*/

type BotRequest struct {
	Update_Id	uint32 `json:"update_id"`
	Message		BotRequestMessage
}

type BotRequestMessage struct {
	From User
	Text		string
	Entities	[]MessageEntity
	Date		uint32
	MessageId	uint32
	Chat		BotRequestChat
	NewComer	User `json:"new_chat_members"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type		string
	Length		int
	Url		string
	MentionedUser	User `json:"user"`
}

// https://core.telegram.org/bots/api#user
type User struct {
	Username	string	`json:"username"`
	First_Name	string	`json:"first_name"`
	Is_Bot		bool	`json:"is_bot"`
	Language_Code	string	`json:"language_code"`
}

type BotRequestChat struct {
	Username	string
	First_Name	string `json:"first_name"`
	Type		string
	Id		uint32
}