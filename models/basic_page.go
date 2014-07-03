/*
Meuh
Copyright (C) 2014 Thomas Silvi

This file is part of Meuh.

GoSimpleConfigLib is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

GoSimpleConfigLib is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Foobar. If not, see <http://www.gnu.org/licenses/>.
*/

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
