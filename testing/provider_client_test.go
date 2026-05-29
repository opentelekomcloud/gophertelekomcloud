package testing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestAuthenticatedHeaders(t *testing.T) {
	p := &golangsdk.ProviderClient{
		TokenID: "1234",
	}
	expected := map[string]string{"X-Auth-Token": "1234"}
	actual := p.AuthenticatedHeaders()
	th.CheckDeepEquals(t, expected, actual)
}

func TestUserAgent(t *testing.T) {
	p := &golangsdk.ProviderClient{}

	p.UserAgent.Prepend("custom-user-agent/2.4.0")
	expected := "custom-user-agent/2.4.0 golangsdk/2.0.0"
	actual := p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)

	p.UserAgent.Prepend("another-custom-user-agent/0.3.0", "a-third-ua/5.9.0")
	expected = "another-custom-user-agent/0.3.0 a-third-ua/5.9.0 custom-user-agent/2.4.0 golangsdk/2.0.0"
	actual = p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)

	p.UserAgent = golangsdk.UserAgent{}
	expected = "golangsdk/2.0.0"
	actual = p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)
}

func TestConcurrentReauth(t *testing.T) {
	var info = struct {
		numreauths int
		mut        *sync.RWMutex
	}{
		0,
		new(sync.RWMutex),
	}

	numconc := 20

	prereauthTok := client.TokenID
	postreauthTok := "12345678"

	p := new(golangsdk.ProviderClient)
	p.UseTokenLock()
	p.SetToken(prereauthTok)
	p.ReauthFunc = func() error {
		time.Sleep(1 * time.Second)
		p.AuthenticatedHeaders()
		info.mut.Lock()
		info.numreauths++
		info.mut.Unlock()
		p.TokenID = postreauthTok
		return nil
	}

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != postreauthTok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		info.mut.RLock()
		hasReauthed := info.numreauths != 0
		info.mut.RUnlock()

		if hasReauthed {
			th.CheckEquals(t, p.Token(), postreauthTok)
		}

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{}`)
	})

	wg := new(sync.WaitGroup)
	reqopts := new(golangsdk.RequestOpts)

	for i := 0; i < numconc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := p.Request("GET", fmt.Sprintf("%s/route", th.Endpoint()), reqopts)
			th.CheckNoErr(t, err)
			if resp == nil {
				t.Errorf("got a nil response")
				return
			}
			if resp.Body == nil {
				t.Errorf("response body was nil")
				return
			}
			defer func() { _ = resp.Body.Close() }()
			actual, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("error reading response body: %s", err)
				return
			}
			th.CheckByteArrayEquals(t, []byte(`{}`), actual)
		}()
	}

	wg.Wait()

	th.AssertEquals(t, 1, info.numreauths)
}

func TestReauthOnForbiddenInvalidAuthToken(t *testing.T) {
	prereauthTok := client.TokenID
	postreauthTok := "12345678"

	p := new(golangsdk.ProviderClient)
	p.SetToken(prereauthTok)

	numreauths := 0
	p.ReauthFunc = func() error {
		numreauths++
		p.SetToken(postreauthTok)
		return nil
	}

	th.SetupHTTP()
	defer th.TeardownHTTP()

	numrequests := 0
	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		numrequests++
		switch r.Header.Get("X-Auth-Token") {
		case prereauthTok:
			w.WriteHeader(http.StatusForbidden)
			_, _ = fmt.Fprint(w, `{"message":"X-Auth-Token is invalid"}`)
		case postreauthTok:
			w.Header().Add("Content-Type", "application/json")
			_, _ = fmt.Fprint(w, `{}`)
		default:
			t.Errorf("unexpected X-Auth-Token: %s", r.Header.Get("X-Auth-Token"))
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	resp, err := p.Request("GET", fmt.Sprintf("%s/route", th.Endpoint()), new(golangsdk.RequestOpts))
	th.CheckNoErr(t, err)
	if resp == nil {
		t.Fatalf("got a nil response")
	}
	defer func() { _ = resp.Body.Close() }()

	actual, err := ioutil.ReadAll(resp.Body)
	th.CheckNoErr(t, err)
	th.CheckByteArrayEquals(t, []byte(`{}`), actual)
	th.AssertEquals(t, 1, numreauths)
	th.AssertEquals(t, 2, numrequests)
}

func TestForbiddenWithoutInvalidAuthTokenDoesNotReauth(t *testing.T) {
	p := new(golangsdk.ProviderClient)
	p.SetToken(client.TokenID)

	numreauths := 0
	p.ReauthFunc = func() error {
		numreauths++
		p.SetToken("12345678")
		return nil
	}

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprint(w, `{"message":"policy does not allow this request"}`)
	})

	resp, err := p.Request("GET", fmt.Sprintf("%s/route", th.Endpoint()), new(golangsdk.RequestOpts))
	if resp == nil {
		t.Fatalf("got a nil response")
	}
	defer func() { _ = resp.Body.Close() }()

	th.AssertEquals(t, 0, numreauths)
	if _, ok := err.(golangsdk.ErrDefault403); !ok {
		t.Fatalf("expected ErrDefault403, got %T: %v", err, err)
	}
}
