import axios from "axios";
import { Template } from "../models/types";

const api = axios.create({
  baseURL: "/api",
  timeout: 2000,
});

export async function getTemplates(): Promise<Template[]> {
  const response = await api.get<Template[]>("/templates");
  return response.data;
}
