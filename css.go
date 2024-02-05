package gocss

import (
	"strconv"
	"strings"
)

type CSS map[string]Stringer
type Selectors map[string]CSS

func (s Selectors) String() string {
	var b strings.Builder
	for k, v := range s {
		b.WriteString(k)
		b.WriteByte('{')
		b.WriteString(v.String())
		b.WriteByte('}')
	}
	return b.String()
}

type Stringer interface {
	String() string
}

type StringerWrapper struct {
	s string
}

func (sw *StringerWrapper) String() string {
	return sw.s
}

func String(s string) Stringer {
	return &StringerWrapper{s: s}
}

func Px(px ...int) Stringer {
	var b strings.Builder
	for i, x := range px {
		b.WriteString(strconv.Itoa(x))
		b.WriteString("px")
		if i < len(px)-1 {
			b.WriteByte(' ')
		}
	}
	return &b
}

func (css CSS) String() string {
	var b strings.Builder
	var selectors string
	for k, v := range css {
		if k == "selectors" {
			if cs, ok := v.(Selectors); ok {
				selectors = cs.String()
			}
		} else {
			b.WriteString(k)
			b.WriteByte(':')
			b.WriteString(v.String())
			b.WriteByte(';')
		}
	}
	b.WriteString(selectors)
	return b.String()
}

// let CAP_LETTER_RE = new RegExp('[A-Z]', 'g');

// let addDash = (m: string) => '-' + m.toLowerCase();

// export let camelTokebab = (s: string): string => {
//   if (s !== s.toLowerCase()) {
//     s = s.replace(CAP_LETTER_RE, addDash);
//   }
//   return s;
// }

// export let cssToString = (css: CSS): string => {
//   let textContent = '';
//   for (let prop in css) {
//     if (prop !== 'selectors') {
//       textContent += camelTokebab(prop) + ':' + css[prop] + ';';
//     }
//   }
//   for (let selector in css.selectors) {
//     textContent += selector + '{' + cssToString(css.selectors[selector]) + '}';
//   }
//   return textContent;
// };

// let id = 0;

// export let generateId = (): string =>  '_' + (id++).toString(32);
