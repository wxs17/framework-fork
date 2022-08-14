package events

import (
	"github.com/goravel/framework/contracts/console"
	eventConsole "github.com/goravel/framework/events/console"
	"github.com/goravel/framework/support/facades"
)

type ServiceProvider struct {
}

//Boot Bootstrap any application services after register.
func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

//Register any application services.
func (receiver *ServiceProvider) Register() {
	facades.Event = &Application{}
}

//registerCommands Register the given commands.
func (receiver *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console.Command{
		&eventConsole.EventMakeCommand{},
		&eventConsole.ListenerMakeCommand{},
	})
}
