package repos

import (
	"errors"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"strings"
)

type Repo struct {
	HostingSite string
	User string
	RepoName string
	Revision string
	Default bool
	//Properties RepoProperties
}

/*type RepoProperties struct {
	Default bool
	Priority int
}*/

func NewRepo(repoPath string) (*Repo, error){

	revSlice := strings.Split(repoPath, "@")

	var repoUrlSlice []string
	var revision string
	if len(revSlice) == 2 { //revision specified
		repoUrlSlice = strings.Split(revSlice[0], "/")
		revision = revSlice[1]
	} else if len(revSlice) == 1{ //no revision specified
		repoUrlSlice = strings.Split(revSlice[0], "/")
		revision = "master"
	} else {
		return &Repo{}, errors.New("Repo path resolution failed")
	}

	return &Repo{repoUrlSlice[0], repoUrlSlice[1],
		repoUrlSlice[2], revision, false}, nil
}

func (r *Repo) GetIdentifier() string {
	identifier := r.HostingSite + "/" + r.User + "/" + r.RepoName + "/"

	return identifier
}

func (r *Repo) GetCompleteIdentifier() string {
	identifier := r.HostingSite + "/" + r.User + "/" + r.RepoName

	if r.Revision != "" {
		identifier += "@" + r.Revision
	}

	return identifier
}

func (r *Repo) GetCloneDir() string {
	cloneDir := viper.GetString("repositoriesUri")
	if cloneDir[len(cloneDir)-1:] != "/" {
		cloneDir += "/"
	}

	cloneDir += r.HostingSite + "/" +
				r.User 		 + "/" +
				r.RepoName

	return cloneDir
}

func (r *Repo) GetCloneParentDirs() []string {
	cleanDir := viper.GetString("repositoriesUri")
	if cleanDir[len(cleanDir)-1:] != "/" {
		cleanDir += "/"
	}

	cleanDirUser := cleanDir +
		r.HostingSite + "/" +
		r.User

	cleanDirHostingSite := cleanDir +
		r.HostingSite

	ret := make([]string, 2)
	ret[0] = cleanDirUser
	ret[1] = cleanDirHostingSite
	return ret
}

func (r *Repo) GetUrl() string {
	return "https://" +
		r.HostingSite + "/" +
		r.User 		  + "/" +
		r.RepoName	  + ".git"
}

func (r *Repo) GetTaskDir() string {
	return r.GetCloneDir() + "/tasks/"
}

func (r *Repo) GetWorkflowDir() string {
	return r.GetCloneDir() + "/workflows/"
}

func (r *Repo) ResolveTaskClassIdentifier(loadTaskClass string) (taskClassIdentifier string) {
	if !strings.Contains(loadTaskClass, "/") {
		taskClassIdentifier = r.HostingSite + "/" + r.User + "/" + r.RepoName + "/tasks/" + loadTaskClass
	} else {
		taskClassIdentifier = loadTaskClass
	}

	if !strings.Contains(loadTaskClass, "@") {
		if r.Revision == "" {
			taskClassIdentifier += "@master"
		} else {
			taskClassIdentifier += "@" + r.Revision
		}
	}

	return
}

func (r *Repo) CheckoutRevision(revision string) error {
	if revision == "" {
		revision = r.Revision
	}

	ref, err := git.PlainOpen(r.GetCloneDir())
	if err != nil {
		return err
	}

	w, err := ref.Worktree()
	if err != nil {
		return err
	}

	//Try remotely as a priority (branches) so that we don't check out old, dangling branch refs (e.g. master)
	newHash, err := ref.ResolveRevision(plumbing.Revision("origin/" + revision))
	if err != nil {
	 	//Try locally (tags + hashes)
		newHash, err = ref.ResolveRevision(plumbing.Revision(revision))
		if err != nil {
			return errors.New("CheckoutRevision: " + err.Error())
		}
	}

	err = w.Checkout(&git.CheckoutOptions{
		Hash:  *newHash,
		Force: true,
	})
	if err != nil {
		return err
	}

	r.Revision = revision //Update repo revision
	return nil
}

func (r *Repo) RefreshRepo() error {

	ref, err := git.PlainOpen(r.GetCloneDir())
	if err != nil {
		return errors.New(err.Error() + ": " + r.GetIdentifier())
	}

	auth := &http.BasicAuth {
		Username: gitAuthUser,
		Password: gitAuthToken,
	}

	err = ref.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Auth: auth,
		Force: true,
	})

	if err != nil && err.Error() != "already up-to-date" {
		return errors.New(err.Error() + ": " + r.GetIdentifier() + " | revision: " + r.Revision)
	}

	err = r.CheckoutRevision("master")
	if err != nil {
		return err
	}

	return nil
}
