package mixpanel

import "testing"

func TestEvent(t *testing.T) {

	mixpanelAPI := New("")
	event := &Event{
		Event: "Test",
		Properties: map[string]string{
			"name": "hello",
		},
	}
	mixpanelAPI.track(event)
}

