package entities

type LinksRequest struct {
	Links []string `json:"links"`
}

type LinksResponse struct {
	Links    map[string]string `json:"links"`
	LinksNum int               `json:"links_num"`
}

type LinksNumRequest struct {
	LinksList []int `json:"links_list"`
}
