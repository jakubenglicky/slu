package cmd

import (
	_ "github.com/sikalabs/slu/cmd/argocd"
	_ "github.com/sikalabs/slu/cmd/argocd/initial_password"
	_ "github.com/sikalabs/slu/cmd/argocd/set_image"
	_ "github.com/sikalabs/slu/cmd/digitalocean"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all/cleanup"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all/delete"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all/list"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/add"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/list"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/rm"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/use_context"
	_ "github.com/sikalabs/slu/cmd/du"
	_ "github.com/sikalabs/slu/cmd/example_server"
	_ "github.com/sikalabs/slu/cmd/expand"
	_ "github.com/sikalabs/slu/cmd/expand/file"
	_ "github.com/sikalabs/slu/cmd/expand/string"
	_ "github.com/sikalabs/slu/cmd/file_templates"
	_ "github.com/sikalabs/slu/cmd/file_templates/editorconfig"
	_ "github.com/sikalabs/slu/cmd/file_templates/gitignore"
	_ "github.com/sikalabs/slu/cmd/file_templates/go_cli_project"
	_ "github.com/sikalabs/slu/cmd/file_templates/incident_response"
	_ "github.com/sikalabs/slu/cmd/generate_docs"
	_ "github.com/sikalabs/slu/cmd/generate_files"
	_ "github.com/sikalabs/slu/cmd/generate_files/tree"
	_ "github.com/sikalabs/slu/cmd/git"
	_ "github.com/sikalabs/slu/cmd/git/commit"
	_ "github.com/sikalabs/slu/cmd/git/commit/terraform_lock"
	_ "github.com/sikalabs/slu/cmd/git/if"
	_ "github.com/sikalabs/slu/cmd/git/if/staged"
	_ "github.com/sikalabs/slu/cmd/git/url"
	_ "github.com/sikalabs/slu/cmd/git/url/get"
	_ "github.com/sikalabs/slu/cmd/git/url/open"
	_ "github.com/sikalabs/slu/cmd/gitlab"
	_ "github.com/sikalabs/slu/cmd/gitlab/prune_environments"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage/remove_skip"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage/show"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage/skip"
	_ "github.com/sikalabs/slu/cmd/go_code"
	_ "github.com/sikalabs/slu/cmd/go_code/version_bump"
	_ "github.com/sikalabs/slu/cmd/install_bin"
	_ "github.com/sikalabs/slu/cmd/install_bin_tool"
	_ "github.com/sikalabs/slu/cmd/ip"
	_ "github.com/sikalabs/slu/cmd/ip_local"
	_ "github.com/sikalabs/slu/cmd/k8s"
	_ "github.com/sikalabs/slu/cmd/k8s/delete_ns"
	_ "github.com/sikalabs/slu/cmd/k8s/get"
	_ "github.com/sikalabs/slu/cmd/k8s/get/configmap"
	_ "github.com/sikalabs/slu/cmd/k8s/get/secret"
	_ "github.com/sikalabs/slu/cmd/k8s/kubeconfig"
	_ "github.com/sikalabs/slu/cmd/k8s/kubeconfig/add"
	_ "github.com/sikalabs/slu/cmd/k8s/lock"
	_ "github.com/sikalabs/slu/cmd/k8s/lock/lock"
	_ "github.com/sikalabs/slu/cmd/k8s/lock/status"
	_ "github.com/sikalabs/slu/cmd/k8s/lock/unlock"
	_ "github.com/sikalabs/slu/cmd/k8s/token"
	_ "github.com/sikalabs/slu/cmd/length"
	_ "github.com/sikalabs/slu/cmd/loggen"
	_ "github.com/sikalabs/slu/cmd/mail"
	_ "github.com/sikalabs/slu/cmd/mail/send"
	_ "github.com/sikalabs/slu/cmd/mysql"
	_ "github.com/sikalabs/slu/cmd/mysql/create"
	_ "github.com/sikalabs/slu/cmd/mysql/drop"
	_ "github.com/sikalabs/slu/cmd/mysql/generate"
	_ "github.com/sikalabs/slu/cmd/mysql/generate/random_data"
	_ "github.com/sikalabs/slu/cmd/mysql/list"
	_ "github.com/sikalabs/slu/cmd/mysql/ping"
	_ "github.com/sikalabs/slu/cmd/postgres"
	_ "github.com/sikalabs/slu/cmd/postgres/create"
	_ "github.com/sikalabs/slu/cmd/postgres/drop"
	_ "github.com/sikalabs/slu/cmd/postgres/list"
	_ "github.com/sikalabs/slu/cmd/postgres/ping"
	_ "github.com/sikalabs/slu/cmd/proxy"
	_ "github.com/sikalabs/slu/cmd/proxy/smtp"
	_ "github.com/sikalabs/slu/cmd/proxy/tcp"
	_ "github.com/sikalabs/slu/cmd/random"
	_ "github.com/sikalabs/slu/cmd/random/password"
	_ "github.com/sikalabs/slu/cmd/random/string"
	_ "github.com/sikalabs/slu/cmd/rmline"
	"github.com/sikalabs/slu/cmd/root"
	_ "github.com/sikalabs/slu/cmd/sleep"
	_ "github.com/sikalabs/slu/cmd/sleep/forever"
	_ "github.com/sikalabs/slu/cmd/sleep/random"
	_ "github.com/sikalabs/slu/cmd/sqlite"
	_ "github.com/sikalabs/slu/cmd/sqlite/read"
	_ "github.com/sikalabs/slu/cmd/static_api"
	_ "github.com/sikalabs/slu/cmd/static_api/version"
	_ "github.com/sikalabs/slu/cmd/time"
	_ "github.com/sikalabs/slu/cmd/time/prefix"
	_ "github.com/sikalabs/slu/cmd/time/unix"
	_ "github.com/sikalabs/slu/cmd/tls"
	_ "github.com/sikalabs/slu/cmd/tls/parse"
	_ "github.com/sikalabs/slu/cmd/tls/parse_file"
	_ "github.com/sikalabs/slu/cmd/tls/parse_k8s_secret"
	_ "github.com/sikalabs/slu/cmd/version"
	_ "github.com/sikalabs/slu/cmd/wait_for"
	_ "github.com/sikalabs/slu/cmd/wait_for/k8s"
	_ "github.com/sikalabs/slu/cmd/wait_for/k8s/job"
	_ "github.com/sikalabs/slu/cmd/wait_for/k8s/pod"
	_ "github.com/sikalabs/slu/cmd/wait_for_it"
	_ "github.com/sikalabs/slu/cmd/wait_for_tls"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
