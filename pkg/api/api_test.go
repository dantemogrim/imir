package api

import (
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFetch_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://mercuryretrogradeapi.com",
		httpmock.NewStringResponder(200, `{"status": "success", "is_retrograde": true}`))

	_, err := Fetch()
	if err != nil {
		t.Errorf("Fetch() failed, expected nil, got %v", err)
	}
}

func TestFetch_HttpGetError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://mercuryretrogradeapi.com",
		httpmock.NewErrorResponder(errors.New("mocked error")))

	_, err := Fetch()
	if err == nil {
		t.Errorf("Fetch() didn't fail, expected an error, got nil")
	}
}

func TestFetch_DecodeError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://mercuryretrogradeapi.com",
		httpmock.NewStringResponder(200, `{"status": "success", "is_retrograde": "not a boolean"}`))

	_, err := Fetch()
	if err == nil {
		t.Errorf("Fetch() didn't fail, expected an error, got nil")
	}
}

func TestFetchByDateInput_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://mercuryretrogradeapi.com?date=2022-12-31",
		httpmock.NewStringResponder(200, `{"status": "success", "is_retrograde": true}`))

	_, err := FetchByDateInput("2022-12-31")
	if err != nil {
		t.Errorf("FetchByDateInput() failed, expected nil, got %v", err)
	}
}

func TestFetchByDateInput_HttpGetError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://mercuryretrogradeapi.com?date=2022-12-31",
		httpmock.NewErrorResponder(errors.New("mocked error")))

	_, err := FetchByDateInput("2022-12-31")
	if err == nil {
		t.Errorf("FetchByDateInput() didn't fail, expected an error, got nil")
	}
}

func TestFetchByDateInput_DecodeError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://mercuryretrogradeapi.com?date=2022-12-31",
		httpmock.NewStringResponder(200, `{"status": "success", "is_retrograde": "not a boolean"}`))

	_, err := FetchByDateInput("2022-12-31")
	if err == nil {
		t.Errorf("FetchByDateInput() didn't fail, expected an error, got nil")
	}
}
