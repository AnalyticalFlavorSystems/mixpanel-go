package mixpanel

import "testing"

func TestEvent(t *testing.T) {

	mixpanelAPI := New("1234")
	event := &Event{
		Event: "Test",
		Properties: map[string]string{
			"name": "hello",
		},
	}
	mixpanelAPI.track(event)
}

func TestSet(t *testing.T) {
	mixpanelAPI := New("1234")
	profile := &Profile{
		DistinctId: "abc",
		Operation: "set",
		Values: map[string]string {
			"test": "hi",
		},
	}
	mixpanelAPI.set(profile)
}

