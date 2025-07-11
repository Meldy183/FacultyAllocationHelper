export const buildQuery = (queryArgs: { [key: string]: string[] }) => {
  const searchParams = new URLSearchParams();

  Object.entries(queryArgs).forEach(([key, values]) => {
    values.forEach((value: string) => {
      searchParams.append(key, value);
    });
  });

  return "?" + searchParams.toString();
}