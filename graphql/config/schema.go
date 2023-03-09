/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package config

func Schema() (string, error) {
	return `
type Config {
	id: ID!
	name: String!
	global: Global!
	routing: Routing!
	dns: Dns!
	selected: Boolean!
	referenceGroups: [String!]!
}

type Function {
	name: String!
	not: Boolean!
	params: [Param!]!
}
type Param {
	key: String!
	val: String!
}

type AndFunctions {
	and: [Function!]!
}

type Plaintext {
	val: String!
}

union AndFunctionsOrPlaintext = AndFunctions | Plaintext
union FunctionOrPlaintext = Function | Plaintext

`, nil
}
