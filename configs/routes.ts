type routeType = {
	routeName: string,
	routePath: string,
}

export const routesAuth: routeType[] = [
	{
		routeName: "Registration",
		routePath: "/auth/registration"
	},
	{
		routeName: "Login",
		routePath: "/auth/login"
	}
]

export const Routes = [...routesAuth];