package httpClient

type GetList []Get

type Get struct {
	path        string
	queryString string
	host        string
}

func NewGet() *Get {
	return &Get{}
}

func (r *Get) SetPath(path string) *Get {
	r.path = path

	return r
}

func (r *Get) SetQueryString(queryString string) *Get {
	r.queryString = queryString

	return r
}

func (r *Get) SetHost(host string) *Get {
	r.host = host

	return r
}

type PostList []Post

type Post struct {
	host    string
	request interface{}
	path    string
}

func NewPost() *Post {
	return &Post{}
}

func (r *Post) SetPath(path string) *Post {
	r.path = path

	return r
}

func (r *Post) SetHost(host string) *Post {
	r.host = host

	return r
}

func (r *Post) SetRequest(request interface{}) *Post {
	r.request = request

	return r
}
