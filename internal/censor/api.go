package censor

import "strings"

func (c *Censor) Check(text string) bool {
	for _, bw := range c.cfg.BadWordPatterns {
		if strings.Contains(text, bw) {
			return false
		}
	}
	return true
}
