package apiv1

type ListOptions struct {
	Offset   int64  `form:"offset" url:"offset,omitempty"`
	Provider string `form:"provider" url:"provider,omitempty"`
	Verified bool   `form:"verified" url:"verified,omitempty"`
}

type SearchOptions struct {
	Q         string `form:"q" url:"q,omitempty"`
	Offset    int64  `form:"offset" url:"offset,omitempty"`
	Provider  string `form:"provider" url:"provider,omitempty"`
	Namespace string `form:"namespace" url:"namespace,omitempty"`
	Verified  bool   `form:"verified" url:"verified,omitempty"`
}

type ListLatestVersionOptions struct {
	Offset int64 `form:"offset" url:"offset,omitempty"`
}
