package apiv1

type ListParams struct {
	Offset   int64  `form:"offset" url:"offset,omitempty"`
	Provider string `form:"provider" url:"provider,omitempty"`
	Verified bool   `form:"verified" url:"verified,omitempty"`
}

type SearchParams struct {
	Q         string `form:"q" url:"q,omitempty"`
	Offset    int64  `form:"offset" url:"offset,omitempty"`
	Provider  string `form:"provider" url:"provider,omitempty"`
	Namespace string `form:"namespace" url:"namespace,omitempty"`
	Verified  bool   `form:"verified" url:"verified,omitempty"`
}

type ListLatestVersionParams struct {
	Offset int64 `form:"offset" url:"offset,omitempty"`
}
