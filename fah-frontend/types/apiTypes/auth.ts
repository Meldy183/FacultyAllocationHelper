// RegisterInterface, LoginInterface, RefreshInterface, LogoutInterface

export type RegisterProcessType = {
	request: {
		email: string,
		password: string,
	},
	response: {
		message: string
	}
}

export type LoginProcessType = {
	request: {
		email: string,
		password: string,
	},
	response: {
		message: string
	}
}

export type RefreshProcessType = {
	request: { },
	response: {
		message: string
	}
}

export type LogoutProcessType = {
	request: { },
	response: {
		message: string
	}
}