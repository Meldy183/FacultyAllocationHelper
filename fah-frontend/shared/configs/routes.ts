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

export const dashboardRoute: routeType = {
	routeName: "Dashboard",
	routePath: "/dashboard"
}

export const Routes = [...routesAuth];