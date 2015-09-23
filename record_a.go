package infoblox

// https://192.168.2.200/wapidoc/objects/record.host.html
func (c *Client) RecordA() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:a",
	}
}

type RecordAObject struct {
	Object
}

func (c *Client) RecordAObject(ref string) *RecordAObject {
	return &RecordAObject{
		Object{
			_ref: ref,
			r:    c.RecordA(),
		},
	}
}
