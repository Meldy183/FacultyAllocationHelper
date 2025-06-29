export const getBorderColorByWorkload = (workload: number): string => {
	if (workload > 0.8) return "#FF0000";
	if (workload > 0.65) return "#FF9D00";
	if (workload > 0.4) return "#FFFB00";
	return "#88FF00";
};