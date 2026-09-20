
package gui

import "fmt"

type Button interface {
	Paint() string
}

type Checkbox interface {
	Paint() string
}


type WindowsButton struct{}

func (WindowsButton) Paint() string { return "Rendering Windows button" }

type WindowsCheckbox struct{}

func (WindowsCheckbox) Paint() string { return "Rendering Windows checkbox" }

// --- macOS family (Concrete Products) ---

type MacOSButton struct{}

func (MacOSButton) Paint() string { return "Rendering macOS button" }

type MacOSCheckbox struct{}

func (MacOSCheckbox) Paint() string { return "Rendering macOS checkbox" }


type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsFactory struct{}

func (WindowsFactory) CreateButton() Button     { return WindowsButton{} }
func (WindowsFactory) CreateCheckbox() Checkbox { return WindowsCheckbox{} }

// MacOSFactory is a Concrete Factory producing the macOS family.
type MacOSFactory struct{}

func (MacOSFactory) CreateButton() Button     { return MacOSButton{} }
func (MacOSFactory) CreateCheckbox() Checkbox { return MacOSCheckbox{} }



func NewFactory(platform string) (GUIFactory, error) {
	switch platform {
	case "WINDOWS":
		return WindowsFactory{}, nil
	case "MACOS":
		return MacOSFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported UI platform %q: expected WINDOWS or MACOS", platform)
	}
}
