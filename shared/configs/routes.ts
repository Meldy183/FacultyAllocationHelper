type routeType = {
	routeName: string,
	routePath: string,
}

export const routesAuth: routeType[] = [
	{
		routeName: "Registration",
		routePath: "/registration"
	},
	{
		routeName: "Login",
		routePath: "/login"
	}
]

export const dashboardRoute: routeType = {
	routeName: "Dashboard",
	routePath: "/dashboard"
}

export const Routes = [...routesAuth];