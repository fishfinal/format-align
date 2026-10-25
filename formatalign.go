// Copyright 2022 fishfinal <contact.fishfinal@gmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formatalign

import (
	"fmt"
	"strconv"
)

type Alignment string

func (alignment Alignment) Is(align Alignment) bool {
	if alignment == align {
		return true
	}

	return false
}

const (
	AlignmentOfRight Alignment = "right"
	AlignmentOfLeft  Alignment = "left"
)

type Option interface {
	apply(*FormatAlign)
}

type FormatAlign struct {
	alignment   Alignment
	placeholder int // field name placeholder length.
	separator   string
}

func (align *FormatAlign) Format(name string, value interface{}) string {
	fieldname := align.fieldname()
	return fmt.Sprintf(fmt.Sprintf(fieldname+" %v", name, value))
}

func (align *FormatAlign) SetSeparator(separator string) *FormatAlign {
	align.separator = separator

	return align
}

func (align *FormatAlign) SetPlaceholder(placeholder int) *FormatAlign {
	align.placeholder = placeholder

	return align
}

func (align *FormatAlign) fieldname() string {
	var indent = "%"
	if align.alignment.Is(AlignmentOfLeft) {
		indent += "-"
	}
	indent += strconv.Itoa(align.placeholder)
	indent += "s"
	indent += align.separator

	return indent
}

func NewFormatAlign(options ...Option) *FormatAlign {
	align := &FormatAlign{
		alignment:   AlignmentOfLeft,
		placeholder: 26,
		separator:   ":",
	}

	for _, o := range options {
		o.apply(align)
	}

	return align
}

type OptionFunc func(formatAlign *FormatAlign)

func (o OptionFunc) apply(align *FormatAlign) {
	o(align)
}

func WithAlignment(alignment Alignment) OptionFunc {
	return func(formatAlign *FormatAlign) {
		formatAlign.alignment = alignment
	}
}

func WithPlaceholder(placeholder int) OptionFunc {
	return func(formatAlign *FormatAlign) {
		formatAlign.placeholder = placeholder
	}
}

func WithSeparator(separator string) OptionFunc {
	return func(formatAlign *FormatAlign) {
		formatAlign.separator = separator
	}
}
