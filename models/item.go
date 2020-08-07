// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
// 
// SPDX-License-Identifier: MIT
package models

type Item struct {
	ID    string
	Name  string
	Price float32
}

func (i Item) Valid() bool {
	return len(i.Name) > 0 && i.Price >= 0
}
