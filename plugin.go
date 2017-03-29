package kong

type RequestTransformer struct {
	Name               string `url:"name"`
	ConsumerId         string `url:"consumer_id,omitempty"`
	RemoveHeaders      string `url:"config.remove.headers,omitempty"`
	RemoveQuerystring  string `url:"config.remove.querystring,omitempty"`
	RemoveBody         string `url:"config.remove.body,omitempty"`
	ReplaceHeaders     string `url:"config.replace.headers,omitempty"`
	ReplaceQuerystring string `url:"config.replace.querystring,omitempty"`
	ReplaceBody        string `url:"config.replace.body,omitempty"`
	AddHeaders         string `url:"config.add.headers,omitempty"`
	AddQuerystring     string `url:"config.add.querystring,omitempty"`
	AddBody            string `url:"config.add.body,omitempty"`
	AppendHeaders      string `url:"config.append.headers,omitempty"`
	AppendQuerystring  string `url:"config.append.querystring,omitempty"`
	AppendBody         string `url:"config.append.body,omitempty"`
}
