package httpclient

import (
	"encoding/json"
	"testing"

	. "github.com/r7kamura/gospel"
)

type (
	payloadJSON struct {
		Text string `json:"text"`
	}
)

func TestDescribe(t *testing.T) {
	Describe(t, "Post JSON test", func() {
		req := new(Requestinfo)
		req.Url = "http://localhost:8080/"
		req.User = ""
		req.Password = ""

		Context("Post Dummy Json", func() {
			mes := &payloadJSON{
				Text: "test",
			}
			body, _ := json.Marshal(mes)
			_, err := PostJsontoTarget(req, body)
			It("err is not nil", func() {
				Expect(err).To(Equal, nil)
			})
		})
	})

}
