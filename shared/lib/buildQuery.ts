export const buildQuery = (queryArgs: { [key: string]: string[] }) => {
  const searchParams = new URLSearchParams();

  Object.entries(queryArgs).forEach(([key, values]) => {
    values.forEach((value: string) => {
      searchParams.append(key, value);
    });
  });

  searchParams.append("year", "2026");

  return "?" + searchParams.toString();
}