package jwt_benchmark_test

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"testing"
	"time"
)

var scopeString string = `
{
  "id": 2126244,
  "node_id": "MDEwOlJlcG9zaXRvcnkyMTI2MjQ0",
  "name": "bootstrap",
  "full_name": "twbs/bootstrap",
  "private": false,
  "owner": {
    "login": "twbs",
    "id": 2918581,
    "node_id": "MDEyOk9yZ2FuaXphdGlvbjI5MTg1ODE=",
    "avatar_url": "https://avatars.githubusercontent.com/u/2918581?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/twbs",
    "html_url": "https://github.com/twbs",
    "followers_url": "https://api.github.com/users/twbs/followers",
    "following_url": "https://api.github.com/users/twbs/following{/other_user}",
    "gists_url": "https://api.github.com/users/twbs/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/twbs/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/twbs/subscriptions",
    "organizations_url": "https://api.github.com/users/twbs/orgs",
    "repos_url": "https://api.github.com/users/twbs/repos",
    "events_url": "https://api.github.com/users/twbs/events{/privacy}",
    "received_events_url": "https://api.github.com/users/twbs/received_events",
    "type": "Organization",
    "site_admin": false
  },
  "html_url": "https://github.com/twbs/bootstrap",
  "description": "The most popular HTML, CSS, and JavaScript framework for developing responsive, mobile first projects on the web.",
  "fork": false,
  "url": "https://api.github.com/repos/twbs/bootstrap",
  "forks_url": "https://api.github.com/repos/twbs/bootstrap/forks",
  "keys_url": "https://api.github.com/repos/twbs/bootstrap/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/twbs/bootstrap/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/twbs/bootstrap/teams",
  "hooks_url": "https://api.github.com/repos/twbs/bootstrap/hooks",
  "issue_events_url": "https://api.github.com/repos/twbs/bootstrap/issues/events{/number}",
  "events_url": "https://api.github.com/repos/twbs/bootstrap/events",
  "assignees_url": "https://api.github.com/repos/twbs/bootstrap/assignees{/user}",
  "branches_url": "https://api.github.com/repos/twbs/bootstrap/branches{/branch}",
  "tags_url": "https://api.github.com/repos/twbs/bootstrap/tags",
  "blobs_url": "https://api.github.com/repos/twbs/bootstrap/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/twbs/bootstrap/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/twbs/bootstrap/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/twbs/bootstrap/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/twbs/bootstrap/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/twbs/bootstrap/languages",
  "stargazers_url": "https://api.github.com/repos/twbs/bootstrap/stargazers",
  "contributors_url": "https://api.github.com/repos/twbs/bootstrap/contributors",
  "subscribers_url": "https://api.github.com/repos/twbs/bootstrap/subscribers",
  "subscription_url": "https://api.github.com/repos/twbs/bootstrap/subscription",
  "commits_url": "https://api.github.com/repos/twbs/bootstrap/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/twbs/bootstrap/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/twbs/bootstrap/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/twbs/bootstrap/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/twbs/bootstrap/contents/{+path}",
  "compare_url": "https://api.github.com/repos/twbs/bootstrap/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/twbs/bootstrap/merges",
  "archive_url": "https://api.github.com/repos/twbs/bootstrap/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/twbs/bootstrap/downloads",
  "issues_url": "https://api.github.com/repos/twbs/bootstrap/issues{/number}",
  "pulls_url": "https://api.github.com/repos/twbs/bootstrap/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/twbs/bootstrap/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/twbs/bootstrap/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/twbs/bootstrap/labels{/name}",
  "releases_url": "https://api.github.com/repos/twbs/bootstrap/releases{/id}",
  "deployments_url": "https://api.github.com/repos/twbs/bootstrap/deployments",
  "created_at": "2011-07-29T21:19:00Z",
  "updated_at": "2021-02-14T14:36:44Z",
  "pushed_at": "2021-02-14T16:02:20Z",
  "git_url": "git://github.com/twbs/bootstrap.git",
  "ssh_url": "git@github.com:twbs/bootstrap.git",
  "clone_url": "https://github.com/twbs/bootstrap.git",
  "svn_url": "https://github.com/twbs/bootstrap",
  "homepage": "https://getbootstrap.com",
  "size": 175400,
  "stargazers_count": 148213,
  "watchers_count": 148213,
  "language": "JavaScript",
  "has_issues": true,
  "has_projects": true,
  "has_downloads": true,
  "has_wiki": false,
  "has_pages": true,
  "forks_count": 72238,
  "mirror_url": null,
  "archived": false,
  "disabled": false,
  "open_issues_count": 324,
  "license": {
    "key": "mit",
    "name": "MIT License",
    "spdx_id": "MIT",
    "url": "https://api.github.com/licenses/mit",
    "node_id": "MDc6TGljZW5zZTEz"
  },
  "forks": 72238,
  "open_issues": 324,
  "watchers": 148213,
  "default_branch": "main",
  "temp_clone_token": null,
  "organization": {
    "login": "twbs",
    "id": 2918581,
    "node_id": "MDEyOk9yZ2FuaXphdGlvbjI5MTg1ODE=",
    "avatar_url": "https://avatars.githubusercontent.com/u/2918581?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/twbs",
    "html_url": "https://github.com/twbs",
    "followers_url": "https://api.github.com/users/twbs/followers",
    "following_url": "https://api.github.com/users/twbs/following{/other_user}",
    "gists_url": "https://api.github.com/users/twbs/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/twbs/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/twbs/subscriptions",
    "organizations_url": "https://api.github.com/users/twbs/orgs",
    "repos_url": "https://api.github.com/users/twbs/repos",
    "events_url": "https://api.github.com/users/twbs/events{/privacy}",
    "received_events_url": "https://api.github.com/users/twbs/received_events",
    "type": "Organization",
    "site_admin": false
  },
  "network_count": 72238,
  "subscribers_count": 7142
}
`

func buildClaim() jwt.MapClaims {
	claim := jwt.MapClaims{}
	claim["iss"] = "test iss"
	claim["sub"] = "test sub"
	claim["aud"] = "test aud"
	t := time.Now()
	claim["iat"] = t.Unix()
	claim["exp"] = t.Add(5 * time.Minute)
	claim["scope"] = scopeString
	return claim
}

func readRsaKey4096() *rsa.PrivateKey {
	bytes, err := ioutil.ReadFile("jwtRS256_4096.key.pem")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
	return key
}

func readRsaPubKey4096() *rsa.PublicKey {
	bytes, err := ioutil.ReadFile("jwtRS256_4096.key.pub.pem")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
	return key
}

func readRsaKey2048() *rsa.PrivateKey {
	bytes, err := ioutil.ReadFile("jwtRS256_2048.key.pem")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
	return key
}

func readRsaPubKey2048() *rsa.PublicKey {
	bytes, err := ioutil.ReadFile("jwtRS256_2048.key.pub.pem")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
	return key
}

func readRsaKey1024() *rsa.PrivateKey {
	bytes, err := ioutil.ReadFile("jwtRS256_1024.key.pem")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
	return key
}

func readRsaPubKey1024() *rsa.PublicKey {
	bytes, err := ioutil.ReadFile("jwtRS256_1024.key.pub.pem")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
	return key
}

func BenchmarkHS256(b *testing.B) {
	claim := buildClaim()
	b.ResetTimer()
	key := []byte("CPxcsmJMj8AX5AGMFEQTwTtKwsGws4Qc")
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseHS256(b *testing.B) {
	claim := buildClaim()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	key := []byte("CPxcsmJMj8AX5AGMFEQTwTtKwsGws4Qc")
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	}
}

func BenchmarkHS384(b *testing.B) {
	claim := buildClaim()
	key := []byte("FXfbmsJY5jyUk974GfJ9ZFf8hWxn5_sDPJu2KPNCwGt2PLBC")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodHS384, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseHS384(b *testing.B) {
	claim := buildClaim()
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claim)
	key := []byte("FXfbmsJY5jyUk974GfJ9ZFf8hWxn5_sDPJu2KPNCwGt2PLBC")
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	}
}

func BenchmarkHS512(b *testing.B) {
	claim := buildClaim()
	b.ResetTimer()
	key := []byte("ke5cXinpGY8QEA8x6SEpAPzCfhS7MtfkxR6p8mtEYmErPn5WNT4wpafLZMWEVfgx")
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseHS512(b *testing.B) {
	claim := buildClaim()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	key := []byte("ke5cXinpGY8QEA8x6SEpAPzCfhS7MtfkxR6p8mtEYmErPn5WNT4wpafLZMWEVfgx")
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	}
}

func BenchmarkRS256_4096(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey4096()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS256_4096(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey4096()
	pubKey := readRsaPubKey4096()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS384_4096(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey4096()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS384, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS384_4096(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey4096()
	pubKey := readRsaPubKey4096()
	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS512_4096(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey4096()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS512_4096(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey4096()
	pubKey := readRsaPubKey4096()
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS256_2048(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey2048()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS256_2048(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey2048()
	pubKey := readRsaPubKey2048()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS384_2048(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey2048()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS384, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS384_2048(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey2048()
	pubKey := readRsaPubKey2048()
	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS512_2048(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey2048()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS512_2048(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey2048()
	pubKey := readRsaPubKey2048()
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS256_1024(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey1024()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS256_1024(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey1024()
	pubKey := readRsaPubKey1024()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS384_1024(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey1024()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS384, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS384_1024(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey1024()
	pubKey := readRsaPubKey1024()
	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}

func BenchmarkRS512_1024(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey1024()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
		token.SignedString(key)
	}
}

func BenchmarkParseRS512_1024(b *testing.B) {
	claim := buildClaim()
	key := readRsaKey1024()
	pubKey := readRsaPubKey1024()
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
	tokenString, _ := token.SignedString(key)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
	}
}
