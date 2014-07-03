package models

type BasicPage struct {
	Title    string
	Info     []string
	Warnings []string
	Errors   []string
	Debug    []string
}

const (
	BasicPageArrayIncr = 5
)

//-------------------------------------------------------------------------------------------------

func (p *BasicPage) Init(title string) {
	p.Title = title
}

//-------------------------------------------------------------------------------------------------

func (p *BasicPage) AddError(item string) {
	n := len(p.Errors)
	if cap(p.Errors) == n {
		p.Errors = make([]string, n+1, n+BasicPageArrayIncr)
	} else {
		p.Errors = p.Errors[0 : n+1]
	}
	p.Errors[n] = item
}

//-------------------------------------------------------------------------------------------------

func (p *BasicPage) AddWarning(item string) {
	n := len(p.Warnings)
	if cap(p.Warnings) == n {
		p.Warnings = make([]string, n+1, n+BasicPageArrayIncr)
	} else {
		p.Warnings = p.Warnings[0 : n+1]
	}
	p.Warnings[n] = item
}

//-------------------------------------------------------------------------------------------------

func (p *BasicPage) AddInfo(item string) {
	n := len(p.Info)
	if cap(p.Info) == n {
		p.Info = make([]string, n+1, n+BasicPageArrayIncr)
	} else {
		p.Info = p.Info[0 : n+1]
	}
	p.Info[n] = item
}

//-------------------------------------------------------------------------------------------------

func (p *BasicPage) AddDebug(item string) {
	n := len(p.Debug)
	if cap(p.Debug) == n {
		p.Debug = make([]string, n+1, n+BasicPageArrayIncr)
	} else {
		p.Debug = p.Debug[0 : n+1]
	}
	p.Debug[n] = item
}

//-------------------------------------------------------------------------------------------------
