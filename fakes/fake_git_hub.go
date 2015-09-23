// This file was generated by counterfeiter
package fakes

import (
	"io"
	"net/url"
	"os"
	"sync"

	"github.com/concourse/github-release-resource"
	"github.com/zachgersh/go-github/github"
)

type FakeGitHub struct {
	ListReleasesStub        func() ([]github.RepositoryRelease, error)
	listReleasesMutex       sync.RWMutex
	listReleasesArgsForCall []struct{}
	listReleasesReturns     struct {
		result1 []github.RepositoryRelease
		result2 error
	}
	GetReleaseByTagStub        func(tag string) (*github.RepositoryRelease, error)
	getReleaseByTagMutex       sync.RWMutex
	getReleaseByTagArgsForCall []struct {
		tag string
	}
	getReleaseByTagReturns struct {
		result1 *github.RepositoryRelease
		result2 error
	}
	CreateReleaseStub        func(release github.RepositoryRelease) (*github.RepositoryRelease, error)
	createReleaseMutex       sync.RWMutex
	createReleaseArgsForCall []struct {
		release github.RepositoryRelease
	}
	createReleaseReturns struct {
		result1 *github.RepositoryRelease
		result2 error
	}
	UpdateReleaseStub        func(release github.RepositoryRelease) (*github.RepositoryRelease, error)
	updateReleaseMutex       sync.RWMutex
	updateReleaseArgsForCall []struct {
		release github.RepositoryRelease
	}
	updateReleaseReturns struct {
		result1 *github.RepositoryRelease
		result2 error
	}
	ListReleaseAssetsStub        func(release github.RepositoryRelease) ([]github.ReleaseAsset, error)
	listReleaseAssetsMutex       sync.RWMutex
	listReleaseAssetsArgsForCall []struct {
		release github.RepositoryRelease
	}
	listReleaseAssetsReturns struct {
		result1 []github.ReleaseAsset
		result2 error
	}
	UploadReleaseAssetStub        func(release github.RepositoryRelease, name string, file *os.File) error
	uploadReleaseAssetMutex       sync.RWMutex
	uploadReleaseAssetArgsForCall []struct {
		release github.RepositoryRelease
		name    string
		file    *os.File
	}
	uploadReleaseAssetReturns struct {
		result1 error
	}
	DeleteReleaseAssetStub        func(asset github.ReleaseAsset) error
	deleteReleaseAssetMutex       sync.RWMutex
	deleteReleaseAssetArgsForCall []struct {
		asset github.ReleaseAsset
	}
	deleteReleaseAssetReturns struct {
		result1 error
	}
	DownloadReleaseAssetStub        func(asset github.ReleaseAsset) (io.ReadCloser, error)
	downloadReleaseAssetMutex       sync.RWMutex
	downloadReleaseAssetArgsForCall []struct {
		asset github.ReleaseAsset
	}
	downloadReleaseAssetReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	GetTarballLinkStub        func(tag string) (*url.URL, error)
	getTarballLinkMutex       sync.RWMutex
	getTarballLinkArgsForCall []struct {
		tag string
	}
	getTarballLinkReturns struct {
		result1 *url.URL
		result2 error
	}
	GetZipballLinkStub        func(tag string) (*url.URL, error)
	getZipballLinkMutex       sync.RWMutex
	getZipballLinkArgsForCall []struct {
		tag string
	}
	getZipballLinkReturns struct {
		result1 *url.URL
		result2 error
	}
}

func (fake *FakeGitHub) ListReleases() ([]github.RepositoryRelease, error) {
	fake.listReleasesMutex.Lock()
	fake.listReleasesArgsForCall = append(fake.listReleasesArgsForCall, struct{}{})
	fake.listReleasesMutex.Unlock()
	if fake.ListReleasesStub != nil {
		return fake.ListReleasesStub()
	} else {
		return fake.listReleasesReturns.result1, fake.listReleasesReturns.result2
	}
}

func (fake *FakeGitHub) ListReleasesCallCount() int {
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	return len(fake.listReleasesArgsForCall)
}

func (fake *FakeGitHub) ListReleasesReturns(result1 []github.RepositoryRelease, result2 error) {
	fake.ListReleasesStub = nil
	fake.listReleasesReturns = struct {
		result1 []github.RepositoryRelease
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) GetReleaseByTag(tag string) (*github.RepositoryRelease, error) {
	fake.getReleaseByTagMutex.Lock()
	fake.getReleaseByTagArgsForCall = append(fake.getReleaseByTagArgsForCall, struct {
		tag string
	}{tag})
	fake.getReleaseByTagMutex.Unlock()
	if fake.GetReleaseByTagStub != nil {
		return fake.GetReleaseByTagStub(tag)
	} else {
		return fake.getReleaseByTagReturns.result1, fake.getReleaseByTagReturns.result2
	}
}

func (fake *FakeGitHub) GetReleaseByTagCallCount() int {
	fake.getReleaseByTagMutex.RLock()
	defer fake.getReleaseByTagMutex.RUnlock()
	return len(fake.getReleaseByTagArgsForCall)
}

func (fake *FakeGitHub) GetReleaseByTagArgsForCall(i int) string {
	fake.getReleaseByTagMutex.RLock()
	defer fake.getReleaseByTagMutex.RUnlock()
	return fake.getReleaseByTagArgsForCall[i].tag
}

func (fake *FakeGitHub) GetReleaseByTagReturns(result1 *github.RepositoryRelease, result2 error) {
	fake.GetReleaseByTagStub = nil
	fake.getReleaseByTagReturns = struct {
		result1 *github.RepositoryRelease
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) CreateRelease(release github.RepositoryRelease) (*github.RepositoryRelease, error) {
	fake.createReleaseMutex.Lock()
	fake.createReleaseArgsForCall = append(fake.createReleaseArgsForCall, struct {
		release github.RepositoryRelease
	}{release})
	fake.createReleaseMutex.Unlock()
	if fake.CreateReleaseStub != nil {
		return fake.CreateReleaseStub(release)
	} else {
		return fake.createReleaseReturns.result1, fake.createReleaseReturns.result2
	}
}

func (fake *FakeGitHub) CreateReleaseCallCount() int {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return len(fake.createReleaseArgsForCall)
}

func (fake *FakeGitHub) CreateReleaseArgsForCall(i int) github.RepositoryRelease {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return fake.createReleaseArgsForCall[i].release
}

func (fake *FakeGitHub) CreateReleaseReturns(result1 *github.RepositoryRelease, result2 error) {
	fake.CreateReleaseStub = nil
	fake.createReleaseReturns = struct {
		result1 *github.RepositoryRelease
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) UpdateRelease(release github.RepositoryRelease) (*github.RepositoryRelease, error) {
	fake.updateReleaseMutex.Lock()
	fake.updateReleaseArgsForCall = append(fake.updateReleaseArgsForCall, struct {
		release github.RepositoryRelease
	}{release})
	fake.updateReleaseMutex.Unlock()
	if fake.UpdateReleaseStub != nil {
		return fake.UpdateReleaseStub(release)
	} else {
		return fake.updateReleaseReturns.result1, fake.updateReleaseReturns.result2
	}
}

func (fake *FakeGitHub) UpdateReleaseCallCount() int {
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	return len(fake.updateReleaseArgsForCall)
}

func (fake *FakeGitHub) UpdateReleaseArgsForCall(i int) github.RepositoryRelease {
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	return fake.updateReleaseArgsForCall[i].release
}

func (fake *FakeGitHub) UpdateReleaseReturns(result1 *github.RepositoryRelease, result2 error) {
	fake.UpdateReleaseStub = nil
	fake.updateReleaseReturns = struct {
		result1 *github.RepositoryRelease
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) ListReleaseAssets(release github.RepositoryRelease) ([]github.ReleaseAsset, error) {
	fake.listReleaseAssetsMutex.Lock()
	fake.listReleaseAssetsArgsForCall = append(fake.listReleaseAssetsArgsForCall, struct {
		release github.RepositoryRelease
	}{release})
	fake.listReleaseAssetsMutex.Unlock()
	if fake.ListReleaseAssetsStub != nil {
		return fake.ListReleaseAssetsStub(release)
	} else {
		return fake.listReleaseAssetsReturns.result1, fake.listReleaseAssetsReturns.result2
	}
}

func (fake *FakeGitHub) ListReleaseAssetsCallCount() int {
	fake.listReleaseAssetsMutex.RLock()
	defer fake.listReleaseAssetsMutex.RUnlock()
	return len(fake.listReleaseAssetsArgsForCall)
}

func (fake *FakeGitHub) ListReleaseAssetsArgsForCall(i int) github.RepositoryRelease {
	fake.listReleaseAssetsMutex.RLock()
	defer fake.listReleaseAssetsMutex.RUnlock()
	return fake.listReleaseAssetsArgsForCall[i].release
}

func (fake *FakeGitHub) ListReleaseAssetsReturns(result1 []github.ReleaseAsset, result2 error) {
	fake.ListReleaseAssetsStub = nil
	fake.listReleaseAssetsReturns = struct {
		result1 []github.ReleaseAsset
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) UploadReleaseAsset(release github.RepositoryRelease, name string, file *os.File) error {
	fake.uploadReleaseAssetMutex.Lock()
	fake.uploadReleaseAssetArgsForCall = append(fake.uploadReleaseAssetArgsForCall, struct {
		release github.RepositoryRelease
		name    string
		file    *os.File
	}{release, name, file})
	fake.uploadReleaseAssetMutex.Unlock()
	if fake.UploadReleaseAssetStub != nil {
		return fake.UploadReleaseAssetStub(release, name, file)
	} else {
		return fake.uploadReleaseAssetReturns.result1
	}
}

func (fake *FakeGitHub) UploadReleaseAssetCallCount() int {
	fake.uploadReleaseAssetMutex.RLock()
	defer fake.uploadReleaseAssetMutex.RUnlock()
	return len(fake.uploadReleaseAssetArgsForCall)
}

func (fake *FakeGitHub) UploadReleaseAssetArgsForCall(i int) (github.RepositoryRelease, string, *os.File) {
	fake.uploadReleaseAssetMutex.RLock()
	defer fake.uploadReleaseAssetMutex.RUnlock()
	return fake.uploadReleaseAssetArgsForCall[i].release, fake.uploadReleaseAssetArgsForCall[i].name, fake.uploadReleaseAssetArgsForCall[i].file
}

func (fake *FakeGitHub) UploadReleaseAssetReturns(result1 error) {
	fake.UploadReleaseAssetStub = nil
	fake.uploadReleaseAssetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitHub) DeleteReleaseAsset(asset github.ReleaseAsset) error {
	fake.deleteReleaseAssetMutex.Lock()
	fake.deleteReleaseAssetArgsForCall = append(fake.deleteReleaseAssetArgsForCall, struct {
		asset github.ReleaseAsset
	}{asset})
	fake.deleteReleaseAssetMutex.Unlock()
	if fake.DeleteReleaseAssetStub != nil {
		return fake.DeleteReleaseAssetStub(asset)
	} else {
		return fake.deleteReleaseAssetReturns.result1
	}
}

func (fake *FakeGitHub) DeleteReleaseAssetCallCount() int {
	fake.deleteReleaseAssetMutex.RLock()
	defer fake.deleteReleaseAssetMutex.RUnlock()
	return len(fake.deleteReleaseAssetArgsForCall)
}

func (fake *FakeGitHub) DeleteReleaseAssetArgsForCall(i int) github.ReleaseAsset {
	fake.deleteReleaseAssetMutex.RLock()
	defer fake.deleteReleaseAssetMutex.RUnlock()
	return fake.deleteReleaseAssetArgsForCall[i].asset
}

func (fake *FakeGitHub) DeleteReleaseAssetReturns(result1 error) {
	fake.DeleteReleaseAssetStub = nil
	fake.deleteReleaseAssetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitHub) DownloadReleaseAsset(asset github.ReleaseAsset) (io.ReadCloser, error) {
	fake.downloadReleaseAssetMutex.Lock()
	fake.downloadReleaseAssetArgsForCall = append(fake.downloadReleaseAssetArgsForCall, struct {
		asset github.ReleaseAsset
	}{asset})
	fake.downloadReleaseAssetMutex.Unlock()
	if fake.DownloadReleaseAssetStub != nil {
		return fake.DownloadReleaseAssetStub(asset)
	} else {
		return fake.downloadReleaseAssetReturns.result1, fake.downloadReleaseAssetReturns.result2
	}
}

func (fake *FakeGitHub) DownloadReleaseAssetCallCount() int {
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	return len(fake.downloadReleaseAssetArgsForCall)
}

func (fake *FakeGitHub) DownloadReleaseAssetArgsForCall(i int) github.ReleaseAsset {
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	return fake.downloadReleaseAssetArgsForCall[i].asset
}

func (fake *FakeGitHub) DownloadReleaseAssetReturns(result1 io.ReadCloser, result2 error) {
	fake.DownloadReleaseAssetStub = nil
	fake.downloadReleaseAssetReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) GetTarballLink(tag string) (*url.URL, error) {
	fake.getTarballLinkMutex.Lock()
	fake.getTarballLinkArgsForCall = append(fake.getTarballLinkArgsForCall, struct {
		tag string
	}{tag})
	fake.getTarballLinkMutex.Unlock()
	if fake.GetReleaseByTagStub != nil {
		return fake.GetTarballLinkStub(tag)
	} else {
		return fake.getTarballLinkReturns.result1, fake.getTarballLinkReturns.result2
	}
}

func (fake *FakeGitHub) GetTarballLinkCallCount() int {
	fake.getTarballLinkMutex.RLock()
	defer fake.getTarballLinkMutex.RUnlock()
	return len(fake.getTarballLinkArgsForCall)
}

func (fake *FakeGitHub) GetTarballLinkArgsForCall(i int) string {
	fake.getTarballLinkMutex.RLock()
	defer fake.getTarballLinkMutex.RUnlock()
	return fake.getTarballLinkArgsForCall[i].tag
}

func (fake *FakeGitHub) GetTarballLinkReturns(result1 *url.URL, result2 error) {
	fake.GetTarballLinkStub = nil
	fake.getTarballLinkReturns = struct {
		result1 *url.URL
		result2 error
	}{result1, result2}
}

func (fake *FakeGitHub) GetZipballLink(tag string) (*url.URL, error) {
	fake.getZipballLinkMutex.Lock()
	fake.getZipballLinkArgsForCall = append(fake.getZipballLinkArgsForCall, struct {
		tag string
	}{tag})
	fake.getZipballLinkMutex.Unlock()
	if fake.GetReleaseByTagStub != nil {
		return fake.GetZipballLinkStub(tag)
	} else {
		return fake.getZipballLinkReturns.result1, fake.getZipballLinkReturns.result2
	}
}

func (fake *FakeGitHub) GetZipballLinkCallCount() int {
	fake.getZipballLinkMutex.RLock()
	defer fake.getZipballLinkMutex.RUnlock()
	return len(fake.getZipballLinkArgsForCall)
}

func (fake *FakeGitHub) GetZipballLinkArgsForCall(i int) string {
	fake.getZipballLinkMutex.RLock()
	defer fake.getZipballLinkMutex.RUnlock()
	return fake.getZipballLinkArgsForCall[i].tag
}

func (fake *FakeGitHub) GetZipballLinkReturns(result1 *url.URL, result2 error) {
	fake.GetZipballLinkStub = nil
	fake.getZipballLinkReturns = struct {
		result1 *url.URL
		result2 error
	}{result1, result2}
}

var _ resource.GitHub = new(FakeGitHub)
