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

Special Event properties (optional
* distinct\_id (string): The value of distinct\_id will be treated as a string, and used to uniquely identify a user associated with your event. If you provide a distinct\_id property with your events, you can track a given user through funnels and distinguish unique users for retention analyses. You should always send the same distinct_id when an event is triggered by the same user.

* time (number): The time an event occurred. If present, the value should be a unix timestamp (seconds since midnight, January 1st, 1970 - UTC). If this property is not included in your request, Mixpanel will use the time the event arrives at the server.

* ip (string): An IP address string (e.g. "127.0.0.1") associated with the event. This is used for adding geolocation data to events, and should only be required if you are making requests from your backend. If "ip" is absent (and ip=1 is not provided as a URL parameter), Mixpanel will ignore the IP address of the request.
