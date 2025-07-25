package webhook_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"

	"github.com/lagzenthakuri/secyourflow/reaction/action/webhook"
	catalystTesting "github.com/lagzenthakuri/secyourflow/testing"
)

func TestWebhook_Run(t *testing.T) {
	t.Parallel()

	server := catalystTesting.NewRecordingServer()

	go http.ListenAndServe("127.0.0.1:12347", server) //nolint:gosec,errcheck

	if err := catalystTesting.WaitForStatus("http://127.0.0.1:12347/health", http.StatusOK, 5*time.Second); err != nil {
		t.Fatal(err)
	}

	type fields struct {
		Headers map[string]string
		URL     string
	}

	type args struct {
		payload string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]any
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			fields: fields{
				Headers: map[string]string{},
				URL:     "http://127.0.0.1:12347/foo",
			},
			args: args{
				payload: "test",
			},
			want: map[string]any{
				"statusCode": 200,
				"headers": map[string]any{
					"Content-Length": []any{"14"},
					"Content-Type":   []any{"application/json; charset=UTF-8"},
				},
				"body":            "{\"test\":true}\n",
				"isBase64Encoded": false,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			a := &webhook.Webhook{
				Headers: tt.fields.Headers,
				URL:     tt.fields.URL,
			}
			got, err := a.Run(ctx, tt.args.payload)
			tt.wantErr(t, err)

			want, err := json.Marshal(tt.want)
			require.NoError(t, err)

			got, err = sjson.DeleteBytes(got, "headers.Date")
			require.NoError(t, err)

			assert.JSONEq(t, string(want), string(got))
		})
	}
}
