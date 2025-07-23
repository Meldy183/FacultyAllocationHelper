type routeType = {
	routeName: string,
	routePath: string,
}

export const routesAuth: routeType[] = [
	{
		routeName: "Registration",
		routePath: "/index/registration"
	},
	{
		routeName: "Login",
		routePath: "/index/login"
	}
]

export const dashboardRoute: routeType = {
	routeName: "Dashboard",
	routePath: "/dashboard"
}

export const Routes = [...routesAuth];