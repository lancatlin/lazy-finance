import axios from "axios";

const api = axios.create({
  baseURL: "/api",
  timeout: 2000,
});

export async function getTemplates(): Promise<Template[]> {
  const response = await api.get("/templates");
  console.log(response.data);
  return response.data;
}
