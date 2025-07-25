package reaction

import (
	"github.com/pocketbase/pocketbase"

	"github.com/lagzenthakuri/secyourflow/reaction/schedule"
	"github.com/lagzenthakuri/secyourflow/reaction/trigger/hook"
	"github.com/lagzenthakuri/secyourflow/reaction/trigger/webhook"
)

func BindHooks(pb *pocketbase.PocketBase, test bool) {
	schedule.Start(pb)
	hook.BindHooks(pb, test)
	webhook.BindHooks(pb)
}
