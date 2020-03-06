package v3

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type HarborMeta struct {
	Public bool `json:"public,omitempty"`
}

type HarborProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ProjectId         int        `json:"project_id,omitempty"`
	CreationTime      string     `json:"creation_time,omitempty"`
	CurrentUserRoleId int        `json:"current_user_role_id,omitempty"`
	RepoCount         int        `json:"repo_count,omitempty"`
	Metadata          HarborMeta `json:"harborMeta,omitempty"`
}

type HarborRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ProjectId    int    `json:"project_id,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
	TagsCount    int    `json:"tags_count,omitempty"`
	PullCount    int    `json:"pull_count,omitempty"`
}

type HarborTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Size          int64    `json:"size,omitempty"`
	Author        string   `json:"author,omitempty"`
	DockerVersion string   `json:"docker_version,omitempty"`
	TagLabels     []string `json:"tag_labels,omitempty"`
	FullName      string   `json:"full_name,omitempty"`
	Repo          string   `json:"repo,omitempty"`
}
