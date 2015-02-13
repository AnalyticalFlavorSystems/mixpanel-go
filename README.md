##This will track events, profile and more.  

To use
```golang
mxipanelAPI := New("PROJECT_ID")
event := &Event{
    Event: "Video Played",
    Properties: map[string]string{
        "Video length": "213",
    }
}
mixpanelAPI.track(event)
```
