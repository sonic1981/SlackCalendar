package SlackCalender

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

type SlackMessage struct {
	Ok       bool `json:"ok"`
	Channels []struct {
		ID                      string        `json:"id"`
		Name                    string        `json:"name"`
		IsChannel               bool          `json:"is_channel"`
		IsGroup                 bool          `json:"is_group"`
		IsIm                    bool          `json:"is_im"`
		IsMpim                  bool          `json:"is_mpim"`
		IsPrivate               bool          `json:"is_private"`
		Created                 int           `json:"created"`
		IsArchived              bool          `json:"is_archived"`
		IsGeneral               bool          `json:"is_general"`
		Unlinked                int           `json:"unlinked"`
		NameNormalized          string        `json:"name_normalized"`
		IsShared                bool          `json:"is_shared"`
		IsOrgShared             bool          `json:"is_org_shared"`
		IsPendingExtShared      bool          `json:"is_pending_ext_shared"`
		PendingShared           []interface{} `json:"pending_shared"`
		ParentConversation      interface{}   `json:"parent_conversation"`
		Creator                 string        `json:"creator"`
		IsExtShared             bool          `json:"is_ext_shared"`
		SharedTeamIds           []string      `json:"shared_team_ids"`
		PendingConnectedTeamIds []interface{} `json:"pending_connected_team_ids"`
		IsMember                bool          `json:"is_member"`
		Topic                   struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"topic"`
		Purpose struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"purpose"`
		PreviousNames []interface{} `json:"previous_names"`
		NumMembers    int           `json:"num_members"`
	} `json:"channels"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

func GetChannelsObj() SlackMessage {

	var returnObj SlackMessage
	byteArr := GetChannels()

	err := json.Unmarshal(byteArr, &returnObj)
	if err != nil {
		panic("Failed to parse byte array json to SlackMessage")
	}

	return returnObj
}

func GetChannels() []byte {

	token := *flag.String("t", "", "The Slack Token required to access slack.")

	req, err := http.NewRequest("GET", "https://slack.com/api/conversations.list", nil)

	req.Header.Set("Authorization", token)

	if err != nil {
		panic(err)
	}

	resp, err := netClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return body
}
