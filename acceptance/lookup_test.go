package acceptance

import (
	"os"
	"testing"

	"github.com/magicmemories/go-jerakia"
	fixtures "github.com/magicmemories/go-jerakia/testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupBasic(t *testing.T) {
	if v := os.Getenv("JERAKIA_ACC"); v == "" {
		t.Skip("JERAKIA_ACC not set")
	}

	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "test",
	}

	actual, err := jerakia.Lookup(client, "cities", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.LookupBasicResult
	assert.Equal(t, expected, *actual)
}

func TestLookupSingleBoolResult(t *testing.T) {
	if v := os.Getenv("JERAKIA_ACC"); v == "" {
		t.Skip("JERAKIA_ACC not set")
	}

	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "test",
	}

	actual, err := jerakia.Lookup(client, "booltrue", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.LookupSingleBoolResult
	assert.Equal(t, expected, *actual)
}

func TestLookupMetadata(t *testing.T) {
	if v := os.Getenv("JERAKIA_ACC"); v == "" {
		t.Skip("JERAKIA_ACC not set")
	}

	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "test",
		Metadata: map[string]string{
			"hostname": "example",
		},
	}

	actual, err := jerakia.Lookup(client, "users", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.LookupMetadataResult
	assert.Equal(t, expected, *actual)
}

func TestLookupHashMerge(t *testing.T) {
	if v := os.Getenv("JERAKIA_ACC"); v == "" {
		t.Skip("JERAKIA_ACC not set")
	}

	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "test",
		Metadata: map[string]string{
			"env": "dev",
		},
		LookupType: "cascade",
		Merge: "hash",
	}

	actual, err := jerakia.Lookup(client, "hash", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.LookupHashMergeResult
	assert.Equal(t, expected, *actual)
}

func TestLookupKeyless(t *testing.T) {
	if v := os.Getenv("JERAKIA_ACC"); v == "" {
		t.Skip("JERAKIA_ACC not set")
	}

	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "keyless",
	}

	actual, err := jerakia.Lookup(client, "", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.LookupKeylessResult
	assert.Equal(t, expected, *actual)
}