import axios from "axios";
import { API_PATH } from "@/configs/constants";

export async function useCheckAuthorized() {
	 const response = await axios.post(API_PATH + "/auth/login", {}, {
		 withCredentials: true
	 });

	console.log(response);

	return response
}